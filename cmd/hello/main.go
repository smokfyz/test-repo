package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/smokfyz/test-repo/pkg/hello"
)

func parseCSVInts(s string) ([]int, error) {
	if strings.TrimSpace(s) == "" {
		return nil, nil
	}
	parts := strings.Split(s, ",")
	res := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			return nil, errors.New("empty value in sum list")
		}
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("invalid integer %q: %w", p, err)
		}
		res = append(res, n)
	}
	return res, nil
}

func main() {
	var name string
	var sumCSV string

	flag.StringVar(&name, "name", "", "name to greet (defaults to 'World')")
	flag.StringVar(&sumCSV, "sum", "", "comma-separated list of integers to sum (e.g. '1,2,3,4')")
	flag.Parse()

	// Greeting
	fmt.Println(hello.Greet(name))

	// Optional sum
	if strings.TrimSpace(sumCSV) != "" {
		nums, err := parseCSVInts(sumCSV)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		fmt.Println("Sum:", hello.Sum(nums...))
	}
}
