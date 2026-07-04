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
	list := flag.Bool("l", false, "List Note Titles")
	search := flag.Bool("s", false, "Search Note Titles")
	delete := flag.Bool("d", false, "Delete Note")

	flag.Parse()
	args := flag.Args()

	editor := getEditorFromEnvVars()
	noteManager, err := notemgr.NewNoteManager(notesDir, editor)
	if err != nil {
		fmt.Printf("Failed to initialize notes: %s\n", err.Error())
		return
	}

	flags := []*bool{edit, all, list, search, delete}
	if err = veryifyFlags(flags); err != nil {
		fmt.Printf("Invalid command: %s\n", err.Error())
		return
	}

	if edit != nil && *edit {
		editNote(noteManager, args)
	} else if all != nil && *all {
		getAllNotes(noteManager)
	} else if list != nil && *list {
		listNotes(noteManager)
	} else if search != nil && *search {
		searchNotes(noteManager, args)
	} else if delete != nil && *delete {
		deleteNote(noteManager, args)
	} else {
		getNote(noteManager, args)
	}
}

func getNotesDir() (string, error) {
	usrHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Unable to get user home: %w", err)
	}
	return usrHomeDir + "/" + NOTES_PATH, nil
}

func veryifyFlags(flags []*bool) error {
	flagSet := false
	for _, flag := range flags {
		if flag == nil || *flag == false {
			continue
		}
		if flagSet {
			return fmt.Errorf("only one flag can be set at a time")
		}
		flagSet = true
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

func getAllNotes(noteManager notemgr.NoteManager) {
	searchAndPrintNotes(noteManager, "")
}

func searchNotes(noteManager notemgr.NoteManager, args []string) {
	noteQuery, err := getNoteNameFromArgs(args)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	searchAndPrintNotes(noteManager, noteQuery)
}

func searchAndPrintNotes(noteManager notemgr.NoteManager, query string) {
	noteSlice, err := noteManager.GetNotes(query)
	if err != nil {
		fmt.Printf("Error fetching all notes: %s\n", err.Error())
		return
	}
	if noteSlice == nil {
		fmt.Printf("Error fetching all notes: result is nil\n")
		return
	}

	notes.SortNotes(noteSlice)
	for _, note := range noteSlice {
		printFormattedNote(note)
	}
}

func listNotes(noteManager notemgr.NoteManager) {
	noteNames, err := noteManager.ListNotes()
	if err != nil {
		fmt.Printf("Error listing notesL %s\n", err.Error())
		return
	}
	for _, note := range noteNames {
		fmt.Printf("%s\n", note)
	}
}

func deleteNote(noteManager notemgr.NoteManager, args []string) {
	noteName, err := getNoteNameFromArgs(args)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	err = noteManager.DeleteNote(noteName)
	if err != nil {
		fmt.Printf("Error deleting note %s: %s\n", noteName, err.Error())
		return
	}
	fmt.Printf("Deleted note %s\n", noteName)
}

func getNote(noteManager notemgr.NoteManager, args []string) {
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
		fmt.Printf("Error opening note %s: note is nil\n", noteName)
	}
	printFormattedNote(*note)
}

func printFormattedNote(note notes.Note) {
	fmt.Println(note.Name)
	contents := strings.TrimRight(note.Contents, "\n")
	for line := range strings.SplitSeq(contents, "\n") {
		fmt.Println("    " + line)
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
