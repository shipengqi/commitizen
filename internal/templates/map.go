package templates

import (
	"github.com/charmbracelet/huh"

	"github.com/shipengqi/commitizen/internal/parameter"
)

type SortedGroupMap struct {
	groups  map[string]*parameter.Group
	fields  map[string][]huh.Field
	ordered []string
}

func NewSortedGroupMap(groups ...*parameter.Group) *SortedGroupMap {
	groupmap := make(map[string]*parameter.Group)
	for _, group := range groups {
		groupmap[group.Name] = group
	}
	return &SortedGroupMap{
		groups:  groupmap,
		fields:  make(map[string][]huh.Field),
		ordered: make([]string, 0),
	}
}

func (m *SortedGroupMap) Length() int {
	return len(m.ordered)
}

func (m *SortedGroupMap) SetFields(group string, fields []huh.Field) {
	if _, ok := m.fields[group]; !ok {
		// save ordered group name only first time
		m.ordered = append(m.ordered, group)
	}
	if _, ok := m.groups[group]; !ok {
		// create group if not exist
		m.groups[group] = parameter.NewGroup(group)
	}
	m.fields[group] = fields
}

func (m *SortedGroupMap) GetFields(group string) ([]huh.Field, bool) {
	exists := true
	if _, ok := m.fields[group]; !ok {
		exists = false
	}
	return m.fields[group], exists
}

func (m *SortedGroupMap) GetGroup(group string) (*parameter.Group, bool) {
	exists := true
	if _, ok := m.groups[group]; !ok {
		exists = false
	}
	return m.groups[group], exists
}

func (m *SortedGroupMap) FormGroups(all map[parameter.FieldKey]huh.Field) []*huh.Group {
	var groups []*huh.Group
	for _, key := range m.ordered {
		groups = append(groups, m.groups[key].Render(all, m.fields[key]))
	}
	return groups
}
