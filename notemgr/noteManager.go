package notemgr

import (
	"fmt"
	"os"
	"terminal-notes/notes"
)

type NoteManager interface {
	ReadNote(noteName string) (*notes.Note, error)
}

type fileNoteManager struct {
	notesRootDir string
}

func NewNoteManager(rootDir string) NoteManager {
	return &fileNoteManager{notesRootDir: rootDir}
}

func (mgr *fileNoteManager) ReadNote(noteName string) (*notes.Note, error) {
	noteFileName := mgr.formatNoteFileName(noteName)
	noteFilePath := mgr.notesRootDir + noteFileName

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

func (nr *fileNoteManager) formatNoteFileName(noteName string) string {
	return noteName + ".md"
}
