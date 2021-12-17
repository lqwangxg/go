package main

import (
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	filename := "./Book2.xlsx"
	if Exists(filename) {
		write2Excel(filename)
	} else {
		createExcel(filename)
	}
}
func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func createExcel(filename string) {
	f := excelize.NewFile()
	// Create a new sheet
	newSheet := f.NewSheet("Sheet2")
	// Set value of a cell
	f.SetCellValue("Sheet2", "A2", "Hello world")
	f.SetCellValue("Sheet1", "A1", "This is a test")
	f.SetActiveSheet(newSheet)
	err := f.SaveAs(filename)
	if err != nil {
		fmt.Println(err)
	}
}
func write2Excel(filename string) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell := f.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)

	// Get all the rows in the Sheet1.
	rows := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	f.SetCellValue("Sheet2", "A3", "Hello world-------TEST")
	err = f.SaveAs(filename)
	if err != nil {
		fmt.Println(err)
	}

}
