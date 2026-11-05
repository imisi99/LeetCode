# Time -> 0(MN) Space -> 0(1)
def findTheDistanceValue(arr1: list[int], arr2: list[int], d: int) -> int:
    count = 0
    for i in arr1:
        addVal = 1
        for j in arr2:
            if abs(i - j) <= d:
                addVal = 0
                break
        count += addVal

    return count


print(findTheDistanceValue([4, 5, 8], [10, 9, 1, 8], 2))
