package databinding

type Species struct {
	Id   int
	Name string
}

func KnownSpecies() []*Species {
	return []*Species{
		{1, "Dog"},
		{2, "Cat"},
		{3, "Bird"},
		{4, "Fish"},
		{5, "Elephant"},
	}
}
