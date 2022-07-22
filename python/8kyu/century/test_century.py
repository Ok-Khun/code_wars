import unittest
from century import century


class TestCentury(unittest.TestCase):
    def test_century(self):
        self.assertEqual(century(1705), 18, "want 18 but got a different result")
        self.assertEqual(century(1900), 19, "want 19 but got a different result")
        self.assertEqual(century(1601), 17, "want 17 but got a different result")
        self.assertEqual(century(2000), 20, "want 20 but got a different result")

if __name__ == "__main__":
    unittest.main()