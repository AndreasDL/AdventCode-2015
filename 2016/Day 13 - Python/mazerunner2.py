import queue
import sys

startPos = [0, [1,1]] #[dist, [x,y]]
favNumber = 1364 #10
maxDist = 50

dxs = [0, 1,  0, -1]
dys = [1, 0, -1,  0]

def formula(x,y, salt=favNumber):
	#return x*x + 3*x + 2*x*y + y + y*y + 10
	return x*(x + 3 + y+y) + (y+1)*y + salt #speedup
	
def wallAtPosition(x, y, formula=formula):
	val = formula(x,y)
	
	counter = 0
	for bit in list(bin(val))[2:]:
		counter += int(bit)

	return counter % 2 == 1

def printField(width,height):
	for y in range(height):
		for x in range(width):
			if moveValid(x,y):#
				print(".", end="")
			else:
				print("#", end="")
		print()

def hashPos(pos):
	return str(pos[1][0]) + "|" +str(pos[1][1])

#field cache
fieldCache = dict() #fieldCache[y][x] = true if there is a wall at that position
def moveValid(x,y): #valid if in positive quadrand and not a wall
	if x < 0 or y < 0:
		return False

	if y not in fieldCache:
		fieldCache[y] = dict()

	if x not in fieldCache[y]:
		posFree = not wallAtPosition(x,y)
		fieldCache[y][x] = posFree
		
		return posFree #avoid relookup of hash value, when in fact we just have it

	return fieldCache[y][x]

#printField(50,50) #will fill up cache
#print()
#print()

seen = dict()
paths = queue.Queue()
currPos = startPos
while currPos is not None:

	if not hashPos(currPos) in seen:
		seen[hashPos(currPos)] = 1

		for (dx, dy) in zip(dxs,dys):
			newX = currPos[1][0] + dx
			newY = currPos[1][1] + dy

			if moveValid(newX, newY) and currPos[0] < maxDist:
				newPos = [
					currPos[0]+1,
					[newX,newY]
				]
				paths.put(newPos)

	if paths.empty(): #concurrent => the get function will just block and wait for input
		currPos = None
	else:
		currPos = paths.get()

		
print(len(seen.keys()))