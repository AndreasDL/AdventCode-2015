import sys
from multiprocessing import Pool

POOL_SIZE = 10
BLOCK_SIZE = 1000000

class Disc:
	startPos = 0
	posCount = 0

	def __init__(self,startPos,posCount):
		self.startPos = startPos
		self.posCount = posCount

	def fallsThroughAt(self, time):
		return 0 == (self.startPos + time) % self.posCount

def parseLine(line):
	instr = line.strip().split(" ")
	return Disc(
		posCount = int(instr[3]),
		startPos = int(instr[-1][:-1])
	)

def getDiscs(fname):
	lines = None
	with open(fname) as f:
		lines = f.readlines()

	discs = list()
	for line in lines:
		discs.append(parseLine(line))

	return discs

def search(start, stop):
	time = start
	length = len(discs)
	
	found = False
	while not found and time < stop:
		j = 0
		while j < length and discs[j].fallsThroughAt(time + j + 1):
			j+=1

		if j == length:
			print(time)
			found = True
			
		else:
			time += 1

		if time % 100000 == 0:
			print(time)

	return time if found else None

def worker(block):
	start = block * BLOCK_SIZE
	stop  = start + BLOCK_SIZE

	return search(start,stop)

discs = getDiscs("realInput.txt") #global variable => readonly use only

if __name__ == "__main__":

	i = 0
	found = False
	while not found:
		
		pool = Pool(processes=POOL_SIZE)
		job_results = pool.map( worker, range(i * POOL_SIZE, (i+1) * POOL_SIZE) )
		pool.close()
		pool.join()

		j = 0
		while j < len(job_results) and job_results[j] is None:
			j+=1

		if j == len(job_results):
			i += 1
		else:
			found = True
			print(job_results[j])
