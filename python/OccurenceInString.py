# Time -> 0(N) Space -> 0(1)
def strStr(haystack: str, needle: str) -> int:
    idx = 0
    while idx < len(haystack):
        if haystack[idx] == needle[0]:
            if idx + len(needle) <= len(haystack):
                if haystack[idx : idx + len(needle)] == needle:
                    return idx
            else:
                break

        idx += 1
    return -1
