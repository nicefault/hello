package hello

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("[v2] Hi, %v, Welcome!", name)
	return message
}
