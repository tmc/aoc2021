package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var flagWindowed = flag.Bool("window", false, "If true, compute windows")

func main() {
	flag.Parse()
	if *flagWindowed {
		windowed()
		return
	}

	s := bufio.NewScanner(os.Stdin)
	n := -1
	t := 0
	for s.Scan() {
		m, _ := strconv.Atoi(strings.TrimSpace(s.Text()))
		if n == -1 {
			n = m
			continue
		}
		if m > n {
			t++
		}
		n = m
	}
	fmt.Println(t)
}

func windowed() {
	s := bufio.NewScanner(os.Stdin)
	n := -1
	t := 0
}
