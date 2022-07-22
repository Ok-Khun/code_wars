import unittest
from basic_op import basic_op

class TestBasicMaths(unittest.TestCase):
    def test_basic_operations(self):
        self.assertEqual(basic_op('+', 4, 7), 11, "case fail")

if __name__ == "__main__":
    unittest.main()