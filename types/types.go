package types

type InputData struct {
	CompanyId string `json:"companyId"`
	PinCode   string `json:"pinCode"`
}

type OutputData struct {
	PersonalNo       string `json:"personalNo"`       // Şəxsiyyət vəsiqəsi seriya no
	GivenName        string `json:"givenName"`        // Ad
	Surname          string `json:"surname"`          // Soyad
	Patronymic       string `json:"patronymic"`       // Ata adı
	PlaceOfBirth     string `json:"placeOfBirth"`     // Doğulduğu yer
	DateOfBirth      string `json:"dateOfBirth"`      // Doğum tarixi
	MaritalStatus    string `json:"maritalStatus"`    // Ailə vəziyyəti
	Gender           string `json:"gender"`           // Cinsi
	Address          string `json:"address"`          // Yaşayış yeri
	IssuingAuthority string `json:"issuingAuthority"` // Vəsiqəni verən orqan
	DateOfIssue      string `json:"dateOfIssue"`      // Vəsiqənin verilmə tarixi
	DateOfExpiry     string `json:"dateOfExpiry"`     // Vəsiqənin etibarlı olduğu tarix
}
