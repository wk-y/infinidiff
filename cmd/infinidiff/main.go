package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/wk-y/infinidiff"
	"github.com/wk-y/infinidiff/internal/util"
)

//go:embed usage.txt
var usage string

func main() {
	flag.CommandLine.Usage = func() {
		fmt.Print(usage)
	}
	flag.Parse()

	// read files
	lines := [][]string{}
	for _, filename := range flag.Args() {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file %s: %v\n", filename, err)
			os.Exit(1)
		}
		fLines, err := util.ReadLines(f)
		f.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filename, err)
			os.Exit(1)
		}
		lines = append(lines, fLines)
	}

	// get lengths
	lengths := make([]int, len(lines))
	for i := range lines {
		lengths[i] = len(lines[i])
	}

	// compute diff
	result := infinidiff.Diff(lengths, func(a, ai, b, bi int) bool {
		return lines[a][ai] == lines[b][bi]
	})

	// print diff
	indices := make([]int, len(lines))
	for _, selection := range result {
		var line string

		for i := range selection {
			if selection[i] {
				fmt.Print(i % 10) // print last digit of file index
				line = lines[i][indices[i]]
				indices[i]++
			} else {
				fmt.Print(" ") // blank space if line is not in the file
			}
		}

		fmt.Print(" ", line)

		if !strings.HasSuffix(line, "\n") {
			// unified-diff style report
			fmt.Println("\n\\ No newline at end of file")
		}
	}
}
