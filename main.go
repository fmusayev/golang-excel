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
	for indx, row := range rows {
		if indx == 0 {
			continue
		}

		inputData := types.InputData{
			CompanyId: row[0],
			PinCode:   row[1],
		}

		result, err := apiService.SendRequest(inputData.PinCode)
		if err != nil {
			fmt.Println(inputData.PinCode, ". Error: ", err)
			continue
		}

		fmt.Println(inputData.PinCode, ". Writing result to excel: ", result)
		mapOutputToCell(result, indx+1, xlsx)
	}

	xlsx.Save()
}

func mapOutputToCell(data types.OutputData, indx int, xlsx *excelize.File) {

	m := map[string]string{}
	m["C"] = data.PersonalNo
	m["D"] = data.GivenName
	m["E"] = data.Surname
	m["F"] = data.Patronymic
	m["G"] = data.PlaceOfBirth
	m["H"] = data.DateOfBirth
	m["I"] = data.MaritalStatus
	m["J"] = data.Gender
	m["K"] = data.Address
	m["L"] = data.IssuingAuthority
	m["M"] = data.DateOfIssue
	m["N"] = data.DateOfExpiry

	for k, v := range m {
		cellIndex := fmt.Sprintf("%s%d", k, indx)
		xlsx.SetCellValue("Input", cellIndex, v)
	}
}
