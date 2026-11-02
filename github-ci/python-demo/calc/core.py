"""Basic arithmetic helpers."""


def add(a: int, b: int) -> int:
    """Return the sum of a and b."""
    return a + b


def sub(a: int, b: int) -> int:
    """Return a minus b."""
    return a - b


def mul(a: int, b: int) -> int:
    """Return the product of a and b."""
    return a * b


def div(a: int, b: int) -> float:
    """Return a divided by b.

    Raises:
        ZeroDivisionError: if b is zero.
    """
    if b == 0:
        raise ZeroDivisionError("divide by zero")
    return a / b
