package meta

type Meta interface {
	PlaceAt() Place
	Directive() string
	Repeatable() bool
}

//Group a group is consists of multiple Meta with the same name
type Group []Meta
