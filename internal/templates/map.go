package templates

import "github.com/charmbracelet/huh"

type SortedGroupMap struct {
	values map[string][]huh.Field
	keys   []string
}

func NewSortedGroupMap() *SortedGroupMap {
	return &SortedGroupMap{
		values: make(map[string][]huh.Field),
		keys:   make([]string, 0),
	}
}

func (m *SortedGroupMap) Set(key string, value []huh.Field) {
	if _, ok := m.values[key]; !ok {
		// save ordered key only first time
		m.keys = append(m.keys, key)
	}
	m.values[key] = value
}

func (m *SortedGroupMap) Get(key string) ([]huh.Field, bool) {
	exists := true
	if _, ok := m.values[key]; !ok {
		exists = false
	}
	return m.values[key], exists
}

func (m *SortedGroupMap) FormGroups() []*huh.Group {
	var groups []*huh.Group
	for _, key := range m.keys {
		groups = append(groups, huh.NewGroup(m.values[key]...))
	}

	return groups
}
