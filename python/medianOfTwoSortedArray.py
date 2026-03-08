def findMedianSortedArrays(nums1, nums2) -> float:
    if len(nums1) > len(nums2):
        nums1, nums2 = nums2, nums1

    m = len(nums1)
    n = len(nums2)
    mid = (m + n + 1) / 2

    x, y = 0, m
    while x <= y:
        i = (x + y) // 2
        j = mid - i

        num1Left = float("-inf")
        num1Right = float("inf")
        num2Left = float("-inf")
        num2Right = float("inf")

        if i > 0:
            num1Left = nums1[i - 1]
        if i < m:
            num1Right = nums1[i]
        if j > 0:
            num2Left = nums2[j - 1]
        if j < n:
            num2Right = nums2[j]

        if num1Left <= num2Right and num2Left <= num1Right:
            maxLeft = max(num1Left, num2Left)
            if (m + n) % 2 == 0:
                minRight = min(num1Right, num2Right)
                return (minRight + maxLeft) / 2
            else:
                return maxLeft

        if num1Left > num2Right:
            y = i - 1
        else:
            x = i + 1

    return 0


print(findMedianSortedArrays([1, 2, 12], [4, 5, 6, 7, 8, 9]))
