package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/xuri/excelize/v2"
)

type ExcelData struct {
	TotalRows int
	Error     error
	File      *excelize.File
	Rows      [][]string
}

type RowData struct {
	Data  interface{}
	Error error
}

const (
	size = 100
)

func readFile(filePaths []string) <-chan ExcelData {
	excel := make(chan ExcelData)

	go func() {
		defer close(excel)
		for _, path := range filePaths {
			file, err := excelize.OpenFile(path)
			if err != nil {
				fmt.Println("Error Open File: ", err.Error())
				excel <- ExcelData{Error: err}
				return
			}

			rows, err := file.GetRows("Trang tính1")
			if err != nil {
				fmt.Println("Error Get Rows: ", err.Error())
				excel <- ExcelData{Error: err}
				return
			}

			excel <- ExcelData{
				TotalRows: len(rows),
				Error:     nil,
				File:      file,
				Rows:      rows,
			}
		}
	}()

	return excel
}

func handleExcelData(data <-chan ExcelData) <-chan []RowData {
	var wg sync.WaitGroup

	out := make(chan []RowData)

	for excel := range data {
		if excel.Error != nil {
			continue
		}
		pages := int(math.Ceil(float64(excel.TotalRows) / float64(size)))

		for i := 0; i < int(pages); i++ {
			startIndex := i * size
			endIndex := (i + 1) * size
			if endIndex > excel.TotalRows {
				endIndex = excel.TotalRows
			}

			rowData := excel.Rows[startIndex:endIndex]
			wg.Add(1)
			go func(rowData [][]string) {
				defer wg.Done()
				rows := make([]RowData, len(rowData))

				for _, row := range rowData {
					fmt.Println("Row: ", row)
					rows = append(rows, RowData{Data: row})
				}

				out <- rows
			}(rowData)
		}

		excel.File.Close()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func collectRowData(data <-chan []RowData) {
	for rows := range data {
		for _, row := range rows {
			if row.Error == nil {
				continue
			}
		}
	}
}

func main() {
	excel := readFile([]string{"input/Test.xlsx"})
	rows := handleExcelData(excel)
	collectRowData(rows)
}
