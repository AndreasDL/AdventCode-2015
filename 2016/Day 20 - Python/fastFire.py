import numpy as np
from math import floor
import sys

FNAME = "realInput.txt"
MAX_VAL = np.uint32(4294967295)

class Restriction:
	def __init__(self,start,stop):
		self.start = np.uint32(start)
		self.stop = np.uint32(stop)

	def isOutside(self, value):
		return self.start > value or self.stop < value

	def toStr(self):
		return "restriction [ ", self.start, " ; ", self.stop , " ]"

def parseLineToRestriction(line):
	args = line.split("-")
	return Restriction(args[0], args[1])

def getRestrictions(fname):
	restrictions = list()
	with open(fname) as f:
		for line in f.readlines():
			restrictions.append(parseLineToRestriction(line))
	
	restrictions.sort(key=lambda x: x.start)

	return restrictions

def sortRestrictions(restrictions):

	restrictions.sort(key = lambda x: x.stop)
	restrictions.sort(key = lambda x: x.start) #stable sorting
	
	return restrictions

	
if __name__ == "__main__":

	restrictions = getRestrictions(FNAME)
	restrictions = sortRestrictions(restrictions) #is actually in place

	i = 0
	length = len(restrictions)
	counter = np.uint32(0)
	startval = restrictions[i].start
	furtestPoint = restrictions[i].stop
	while i < length and furtestPoint < MAX_VAL:

		#quickly run over ovelapping ranges
		while i < length and restrictions[i].start <= furtestPoint: 
			furtestPoint = furtestPoint if furtestPoint > restrictions[i].stop else restrictions[i].stop
			print(restrictions[i].toStr(), " ==>> ", startval, " -> ", furtestPoint)
			i += 1

		if i < length:
			counter += np.uint32(restrictions[i].start - furtestPoint - 1) #-2 to correct for edges inclusive

			print("\tGap: ", restrictions[i].start, " -> " , furtestPoint, " = ", counter)
			
			startval = restrictions[i].start
			furtestPoint = restrictions[i].stop
		

	print(counter)

