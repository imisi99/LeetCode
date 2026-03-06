def contWithMostWater(array: list[int], idx: int) -> int:
    if idx == len(array) - 2:
        area = min(array[len(array) - 1], array[len(array) - 2])
        return area

    area = findMaxIdx(array, idx)
    return max(area, contWithMostWater(array, idx + 1))


def findMaxIdx(array: list, i: int) -> int:
    max = 0
    j = i
    while i < len(array):
        width = i - j
        height = min(array[j], array[i])
        area = width * height
        if area > max:
            max = area
        i += 1

    return max


# Time -> 0(N^2) Sapce -> 0(N)
def solution(array) -> int:
    if len(array) <= 1:
        return array[0] if len(array) == 1 else 0
    return contWithMostWater(array, 0)


def solution2(array) -> int:
    area = 0
    first, end = 0, len(array) - 1

    while first < end:
        currArea = min(array[first], array[end]) * (end - first)
        area = max(area, currArea)

        if array[first] < array[end]:
            first += 1
        else:
            end -= 1

    return area


print(solution([1, 8, 6, 2, 5, 4, 8, 3, 7]))
print(solution2([1, 8, 6, 2, 5, 4, 8, 3, 7]))
print(solution([1, 2]))
