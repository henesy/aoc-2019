package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
)

// Calcuate f=int(m/3)-2 with one f per line
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	for {
		s, err := in.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			fatal("err: could not read input line - ", err)
		}

		if len(s) < 2 {
			fatal("err: empty line")
		}

		s = s[:len(s)-1]

		n, err := strconv.Atoi(s)
		if err != nil {
			fatal("err: invalid number for mass - ", err)
		}

		fuelReq := calcFuel(n)

		out.WriteString(strconv.Itoa(fuelReq) + "\n")
	}

	out.Flush()
}

// Calculate f=int(m/3)-2
func calcFuel(mass int) int {
	return int(mass / 3) - 2
}

// Fatal - end program with an error message and newline
func fatal(s ...interface{}) {
	fmt.Fprintln(os.Stderr, s...)
	os.Exit(1)
}
