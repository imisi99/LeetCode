def intToRoman(num: int) -> str:
    result = []
    while num > 0:
        if (num // 1000) >= 1:
            result.append("M")
            num -= 1000
        elif (num // 100) >= 1:
            if num >= 400 and num <= 499:
                result.append("CD")
                num -= 400
            elif num >= 900 and num <= 999:
                result.append("CM")
                num -= 900
            elif num >= 500:
                result.append("D")
                num -= 500
            else:
                result.append("C")
                num -= 100

        elif (num // 10) >= 1:
            if num >= 40 and num <= 49:
                result.append("XL")
                num -= 40
            elif num >= 90 and num <= 99:
                result.append("XC")
                num -= 90
            elif num >= 50:
                result.append("L")
                num -= 50
            else:
                result.append("X")
                num -= 10

        else:
            if num == 4:
                result.append("IV")
                num -= 4
            elif num == 9:
                result.append("IX")
                num -= 9
            elif num >= 5:
                result.append("V")
                num -= 5
            else:
                result.append("I")
                num -= 1

    return "".join(result)


print(intToRoman(3979))
