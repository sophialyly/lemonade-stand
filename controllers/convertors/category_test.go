package convertors

import (
	"testing"

	"github.com/SteveAzz/lemonade-stand/models"
)

func Test_ConvertsCategoryToViewModel(t *testing.T) {
	category := models.Category{}
	category.SetID(1)
	category.SetImageURL("the image url")
	category.SetTitle("title")
	category.SetDescription("description")
	category.SetIsOrientRight(true)

	result := ConvertCategoryToViewModel(category)

	if result.ID != category.ID() {
		t.Log("Id is not converting properly")
		t.Fail()
	}

	if result.ImageURL != category.ImageURL() {
		t.Log("ImageURL is not converting properly")
		t.Fail()
	}

	if result.Description != category.Description() {
		t.Log("Description is not converting properly")
		t.Fail()
	}

	if result.Title != category.Title() {
		t.Log("Title is not converting properly")
		t.Fail()
	}

	if result.IsOrientRight != category.IsOrientRight() {
		t.Log("IsOrientRight is not converting properly")
		t.Fail()
	}
}
