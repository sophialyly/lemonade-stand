package viewmodels

// Products is all the information that will be displayed to the user.
type Products struct {
	Title    string
	Active   string
	Products []Product
}

// GetProducts creates all of the needed data
func GetProducts(ID int) Products {
	var result Products

	var shopName string
	switch ID {
	case 1:
		shopName = "Juice"
	case 2:
		shopName = "Supply"
	case 3:
		shopName = "Advertising"
	}
	result.Title = "Lemonade Stand Society - " + shopName + " Shop"
	result.Active = "shop"

	if ID == 1 {
		result.Products = getProductList()
	}

	return result
}

func getProductList() []Product{
		lemonJuice := MakeLemonJuiceProduct()
		appleJuice := MakeAppleJuiceProduct()
		watermelonJuice := MakeWatermelonJuiceProduct()
		kiwiJuice := MakeKiwiJuiceProduct()
		mangosteenJuice := MakeMangosteenJuiceProduct()
		orangeJuice := MakeOrangeJuiceProduct()
		pineappleJuice := MakePineappleJuiceProduct()
		strawberryJuice := MakeStrawberryJuiceProduct()

		return []Product{
			lemonJuice,
			appleJuice,
			watermelonJuice,
			kiwiJuice,
			mangosteenJuice,
			orangeJuice,
			pineappleJuice,
			strawberryJuice,
		}
}

// ProductVM is the viewmodel
type ProductVM struct {
	Title   string
	Active  string
	Product Product
}

// GetProduct creates a new ProductVM
func GetProduct(ID int) ProductVM {

	productList := getProductList()

	var product Product

	for _, p := range productList {
		if p.ID == ID {
			product = p
			break;
		}
	}

	return ProductVM{
		Title:   "Lemonade Stand Society - " + product.Name,
		Active:  "shop",
		Product: product,
	}
}

// Product is all the information about a specific product
type Product struct {
	Name             string
	DescriptionShort string
	DescriptionLong  string
	PricePerLiter    float32
	PricePer10Liter  float32
	Origin           string
	IsOrganic        bool
	ImageURL         string
	ID               int
}

// MakeLemonJuiceProduct creates a new Product with lemon juice information
func MakeLemonJuiceProduct() Product {
	result := Product{
		Name:             "Lemon Juice",
		DescriptionShort: "Made from fresh, organic California lemons.",
		DescriptionLong: `Made from premium, organic Meyer lemons. These fruit are left on the tree until they reach the peak of ripeness and then juiced within 8 hours of being picked.
			<br/>
			Lemonade made from our premium juice is sure to make your stand the most popular in the neighborhood.`,
		PricePerLiter:   1.09,
		PricePer10Liter: 1.04,
		Origin:          "California",
		IsOrganic:       true,
		ImageURL:        "lemon.png",
		ID:              1,
	}

	return result
}

func MakeAppleJuiceProduct() Product {
	result := Product{
		Name:             "Apple Juice",
		DescriptionShort: "The perfect blend of Washington apples.",
		DescriptionLong:  "The perfect blend of Washington apples.",
		PricePerLiter:    0.89,
		PricePer10Liter:  0.85,
		Origin:           "Ohio",
		IsOrganic:        true,
		ImageURL:         "apple.png",
		ID:               2,
	}

	return result
}

func MakeWatermelonJuiceProduct() Product {
	return Product{
		Name:             "Watermelon Juice",
		DescriptionShort: "From sun-drenched fields in FlorIDa.",
		DescriptionLong:  "From sun-drenched fields in FlorIDa.",
		PricePerLiter:    3.99,
		PricePer10Liter:  3.84,
		Origin:           "FlorIDa",
		IsOrganic:        true,
		ImageURL:         "watermelon.png",
		ID:               3,
	}
}

func MakeKiwiJuiceProduct() Product {
	return Product{
		Name:             "Kiwi Juice",
		DescriptionShort: "California sunshine and rain distilled into sweet goodness",
		DescriptionLong:  "California sunshine and rain distilled into sweet goodness",
		PricePerLiter:    5.99,
		PricePer10Liter:  5.54,
		Origin:           "California",
		IsOrganic:        false,
		ImageURL:         "kiwi.png",
		ID:               4,
	}
}

func MakeMangosteenJuiceProduct() Product {
	return Product{
		Name:             "Mangosteen Juice",
		DescriptionShort: "Our latest taste sensation, imported directly from Hawaii",
		DescriptionLong:  "Our latest taste sensation, imported directly from Hawaii",
		PricePerLiter:    6.87,
		PricePer10Liter:  6.79,
		Origin:           "Hawaii",
		IsOrganic:        false,
		ImageURL:         "mangosteen.png",
		ID:               5,
	}
}

func MakeOrangeJuiceProduct() Product {
	return Product{
		Name:             "Orange Juice",
		DescriptionShort: "Fresh squeezed from FlorIDa's best oranges.",
		DescriptionLong:  "Fresh squeezed from FlorIDa's best oranges.",
		PricePerLiter:    1.67,
		PricePer10Liter:  1.63,
		Origin:           "FlorIDa",
		IsOrganic:        false,
		ImageURL:         "orange.png",
		ID:               6,
	}
}

func MakePineappleJuiceProduct() Product {
	return Product{
		Name:             "Pineapple Juice",
		DescriptionShort: "An exotic and refreshing offering. Straight from Hawaii.",
		DescriptionLong:  "An exotic and refreshing offering. Straight from Hawaii.",
		PricePerLiter:    2.48,
		PricePer10Liter:  2.27,
		Origin:           "Hawaii",
		IsOrganic:        false,
		ImageURL:         "pineapple.png",
		ID:               7,
	}
}

func MakeStrawberryJuiceProduct() Product {
	return Product{
		Name:             "Strawberry Juice",
		DescriptionShort: "MThe perfect balance of sweet and tart.",
		DescriptionLong:  "The perfect balance of sweet and tart.",
		PricePerLiter:    4.36,
		PricePer10Liter:  4.27,
		Origin:           "California",
		IsOrganic:        false,
		ImageURL:         "strawberry.png",
		ID:               8,
	}
}
