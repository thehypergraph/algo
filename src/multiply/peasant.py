# Peasant Multiplication
import math


def peasant_multiply(x: int, y: int) -> int:
    prod = 0
    while x > 0:
        if x % 2 == 1:
            prod = prod + y
        x = math.floor(x / 2)
        y = y + y
    return prod
