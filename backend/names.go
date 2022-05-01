package main

import ()

var current = -1

var names = []string{
	"Viggo Kahn",
	"Stefan Nilsson",
	"Benoit Baudry",
	"Martin Monperrus",
}

func nextName() string {
	if current == -1 {
		current = 0
		return "Arvid Gotthard"
	}
	name := names[current]
	current += 1
	current %= len(names)
	return name
}
