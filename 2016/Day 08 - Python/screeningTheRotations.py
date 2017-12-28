OFF = "."
ON = "#"

class field:
	width, height = None, None
	field = None

	def __init__(self, width, height):
		self.width = width
		self.height = height
		
		self.field = [[OFF for i in range(self.width)] for j in range(self.height)]

	def print(self):
		for line in self.field:
			for char in line:
				print(char,end="")
			print()
		print()

	def rect(self,x,y):
		for j in range(y):
			for i in range(x):
				self.field[j][i] = ON

	def rotateColumn(self,x,b):
		oldCol = [self.field[j][x] for j in range(self.height)]
		
		for j in range(self.height):
			self.field[(j+b)%self.height][x] = oldCol[j]

	def rotateRow(self,y,b):
		oldRow = self.field[y][:]
		for i in range(self.width):
			self.field[y][(i+b)%self.width] = oldRow[i]

	def countON(self):
		count = 0
		for j in range(self.height):
			for i in range(self.width):
				if self.field[j][i] == ON:
					count += 1
		return count


def testInput():
	f = field(7,3)
	f.rect(3,2)
	f.print()
	f.rotateColumn(1,1)
	f.print()
	f.rotateRow(0,4)
	f.print()
	f.rotateColumn(1,1)
	f.print()

def realInput():
	f = field(width=50, height=6)
	with open("realInput.txt") as file:
		lines = file.readlines()

		for line in lines:
			instructions = line.split(" ")

			if instructions[0] == "rect":
				args = instructions[1].split("x")
				f.rect(int(args[0]), int(args[1]))

			elif instructions[0] == "rotate":
				index = int(instructions[2].split("=")[1])
				by = int(instructions[4])

				if instructions[1] == "row":
					f.rotateRow(index,by)
				elif instructions[1] == "column":
					f.rotateColumn(index,by)

			else:
				print(instructions[0], " not found !", line)
	f.print()
	print(f.countON())

realInput()