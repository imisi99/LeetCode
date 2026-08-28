# Time -> 0(M+N), Space -> 0(min(M, N))
def intersection(nums1: list[int], nums2: list[int]) -> list[int]:
    lookup = set()

    if len(nums1) > len(nums2):
        nums1, nums2 = nums2, nums1

    for num in nums1:
        lookup.add(num)

    result = []

    for num in nums2:
        if num in lookup:
            result.append(num)
            lookup.remove(num)

    return result


print(intersection([1, 2, 3, 2, 4, 9], [9, 4, 2]))
