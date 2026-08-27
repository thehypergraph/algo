import unittest

from series import fib


class TestFib(unittest.TestCase):
    def test_fib(self):
        self.assertEqual(fib(10), 55)


if __name__ == "__main__":
    unittest.main()
