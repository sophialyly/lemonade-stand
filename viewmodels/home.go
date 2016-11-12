package viewmodels

// Home is where all the information
// about the home page needs to be shown
type Home struct {
	Title  string
	Active string
}

// GetHome will return the data for
// the home page
func GetHome() Home {
	result := Home{
		Title:  "Lemonade Stand",
		Active: "home",
	}

	return result
}

type Login struct {
	Title  string
	Active string
}

func GetLogin() Login {
	return Login{
		Title:  "Lemonade Stand Society - Login",
		Active: "",
	}
}
