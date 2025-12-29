package main

import (
	"flag"
	"fmt"
	"strings"
	"terminal-notes/notemgr"
	"terminal-notes/notes"
)

const (
	NOTES_DIR = "./note_files/"
	MAX_ARGS  = 1
)

func main() {

	noteManager := notemgr.NewNoteManager(NOTES_DIR)

	edit := flag.Bool("e", false, "Edit target note")
	all := flag.Bool("a", false, "Return all Notes")

	flag.Parse()

	args := flag.Args()

	if edit != nil && *edit {
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
	} else if all != nil && *all {
		noteSlice, err := noteManager.ReadNotesTitleContains("")
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
	} else {
		noteName, err := getNoteNameFromArgs(args)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		note, err := noteManager.ReadNote(noteName)
		if err != nil {
			fmt.Printf("Error opening note %s: %s\n", noteName, err.Error())
			return
		}
		if note == nil {
			fmt.Printf("Error opening note %s: note is nil", noteName)
		}
		printFormattedNote(*note)
	}
}

func verifyFlags(args []string, edit, all *bool) error {
	if edit != nil && all != nil && *edit && *all {
		return fmt.Errorf("incompatible flags '-e' and '-a")
	}
	return nil
}

func getNoteNameFromArgs(args []string) (string, error) {
	if len(args) <= 0 || len(args) > MAX_ARGS {
		return "", fmt.Errorf("Invalid number of args, max is %d, min is 0\n", MAX_ARGS)
	}
	noteName := args[0]
	return noteName, nil
}

func verifyArgsNil(args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("Invalid number of args, must be zero")
	}
	return nil
}

func printFormattedNote(note notes.Note) {
	fmt.Println(note.Name)
	for line := range strings.SplitSeq(note.Contents, "\n") {
		fmt.Println("\t" + line)
	}
}
