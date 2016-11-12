package models

type Category struct {
	id            int
	imageURL      string
	title         string
	description   string
	isOrientRight bool
}

func (category *Category) ID() int {
	return category.id
}

func (category *Category) ImageURL() string {
	return category.imageURL
}

func (category *Category) Title() string {
	return category.title
}

func (category *Category) Description() string {
	return category.description
}

func (category *Category) IsOrientRight() bool {
	return category.isOrientRight
}

func (category *Category) SetID(value int) {
	category.id = value
}

func (category *Category) SetImageURL(value string) {
	category.imageURL = value
}

func (category *Category) SetTitle(value string) {
	category.title = value
}

func (category *Category) SetDescription(value string) {
	category.description = value
}

func (category *Category) SetIsOrientRight(value bool) {
	category.isOrientRight = value
}

func GetCategories() []Category {
	return []Category{
		Category{
			id:       1,
			imageURL: "lemon.png",
			title:    "Juices and Mixes",
			description: `Explore our wide assortment of juices and mixes expected by
                                today's lemonade stand clientelle. Now featuring a full line of
                                organic juices that are guaranteed to be obtained from trees that
                                have never been treated with pesticides or artificial
                                fertilizers.`,
			isOrientRight: false,
		}, Category{
			id:       2,
			imageURL: "kiwi.png",
			title:    "Cups, Straws, and Other Supplies",
			description: `From paper cups to bio-degradable plastic to straws and
                            napkins, LSS is your source for the sundries that keep your stand
                            running smoothly.`,
			isOrientRight: true,
		}, Category{
			id:       3,
			imageURL: "pineapple.png",
			title:    "Signs and Advertising",
			description: `Sure, you could just wait for people to find your stand
                            along the side of the road, but if you want to take it to the next
                            level, our premium line of advertising supplies.`,
			isOrientRight: false,
		},
	}
}
