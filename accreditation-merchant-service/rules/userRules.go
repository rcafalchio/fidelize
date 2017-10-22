package rules

import (
	"fidelize/accreditation-merchant-service/models"
)

//ValidateInsetion validade data created by frontend
func ValidateInsetion(merchant *models.Merchant) (bool, string) {

	if merchant.Name == "" {
		return false, "Name is mandatory"
	}

	if merchant.CNPJ == "" {
		return false, "CNPJ is mandatory"
	}

	if merchant.Mail == "" {
		return false, "Email is mandatory"
	}

	return true, ""
}
