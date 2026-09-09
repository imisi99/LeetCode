# Time -> 0(1) Space -> 0(1)
def countCommas(n: int) -> int:
    return (
        max(n - 999, 0)
        + max(n - 999999, 0)
        + max(n - 999999999, 0)
        + max(n - 999999999999, 0)
        + max(n - 999999999999999, 0)
    )


print(countCommas(1992))
print(countCommas(1004590))
