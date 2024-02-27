package hello // package name

import "strings"

func SayHello(names []string) string {
	if len(names) == 0 {
		names = append(names, "World")
	}

	return "Hello " + strings.Join(names, ", ") + "!"
}
