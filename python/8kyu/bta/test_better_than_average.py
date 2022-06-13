import unittest
import bta

class TestBta(unittest.TestCase):
    def test_bta(self):
        self.assertEqual(bta.better_than_average([2, 3], 5), True, "expected the result to be true")
        self.assertEqual(bta.better_than_average([29, 55, 74, 60, 11, 90, 67, 28], 21), False, "expected the result to be false")

if __name__ == "__main__":
    unittest.main()