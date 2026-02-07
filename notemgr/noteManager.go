package notemgr

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"terminal-notes/notes"
)

const fileTypeSuffix = ".md"

var ErrInvalidFileType = errors.New("file is not a supported note file type")

type NoteManager interface {
	GetNote(noteName string) (*notes.Note, error)
	GetNotes(noteNameSubstr string) ([]notes.Note, error)
	EditNote(noteName string) error
}

type fileNoteManager struct {
	notesRootDir   string
	fileTypeSuffix string
	editor         string
}

func NewNoteManager(rootDir, editor string) (NoteManager, error) {

	err := os.MkdirAll(rootDir, 0700)
	if err != nil {
		return nil, fmt.Errorf("Unable to open directory %s: %w", rootDir, err)
	}

	if rootDir[len(rootDir)-1] != '/' {
		rootDir = rootDir + "/"
	}

	return &fileNoteManager{
		notesRootDir:   rootDir,
		fileTypeSuffix: fileTypeSuffix,
		editor:         editor,
	}, nil
}

func (mgr *fileNoteManager) GetNote(noteName string) (*notes.Note, error) {
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

func (mgr *fileNoteManager) GetNotes(noteNameSubstr string) ([]notes.Note, error) {
	notesDir, err := os.ReadDir(mgr.notesRootDir)
	if err != nil {
		return nil, fmt.Errorf("unable to open notes directory at %s: %s", mgr.notesRootDir, err.Error())
	}

	noteFileNames := []string{}
	for _, entry := range notesDir {
		if !entry.IsDir() && strings.Contains(entry.Name(), noteNameSubstr) {
			noteName, err := mgr.getNoteNameFromFileName(entry.Name())
			if err != nil {
				if errors.Is(err, ErrInvalidFileType) {
					continue
				}
				return nil, fmt.Errorf("error getting note name from file name for file %s: %s", entry.Name(), err.Error())
			}
			noteFileNames = append(noteFileNames, noteName)
		}
	}

	notes := []notes.Note{}
	for _, noteFileName := range noteFileNames {
		note, err := mgr.GetNote(noteFileName)
		if err != nil {
			return nil, fmt.Errorf("error opening note %s: %s", noteFileName, err.Error())
		}
		if note == nil {
			// Should never happen
			return nil, fmt.Errorf("eil note found for note %s", noteFileName)
		}
		notes = append(notes, *note)
	}
	return notes, nil
}

func (mgr *fileNoteManager) EditNote(noteName string) error {
	noteFilePath := mgr.getNoteFilePath(noteName)

	// TODO: Make this more secure (See os.root maybe?)
	cmd := exec.Command(mgr.editor, noteFilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("unable to edit file %s: %s", noteFilePath, err)
	}
	return nil
}

func (mgr *fileNoteManager) getNoteFilePath(noteName string) string {
	noteFileName := noteName + ".md"
	return mgr.notesRootDir + noteFileName
}

func (mgr *fileNoteManager) getNoteNameFromFileName(fileName string) (string, error) {
	noteName, found := strings.CutSuffix(fileName, mgr.fileTypeSuffix)
	if !found {
		return "", ErrInvalidFileType
	}
	return noteName, nil
}
