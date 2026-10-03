# Time -> 0(N^2)  Space -> 0(N)
def topKFrequent(nums: list[int], k: int) -> list[int]:
    lookup = {}

    for num in nums:
        lookup[num] = lookup.get(num, 0) + 1

    result = []

    for _ in range(k):
        max_key, max_val = 0, 0
        for key, val in lookup.items():
            if val > max_val:
                max_val = val
                max_key = key

        result.append(max_key)
        lookup.pop(max_key)

    return result


print(topKFrequent([1, 1, 1, 2, 2, 3], 2))
