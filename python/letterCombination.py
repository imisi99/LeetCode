def letterCombinations(digits: str) -> list[str]:
    result = []

    numbers = {
        "2": ["a", "b", "c"],
        "3": ["d", "e", "f"],
        "4": ["g", "h", "i"],
        "5": ["j", "k", "l"],
        "6": ["m", "n", "o"],
        "7": ["p", "q", "r", "s"],
        "8": ["t", "u", "v"],
        "9": ["w", "x", "y", "z"],
    }

    traverse([], 0, digits, numbers, result)

    return result


def traverse(builder: list, idx: int, digits: str, nums: dict, result: list):
    if idx == len(digits):
        return

    for digit in nums[digits[idx]]:
        newBuilder = []
        newBuilder.extend(builder)
        newBuilder.append(digit)
        traverse(newBuilder, idx + 1, digits, nums, result)
        if idx == len(digits) - 1:
            result.append("".join(newBuilder))


print(letterCombinations("23"))
