import numpy as np
from math import floor
import sys

FNAME = "realInput.txt"
MAX_VAL = np.uint32(4294967295)
GRANULARITY1 = 1000000
GRANULARITY2 = 1000

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

def optimizeRestrictions(restrictions, granularity):
	opt_restrictions = dict() #restrictions[r.start % 1 000 000] => list with restrictions that apply in 1000 000 - 1 999 999

	for restriction in restrictions:

		start_index = floor(restriction.start / granularity)
		stop_index  = floor(restriction.stop  / granularity)

		while start_index <= stop_index:

			if start_index not in opt_restrictions:
				opt_restrictions[start_index] = list()

			opt_restrictions[start_index].append(restriction)

			start_index += 1

	return opt_restrictions

if __name__ == "__main__":

	restrictions = getRestrictions(FNAME)

	#first level speedup
	opt_restrictions = optimizeRestrictions(restrictions, GRANULARITY1)
	
	for key in opt_restrictions:
		for r in opt_restrictions[key]:
			print(key, " -> ", r.toStr())

	found = False
	counter = np.uint32(0)
	fl_index = 0 # = floor(counter / GRANULARITY1)
	while counter < MAX_VAL and not found:

		if fl_index in opt_restrictions:
			restrictions_apply = opt_restrictions[fl_index]

			j = 0
			while j < len(restrictions_apply) and restrictions_apply[j].isOutside(counter):
				j += 1

			if j < len(restrictions_apply):
				counter += 1
				if counter % GRANULARITY1 == 0:
					fl_index += 1

			else:
				found = True

		else:
			found = True

		if counter % 100000 == 0:
			print(counter)
	print(counter)