package convertors

import (
	"github.com/SteveAzz/lemonade-stand/models"
	"github.com/SteveAzz/lemonade-stand/viewmodels"
)

func ConvertCategoryToViewModel(category models.Category) viewmodels.Category {
	return viewmodels.Category{
		ID:            category.ID(),
		ImageURL:      category.ImageURL(),
		Title:         category.Title(),
		Description:   category.Description(),
		IsOrientRight: category.IsOrientRight(),
	}
}
