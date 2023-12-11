package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Wire struct {
	operator, src1, src2 string
}

var (
	wires          = make(map[string]Wire)
	memorizedWires = make(map[string]uint16)
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	input := make([]string, 0)

	for r.Scan() {
		line := r.Text()
		input = append(input, line)

		var (
			operator, src1, src2, dst string
		)

		if n, _ := fmt.Sscanf(line, "%s -> %s\n", &src1, &dst); n == 2 {
			operator = "MOV"
		} else if n, _ := fmt.Sscanf(line, "%s AND %s -> %s\n", &src1, &src2, &dst); n == 3 {
			operator = "AND"
		} else if n, _ := fmt.Sscanf(line, "%s OR %s -> %s\n", &src1, &src2, &dst); n == 3 {
			operator = "OR"
		} else if n, _ := fmt.Sscanf(line, "%s RSHIFT %s -> %s\n", &src1, &src2, &dst); n == 3 {
			operator = "RSHIFT"
		} else if n, _ := fmt.Sscanf(line, "%s LSHIFT %s -> %s\n", &src1, &src2, &dst); n == 3 {
			operator = "LSHIFT"
		} else if n, _ := fmt.Sscanf(line, "NOT %s -> %s\n", &src1, &dst); n == 2 {
			operator = "NOT"
		} else {
			panic(line)
		}

		wires[dst] = Wire{operator, src1, src2}

	}

	memorizedWires["b"] = uint16(46065)
	fmt.Println(value("a"))

}

func value(dst string) uint16 {
	if v, err := strconv.ParseUint(dst, 10, 16); err == nil {
		return uint16(v)
	}
	if v, ok := memorizedWires[dst]; ok {
		return v
	}

	var w Wire = wires[dst]
	var v uint16

	switch w.operator {
	case "MOV":
		v = value(w.src1)
	case "AND":
		v = value(w.src1) & value(w.src2)
	case "OR":
		v = value(w.src1) | value(w.src2)
	case "LSHIFT":
		v = value(w.src1) << value(w.src2)
	case "RSHIFT":
		v = value(w.src1) >> value(w.src2)
	case "NOT":
		v = ^value(w.src1)
	}

	memorizedWires[dst] = v
	return v
}
