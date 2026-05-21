# Time -> 0(N) Space -> (N)
def wordPattern(pattern: str, s: str) -> bool:
    words = s.split(" ")
    p_map, s_map = {}, {}

    for pat in pattern:
        p_map[pat] = p_map.get(pat, 0) + 1
    for word in words:
        s_map[word] = s_map.get(word, 0) + 1

    w_map, pat_map = {}, {}

    for _, val in p_map.items():
        pat_map[val] = pat_map.get(val, 0) + 1
    for _, val in s_map.items():
        w_map[val] = w_map.get(val, 0) + 1

    if len(w_map) != len(pat_map):
        return False

    for key, val in w_map:
        if pat_map.get(key, 0) == 0 or pat_map[key] != val:
            return False

    return True


# Time -> 0(N) Space -> (N)
def wordPatternII(pattern: str, s: str) -> bool:
    words = s.split(" ")

    if len(pattern) != len(words):
        return False

    letter_map = {}

    for pat in range(len(pattern)):
        curr_pat = pattern[pat]
        if (
            letter_map.get(curr_pat, None) is not None
            and letter_map[curr_pat] != words[pat]
        ):
            return False
        letter_map[curr_pat] = words[pat]

    return True


print(wordPattern("stfj", "ball call car car"))
print(wordPatternII("aba", "dog cat cat"))
