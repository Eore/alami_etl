package main

import (
	"alami/internal/processor"
	"flag"
)

func main() {

	in := flag.String("in", "", "input file in csv format (required)")
	out := flag.String("out", "out.csv", "output file in csv format")

	flag.Parse()

	if *in == "" || *out == "" {
		flag.CommandLine.Usage()
	}

	processor.New(*in).Run(*out, processor.WithBonus(1000, 100))

}
