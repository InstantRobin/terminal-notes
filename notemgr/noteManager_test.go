package notemgr

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NoteManagerTestSuite struct {
	suite.Suite
	mgr     NoteManager
	tempDir string
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(NoteManagerTestSuite))
}

func (s *NoteManagerTestSuite) SetupTest() {
	tempDir, err := os.MkdirTemp(".", "test-notes-")
	if err != nil {
		s.FailNow("Error creating temp directory %s", err)
		return
	}
	s.tempDir = tempDir
	s.mgr, err = NewNoteManager(s.tempDir, "vim")
	s.Require().NoError(err)
}

func (s *NoteManagerTestSuite) TearDownTest() {
	err := os.RemoveAll("./" + s.tempDir)
	if err != nil {
		s.FailNowf("Failed to delete directory %s. Please delete manually.", s.tempDir)
		return
	}
}

// The test suite can't use vim or other editing software
// So we have to edit notes manually
func (s *NoteManagerTestSuite) WriteNote(name, content string) {
	err := os.WriteFile(s.tempDir+"/"+name+".md", []byte(content), 0666)
	s.Require().NoError(err)
}

func (s *NoteManagerTestSuite) TestGetNote() {
	noteA := "noteA"
	noteB := "noteB"

	noteAContent := "A Content"
	noteBContent := "B Content"

	s.WriteNote(noteB, noteBContent)
	s.WriteNote(noteA, noteAContent)

	note, err := s.mgr.GetNote(noteB)
	s.Require().NoError(err)

	s.Equal(noteBContent, note.Contents)
}

func (s *NoteManagerTestSuite) TestGetNoteComplex() {
	noteA := "noteA"

	noteAContent := "A Content\nSome other stuff\n漢字\r\n其他的東西\r\n"

	s.WriteNote(noteA, noteAContent)

	note, err := s.mgr.GetNote(noteA)
	s.Require().NoError(err)

	s.Equal(noteAContent, note.Contents)
}

func (s *NoteManagerTestSuite) TestGetNotes() {
	noteA := "noteA"
	noteB := "noteB"
	noteC := "some other thing"

	noteAContent := "A Content"
	noteBContent := "B Content"
	noteCContent := "unrelated contents"

	s.WriteNote(noteB, noteBContent)
	s.WriteNote(noteA, noteAContent)
	s.WriteNote(noteC, noteCContent)

	notes, err := s.mgr.GetNotes("note")
	s.Require().NoError(err)

	s.Equal(noteA, notes[0].Name)
	s.Equal(noteAContent, notes[0].Contents)
	s.Equal(noteB, notes[1].Name)
	s.Equal(noteBContent, notes[1].Contents)
}

func (s *NoteManagerTestSuite) TestGetAllNotes() {
	noteA := "noteA"
	noteB := "noteB"
	noteC := "some other thing"

	noteAContent := "A Content"
	noteBContent := "B Content"
	noteCContent := "unrelated contents"

	s.WriteNote(noteB, noteBContent)
	s.WriteNote(noteA, noteAContent)
	s.WriteNote(noteC, noteCContent)

	notes, err := s.mgr.GetNotes("")
	s.Require().NoError(err)

	s.Equal(noteA, notes[0].Name)
	s.Equal(noteAContent, notes[0].Contents)
	s.Equal(noteB, notes[1].Name)
	s.Equal(noteBContent, notes[1].Contents)
	s.Equal(noteC, notes[2].Name)
	s.Equal(noteCContent, notes[2].Contents)
}

func (s *NoteManagerTestSuite) TestListNoNotes() {
	notes, err := s.mgr.ListNotes()
	s.Require().NoError(err)
	s.Empty(notes)
}

func (s *NoteManagerTestSuite) TestListOneNote() {
	noteName := "test"
	s.WriteNote(noteName, "contents")
	notes, err := s.mgr.ListNotes()
	s.Require().NoError(err)

	expected := []string{noteName}
	s.Equal(expected, notes)
}

func (s *NoteManagerTestSuite) TestListTwoNotesAlphabetical() {
	noteA := "noteA"
	noteB := "noteB"

	s.WriteNote(noteB, "contents")
	s.WriteNote(noteA, "contents")
	notes, err := s.mgr.ListNotes()
	s.Require().NoError(err)

	expected := []string{noteA, noteB}
	s.Equal(expected, notes)
}

func (s *NoteManagerTestSuite) TestDeleteNote() {
	noteA := "noteA"
	noteB := "noteB"

	s.WriteNote(noteB, "contents")
	s.WriteNote(noteA, "contents")
	notes, err := s.mgr.ListNotes()
	s.Require().NoError(err)

	expected := []string{noteA, noteB}
	s.Equal(expected, notes)

	err = s.mgr.DeleteNote(noteA)
	s.Require().NoError(err)

	notes, err = s.mgr.ListNotes()
	s.Require().NoError(err)

	expected = []string{noteB}
	s.Equal(expected, notes)
}

func (s *NoteManagerTestSuite) TestDeleteOnlyNote() {
	noteName := "test"
	s.WriteNote(noteName, "contents")
	notes, err := s.mgr.ListNotes()
	s.Require().NoError(err)

	expected := []string{noteName}
	s.Equal(expected, notes)

	err = s.mgr.DeleteNote(noteName)
	s.Require().NoError(err)

	notes, err = s.mgr.ListNotes()
	s.Require().NoError(err)
	s.Empty(notes)
}
