package leetcode

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := make([]string, 0)

	path = strings.ReplaceAll(path, "//", "/")

	path = strings.ReplaceAll(path, "//", "/")

	path = strings.ReplaceAll(path, "//", "/")

	path = strings.ReplaceAll(path, "//", "/")

	var res, p string

	fmt.Println(path)

	for i := 0; i < len(path); i++ {
		if string(path[i]) == "/" {
			if i == 0 {
				continue
			}

			stack = append(stack, p)

			p = ""

			continue
		} else if string(path[i]) == "." {
			if p != "" {
				p += string(path[i])
				continue
			}
			if i < len(path)-2 && string(path[i+1]) == "." && string(path[i+2]) == "." {
				p += "..."
				i += 2

				continue
			}

			if len(stack) != 0 && i < len(path)-2 && string(path[i+1]) == "." && string(path[i+2]) == "/" {
				stack = stack[:len(stack)-1]
				i += 2

				continue
			} else if len(stack) == 0 && i < len(path)-2 && string(path[i+1]) == "." && string(path[i+2]) == "/" {
				i += 2

				continue
			} else if i < len(path)-1 && string(path[i+1]) == "/" {
				i++

				continue
			} else if len(stack) != 0 && i == len(path)-2 && string(path[i+1]) == "." {
				stack = stack[:len(stack)-1]
				i += 2

				continue
			} else if i == len(path)-1 {
				i++

				continue
			} else if i == len(path)-2 && string(path[i+1]) == "." {
				i++
				continue
			}

		}

		p += string(path[i])
	}

	if p != "" {
		stack = append(stack, p)
	}

	if len(stack) == 0 {
		return "/"
	}

	for _, v := range stack {
		res += "/" + v
	}

	return res
}
