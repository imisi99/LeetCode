# Time -> 0(N) Space -> 0(N)
def resultArray(nums: list[int]) -> list[int]:
    arr1 = []
    arr2 = []

    arr1.append(nums[0])
    arr2.append(nums[1])
    idx = 2

    while idx < len(nums):
        if arr1[len(arr1) - 1] > arr2[len(arr2) - 1]:
            arr1.append(nums[idx])
        else:
            arr2.append(nums[idx])
        idx += 1

    arr1.extend(arr2)

    return arr1


print(resultArray([2, 1, 3]))
