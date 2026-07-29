# Time -> 0(N) Space -> 0(N)
def firstUniqChar(s: str) -> int:
    occur = {}
    for i in range(len(s)):
        occur[s[i]] = occur.get(s[i], 0) + 1
    for i in range(len(s)):
        if occur[s[i]] == 1:
            return i
    return -1


print(firstUniqChar("leetcode"))
print(firstUniqChar("loveleetcode"))
