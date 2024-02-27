package main

import "strings"

func Search(q []string, c Comics) Comic {

	for i := range c.Comics {
		occ := 0
		for j := range q {
			if strings.Contains(c.Comics[i].Transcript, q[j]) {
				occ++
			}
		}
		if occ == len(q) {
			return c.Comics[i]
		}
	}
	return Comic{}
}
