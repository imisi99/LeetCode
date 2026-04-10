# Time -> 0(N) Space -> 0(1)
def lengthOfLongestSubstring(s: str) -> int:
    if len(s) <= 1:
        return len(s)

    hastTable: dict[str, list[int]] = {}

    for char in s:
        hastTable[char] = [0, 0]

    bestLen = 0
    startIdx = 0
    for idx in range(len(s)):
        val = hastTable.get(s[idx], [0, 0])
        if val[0] > 0 and val[1] >= startIdx:
            newLen = idx - startIdx
            if newLen > bestLen:
                bestLen = newLen
            startIdx = val[1] + 1
        hastTable[s[idx]] = [1, idx]

    if len(s) - startIdx > bestLen:
        return len(s) - startIdx
    return bestLen


print(lengthOfLongestSubstring("bbbbb"))
