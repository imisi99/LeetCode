# Time -> 0(N) Space -> 0(1)
def linearSearchRange(nums: list[int], target: int) -> list[int]:
    result = [-1, -1]

    for i in range(len(nums)):
        if nums[i] == target:
            if result[0] == -1:
                result[0] = i
            result[1] = i

    return result


# Time -> 0(log N) Space -> 0(1)
def binarySearchRange(nums: list[int], target: int) -> list[int]:
    result = [-1, -1]
    start, end = 0, len(nums) - 1

    while start <= end:
        mid = (start + end) // 2

        if nums[mid] > target:
            end = mid - 1
        elif nums[mid] < target:
            start = mid + 1
        else:
            result[0], result[1] = mid, mid
            tmpend, tmpstart = mid - 1, mid + 1

            while start <= tmpend:
                mid = (start + tmpend) // 2
                if nums[mid] < target:
                    start = mid + 1
                else:
                    result[0] = mid
                    tmpend = mid - 1

            while tmpstart <= end:
                mid = (tmpstart + end) // 2
                if nums[mid] > target:
                    end = mid - 1
                else:
                    result[1] = mid
                    tmpstart = mid + 1

            break

    return result


array = [5, 7, 7, 8, 8, 10]
print(linearSearchRange(array, 8))
print(binarySearchRange(array, 8))
