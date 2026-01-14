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

        if idxToSwap != idx and array[idxToSwap] == array[idx]:
            break

        array[idxToSwap] = array[idx]
        idxToSwap += 1
        prev = idx
        idx += 1
        k += 1

    return k


array = [1, 2, 3, 4, 5]
print("Array -> ", array)
k = removeDuplicate(array)
print("Array -> ", array, "num of unique elem -> ", k)
