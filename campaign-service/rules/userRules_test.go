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
	u := models.User{}

	if ValidateInsetion(&u) == "" {
		t.Fail()
	}
}
