package processor

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type FileReader struct {
	f *os.File
}

func NewFileReader(file string) *FileReader {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	return &FileReader{
		f: f,
	}
}

func (f *FileReader) Read(callback func(data Data)) {
	reader := csv.NewReader(f.f)
	reader.Comma = ';'

	skipHeader := true

	for {
		row, err := reader.Read()
		if err == io.EOF {
			return
		}
		if skipHeader {
			skipHeader = false
			continue
		}
		callback(mapData(row))
	}
}

type FileWriter struct {
	f *os.File
}

func NewFileWriter(outFile string) *FileWriter {
	f, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}

	return &FileWriter{
		f: f,
	}
}

func (f *FileWriter) Write(record string) {
	f.f.WriteString(record)
}

func (f *FileWriter) Close() {
	f.f.Close()
}

func mapData(row []string) Data {
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
