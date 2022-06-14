package meta

type Meta interface {
	Target() Type
	Name() string
	Repeatable() bool
}

//MetaGroup multiple meta of same name for a type
type MetaGroup []Meta

func (m MetaGroup) Target() Type {
	var target Type
	for _, meta := range m {
		target = target | meta.Target()
	}
	return target
}
