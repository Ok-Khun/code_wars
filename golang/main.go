package main

import (
	eightKyu "codewars/8kyu"
	"fmt"
)

func main() {
	// s := []string{
	// 	"JwgxlUIyKByrGRYhPvQdYDLmkGdBGYITLgYbk", "fBLjjVGEtsPwZYfFYUqxPKZynSLYpVWTiP", "rODfCMXNrXfHdiCDR", "byPRFnpTtxqAkZMsOvpTmiGUMkWzOHSlEdvWZpI", "zltROqhRDopmGKVQeVxGkxFRbQxhLyyS", "JMVLZldRSreSLoqecHHhhlkAHymq", "LwgkGoESJaKlYZuOiMQSnpvYJmvqRCbT", "RaCsSqNdkuq", "yJunGyYZhakvBQOHeeWfOjEYvRGHAvWGpOBkBxZ", "SDQwyAyCaiHShExD", "WoEKrNnKzsIyzRzh",
	// }
	s := []string{
		"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps",
	}
	result := eightKyu.TwoSort(s)
	fmt.Println(result)
}
