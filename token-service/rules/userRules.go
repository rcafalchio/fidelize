package rules

import (
	"fidelize/accreditation-merchant-service/models"
	"fmt"
)

//ValidateInsetion validade data created by frontend
func ValidateInsetion(user *models.User) (bool, string) {

	if user.Name == "" {
		return false, "Name is mandatory"
	}

	fmt.Println("user.CPF ==> ", user.CPF)

	if user.CPF == "" {
		return false, "CPF is mandatory"
	}

	if user.Mail == "" {
		return false, "User email is mandatory"
	}

	return true, ""
}
