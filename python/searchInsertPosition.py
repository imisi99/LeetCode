# Time -> 0(log(N)) Space -> 0(1)
def searchInsert(nums: list[int], target: int) -> int:
    start, end = 0, len(nums) - 1

    while start <= end:
        mid = (start + end) // 2

        if nums[mid] == target:
            return mid

        if target > nums[mid]:
            start = mid + 1
        else:
            end = mid - 1

        if start > end:
            return start

    return 0
