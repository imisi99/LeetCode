# Time => 0((k1+k2)*N) Space => 0(1)


def minSumSquareDiff(nums1: list[int], nums2: list[int], k1: int, k2: int) -> int:
    while k1 != 0 or k2 != 0:
        maxIdx = 0
        maxDiff = 0

        for idx in range(len(nums1)):
            diff = nums1[idx] - nums2[idx]
            diff *= diff

            if diff > maxDiff:
                maxDiff = diff
                maxIdx = idx

        if maxDiff == 0:
            return 0

        if k1 != 0:
            if nums1[maxIdx] > nums2[maxIdx]:
                nums1[maxIdx] -= 1
            else:
                nums1[maxIdx] += 1
            k1 -= 1
        elif k2 != 0:
            if nums2[maxIdx] > nums1[maxIdx]:
                nums2[maxIdx] -= 1
            else:
                nums2[maxIdx] += 1
            k2 -= 1

    sumDiff = 0
    for idx in range(len(nums1)):
        diff = nums1[idx] - nums2[idx]
        diff *= diff
        sumDiff += diff

    return sumDiff


# Time => 0(NlgoN) Space => 0(N)
def minSumSquareDiffI(nums1: list[int], nums2: list[int], k1: int, k2: int) -> int:
    d = [0] * (len(nums1) + 1)
    total = 0
    for idx in range(len(nums1)):
        diff = nums1[idx] - nums2[idx]
        if diff < 0:
            diff = -1 * diff
        d[idx] = diff
        total += diff * diff

    k = k1 + k2

    if total <= k:
        return 0

    d.sort(reverse=True)

    for i in range(len(nums1)):
        group = i + 1
        cost = group * (d[i] - d[i + 1])

        if k >= cost:
            k -= cost
            continue
        else:
            q, r = k // group, k % group
            v = d[i] - q

            diff = ((group - r) * v * v) + ((v - 1) * (v - 1) * r)
            j = i + 1
            while j < len(d):
                diff += d[j] * d[j]
                j += 1

            return diff
    return 0


print(minSumSquareDiff([1, 2, 3, 4], [2, 10, 20, 19], 0, 0))
print(minSumSquareDiff([1, 4, 10, 12], [5, 8, 6, 9], 1, 1))
print(minSumSquareDiffI([1, 2, 3, 4], [2, 10, 20, 19], 0, 0))
print(minSumSquareDiffI([1, 4, 10, 12], [5, 8, 6, 9], 1, 1))
