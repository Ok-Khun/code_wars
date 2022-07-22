import unittest
from maps import maps

class TestMaps(unittest.TestCase):
    def test_maps(self):
        self.assertEqual(maps([1, 2, 3]), [2, 4, 6], "test fail")
        self.assertEqual(maps([0, 1, 2, 3, 4, 5, 6, 7, 8, 9]), [0, 2, 4, 6, 8, 10, 12, 14, 16, 18], "test fail")
        self.assertEqual(maps([]), [], "test fails")

if __name__ == "__main__":
    unittest.main()

