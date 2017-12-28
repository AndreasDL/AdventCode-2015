#registers
registers = {
	'a': 7,
	'b': 0,
	'c': 0,
	'd': 0
}

class Cpy:
	def __init__(self,value,regName):
		self.value = value
		self.regName = regName

	def execute(self):
		if isDigitNegative(self.regName):
			return

		if isDigitNegative(self.value):
			registers[str(self.regName)] = int(self.value)
		else:
			registers[str(self.regName)] = registers[str(self.value)]
		return 1

	def toggle(self):
		return Jnz(self.value, self.regName)

	def toStr(self):
		return "[Cpy] - " + str(self.value) + " - " + str(self.regName)

class Inc:
	def __init__(self,regName):
		self.regName = regName

	def execute(self):
		if isDigitNegative(self.regName):
			return

		registers[str(self.regName)] += 1
		return 1

	def toggle(self):
		return Dec(self.regName)
	def toStr(self):
		return "[Inc] - " + str(self.regName)

class Dec:
	def __init__(self,regName):
		self.regName = regName

	def execute(self):
		if isDigitNegative(self.regName):
			return
		registers[str(self.regName)] -= 1
		return 1

	def toggle(self):
		return Inc(self.regName)

	def toStr(self):
		return "[Dec] - " + str(self.regName)

class Jnz:
	def __init__(self, value, regName):
		self.value = value
		self.regName = regName
		
	def execute(self):
		valueToComp = int(self.value) if isDigitNegative(self.value) else registers[str(self.value)]
		if valueToComp != 0:
			return int(self.regName) if isDigitNegative(self.regName) else registers[str(self.regName)]
		else:
			return 1

	def toggle(self):
		return Cpy(self.value, self.regName)

	def toStr(self):
		return "[Jnz] - " + str(self.value) + " - " + str(self.regName)

class Tgl:
	def __init__(self,value, index):
		self.value = value
		self.index = index

	def execute(self):
		valToUse = self.value
		if not isDigitNegative(valToUse):
			valToUse = registers[valToUse]

		valToUse += self.index

		if valToUse == self.index:	#If tgl toggles itself (for example, if a is 0, tgl a would target itself and become inc a), 
			return  1			#the resulting instruction is not executed until the next time it is reached.

		if valToUse >= len(instructions) or valToUse < 0: #If an attempt is made to toggle an instruction outside the program, nothing happens.
			return 1
		
		instructions[valToUse] = instructions[valToUse].toggle()
		return 1
	def toggle(self):
		return Inc(self.value)

	def toStr(self):
		return "[Tgl] - " + str(self.value) + " - " + str(self.index)

def loadInstructions(fname):
	#read instructions
	instructions = list()
	with open(fname) as f:
		for line in f.readlines():
			instr = line.strip().split(" ")

			if   instr[0] == 'inc':
				instructions.append(Inc(instr[1]))

			elif instr[0] == 'dec':
				instructions.append(Dec(instr[1]))

			elif instr[0] == 'cpy':
				instructions.append(Cpy(instr[1],instr[2]))

			elif instr[0] == 'jnz':
				instructions.append(Jnz(instr[1],instr[2]))

			elif instr[0] == 'tgl':
				instructions.append(Tgl(instr[1],len(instructions)))

	return instructions
def isDigitNegative(value):
	try:
		int(value)
		return True
	except:
		return False
def printInstructions():
	for i in instructions:
		print("\t",i.toStr())




instructions = loadInstructions("realInput.txt") #global list

if __name__ == '__main__':

	

	i = 0
	while i < len(instructions):
		instr = instructions[i]
		#print("executing instruction ", i, " : ", instr.toStr())
		
		i += instr.execute()

		#print(registers['a'], " | ", registers['b'], " | ", registers['c'], " | ", registers['d'])
		#printInstructions()

		#print()
		#print()



#print()
#print()
print(registers['a'], " | ", registers['b'], " | ", registers['c'], " | ", registers['d'])