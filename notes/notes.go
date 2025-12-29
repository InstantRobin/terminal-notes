package notes

import "sort"

type Note struct {
	Name string
	Contents string
}

func SortNotes(noteSlice []Note) {
	sort.Slice(noteSlice, func(i, j int) bool {
		return noteSlice[i].Name < noteSlice[j].Name
	})
}