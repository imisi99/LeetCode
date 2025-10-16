from math import gcd


def rotateArray(array: list[int], k: int):
    approachII(array, k)


# Time -> 0(KN) space -> 0(1)
def approachI(array: list[int], k):
    for _ in range(k):
        last = array[len(array) - 1]
        start = len(array) - 2
        while start >= 0:
            array[start], array[start + 1] = array[start + 1], array[start]
            start -= 1
        array[0] = last


# Time -> 0(N) space -> 0(1)
def approachII(array: list[int], k: int):
    k = k % len(array)
    n = len(array)
    cycles = gcd(n, k)
    per_cycle = n // cycles

    for start in range(cycles):
        next, val = start, array[start]
        for _ in range(per_cycle):
            pos = (next + k) % n
            tmp = array[pos]
            array[pos] = val
            next = pos
            val = tmp


array = [1, 2, 3, 4, 5, 6, 7]
arrayI = [1, 2, 3, 4, 5, 6, 7]

approachI(array, 2)
approachII(arrayI, 2)
print(array)
print(arrayI)
