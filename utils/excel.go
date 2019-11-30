package util

import (
	"fmt"
	"reflect"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/fmusayev/golang-excel/types"
	log "github.com/sirupsen/logrus"
)

// Opening excel file and returning reference to file
func OpenExcel(fileName string) *excelize.File {
	log.Info("Loding excel file...")
	xlsx, err := excelize.OpenFile(fileName)
	if err != nil {
		log.Error(err)
	}
	return xlsx
}

// Reading all rows from excel
func ReadData(xlsx *excelize.File, sheetName string) [][]string {
	log.Info("Reading rows...")
	return xlsx.GetRows(sheetName)
}

// Writing data to excel using reflection to get column tag from input struct
// If type is extended String then calling method of that type
func WriteData(data types.ExcelForm, indx int, xlsx *excelize.File) {
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		res := fieldValue.String()
		if !fieldValue.IsZero() {
			if fieldType.Type.Name() == "String" {
				res = fieldValue.MethodByName("Transliterate").Call([]reflect.Value{})[0].String()
			}

			cellIndex := fmt.Sprintf("%s%d", fieldType.Tag.Get("col"), indx)
			xlsx.SetCellValue("Sheet1", cellIndex, res)
		}
	}
}
