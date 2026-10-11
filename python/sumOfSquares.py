# Time => 0(N) Space => 0(1)
def sumOfSquares(nums: list[int]) -> int:
    sumSquares = 0
    n = len(nums)
    for i in range(len(nums)):
        if n % (i + 1) == 0:
            sumSquares += nums[i] ** 2

    return sumSquares


print(sumOfSquares([1, 2, 3, 4]))
print(sumOfSquares([2, 7, 1, 19, 18, 3]))
