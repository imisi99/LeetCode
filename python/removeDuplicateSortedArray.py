# Time -> 0(N) Space -> 0(1)
def removeDuplicate(array: list[int]) -> int:
    k = 1
    idxToSwap, idx, prev = 1, 1, 0

    while idx < len(array):
        while array[idx] == array[prev]:
            if idx == len(array) - 1:
                break
            prev = idx
            idx += 1

        array[idxToSwap] = array[idx]
        idxToSwap += 1
        prev = idx
        idx += 1
        k += 1

    return k


array = [1, 2, 3, 3, 4, 4, 4, 7, 11]
print("Array -> ", array)
k = removeDuplicate(array)
print("Array -> ", array, "num of unique elem -> ", k)
