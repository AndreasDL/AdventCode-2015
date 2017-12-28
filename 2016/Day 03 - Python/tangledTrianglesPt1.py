import numpy

def isTriangle(numbers):
	return (numbers[0] + numbers[1]) > numbers[2]

with open("triangleInput.txt") as f:
	counter = 0
	for line in f.readlines():
		
		numbers = [int(x) for x in line.strip().split()]
		numbers.sort()

		if isTriangle(numbers):
			counter += 1

	print(counter)