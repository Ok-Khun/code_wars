import math
def century(year):
    # check if the year is divisible by 100
    if year % 100 == 0:
        return math.trunc(year / 100)
    return math.trunc(year / 100) + 1

print(century(2000))