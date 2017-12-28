#registers
registers = {
	'a': 0,
	'b': 0,
	'c': 1, #0 for part 1 and 1 for part 2
	'd': 0
}

#functions
#arguments in order that they occur in text
#return value => jump instructions
def cpy(value, regName):
	if value.isdigit():
		registers[str(regName)] = int(value)
	else:
		registers[str(regName)] = registers[str(value)]
	return 1

def inc(regName, value = 1):
	registers[str(regName)] += value
	return 1

def dec(regName, value = 1):
	registers[str(regName)] -= value
	return 1

def jnz(regName, value):
	valueToComp = int(regName) if regName.isdigit() else registers[str(regName)]
	if valueToComp != 0:
		return int(value)
	else:
		return 1

#function mappers, map the assembunny code to the functions
functions = {
	'cpy': cpy,
	'inc': inc,
	'dec': dec,
	'jnz': jnz
}

#read instructions
instructions = list()
with open("realInput.txt") as f:
	for line in f.readlines():
		instructions.append(line.strip().split(" "))

#execute
i = 0
while i < len(instructions):
	instr = instructions[i]
	print(instr)

	print(registers['a'], " | ", registers['b'], " | ", registers['c'], " | ", registers['d'],
		"  ==[", i , "]==>>  ",
		registers['a'], " | ", registers['b'], " | ", registers['c'], " | ", registers['d'])

	i += functions[instr[0]](
		instr[1], 
		instr[2] if len(instr) >= 3 else 1
	)

print()
print()
print(registers['a'], " | ", registers['b'], " | ", registers['c'], " | ", registers['d'])