package main

import (
	"flag"
	"fmt"
	"terminal-notes/reader"
)

const (
	NOTES_DIR = "./note_files/"
	MAX_ARGS  = 1
)

func main() {

	noteReader := reader.NewNoteReader(NOTES_DIR)

	flag.Parse()

	args := flag.Args()

	if len(args) <= 0 || len(args) > MAX_ARGS {
		fmt.Printf("Invalid number of args, max is %d, min is 0\n", MAX_ARGS)
		return
	}

	noteName := args[0]

	note, err := noteReader.ReadNote(noteName)
	if err != nil {
		fmt.Printf("Error opening note for %s: %s\n", noteName, err.Error())
		return
	}

	fmt.Println(note.Contents)
}
