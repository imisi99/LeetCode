# Time -> 0(N^3)
def longestPalindrome(s: str) -> str:
    longest = [0, 0]

    for idx in range(len(s)):
        lasts = []
        j = len(s) - 1
        while j > idx:
            if s[j] == s[idx]:
                lasts.append(j)
            j -= 1

        for last in lasts:
            start, stop = idx, last

            while start < stop and s[start] == s[stop]:
                start += 1
                stop -= 1

            if start == stop or s[start] == s[stop]:
                if (last - idx) > (longest[1] - longest[0]):
                    longest[0], longest[1] = idx, last

    return s[longest[0] : longest[1] + 1]


# Time -> 0(N^2)
def expandCenterPalindrome(s: str) -> str:
    def expand(i: int, j: int) -> tuple[int, int]:
        while i >= 0 and j < len(s) and s[i] == s[j]:
            i -= 1
            j += 1
        return i + 1, j - 1

    start, stop = 0, 0
    for i in range(len(s)):
        l1, r1 = expand(i, i)
        l2, r2 = expand(i, i + 1)

        if r1 - l1 > stop - start:
            stop, start = r1, l1
        if r2 - l2 > stop - start:
            stop, start = r2, l2

    return s[start : stop + 1]


print(longestPalindrome("Gbababb"))
print(expandCenterPalindrome("Gbababb"))
