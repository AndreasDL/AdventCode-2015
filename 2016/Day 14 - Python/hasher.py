import hashlib
import sys

salt = "qzyelonm" #"abc"
KEYS_LEFT = 64


def hash(counter, salt):
    m = hashlib.md5()
    m.update(salt.encode('utf-8'))
    m.update(str(counter).encode('utf-8'))
    return str(m.hexdigest())

def findTripplet(str):
	i = 0
	length = len(str) - 2
	while i < length and (str[i] != str[i+1] or str[i] != str[i+2]):
		i+=1

	if i == length:
		return None
	else:
		return str[i]

def findQuintet(str):
	i = 0
	length = len(str) - 4
	while i < length and (str[i] != str[i+1] or str[i] != str[i+2] or str[i] != str[i+3] or str[i] != str[i+4]):
		i+=1

	if i == length:
		return None
	else:
		return str[i]

def markCandidates(candidates, index, char):
	for cand in candidates:
		if cand.char == char and cand.fl_index + 1000 > index:
			cand.sl_index = index

class Candidate:
	fl_index = -1
	sl_index = -1
	char = ""

	def __init__(self,fl_index, char):
		self.char = char
		self.fl_index = int(fl_index)

	def isInRange(self, index):
		return self.fl_index + 1000 > index

	def isKey(self):
		return self.sl_index != -1

	def toStr(self):
		return "Candidate: " + self.char + " fl: " + str(self.fl_index) + " sl: " + str(self.sl_index)


i = 0
keys = list()
fl_candidates = list()
while len(keys) < KEYS_LEFT:

	hval = hash(i, salt)

	tripplet = findTripplet(hval)
	if tripplet != None:
			
		quintet = findQuintet(hval)
		if quintet != None:
			markCandidates(fl_candidates, i, quintet)

		fl_candidates.append(Candidate(i, tripplet))

		#prune candidates that are out of range
		while len(fl_candidates) > 0 and not fl_candidates[0].isInRange(i) and not fl_candidates[0].isKey():
			fl_candidates.remove(fl_candidates[0])


		#move all leading fl_candidates to keys
		while len(fl_candidates) > 0 and fl_candidates[0].isKey():
			keys.append(fl_candidates[0])
			fl_candidates.remove(fl_candidates[0])

	i += 1

for k in keys:
	print(k.toStr())