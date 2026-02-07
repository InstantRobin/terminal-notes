package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"terminal-notes/notemgr"
	"terminal-notes/notes"
)

const (
	NOTES_PATH     = ".local/share/terminal-notes"
	MAX_ARGS       = 1
	DEFAULT_EDITOR = "vi"
)

func main() {

	notesDir, err := getNotesDir()
	if err != nil {
		fmt.Printf("Error finding notes directory: %s", err.Error())
		return
	}

	edit := flag.Bool("e", false, "Edit or Create target note")
	all := flag.Bool("a", false, "Return all Notes")

	flag.Parse()
	args := flag.Args()

	editor := getEditorFromEnvVars()
	noteManager, err := notemgr.NewNoteManager(notesDir, editor)
	if err != nil {
		fmt.Printf("Failed to initialize notes: %s\n", err.Error())
		return
	}

	if err = verifyArgsAndFlags(args, edit, all); err != nil {
		fmt.Printf("Invalid command: %s\n", err.Error())
		return
	}

	if edit != nil && *edit {
		editNote(noteManager, args)
	} else if all != nil && *all {
		fetchAndPrintAllNotes(noteManager)
	} else {
		fetchAndPrintNote(noteManager, args)
	}
}

func getNotesDir() (string, error) {
	usrHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Unable to get user home: %w", err)
	}
	return usrHomeDir + "/" + NOTES_PATH, nil
}

func verifyArgsAndFlags(args []string, edit, all *bool) error {
	if edit != nil && all != nil && *edit && *all {
		return fmt.Errorf("incompatible flags '-e' and '-a")
	}
	if all != nil && *all {
		if len(args) != 0 {
			return fmt.Errorf("invalid number of args, must be zero")
		}
	}
	return nil
}

func getNoteNameFromArgs(args []string) (string, error) {
	if len(args) <= 0 || len(args) > MAX_ARGS {
		return "", fmt.Errorf("Invalid number of args, max is %d, min is 0\n", MAX_ARGS)
	}
	noteName := strings.ToLower(args[0])
	return noteName, nil
}

func editNote(noteManager notemgr.NoteManager, args []string) {
	noteName, err := getNoteNameFromArgs(args)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	err = noteManager.EditNote(noteName)
	if err != nil {
		fmt.Printf("Error editing note %s: %s\n", noteName, err.Error())
		return
	}
	fmt.Printf("Updated note %s\n", noteName)
}

func fetchAndPrintAllNotes(noteManager notemgr.NoteManager) {
	noteSlice, err := noteManager.GetNotes("")
	if err != nil {
		fmt.Printf("Error fetching all notes: %s", err.Error())
		return
	}
	if noteSlice == nil {
		fmt.Printf("Error fetching all notes: result is nil")
		return
	}

	notes.SortNotes(noteSlice)
	for _, note := range noteSlice {
		printFormattedNote(note)
	}
}

func fetchAndPrintNote(noteManager notemgr.NoteManager, args []string) {
	noteName, err := getNoteNameFromArgs(args)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	note, err := noteManager.GetNote(noteName)
	if err != nil {
		fmt.Printf("Error opening note %s: %s\n", noteName, err.Error())
		return
	}
	if note == nil {
		fmt.Printf("Error opening note %s: note is nil", noteName)
	}
	printFormattedNote(*note)
}

func printFormattedNote(note notes.Note) {
	fmt.Println(note.Name)
	for line := range strings.SplitSeq(note.Contents, "\n") {
		fmt.Println("\t" + line)
	}
}

func getEditorFromEnvVars() string {
	visualEnv := os.Getenv("VISUAL")
	editorEnv := os.Getenv("EDITOR")
	editorCommand := DEFAULT_EDITOR

	if visualEnv != "" {
		editorCommand = visualEnv
	} else if editorEnv != "" {
		editorCommand = editorEnv
	}
	return editorCommand
}
