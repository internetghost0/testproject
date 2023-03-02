package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	// reader for input a whole line
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = strings.Trim(line, "\n")
		// split line by 3 parts
		parts := strings.Split(line, " ")

		if len(parts) != 3 {
			fmt.Println("Error: calculator expected ONLY 3 arguments, example: `1 + 1`")
			os.Exit(1)
		}
		// 1st operand (string)
		a := parts[0]
		// operator 
		op := parts[1]
		// 2nd operand (string)
		b := parts[2]

		// flag for printing in roman or arabic form
		isRomanResult := false

		// uninitialized 1st operand (int) 
		x := 0
		// uninitialized 2nd operand (int) 
		y := 0
		// uninitializes result
		result := 0

		// check if a & b is arabic number
		if isArabNumber(a) && isArabNumber(b) {
			x, err = strconv.Atoi(a)
			if err != nil {
				fmt.Printf("Error: Cannot parse `%s` as integer\n", a)
				os.Exit(1)
			}
			y, err = strconv.Atoi(b)
			if err != nil {
				fmt.Printf("Error: Cannot parse `%s` as integer\n", b)
				os.Exit(1)
			}
		// check if a & b is roman number
		} else if isRomanNumber(a) && isRomanNumber(b) {
			x = RomanToNumber(a)
			y = RomanToNumber(b)
			isRomanResult = true
		// neither roman or arabic, printing error...
		} else {
			fmt.Printf("Error: Cannot parse your input: `%s`\n", line)
			fmt.Printf("Use ONLY Arabic or Roman form for 2 numbers, example: `1 + 1` or `VI - II`\n")
			os.Exit(1)
		}
		// set limit (dont know why)
		if (x < 1 || x > 10) || (y < 1 || y > 10) {
			fmt.Printf("Error: Calculator accepts only 1,2,3..10 range of numbers\n")
			os.Exit(1)
		}

		switch op {
		case "+":
			result = x + y
		case "-":
			result = x - y
		case "/":
			result = x / y
		case "*":
			result = x * y
		default:
			log.Fatalf("Can't use this `%s` as operator\n", op)
		}
		// if flag is set, then print result in roman form
		if isRomanResult {
			if result > 0 {
				fmt.Println(NumberToRoman(result))
			} else {
				// todo: print to stderr not stdout
				fmt.Println("Error: result < 0")
			}
		// print result in arabic form
		} else {
			fmt.Println(result)
		}
	}
}

func isArabNumber(s string) bool {
	for _, ch := range s {
		if !strings.ContainsRune("0123456789-", ch) {
			return false
		}
	}
	return true
}
func isRomanNumber(s string) bool {
	for _, ch := range s {
		if !strings.ContainsRune("IVXLCDM", ch) {
			return false
		}
	}
	return true
}

func RomanToNumber(s string) int {
	var num = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	out := 0
	ln := len(s)
	for i := 0; i < ln; i++ {
		c := string(s[i])
		vc := num[c]
		if i < ln-1 {
			cnext := string(s[i+1])
			vcnext := num[cnext]
			if vc < vcnext {
				out += vcnext - vc
				i++
			} else {
				out += vc
			}
		} else {
			out += vc
		}
	}
	return out
}
func NumberToRoman(n int) string {
	var numInv = map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += numInv[v]
		n -= v
	}
	return out
}
func highestDecimal(n int) int {
	var maxTable = []int{
		1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}

	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}
