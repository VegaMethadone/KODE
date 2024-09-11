package postgres

import (
	"database/sql"
	"errors"
	"kode/config"
	"kode/internal/structs/notes"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Postgres struct{}

func (p *Postgres) NewUser(conf *config.Config, login, mail, password, role string) error {
	db, err := sql.Open("postgres", getConnStr(conf))
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		`INSERT INTO users (login, mail, password, role) VALUES ($1, $2, $3, $4)`,
		login, mail, password, role,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUser(conf *config.Config, login, password string) (int, string, string, error) {
	db, err := sql.Open("postgres", getConnStr(conf))
	if err != nil {
		return 0, "", "", err
	}
	defer db.Close()

	var (
		gotId    int
		gotLogin string
		gotMail  string
		gotHash  string
	)
	err = db.QueryRow(
		`SELECT id, login, mail, password FROM users WHERE login=$1`,
		login,
	).Scan(&gotId, &gotLogin, &gotMail, &gotHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, "", "", errors.New("no such user")
		}
		return 0, "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(gotHash), []byte(password))
	if err != nil {
		return 0, "", "", errors.New("incorrect password")
	}

	return gotId, gotLogin, gotMail, nil

}

func (p *Postgres) NewNote(conf *config.Config, owner int, title, body string) error {
	db, err := sql.Open("postgres", getConnStr(conf))
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(
		`INSERT INTO notes (owner, title, body, date) VALUES ($1, $2, $3, $4)`,
		owner, title, body, time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetNotes(conf *config.Config, owner int) ([]notes.Note, error) {
	db, err := sql.Open("postgres", getConnStr(conf))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(
		`SELECT id, title, body, date FROM notes WHERE owner=$1`,
		owner,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no notes")
		}
		return nil, err
	}
	defer rows.Close()

	arr := []notes.Note{}
	for rows.Next() {
		var note notes.Note
		note.Owner = owner
		err := rows.Scan(&note.Id, &note.Title, &note.Body, &note.Date)
		if err != nil {
			return nil, err
		}

		arr = append(arr, note)
	}

	return arr, nil
}

func (p *Postgres) GetNote(conf *config.Config, idNote, owner int) (*notes.Note, error) {
	db, err := sql.Open("postgres", getConnStr(conf))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var note notes.Note

	row := db.QueryRow(
		`SELECT id, owner, title, body, date FROM notes WHERE id = $1`,
		idNote,
	)

	err = row.Scan(&note.Id, &note.Owner, &note.Title, &note.Body, &note.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("note not found")
		}
		return nil, err
	}

	if note.Owner != owner {
		return nil, errors.New("access denied")
	}

	return &note, nil
}

func GetUserId(conf *config.Config, login string) (int, error) {
	db, err := sql.Open("postgres", getConnStr(conf))
	if err != nil {
		return 0, err
	}
	defer db.Close()

	row := db.QueryRow(
		`SELECT id FROM users WHERE login = $1`,
		login,
	)

	var userId int
	err = row.Scan(&userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user not found")
		}
		return 0, err
	}

	return userId, nil
}
