import hashlib
import sys
import numpy as np
from multiprocessing import Pool

#constants
POOL_SIZE = 10
BLOCK_SIZE = 1000000
START_VAL = 0 #3231929
LENGTH = 8

POS_MAPPER = {"a" : 10, "b": 11, "c": 12, "d": 13, "e": 14, "f": 15} 
for i in range(10):
    POS_MAPPER[str(i)] = i

#inputs
testInput  = "abc"
testPassword = "18f47a30"

realInput = "ffykfhsq"

#HARDCODED SHIZZLE
ID = realInput

def hash(counter):
    m = hashlib.md5()
    m.update(ID.encode('utf-8'))
    m.update(str(counter).encode('utf-8'))
    return str(m.hexdigest())

def worker(block):
    #init
    found = list()
    start = block * BLOCK_SIZE + START_VAL
    stop = (block + 1) * BLOCK_SIZE + START_VAL

    hval = hash(start)

    i = start
    while (i < stop):

        if hval[:5] == "00000":
            found.append( hval )

        i += 1
        hval = hash(i)

    return found

def main():

    i = 0
    password = [None for i in range(8)]
    charLeft = LENGTH

    while(charLeft > 0):
        print("i: ", i, " -> " ,end="")

        pool = Pool(processes=POOL_SIZE)
        job_results = pool.map( worker, range(i * POOL_SIZE, (i+1) * POOL_SIZE) )
        pool.close()
        pool.join()

        #flatten array
        results = list()
        for arr in job_results:
            if len(arr) > 0:
                results.extend(arr)  

        for arr in results:
            if len(arr) > 0:
                pos = POS_MAPPER[arr[5]]
                if pos < LENGTH and password[pos] is None: #note to self None == None is always False
                    password[pos] = arr[6]
                    charLeft -= 1

        print(password)

        i+=1

    code = ""
    for i in range(LENGTH):
        code += password[i]

    print(code)



if __name__ == "__main__":
    main()