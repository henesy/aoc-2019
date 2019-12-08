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
	opWidth	= 4			// Opcodes are 4 ints wide
	magic	= 19690720	// Magic index 0 value from aoc
)

var (
	chatty = false
)


// Process Intcode input - emit final table
func main() {
	flag.BoolVar(&chatty, "D", false, "Enable debug output")
	flag.Parse()

	one, two := 0, 0

	dat, err := ioutil.ReadFile("/dev/stdin")
	if err != nil {
		fatal("err: could not read input - ", err)
	}

	// The file has a trailing newline, remove it
	datStr := strings.Replace(string(dat), "\n", "", -1)

	codeStr := strings.Split(datStr, ",")

	codes := make([]int, len(codeStr))

	// Populate `codes`
	popCodes := func() {
		for i, _ := range codeStr {
			// The `codes` slice will be our op table
			n, err := strconv.Atoi(codeStr[i])
			if err != nil {
				fatal("err: invalid number - ", err)
			}

			codes[i] = n
		}
	}

	popCodes()

	// Nightmare of substitution jumps
	if one == 0 && two == 0 {
		goto NOSUB
	}

	SUBSTITUTE:
	popCodes()

	if one >= 99 && two >= 99 {
		fatal("fatal: no combination found of inputs, hit 99")
	}

	// TODO - wrap everything in an O(n^2)
	/*	for{for{
			codes = make()
			compute(codes)
			// Check output
			// If magic, break, emit {one, two}
		}}
	*/

	if one <= 99 && two >= 99 {
		one++
		two = -1
	}

	if two <= 99 {
		two++
	}

	debug(one, two)

	codes[1] = one
	codes[2] = two

	NOSUB:
		;

	// Process opcodes
	for i := 0; i < len(codes)-1; i += opWidth {
		op := codes[i]

		// Position in `codes` of the two inputs
		args := codes[i+1:i+3]

		// Position in `codes` to store the output
		out := codes[i+3]

		v := 0

		debug("Processing →", op, args, out)

		for _, v := range codes[i+1:] {
			if v >= len(codes) {
				//fatal("err: argument out of range - ", v)
				goto END
			}
		}

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

		END:
			;
	}

	if codes[0] != magic {
		goto SUBSTITUTE
	}

	fmt.Printf("→ %02d%02d\n", one, two)
	fmt.Println(codes[0])

	debug(opStringify(codes))
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
