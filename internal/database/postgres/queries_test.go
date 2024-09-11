package postgres_test

import (
	"fmt"
	"kode/config"
	"kode/internal/database/postgres"
	"kode/internal/structs/notes"
	"kode/internal/structs/user"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var testCase = &user.User{
	Login:    "Vega",
	Mail:     "omega@gmail.com",
	Password: "qweqwe123",
}

var testNote = &notes.Note{
	Id:    1,
	Owner: 100,
	Date:  time.Now(),
	Title: "Pushkin",
	Body: `Я вас любил: любовь еще, быть может,
В душе моей угасла не совсем;
Но пусть она вас больше не тревожит;
Я не хочу печалить вас ничем.
Я вас любил безмолвно, безнадежно,
То робостью, то ревностью томим;
Я вас любил так искренно, так нежно,
Как дай вам бог любимой быть другим.`,
}

var conf = config.NewConfig()
var ps = postgres.NewPostgres()

func TestNewUser(t *testing.T) {

	encriptPassword, err := bcrypt.GenerateFromPassword([]byte(testCase.Password), 10)
	if err != nil {
		t.Fatal(err)
		return
	}

	err = ps.NewUser(conf, testCase.Login, testCase.Mail, string(encriptPassword), "user")
	if err != nil {
		t.Fatal(err)
		return
	}

}

func TestGetUser(t *testing.T) {

	id, login, mail, err := ps.GetUser(conf, testCase.Login, testCase.Password)
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Printf("Id: %d\nLogin: %s\nMail: %s\n", id, login, mail)

}

func TestGetUserId(t *testing.T) {

	id, err := postgres.GetUserId(conf, "Vega")
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Printf("ID:  %d\n", id)
}

func TestNewNote(t *testing.T) {

	id, err := postgres.GetUserId(conf, testCase.Login)
	if err != nil {
		t.Fatal(err)
		return
	}

	err = ps.NewNote(conf, id, testNote.Title, testNote.Body)
	if err != nil {
		t.Fatal(err)
		return
	}

}

func TestGetNotes(t *testing.T) {

	id, err := postgres.GetUserId(conf, testCase.Login)
	if err != nil {
		t.Fatal(err)
		return
	}

	arr, err := ps.GetNotes(conf, id)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, note := range arr {
		fmt.Printf("Id: %d, Owner: %d\n\"%s\"\t%s\n\n", note.Id, note.Owner, note.Title, note.Body)
	}

	current := arr[0]
	newNote, err := ps.GetNote(conf, current.Id, id)
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println("getNode")
	fmt.Printf("Id: %d, Owner: %d\n\"%s\"\t%s\n\n", newNote.Id, newNote.Owner, newNote.Title, newNote.Body)

}
