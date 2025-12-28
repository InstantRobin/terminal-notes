package reader

import (
	"fmt"
	"os"
	"terminal-notes/notes"
)

type NoteReader interface {
	ReadNote(noteName string) (*notes.Note, error)
}

type fileNoteReader struct {
	notesRootDir string
}

func NewNoteReader(rootDir string) NoteReader {
	return &fileNoteReader{notesRootDir: rootDir}
}

func (nr *fileNoteReader) ReadNote(noteName string) (*notes.Note, error) {
	noteFileName := nr.formatNoteFileName(noteName)
	noteFilePath := nr.notesRootDir + noteFileName

	noteFile, err := os.ReadFile(noteFilePath)
	if err != nil {
		return nil, fmt.Errorf("unable to open note file at %s: %s", noteFilePath, err.Error())
	}

	contents := string(noteFile)
	note := notes.Note{
		Name:     noteName,
		Contents: contents,
	}
	return &note, nil
}

func (nr *fileNoteReader) formatNoteFileName(noteName string) string {
	return noteName + ".md"
}
