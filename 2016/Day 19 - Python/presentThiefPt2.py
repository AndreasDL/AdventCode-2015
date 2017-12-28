import sys

NO_ELVES = 3005290

#1->3 0->2
#2->5 1->4
#4->1 3->0
#2->4 1->3
#2 1 wins


elfAtPosition = [ i for i in range(NO_ELVES)]

#init
currIndex = 0
while len(elfAtPosition) > 1:
	
	nextIndex = int(currIndex + len(elfAtPosition)/2) % len(elfAtPosition)
	
	currElf = elfAtPosition[currIndex]
	nextElf = elfAtPosition[nextIndex]
	#print(currElf, " steals from ", nextElf)

	del elfAtPosition[nextIndex]
	if len(elfAtPosition) % 1000 == 0:
		print("pickpockets left: ", len(elfAtPosition))
	
	if currElf < nextElf:
		currIndex += 1

	currIndex %= len(elfAtPosition)

print(elfAtPosition) #this position is zero indexed! +1