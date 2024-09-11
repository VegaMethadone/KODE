package logic

import (
	"encoding/json"
	"fmt"
	"kode/config"
	"kode/internal/database"
	"kode/internal/token"
	"kode/internal/yandex"
	"kode/loger"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var db = database.NewDataBase()
var conf = config.NewConfig()

func LogigNewUser(login, mail, password string) error {

	loger.Logger.Info("Encrypt password")
	encodedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	loger.Logger.Info("Add new user to database")
	err = db.NewUser(conf, login, mail, string(encodedPassword), "user")
	if err != nil {
		return err
	}

	return nil

}

func LogicGetUser(login, password string) (string, error) {

	loger.Logger.Info("Check user in database")
	id, login, mail, err := db.GetUser(conf, login, password)
	if err != nil {
		return "", err
	}

	var init time.Time = time.Now()
	var expire = init.Add(time.Hour * 24)

	loger.Logger.Info("Generate JWT Token")
	tokeString, err := token.NewJwtToken(id, login, mail, init, expire)
	if err != nil {
		return "", err
	}

	return tokeString, nil

}

func LogicNewNote(tokenString, title, body string) error {

	loger.Logger.Info("Validate JWT Token")
	userId, err := token.ValidateTokenAndId(tokenString)
	if err != nil {
		return err
	}

	loger.Logger.Info("Validate text")
	newbody, err := yandex.ValidateBody(body)
	if err != nil {
		return err
	}

	loger.Logger.Info(fmt.Sprintf("Add new note in database by user %d", userId))
	err = db.NewNote(conf, userId, title, string(newbody))
	if err != nil {
		return err
	}

	return nil
}

func LogicGetNote(tokenString string, noteId int) ([]byte, error) {

	loger.Logger.Info("Validate JWT Token")
	userId, err := token.ValidateTokenAndId(tokenString)
	if err != nil {
		return nil, err
	}

	loger.Logger.Info(fmt.Sprintf("Get note from database by user %d", userId))
	note, err := db.GetNote(conf, noteId, userId)
	if err != nil {
		return nil, err
	}

	loger.Logger.Info("Convert data to json")
	jsonData, err := json.Marshal(note)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func LogicGetNotes(tokenString string) ([]byte, error) {

	loger.Logger.Info("Validate JWT Token")
	userId, err := token.ValidateTokenAndId(tokenString)
	if err != nil {
		return nil, err
	}

	loger.Logger.Info(fmt.Sprintf("Get all notes from database by user %d", userId))
	arr, err := db.GetNotes(conf, userId)
	if err != nil {
		return nil, err
	}

	loger.Logger.Info("Convert data to json")
	jsonData, err := json.Marshal(arr)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
