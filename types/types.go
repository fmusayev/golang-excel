package types

type ExcelForm struct {
	ID           string `json:"id" col:"A"`
	FirstName    string `json:"firstName" col:"B"`
	LastName     string `json:"lastName" col:"C"`
	Username     String `json:"username" col:"D"`
	Email        string `json:"email" col:"E"`
	AddressLine  String `json:"addressLine" col:"F"`
	ContactPhone string `json:"contactPhone" col:"G"`
}
