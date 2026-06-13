# Time -> 0(N) Space -> 0(N)
def canWinNim(n: int) -> bool:
    memo = {}

    def playOptimal(n: int) -> bool:
        if n <= 0:
            return False

        if n in memo:
            return memo[n]

        result = (
            not playOptimal(n - 1) or not playOptimal(n - 2) or not playOptimal(n - 3)
        )

        memo[n] = result

        return result

    return playOptimal(n)


# Time -> 0(1) Space -> 0(1)
def canWinNimI(n: int) -> bool:
    return n % 4 != 0


print(canWinNim(9))
print(canWinNimI(9))
