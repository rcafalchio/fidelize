package rules

import (
	"fidelize/accreditation-merchant-service/models"
)

//ValidateInsetion validade data created by frontend
func ValidateInsetion(merchant *models.Merchant) (bool, string) {

	if merchant.Name == "" {
		return false, "Name is mandatory"
	}

	println("TESTE RICARDO - merchant.CNPJ ", merchant.CNPJ)

	if merchant.CNPJ == 0 {
		return false, "CNPJ is mandatory"
	}

	if merchant.Mail == "" {
		return false, "Email is mandatory"
	}

	return true, ""
}

//FindSegment find segment by id informed
func FindSegment(merchant *models.Merchant) {

	if merchant.Segment.ID == 1 {
		merchant.Segment.Name = "ESTETICA E BELEZA"
	} else if merchant.Segment.ID == 2 {
		merchant.Segment.Name = "PADARIA"
	} else if merchant.Segment.ID == 3 {
		merchant.Segment.Name = "ACADEMIA"
	} else if merchant.Segment.ID == 4 {
		merchant.Segment.Name = "RESTAURANTE"
	} else if merchant.Segment.ID == 5 {
		merchant.Segment.Name = "SAUDE"
	} else {
		merchant.Segment.ID = 99
		merchant.Segment.Name = "OUTROS"
	}

}
