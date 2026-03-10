# Time -> 0(N) Space -> 0(1)
def removeDuplicate(array: list[int]) -> int:
    k = 1
    idxToSwap, lastVal = 1, array[0]

    for i in range(1, len(array)):
        if array[i] != lastVal:
            array[idxToSwap] = array[i]
            lastVal = array[i]
            idxToSwap += 1
            k += 1

    return k


array = [1, 2, 3, 4, 5]
print("Array -> ", array)
k = removeDuplicate(array)
print("Array -> ", array, "num of unique elem -> ", k)
