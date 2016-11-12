package models

import "testing"

func Test_ReturnsNonEmptySlice(t *testing.T) {
	categories := GetCategories()

	if len(categories) == 0 {
		t.Log("Non Empty slice returned")
		t.Fail()
	}
}
