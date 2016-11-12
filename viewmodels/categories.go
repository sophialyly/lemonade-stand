package viewmodels

// Categories is all the information
// that need to be displayed
// in the categories page
type Categories struct {
	Title      string
	Active     string
	Categories []Category
}

// Category is the information
// about a specific
type Category struct {
	ID            int
	ImageURL      string
	Title         string
	Description   string
	IsOrientRight bool
}

// GetCategories creates all of the information
// needed to be displayed in the categories page.
func GetCategories() Categories {
	result := Categories{
		Title:  "Lemonade Stand Society - Shop",
		Active: "shop",
	}

	return result
}
