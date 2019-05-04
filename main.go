package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/fmusayev/golang-excel/api"
	"github.com/fmusayev/golang-excel/types"
)

func main() {
	xlsx, err := excelize.OpenFile("./resources/EXAMPLE.XLSX")
	if err != nil {
		log.Fatal(err)
		return
	}

	apiService := api.API{}

	rows := xlsx.GetRows("Input")
	for _, row := range rows {
		inputData := types.InputData{
			CompanyId: row[0],
			PinCode:   row[1],
		}
		fmt.Println(inputData)

		result, err := apiService.SendRequest(inputData.PinCode)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(result)
	}

}
