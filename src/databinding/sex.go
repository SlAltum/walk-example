package databinding

type Sex byte

const (
	SexMale Sex = 1 + iota
	SexFemale
	SexHermaphrodite
)
