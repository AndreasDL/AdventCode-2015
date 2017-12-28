
import re

def getChecksum(name):
		
	name = name.replace("-","") #remove -
	
	#count & keep track of indexes
	occurences = dict()
	for i, char in enumerate(list(name)):		
		if char not in occurences:
			occurences[char] = 0 #occurences[char] = count

		occurences[char] += 1

	#get first most common occurence
	#sort is guaranteed to be stable in python !!
	keys = list(occurences.keys())
	keys.sort()
	keys.sort(key=lambda k: occurences[k], reverse=True)

	checksum = ''.join( keys[:5] )
	return checksum

def isDecoy(name, checksum):
	shouldBe = getChecksum(name)
	return shouldBe == checksum

def getChar(char, rot):
	return chr( 97 + ( ord(char.lower()) -97 +rot )%26 )

def decrypt(name,sectorId):
	#get indexes of '-'
	plain = ""
	for char in list(name.lower()):
		plain += (" " if char == "-" else getChar(char,sectorId))

	return plain


#print(decrypt("qzmt-zixmtkozy-ivhz",343))

sectorIdSum= 0
regex = re.compile("([^0-9]+)-([0-9]+)\[([a-zA-Z]+)\]")
with open("roomInput.txt") as f:
	
	for line in f.readlines():
		room = regex.findall(line.strip())[0]
		
		if isDecoy(room[0],room[2]):
			sectorIdSum += int(room[1])
			print(room[1], " -> ", decrypt(room[0], int(room[1])) )
