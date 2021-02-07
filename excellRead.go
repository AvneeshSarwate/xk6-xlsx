package main

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/loadimpact/k6/js/modules"
)

func main() {
	// f, err := excelize.OpenFile("DEMO_APIServices.xlsx")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // Get value from cell by given worksheet name and axis.
	// cell, err := f.GetCellValue("services", "B2")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(cell)
	// // Get all the rows in the Sheet1.
	// rows, err := f.GetRows("services")
	// for _, row := range rows {
	// 	for _, colCell := range row {
	// 		fmt.Print(colCell, "\t")
	// 	}
	// 	fmt.Println()
	// }

	sheetVals, err := getSheetMaps("DEMO_APIServices.xlsx", "services")
	// jsonString, err := json.Marshal(sheetVals)
	// fmt.Println(err)
	if err == nil {
		fmt.Println(sheetVals)
	} else {
		fmt.Println(err)
	}

}

func init() {
	modules.Register("k6/x/redis", new(SheetReader))
}

type SheetReader struct{}

func (r *SheetReader) Get(fileName, sheetName string) (map[string]map[string]string, error) {
	return getSheetMaps(fileName, sheetName)
}

var fileSheetMap = make(map[string]map[string]map[string]string)

func getSheetMaps(fileName, sheetName string) (map[string]map[string]string, error) {
	fileSheetKey := fileName + "-" + sheetName
	if val, ok := fileSheetMap[fileSheetKey]; ok {
		return val, nil
	} else {
		f, err := excelize.OpenFile("DEMO_APIServices.xlsx")
		if err != nil {
			return nil, err
		}

		rows, err := f.GetRows("services")
		if err != nil {
			return nil, err
		}

		sheetValues := make(map[string]map[string]string)

		columnNames := rows[0][1:] //column names are in first row, columns 1-onwards
		dataRows := rows[1:]       //rows with data about services are in rows 1-onwwards
		for _, row := range dataRows {

			if strings.TrimSpace(row[0]) != "" {
				// fmt.Println(row)
				rowMap := make(map[string]string)
				for i, colVal := range row[1:] { //iterate over all of the data values in a row (skip service name)
					rowMap[columnNames[i]] = colVal
				}
				sheetValues[row[0]] = rowMap
			}
		}

		fileSheetMap[fileSheetKey] = sheetValues
		// fmt.Println()
		return sheetValues, nil
	}
}
