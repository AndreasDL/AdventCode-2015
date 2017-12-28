from enum import Enum
import sys

class destination(Enum):
	bot = 1
	output = 2	

class outputBin:
	values = None # list(int)

	def __init__(self):
		self.values = list()

	def addValue(self,value):
		self.values.append(value)

class bot:
	lower  = None # tuple(destination, index)
	higher = None # tuple(destination, index)
	values = None # list(val1, val2)

	def __init__(self):
		self.values = list()

	def addValue(self, value):
		self.values.append(value)

	def setLowerHigher(self,lower,higher):
		self.lower = lower
		self.higher = higher

	def distributeIfPossible(self, bots, outputBins):
		retLower = None
		retHigher = None

		if self.canDistribute():
			self.values.sort()
			
			if self.lower[0] == destination.bot: 
				bots[self.lower[1]].addValue(self.values[0])
				retLower = self.lower[1]
			else:
				outputBins[self.lower[1]].addValue(self.values[0])

			if self.higher[0] == destination.bot: 
				bots[self.higher[1]].addValue(self.values[1])
				retHigher = self.higher[1]
			else:
				outputBins[self.higher[1]].addValue(self.values[1])

			#cleanup
			self.higher = None
			self.lower = None
			self.values = list()

		return retLower, retHigher

	def canDistribute(self):
		return len(self.values) >= 2
			

#get begin state
bots = dict() #bots[index] = bot
outputBins = dict() #outputBins[index] = outputBin
tasks = list()
with open("realInput.txt") as f:
	lines = f.readlines()

	for line in lines:
		args = line.strip().lower().split(" ")

		if args[0] == "value": # e.g. value 5 goes to bot 2
			index = int(args[5])
			value = int(args[1])
			
			if index not in bots: bots[index] = bot()
			
			bots[index].addValue(value)

			if bots[index].canDistribute: tasks.append(index)

		elif args[0] == "bot": # e.g. bot 2 gives low to bot 1 and high to bot 0 or to output
			index = int(args[1])
			lowerIndex  = int(args[6])
			higherIndex = int(args[11])

			lower, higher = None, None

			if index not in bots: bots[index] = bot()


			if args[5] == "output":
				if lowerIndex not in outputBins: outputBins[lowerIndex] = outputBin()
				lower = (destination.output, lowerIndex)
			else:
				lower = (destination.bot, lowerIndex)

			if args[10] == "output":
				if higherIndex not in outputBins: outputBins[higherIndex] = outputBin()
				higher = (destination.output, higherIndex)
			else:
				higher = (destination.bot, higherIndex)


			bots[index].setLowerHigher(lower, higher)

while len(tasks) > 0:
	index = tasks.pop()

	bot = bots[index]

	if 61 in bot.values and 17 in bot.values:
		print(index)

	lower, higher = bot.distributeIfPossible(bots,outputBins)

	if lower is not None: tasks.append(lower)
	if higher is not None: tasks.append(higher)

