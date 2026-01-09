# Time -> 0(N*M) where M is the length of the longest word Space -> 0(1)
def longestCommonPrefix(array: list[str]) -> str:
    prefix = array[0]
    for word in array[1:]:
        for i in range(len(prefix)):
            if i == len(word) or prefix[i] != word[i]:
                prefix = prefix[:i]
                break
    return prefix


array = ["profile", "prolific", "pro", "property"]
prefix = longestCommonPrefix(array)
print(prefix)
