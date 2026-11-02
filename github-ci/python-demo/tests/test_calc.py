import pytest

from calc import add, div, mul, sub


def test_add():
    assert add(2, 3) == 5


def test_sub():
    assert sub(7, 4) == 3


def test_mul():
    assert mul(6, 7) == 42


def test_div():
    assert div(10, 2) == 5


def test_div_by_zero():
    with pytest.raises(ZeroDivisionError):
        div(1, 0)
