# Time -> 0(N) Space -> 0(1)
def findTheDifference(s: str, t: str) -> str:
    lookup = {}

    for val in s:
        lookup[val] = lookup.get(val, 0) + 1

    for val in t:
        if val in lookup and lookup[val] > 0:
            lookup[val] -= 1
        else:
            return val

    return ""
