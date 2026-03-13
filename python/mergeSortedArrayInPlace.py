# Time -> 0(M+N) Space -> 0(1)
def merge(nums1: list[int], m: int, nums2: list[int], n: int):
    end = m + n - 1
    m -= 1
    n -= 1

    while end >= 0:
        if n < 0:
            break
        if m < 0:
            while end >= 0:
                nums1[end] = nums2[n]
                end -= 1
                n -= 1
            break

        if nums1[m] >= nums2[n]:
            nums1[end] = nums1[m]
            end -= 1
            m -= 1
        else:
            nums1[end] = nums2[n]
            end -= 1
            n -= 1


array1 = [2, 3, 5, 0, 0, 0]
array2 = [1, 4, 6]
merge(array1, len(array1) - len(array2), array2, len(array2))
print(array1)
