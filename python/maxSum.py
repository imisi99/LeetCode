# kj Time => 0(N) Space => 0(1)
def maxSum(nums: list[int], k: int) -> int:
    coins = [0] * 31
    for num in nums:
        for i in range(len(coins)):
            coins[i] += (num >> i) & 1

    total = 0
    while k > 0:
        k -= 1
        curr = 0
        for i in range(len(coins)):
            if coins[i] > 0:
                curr |= 1 << i
                coins[i] -= 1

        total = (total + (curr**2)) % (10**9 + 7)
    return total


print(maxSum([2, 6, 5, 8], 2))
print(maxSum([4, 5, 4, 7], 3))
