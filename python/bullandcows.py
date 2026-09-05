# Time -> 0(N) Space -> 0(1)
def getHint(secret: str, guess: str) -> str:
    lookup = {}

    for val in secret:
        lookup[val] = lookup.get(val, 0) + 1

    bull, cow = 0, 0

    for idx in range(len(guess)):
        val = guess[idx]
        if val == secret[idx]:
            lookup[val] -= 1
            bull += 1

    for idx in range(len(guess)):
        if guess[idx] == secret[idx]:
            continue

        val = guess[idx]
        if lookup.get(val, 0) >= 1:
            lookup[val] -= 1
            cow += 1

    return f"{bull}A{cow}B"


print(getHint("1807", "7810"))
print(getHint("1123", "0111"))
print(getHint("567", "567"))
print(getHint("11", "10"))
