package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	switch os.Getenv("GO_TEST_MODE") {
	default:
		os.Exit(m.Run())

	case "1":
		main()
	}
}

var uniqueLocationTests = []struct {
	in   string
	out  int64
	name string
}{
	{"", 0, "No movements"},
	{"N", 1, "Single direction"},
	{"E", 1, "Single direction"},
	{"S", 1, "Single direction"},
	{"W", 1, "Single direction"},
	{"NN", 2, "Twice in single direction"},
	{"NS", 2, "Out and back"},
	{"NESW", 4, "Basic square"},
	{"NESSWWNNES", 9, "Orbit origin and return"},
}

func TestUniqueLocations(t *testing.T) {
	for _, test := range uniqueLocationTests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := run(test.in)
			if err != nil {
				t.Fatal(err)
			}
			if actual != test.out {
				t.Errorf("Expected %d, but was %d", test.out, actual)
			}
		})
	}
}

func run(input string) (int64, error) {
	args := []string{input}
	args = append(args, os.Args[1:]...)

	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = []string{"GO_TEST_MODE=1"}
	out, err := cmd.Output()
	if err != nil {
		return 0, err
	}
	number, err := strconv.ParseInt(strings.Trim(string(out), "\n"), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Output was not a string: %s", string(out))
	}
	return number, err
}
