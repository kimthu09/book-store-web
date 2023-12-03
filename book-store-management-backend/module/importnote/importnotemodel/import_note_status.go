package importnotemodel

type ImportNoteStatus string

const (
	InProgress ImportNoteStatus = "InProgress"
	Done       ImportNoteStatus = "Done"
	Cancel     ImportNoteStatus = "Cancel"
)
