package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dowts/go-packages/UsageGenerator"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		scenario, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid scenario id: %v\n", err)
			os.Exit(1)
		}
		UsageGenerator.Auto(scenario)
		return
	}

	if len(args) == 6 {
		in1, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid input1: %v\n", err)
			os.Exit(1)
		}
		in2, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid input2: %v\n", err)
			os.Exit(1)
		}
		in3, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid input3: %v\n", err)
			os.Exit(1)
		}
		in4, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid input4: %v\n", err)
			os.Exit(1)
		}
		in5, err := strconv.ParseFloat(args[4], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid input5: %v\n", err)
			os.Exit(1)
		}
		in6, err := strconv.ParseFloat(args[5], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid input6: %v\n", err)
			os.Exit(1)
		}

		UsageGenerator.Start(in1, in2, in3, in4, in5, in6)
		return
	}

	printUsage()
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "DoWTS (Dec 2025 pricing)\n")
	fmt.Fprintf(os.Stderr, "\nAuto mode (recommended):\n")
	fmt.Fprintf(os.Stderr, "  go run dowts.go <scenario_id>\n")
	fmt.Fprintf(os.Stderr, "\nManual mode (advanced):\n")
	fmt.Fprintf(os.Stderr, "  go run dowts.go <input1> <input2> <input3> <input4> <input5> <input6>\n")
	fmt.Fprintf(os.Stderr, "\nSee README for parameter details.\n")
	os.Exit(1)
}
