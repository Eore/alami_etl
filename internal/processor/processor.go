package processor

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type DataInterface interface {
}

type Processor struct {
	sourceFile string
}

func New(sourceFile string) *Processor {
	return &Processor{
		sourceFile: sourceFile,
	}
}

func (p *Processor) Run(out string, opts ...Option) {
	f, err := os.Create(out)
	if err != nil {
		panic(err)
	}

	// write header to csv file
	f.WriteString("id;Nama;Age;Balanced;No 2b Thread-No;No 3 Thread-No;Previous Balanced;Average Balanced;No 1 Thread-No;Free Transfer\n")

	od := optionData{}

	for _, opt := range opts {
		opt(&od)
	}

	i := 1
	for data := range p.streamReadData() {
		if i <= od.countUser {
			data.AddBalance(od.bonusBudget / float64(od.countUser))
		}
		data.CalculateAverage()
		data.CalculateBenefit()
		f.WriteString(data.ToRowCSVString())
		i++
	}
}

type optionData struct {
	bonusBudget float64
	countUser   int
}

type Option func(optData *optionData)

func WithBonus(budget float64, countUser int) Option {
	return func(optionData *optionData) {
		optionData.bonusBudget = budget
		optionData.countUser = countUser
	}
}

func (p *Processor) streamReadData() chan Data {
	ch := make(chan Data)

	go func(ch chan Data) {
		defer close(ch)

		file, err := os.Open(p.sourceFile)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		reader := csv.NewReader(file)
		reader.Comma = ';'
		skipHeader := true

		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if skipHeader {
				skipHeader = false
				continue
			}
			ch <- p.mapData(row)
		}
	}(ch)

	return ch
}

func (p *Processor) mapData(row []string) Data {
	id, _ := strconv.Atoi(row[0])
	name := row[1]
	age, _ := strconv.Atoi(row[2])
	balanced, _ := strconv.ParseFloat(row[3], 64)
	previousBalanced, _ := strconv.ParseFloat(row[4], 64)
	freeTransfer, _ := strconv.Atoi(row[6])

	data := Data{
		ID:               id,
		Name:             name,
		Age:              age,
		Balanced:         balanced,
		PreviousBalanced: previousBalanced,
		FreeTransfer:     freeTransfer,
	}

	return data
}
