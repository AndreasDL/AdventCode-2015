fileName = "realInput.txt"


freqCols = None

with open(fileName) as f:
	lines = f.readlines()

	#init columns
	freqCols = [ {} for i in range(len(lines[0]))]

	for line in lines:
		for (i,c) in enumerate(line):
			if c not in freqCols[i]:
				freqCols[i][c] = 0

			freqCols[i][c] += 1

code = ""
for freq in freqCols:
	
	minChar = list(freq.keys())[0]
	minVal = freq[minChar]

	for key in freq:
		if freq[key] < minVal:
			minChar = key
			minVal = freq[key]

	code += minChar
print(code)

