TRAPCHAR = "^"
SAFECHAR = "."


def strToLine(str):
	line = list()
	line.append(False)

	for char in str:
		line.append( char == TRAPCHAR)

	line.append(False)

	return line

def lineToStr(line):
	str = ""
	for i in range(1, len(line) -1):
		if line[i]:
			str += TRAPCHAR
		else:
			str += SAFECHAR
	return str

def getNextTile(left, center, right):
	return (left and not right) or (not left and right)

def getNextLine(currLine):

	nextLine = list()
	nextLine.append(False)

	for i in range(1, len(currLine)-1):
		nextLine.append( getNextTile(currLine[i-1], currLine[i], currLine[i+1]) )

	nextLine.append(False)
	return nextLine

def count(line):
	counter = 0
	for i in range(1, len(line) - 1):
		if not line[i]:
			counter += 1

	return counter


if __name__ == "__main__":
	
	counter = 0
	ROWCOUNT = 400000#40
	currLine = strToLine(".^^.^^^..^.^..^.^^.^^^^.^^.^^...^..^...^^^..^^...^..^^^^^^..^.^^^..^.^^^^.^^^.^...^^^.^^.^^^.^.^^.^.")
	#strToLine(".^^.^.^^^^") #"..^^."
	
	#print(lineToStr(currLine))
	counter += count(currLine)
	for i in range(ROWCOUNT - 1):
		currLine = getNextLine(currLine)
		#print(lineToStr(currLine))
		counter += count(currLine)

	print(counter)

