# Time -> 0(N) Space -> 0(N)
def majorityElement(nums: list[int]) -> int:
    counter = {}
    for num in nums:
        counter[num] = counter.get(num, 0) + 1
        if counter[num] > (len(nums) // 2):
            return num
    return 0


print(majorityElement([1, 2, 2, 2]))
