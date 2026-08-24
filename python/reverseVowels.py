# Time -> 0(N) Space -> 0(N)
def reverseVowels(s: str) -> str:
    vowels = []
    for char in s:
        if char in "aeiouAEIOU":
            vowels.append(char)

    reversed = []

    for char in s:
        if char in "aeiouAEIOU":
            reversed.append(vowels.pop())
        else:
            reversed.append(char)
    return "".join(reversed)


print(reverseVowels("IceCreAm"))
