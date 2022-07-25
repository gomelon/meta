package meta

type Meta struct {
	name       string
	properties map[string]string
}

func New(name string) *Meta {
	return &Meta{
		name:       name,
		properties: map[string]string{},
	}
}

func (m *Meta) Name() string {
	return m.name
}

func (m *Meta) Property(key string) string {
	return m.properties[key]
}

func (m *Meta) SetProperty(key, value string) *Meta {
	m.properties[key] = value
	return m
}

func (m *Meta) Properties() map[string]string {
	return m.properties
}

func (m *Meta) SetProperties(properties map[string]string) *Meta {
	for k, v := range properties {
		m.properties[k] = v
	}
	return m
}

//Group a group is consists of multiple Meta with the same name
type Group []*Meta
