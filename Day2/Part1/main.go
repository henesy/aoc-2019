package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"flag"
)


const (
	opWidth = 4		// Opcodes are 4 ints wide
)

var (
	chatty = false
)


// Process Intcode input - emit final table
func main() {
	flag.BoolVar(&chatty, "D", false, "Enable debug output")
	flag.Parse()

	dat, err := ioutil.ReadFile("/dev/stdin")
	if err != nil {
		fatal("err: could not read input - ", err)
	}

	// The file has a trailing newline, remove it
	datStr := strings.Replace(string(dat), "\n", "", -1)

	codeStr := strings.Split(datStr, ",")

	codes := make([]int, len(codeStr))

	for i, _ := range codeStr {
		// The `codes` slice will be our op table
		n, err := strconv.Atoi(codeStr[i])
		if err != nil {
			fatal("err: invalid number - ", err)
		}

		codes[i] = n
	}

	// Process opcodes
	for i := 0; i < len(codes)-1; i += opWidth {
		op := codes[i]

		// Position in `codes` of the two inputs
		args := codes[i+1:i+3]

		// Position in `codes` to store the output
		out := codes[i+3]

		v := 0

		debug("Processing →", op, args, out)

		switch op {
		case 1:
			// Addition
			v = codes[args[0]] + codes[args[1]]
		case 2:
			// Multiplication
			v = codes[args[0]] * codes[args[1]]
		case 99:
			// Halt
			debug("\t→ Halting…")
			goto END
		}

		// Output
		codes[out] = v

		// DEBUG
		debug("\t→", v)

		END:;
	}

	fmt.Println(opStringify(codes))
}

// Solution-format the opcode slice
func opStringify(c []int) string {
	s := strconv.Itoa(c[0])

	for i := 1; i < len(c); i++ {
		s += ","
		s += strconv.Itoa(c[i])
	}

	return s
}

// Fatal - end program with an error message and newline
func fatal(s ...interface{}) {
	fmt.Fprintln(os.Stderr, s...)
	os.Exit(1)
}

// Debug output
func debug(s ...interface{}) {
	if chatty {
		fmt.Fprintln(os.Stderr, s...)
	}
}
