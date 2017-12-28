def isTriangle(numbers):
	numbers.sort()
	return (numbers[0] + numbers[1]) > numbers[2]


drawingsOnTheWall = []
with open("triangleInput.txt") as f:
	
	for line in f.readlines():		
		numbers = [int(x) for x in line.strip().split()]
		drawingsOnTheWall.append(numbers)

print(drawingsOnTheWall)


counter = 0
for y in range(0,len(drawingsOnTheWall),3):
	for x in range(len(drawingsOnTheWall[0])):
		numbers = [drawingsOnTheWall[y][x], drawingsOnTheWall[y+1][x], drawingsOnTheWall[y+2][x]]
		if isTriangle(numbers):
			counter += 1

print(counter)
