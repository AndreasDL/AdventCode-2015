import re

def getChecksum(name):
		
	name = name.replace("-","") #remove -
	
	#count & keep track of indexes
	occurences = dict()
	for char in list(name):
		if char not in occurences:
			occurences[char] = 0 #init occurences[char] = count

		occurences[char] += 1

	#get first most common occurence
	#sort is guaranteed to be stable in python !!
	keys = list(occurences.keys())
	keys.sort()
	keys.sort(key=lambda k: occurences[k], reverse=True)

	#[ print("[", key, "] => ", occurences[key]) for key in keys ]

	checksum = ''.join( keys[:5] )

	return checksum

def isDecoy(name, checksum):
	shouldBe = getChecksum(name)
	#print(shouldBe, " == ", checksum)
	return shouldBe == checksum


sectorIdSum= 0
regex = re.compile("([^0-9]+)-([0-9]+)\[([a-zA-Z]+)\]")
with open("roomInput.txt") as f:
	
	for line in f.readlines():
		room = regex.findall(line.strip())[0]
		
		#print(room)
		if isDecoy(room[0],room[2]):
			sectorIdSum += int(room[1])

print(sectorIdSum)