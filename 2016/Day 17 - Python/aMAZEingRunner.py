import hashlib
import queue


WIDTH = 4
HEIGHT = 4
startPos = [0,3]

SALT = "qtetzkpl"
#"ulqzkmiv" #"kglvqrro" #"ihgpwlah"

class State:
	position = None #position = [x,y]
	path = ""
	hval = ""

	def __init__(self,position,path):
		self.position = position
		self.path = path
		self.hval = self.hash()

	def pathLength(self):
		return len(self.path)

	def isGoalState(self):
		return self.position[0] == 3 and self.position[1] == 0

	def nextStates(self):
		nextStates = list()
		
		#up
		if self.position[1] < 3 and self.hval[0] in ["b","c","d","e","f"]: #open
			nextStates.append(State([self.position[0], self.position[1] + 1], self.path + "U"))
		#down
		if self.position[1] > 0 and self.hval[1] in ["b","c","d","e","f"]: #open
			nextStates.append(State([self.position[0], self.position[1] - 1], self.path + "D"))
		#left
		if self.position[0] > 0 and self.hval[2] in ["b","c","d","e","f"]: #open
			nextStates.append(State([self.position[0] - 1, self.position[1]], self.path + "L"))
		#right
		if self.position[0] < 3 and self.hval[3] in ["b","c","d","e","f"]: #open
			nextStates.append(State([self.position[0] + 1, self.position[1]], self.path + "R"))

		return nextStates

	def hash(self):
	    m = hashlib.md5()
	    m.update(SALT.encode('utf-8'))
	    m.update(self.path.encode('utf-8'))
	    return str(m.hexdigest())

	def toStr(self):
		return "State at: ", self.position, " path: ", self.path, " hash: ", self.hash()[:4]

currState = State(startPos, "")

seen = dict()
paths = queue.Queue()
while currState is not None and not currState.isGoalState():

    #print(currState.toStr())

    if not currState.hval in seen:
        seen[currState.hash()] = 1

        for state in currState.nextStates():
            #print("\t", state.toStr())
            paths.put(state)

    if paths.empty(): #concurrent => the get function will just block and wait for input
        currState = None
    else:
        currState = paths.get()

if currState is not None:
	print(currState.toStr())
else:
	print("something went wrong")