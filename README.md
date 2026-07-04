# terminal-notes
A simple command line utility to take notes

Inspired by coming back to a microservice I haven't touched in 6 months and forgetting every command I used last time.

Notes are not encrypted, please avoid storing any credentials in your notes.

Editor is set by `$VISUAL` and `$EDITOR` env vars, uses `vi` by default.

Requires Go version 1.25 or greater.

## Building

Build with `go build -o terminal-notes`

## Installation

Install with `sudo ./install.sh`

It is recommended to alias `terminal-notes` into something more quickly typeable, like `tn` or `tnotes`

You can do so by inserting into your `.bashrc` file `alias {command}=terminal-notes` 

For example:

Add `alias tn=terminal-notes`

This will allow you to run `tn {noteName}`

## Uninstallation

Uninstall with `sudo ./uninstall.sh`

You will be prompted as to whether to delete your notes or not. 

## Commands

### `terminal-notes {note name}`

Fetches a note by name

```
~$ terminal-notes tnotes
tnotes
    A useful terminal based note program
```

### `terminal-notes -a`

Fetches all notes, sorted alphabetically.

```
~$ terminal-notes -a
noted
    Past tense for of verb "to note" 
tnotes
    note manager
```

### `terminal-notes -l`

Fetches names of all notes, sorted alphabetically.

```
~$ terminal-notes -l
noted
tnotes
```

### `terminal-notes -s {query}`

Search all notes titles for the given query. 
Returns all matching notes, sorted alphabetically. 

```
~$ terminal-notes -s note
noted
    Past tense for of verb "to note" 
tnotes
    A useful terminal based note program
```

### `terminal-notes -e {note name}`

Edits a note, or creates it if it does not already exist. 

If set, uses the editor definied in `$VISUAL` or `$EDITOR`, defaults to `vi`


```
~$ terminal-notes -e tnotes
```

```vi
A useful terminal based note program
```

### `terminal-notes -d {note name}`

Deletes a note. Warning: Notes are deleted permanently.

```
~$ terminal notes -d tnotes
```
