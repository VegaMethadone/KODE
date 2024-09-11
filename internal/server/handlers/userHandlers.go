package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"kode/internal/logic"
	"kode/internal/structs/user"
	"kode/loger"
	"net/http"
	"time"
)

func pingServer(w http.ResponseWriter, r *http.Request) {
	loger.Logger.Info(fmt.Sprintf("Request ping from ip: %s", r.RemoteAddr))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
	loger.Logger.Info("Success")
}

func newUser(w http.ResponseWriter, r *http.Request) {
	loger.Logger.Info(fmt.Sprintf("Request add new user from ip: %s", r.RemoteAddr))

	// Читаю тело запроса в json формате
	body, err := io.ReadAll(r.Body)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Складывю из  json в структуру
	var user user.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// бизнес логика
	err = logic.LogigNewUser(user.Login, user.Mail, user.Password)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// конец
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	ans := map[string]string{
		"info": "User created successfully",
	}
	response, err := json.Marshal(ans)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	loger.Logger.Info("Success")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	loger.Logger.Info(fmt.Sprintf("Request get user from ip: %s", r.RemoteAddr))

	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")

	// проверяю наличие логина и пароля в строке  запроса
	if login == "" || password == "" {
		loger.Logger.Info("No login or password")
		http.Error(w, "Missing login or password", http.StatusBadRequest)
		return
	}

	// бизнес логика
	token, err := logic.LogicGetUser(login, password)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "No such user", http.StatusBadRequest)
		return
	}

	// cookies
	expiration := time.Now().Add(time.Hour * 24)
	cookie := &http.Cookie{
		Name:     "kode",
		Value:    token,
		Path:     "/",
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	// конец
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	ans := map[string]string{
		"info": "User authorized successfully",
	}
	response, err := json.Marshal(ans)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
		return
	}
	w.Write(response)
	loger.Logger.Info("Success")
}
