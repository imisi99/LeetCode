# Time -> 0(M+N), Space -> 0(min(M, N))
def intersection(nums1: list[int], nums2: list[int]) -> list[int]:
    lookup = {}

    if len(nums1) > len(nums2):
        nums1, nums2 = nums2, nums1

    for num in nums1:
        lookup[num] = lookup.get(num, 0) + 1

    result = []

    for num in nums2:
        if num in lookup and lookup[num] > 0:
            result.append(num)
            lookup[num] -= 1

    return result


print(intersection([1, 2, 3, 2, 4, 9], [9, 4, 2, 2, 2]))
