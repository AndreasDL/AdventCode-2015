from itertools import combinations

class Node:
	x, y = 0, 0
	size = 0
	used = 0
	available = 0
	usePercentage = 0

	def __init__(self,line):
		line = line.strip().split()

		name = line[0].split("-")
		self.x = int(name[1][1:])
		self.y = int(name[2][1:])

		self.size = int(line[1][:-1])
		self.used = int(line[2][:-1])
		self.available = int(line[3][:-1])
		self.usedPercentage = int(line[4][:-1])

	def formsViablePair(self,node):
		if node.x == self.x and node.y == self.y: #not same
			return False

		if self.used == 0: #not empty
			return False

		return node.available >= self.used #would fit

def getNodes(fname):
	lines = None
	with open(fname) as f:
		lines = f.readlines()
	lines = lines[2:]

	nodes = list()
	for line in lines:
		nodes.append(Node(line))

	return nodes

def countViablePairs(nodes):
	counter = 0
	for (nodeA, nodeB) in combinations(nodes,2):
		if nodeA.formsViablePair(nodeB) or nodeB.formsViablePair(nodeA):
			counter += 1

	return counter

def partOne():
	nodes = getNodes("realInput.txt")
	print(countViablePairs(nodes))

def partTwo():
	startPosX, startPosY = 0, 0
	goalPosX, goalPosY = 36, 0
	
	


if __name__ == "__main__":
	partTwo()
	