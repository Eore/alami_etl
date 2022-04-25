package main

import (
	"alami/internal/processor"
	"alami/internal/worker"
	"flag"
)

const (
	header            = "id;Nama;Age;Balanced;No 2b Thread-No;No 3 Thread-No;Previous Balanced;Average Balanced;No 1 Thread-No;Free Transfer\n"
	budget    float64 = 1000
	budgetFor int     = 100
)

func main() {
	in := flag.String("in", "", "input file in csv format (required)")
	out := flag.String("out", "out.csv", "output file in csv format")

	flag.Parse()

	if *in == "" || *out == "" {
		flag.CommandLine.Usage()
	}

	w := worker.New(8)

	count := 1
	writer := processor.NewFileWriter(*out)
	reader := processor.NewFileReader(*in)

	writer.Write(header)
	reader.Read(func(data processor.Data) {
		w.Submit(func(workerID int) error {
			data.CalculateAverage()
			data.CalculateBenefit()
			if count <= budgetFor {
				data.AddBalance(budget / float64(budgetFor))
			}
			writer.Write(data.ToRowCSVString(workerID))

			return nil
		})
	})
}
