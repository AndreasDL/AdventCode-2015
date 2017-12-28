import sys
import itertools
#Some of these functions might perform in place operations and therefor corrupt change the input, but hey I don't care

def swapPosition(line, x, y):
	char = line[x]
	line[x] = line[y]
	line[y] = char

	return line

def swapLetter(line, x, y):
	for i in range(len(line)):
		if line[i] == x:
			line[i] = y
		elif line[i] == y:
			line[i] = x
	return line

def rotatePosition(line, posToLeft):
	newline = list()
	length = len(line)
	
	for i in range(length):
		newline.append( line[ (i + posToLeft) % length] )

	return newline

def rotateLetter(line, char):
	by = line.index(char)
	if by >= 4:
		by += 1

	by += 1
	return rotatePosition(line,-by)

def reverseRotateLetter(line,char):
	# origineel = 0, 1, 2, 3, 4, 5, 6, 7
	# shift = positie + 1, +2 als positie groter is dan 4 = 1, 2, 3, 4, 6, 7, 8, 9
	# positie na shift = 1, 3, 5, 7, 10, 12, 14, 16
	# normalized positie = 1, 3, 5, 7, 2, 4, 6, 0
	# THANKS GERT

	index = line.index(char)
	print(line, " ", index)
	reverseMap = [1,3,5,7,2,4,6,0]	

	#werkt niet altijd om ongekende reden </3
	return rotatePosition(line, - (reverseMap.index(index) - index))

def reversePositions(line, x, y):

	while x < y:
		swapPosition(line,x,y)

		x += 1
		y -= 1

	return line

def movePosition(line,x,y):
	char = str(line[x])
	line.remove(line[x])
	line.insert(y,char)

	return line

def test():
	line = list("abcde")
	print(line)

	line = swapPosition(line,4,0) #swap position 4 with position 0 swaps the first and last letters, producing the input for the next step, ebcda.
	print(line)
	line = swapLetter(line,"d","b") #swap letter d with letter b swaps the positions of d and b: edcba.
	print(line)
	line = reversePositions(line,0,4) #reverse positions 0 through 4 causes the entire string to be reversed, producing abcde.
	print(line)
	print()
	line = rotatePosition(line,1) #rotate left 1 step shifts all letters left one position, causing the first letter to wrap to the end of the string: bcdea.
	print(line)
	line = movePosition(line,1,4) #move position 1 to position 4 removes the letter at position 1 (c), then inserts it at position 4 (the end of the string): bdeac.
	print(line)
	line = movePosition(line,3,0) #move position 3 to position 0 removes the letter at position 3 (a), then inserts it at position 0 (the front of the string): abdec.
	print(line)
	line = rotateLetter(line,"b") #rotate based on position of letter b finds the index of letter b (1), then rotates the string right once plus a number of times equal to that index (2): ecabd.
	print(line)
	line = rotateLetter(line,"d") #rotate based on position of letter d finds the index of letter d (4), then rotates the string right once, plus a number of times equal to that index, plus an additional time because the index was at least 4, for a total of 6 right rotations: decab.
	print(line)

	#After these steps, the resulting scrambled password is decab.
def forward(line, value):
	instr = line.strip().split(" ")

	if instr[0] == "rotate":
		if instr[1] == "based": #rotate right 4 steps
			value = rotateLetter(value, str(instr[6]))
		else:
			pos = int(instr[2])
			if instr[1] == "right":
				pos = -pos
			value = rotatePosition(value, pos)

	elif instr[0] == "swap":
		if instr[1] == "letter":
			value = swapLetter(value, str(instr[2]), str(instr[5]))
		else:
			value = swapPosition(value, int(instr[2]), int(instr[5]))

	elif instr[0] == "reverse":
		value = reversePositions(value, int(instr[2]), int(instr[4]))

	elif instr[0] == "move":
		value = movePosition(value, int(instr[2]), int(instr[5]))

	else:
		print(line)
		sys.exit()
	return value

if __name__ == "__main__":
	value = list("abcdefgh")

	instructions = list()
	with open("input.txt") as f:
		instructions =  f.readlines()

	for line in instructions:
		value = forward(line,value)

	print("forward :", ''.join(value))


	value = list("fbgdceah")
	for key in itertools.permutations(list("abcdefgh")):
		scrambledVal = key
		for line in instructions:
			scrambledVal = forward(line,scrambledVal)

		if scrambledVal == value:
			print(''.join(key))
			sys.exit()
