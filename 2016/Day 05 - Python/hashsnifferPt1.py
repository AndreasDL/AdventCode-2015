import hashlib
import sys
import numpy as np
from multiprocessing import Pool

#constants
POOL_SIZE = 10
BLOCK_SIZE = 1000000
START_VAL = 0 #3231929

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
    results = list()

    while(len(results) < 8):
        print("i: ", i, " -> " ,end="")

        pool = Pool(processes=POOL_SIZE)
        job_results = pool.map( worker, range(i * POOL_SIZE, (i+1) * POOL_SIZE) )
        pool.close()
        pool.join()

        #flatten array
        for arr in job_results:
            if len(arr) > 0:
                results.extend(arr)            
        i+=1

    solution = ""
    for entry in results:
        if len(entry) > 5:
            solution += entry[5]

    print(solution[:8])



if __name__ == "__main__":
    main()