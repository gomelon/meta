package meta

type Meta interface {
	Target() Type
	Directive() string
	Repeatable() bool
}

//Group a group is consists of multiple Meta with the same name
type Group []Meta
