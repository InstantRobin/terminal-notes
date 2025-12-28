package main

import (
	"flag"
	"fmt"
	"terminal-notes/notemgr"
)

const (
	NOTES_DIR = "./note_files/"
	MAX_ARGS  = 1
)

func main() {

	noteManager := notemgr.NewNoteManager(NOTES_DIR)

	edit := flag.Bool("e", false, "Edit target note")

	flag.Parse()

	args := flag.Args()

	if len(args) <= 0 || len(args) > MAX_ARGS {
		fmt.Printf("Invalid number of args, max is %d, min is 0\n", MAX_ARGS)
		return
	}

	noteName := args[0]

	if edit != nil && *edit {
		err := noteManager.EditNote(noteName)
		if err != nil {
			fmt.Printf("Error editing note %s: %s\n", noteName, err.Error())
			return
		}
		fmt.Printf("Updated note %s\n", noteName)
	} else {
		note, err := noteManager.ReadNote(noteName)
		if err != nil {
			fmt.Printf("Error opening note %s: %s\n", noteName, err.Error())
			return
		}
		fmt.Println(note.Contents)
	}
}
