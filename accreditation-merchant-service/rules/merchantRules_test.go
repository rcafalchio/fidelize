package rules

// code based on technique explained here:
// https://www.elastic.co/blog/code-coverage-for-your-golang-system-tests
// you can look there if you want to see how not to execute this test
// when running unit test etc.

// This file is mandatory as otherwise the packetbeat.test binary is not generated correctly.

import (
	"fidelize/accreditation-merchant-service/models"
	"testing"
)

// Test started when the test binary is started. Only calls main.
func TestEmptyName(t *testing.T) {
	m := models.Merchant{}
	m.CNPJ = 1027058000191 //"01.027.058/0001-91"
	m.Name = "TESTE"
	m.Mail = "teste@teste"
	result, _ := ValidateInsetion(&m)

	if !result {
		t.Fail()
	}
}
