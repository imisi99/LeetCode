# Time -> 0(N) Space -> 0(N)
def isIsomorphic(s: str, t: str) -> bool:
    mapped = {}

    if len(s) != len(t):
        return False

    for i in range(len(s)):
        if s[i] in mapped:
            if t[i] != mapped[s[i]]:
                return False
        mapped[s[i]] = t[i]

    return True


print(isIsomorphic("add", "egg"))
print(isIsomorphic("add", "egd"))
