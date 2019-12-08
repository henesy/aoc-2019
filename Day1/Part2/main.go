package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
)

// Calculate total sum fuel required for take-off
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	sum := 0

	for {
		// Read line
		s, err := in.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			fatal("err: could not read input line - ", err)
		}

		// Strip trailing newline
		if len(s) < 2 {
			fatal("err: empty line")
		}

		s = s[:len(s)-1]

		// Extract number
		n, err := strconv.Atoi(s)
		if err != nil {
			fatal("err: invalid number for mass - ", err)
		}

		fuelReq := calcFuel(n)

		// Calculate the total
		for f := fuelReq; f > 0; f = calcFuel(f) {
			sum += f
		}

	}

	out.WriteString(strconv.Itoa(sum) + "\n")
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
