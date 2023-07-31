package main

import (
	"os"
	"strconv"

	"tuanm.dev/golab/mod01"
	"tuanm.dev/golab/mod04"
	"tuanm.dev/golab/mod05"
)

func main() {
	args := os.Args[1:]
	var mod int
	var lab int

	read(&mod, 0, args...)
	read(&lab, 1, args...)

	switch mod {
	case 1:
		mod01.Run(lab)
	case 4:
		mod04.Run(lab)
	case 5:
		mod05.Run(lab)
	default:
		println("no module found")
	}
}

func read(n *int, i int, args ...string) {
	if len(args) > i {
		v, err := strconv.Atoi(args[i])
		if err != nil {
			panic(err)
		}
		*n = v
	} else {
		panic("index of of bound")
	}
}
