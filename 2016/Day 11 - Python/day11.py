import queue
import sys
from copy import deepcopy
from itertools import combinations


FLOOR_COUNT = 4

class State:
    currFloor = 0
    distance = 0
    pairs = list() #[[0,1],[0,2]] corresponds to start of test state; for each element keep floors in pair: (microchip, generator)

    def __init__(self,currFloor,distance,pairs):
        self.currFloor = currFloor
        self.distance = distance

        #sort in python are guaranteed to be stable :D
        self.pairs = sorted(pairs, key= lambda x: x[0])
        self.pairs = sorted(self.pairs, key= lambda x: x[1])

    def isGoalState(self):
        if self.currFloor != FLOOR_COUNT - 1:#goal state will be reached when elevator is on upper floor
            return False;

        for element in self.pairs:
            for item in element:
                if item != FLOOR_COUNT - 1:
                    return False

        return True

    def isValid(self):
        # a chip can only be with another chips if it is connected to its own generator:
        #=> invalid : find a chip that is with another generator
        i = 0
        while i < len(self.pairs):
            j = 0
            while j < len(self.pairs):
                if j != i:
                    if self.pairs[i][0] != self.pairs[i][1] \
                        and self.pairs[i][0] == self.pairs[j][1]:
                        return False

                j+=1
            i+=1

        return True

    def getAdjacent(self):
        #get stuffz at current floor
        currGenerators = list()
        currChips = list()
        for el, pair in enumerate(self.pairs):
            if pair[0] == self.currFloor: currChips.append(el)
            if pair[1] == self.currFloor: currGenerators.append(el)

        adjStates = list()
        for c in range(len(currChips) + 1):
            for g in range(len(currGenerators) + 1):
            
                if g+c > 0 and g+c <= 2: #elevator must carry at least 1 item and cannot carry more than 2 items
                    
                    #add new states
                    for chosenChips in combinations(currChips,c):
                        for chosenGens in combinations(currGenerators, g):
                            
                            #print("c:", chosenChips , " - g:", chosenGens)

                            newPairs = deepcopy(self.pairs)
                            if self.currFloor < FLOOR_COUNT - 1: #go up
                                for cc in chosenChips:
                                    newPairs[cc][0] += 1

                                for gg in chosenGens:
                                    newPairs[gg][1] += 1

                                #go up 
                                newState = State(
                                    currFloor = self.currFloor + 1,
                                    distance  = self.distance  + 1,
                                    pairs     = newPairs
                                )

                                #newState.dump()

                                if newState.isValid():
                                    adjStates.append(newState)

                            newPairs = deepcopy(self.pairs)
                            if self.currFloor > 0: #go down
                                for cc in chosenChips:
                                    newPairs[cc][0] -= 1

                                for gg in chosenGens:
                                    newPairs[gg][1] -= 1

                                #go up 
                                newState = State(
                                    currFloor = self.currFloor - 1,
                                    distance  = self.distance  + 1,
                                    pairs     = newPairs
                                )
                            
                                #newState.dump()

                                if newState.isValid():
                                    adjStates.append(newState)

        return adjStates

    def hash(self):
        var = str(self.currFloor) + "|"
        for pair in self.pairs:
            var+= str(pair[0]) + str(pair[1]) + "|"
        return var

    def dump(self):
        print(self.currFloor,  " -> ", self.pairs, " {", self.distance, "} ==>> ", self.hash(), " ==>> ", self.isValid())


#hardcoded start state
testStartState = State(0,0, [[0,1],[0,2]])
examplesStates = [
    State(0,0, [[0,1],[0,2]]),
    State(1,1, [[1,1],[0,2]]),
    State(2,2, [[2,2],[0,2]]),
    State(1,3, [[1,2],[0,2]]),
    State(0,4, [[0,2],[0,2]]),
    State(1,5, [[1,2],[1,2]]),
    State(2,6, [[2,2],[2,2]]),
    State(3,7, [[3,2],[3,2]]),
    State(2,8, [[2,2],[3,2]]),
    State(3,9, [[2,3],[3,3]]),
    State(2,10,[[2,3],[2,3]]),
    State(3,11,[[3,3],[3,3]])
]

#realinput:
#The first floor contains
#a polonium generator, 
#a thulium generator, a thulium-compatible microchip, 
#a promethium generator, 
#a ruthenium generator, a ruthenium-compatible microchip, 
#a cobalt generator, cobalt-compatible microchip.


#part two:
#An elerium generator.
#An elerium-compatible microchip.
#A dilithium generator.
#A dilithium-compatible microchip.

#The second floor contains 
#a polonium-compatible microchip 
#a promethium-compatible microchip.

#The third floor contains nothing relevant.
#The fourth floor contains nothing relevant.

#realInput = State(0,0, [[0,1],[0,0],[0,1],[0,0],[0,0]]) #Part1
realInput = State(0,0, [[0,1],[0,0],[0,1],[0,0],[0,0],[0,0],[0,0]]) #Part 2




#for s in examplesStates:
#    print(s.hash())

#for i in range(len(examplesStates)-1):
#    state = examplesStates[i]
#    nextState = examplesStates[i+1].hash()
#
#    print("state ", i , " -> ", state.hash())
#
#    print("nextSate ", nextState)
#
#    print("adjacentStates")
#    adjHashes = [s.hash() for s in state.getAdjacent()]
#    
#    print(nextState in adjHashes)
#    
#    for h in adjHashes:
#        print(h)
#
#    print()
#    print()
#
#    input()
#sys.exit()


seen = dict()
paths = queue.Queue()
currentState = realInput #testStartState
while currentState is not None and not currentState.isGoalState():
#    print(currentState.hash(), end="")

    if not currentState.hash() in seen:
#        print(" -> Not Seen -> ")

        seen[currentState.hash()] = 1

        for state in currentState.getAdjacent():
            paths.put(state)
#           print("\t", state.hash())

#    else:
#        print(" -> Seen -> Skip")

#    print("Queue size: ", paths.qsize(), " - Seen size: ", len(seen.keys()))
#    print()
#    print()
#    input()

    if paths.empty(): #concurrent => the get function will just block and wait for input
        currentState = None
    else:
        currentState = paths.get()

print(currentState.hash(), " found at ", currentState.distance)