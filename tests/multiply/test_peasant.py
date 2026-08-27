import unittest

from multiply import peasant_multiply


class TestFib(unittest.TestCase):
    def test_peasant(self):
        self.assertEqual(peasant_multiply(10, 20), 200)


if __name__ == "__main__":
    unittest.main()
