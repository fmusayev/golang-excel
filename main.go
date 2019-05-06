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
	m["C"] = data.DocumentNumber
	m["D"] = data.Name
	m["E"] = data.Surname
	m["F"] = data.Patronymic
	m["G"] = data.BirthCountryName + ", " + data.BirthCity
	m["H"] = data.BirthDate
	m["I"] = data.MaritalStatus
	m["J"] = data.Gender
	m["K"] = data.RegistrationAddress
	m["L"] = data.DocGivenOrganization
	m["M"] = data.DocGivenDate
	m["N"] = data.ExpireDate

	for k, v := range m {
		cellIndex := fmt.Sprintf("%s%d", k, indx)
		xlsx.SetCellValue("Input", cellIndex, v)
	}
}
