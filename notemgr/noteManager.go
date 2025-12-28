package notemgr

import (
	"fmt"
	"os"
	"os/exec"
	"terminal-notes/notes"
)

type NoteManager interface {
	ReadNote(noteName string) (*notes.Note, error)
	EditNote(noteName string) error
}

type fileNoteManager struct {
	notesRootDir string
}

func NewNoteManager(rootDir string) NoteManager {
	return &fileNoteManager{notesRootDir: rootDir}
}

func (mgr *fileNoteManager) ReadNote(noteName string) (*notes.Note, error) {
	noteFilePath := mgr.getNoteFilePath(noteName)

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

func (mgr *fileNoteManager) EditNote(noteName string) error {
	noteFilePath := mgr.getNoteFilePath(noteName)

	cmd := exec.Command("vim", noteFilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Unable to edit file %s: %s", noteFilePath, err)
	}
	return nil
}

func (mgr *fileNoteManager) getNoteFilePath(noteName string) string {
	noteFileName := noteName + ".md"
	return mgr.notesRootDir + noteFileName
}
