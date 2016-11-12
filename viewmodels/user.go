package viewmodels

type Profile struct {
	Title  string
	Active string
	User   User
}

type User struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Stand     Stand
}

type Stand struct {
	Address string
	City    string
	State   string
	Zip     string
}

func GetProfile() Profile {
	return Profile{
		Title: "Lemonade Stand Supply - Profile",
	}
}
