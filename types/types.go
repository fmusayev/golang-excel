package types

type InputData struct {
	CompanyId string `json:"companyId"`
	PinCode   string `json:"pinCode"`
}

type OutputData struct {
	DocumentNumber       string `json:"documentNumber"`       // Şəxsiyyət vəsiqəsi seriya no
	Name                 string `json:"name"`                 // Ad
	Surname              string `json:"surname"`              // Soyad
	Patronymic           string `json:"patronymic"`           // Ata adı
	BirthDate            string `json:"birthDate"`            // Doğum tarixi
	BirthCountryName     string `json:"birthCountryName"`     // Doğulduğu olke
	BirthCity            string `json:"birthCity"`            // Doğulduğu sheher
	MaritalStatus        string `json:"maritalStatus"`        // Ailə vəziyyəti
	Gender               string `json:"gender"`               // Cinsi
	RegistrationAddress  string `json:"registrationAddress"`  // Yaşayış yeri
	DocGivenOrganization string `json:"docGivenOrganization"` // Vəsiqəni verən orqan
	DocGivenDate         string `json:"docGivenDate"`         // Vəsiqənin verilmə tarixi
	ExpireDate           string `json:"expireDate"`           // Vəsiqənin etibarlı olduğu tarix
}
