package database

import (
	"kode/config"
	"kode/internal/database/postgres"
	"kode/internal/structs/notes"
)

// Реализую паттерн состояние для того, чтобы можно было расширять приложение и подключать другие бд, если надо

type State interface {
	NewUser(*config.Config, string, string, string, string) error
	GetUser(*config.Config, string, string) (int, string, string, error)
	NewNote(*config.Config, int, string, string) error
	GetNote(*config.Config, int, int) (*notes.Note, error)
	GetNotes(*config.Config, int) ([]notes.Note, error)
}

type Database struct {
	postgres State
	//  mongo State
	//  redis State

	current State
}

func NewDataBase() *Database {
	ps := postgres.NewPostgres()
	newDataBase := &Database{
		postgres: ps,
		// mongo: mg
		// redis: rd
		current: ps,
	}

	return newDataBase
}

func (db *Database) SetState(state State) { db.current = state }

func (db *Database) NewUser(conf *config.Config, login, mail, password, role string) error {
	return db.current.NewUser(conf, login, mail, password, role)
}
func (db *Database) GetUser(conf *config.Config, login, password string) (int, string, string, error) {
	return db.current.GetUser(conf, login, password)
}
func (db *Database) NewNote(conf *config.Config, owner int, title, body string) error {
	return db.current.NewNote(conf, owner, title, body)
}
func (db *Database) GetNotes(conf *config.Config, owner int) ([]notes.Note, error) {
	return db.current.GetNotes(conf, owner)
}
func (db *Database) GetNote(conf *config.Config, idNote, owner int) (*notes.Note, error) {
	return db.current.GetNote(conf, idNote, owner)
}
