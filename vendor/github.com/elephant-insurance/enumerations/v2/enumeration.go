package enumerations

/* THIS FILE IS OBSOLETE !!!
I am leaving it here so that we can compile while we wait to get the Occupations and military stuff straigtened out.
This file should be deleted in its entirety as soon as those enums are updated to the latest standard.
*/

// Enumeration handles reference data/lookups.  It's a collection of EnumerationItems
type Enumeration struct {
	Name  string
	Desc  string
	Items []EnumerationItem
}

// EnumerationItem represents each, distinct enumeration key/description
type EnumerationItem struct {
	ID        string `structs:"Value" json:"Value"`
	Desc      string
	Meta      map[string]string `json:",omitempty"`
	SortOrder int
}

// HasID => searches through a map of string keys and
//           1. returns true and index if found
//           2. returns false and -1 if not found
func (enum Enumeration) HasID(id string) (bool, int) {
	for idx, ei := range enum.Items {
		if ei.ID == id {
			return true, idx
		}
	}

	return false, -1
}

// make a copy of an enum item
func (i EnumerationItem) cloneItem() EnumerationItem {
	newMeta := make(map[string]string)
	for k, v := range i.Meta {
		newMeta[k] = v
	}
	return EnumerationItem{
		ID:        i.ID,
		Desc:      i.Desc,
		Meta:      newMeta,
		SortOrder: i.SortOrder,
	}
}

// CloneEnum creates a copy of an Enumeration struct
// which can be altered without causing side-effects
func CloneEnum(e *Enumeration) *Enumeration {
	newItems := make([]EnumerationItem, len(e.Items))
	for idx := 0; idx < len(e.Items); idx++ {
		newItems[idx] = e.Items[idx].cloneItem()
	}
	rtn := Enumeration{
		Name:  e.Name,
		Desc:  e.Desc,
		Items: newItems,
	}

	return &rtn
}
