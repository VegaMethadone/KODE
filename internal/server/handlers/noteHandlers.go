package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"kode/internal/logic"
	"kode/loger"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type JsonNote struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func newNote(w http.ResponseWriter, r *http.Request) {
	loger.Logger.Info(fmt.Sprintf("Request add new note from ip: %s", r.RemoteAddr))

	// Проверяю наличие куки у пользователя
	cookie, err := r.Cookie("kode")
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Unable to find cookie", http.StatusForbidden)
		return
	}

	// Читая тело запроса, которые содержит json
	body, err := io.ReadAll(r.Body)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Распаковываю его в структуру
	var jsNote JsonNote
	err = json.Unmarshal(body, &jsNote)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Отправляю на уровень бизне слогики и затем к бд
	err = logic.LogicNewNote(cookie.Value, jsNote.Title, jsNote.Body)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	// конец
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	ans := map[string]string{
		"info": "Note created successfully",
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

func getNote(w http.ResponseWriter, r *http.Request) {
	loger.Logger.Info(fmt.Sprintf("Request get note from ip: %s", r.RemoteAddr))

	// Проверяю наличие куки у пользователя
	cookie, err := r.Cookie("kode")
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Unable to find cookie", http.StatusForbidden)
		return
	}

	// Вытаскиваю ID записи  из строки запроса
	vars := mux.Vars(r)
	notId, err := strconv.Atoi(vars["noteId"])
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	// Отправляю на уровень бизне слогики и затем к бд
	response, err := logic.LogicGetNote(cookie.Value, notId)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Failed to retrieve note", http.StatusInternalServerError)
		return
	}

	//  конец
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
		return
	}
	loger.Logger.Info("Success")
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	loger.Logger.Info(fmt.Sprintf("Request get all notes from ip: %s", r.RemoteAddr))

	// Проверяю наличие куки у пользователя
	cookie, err := r.Cookie("kode")
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Unable to find cookie", http.StatusForbidden)
		return
	}

	// Отправляю на уровень бизне слогики и затем к бд
	response, err := logic.LogicGetNotes(cookie.Value)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Failed to retrieve notes", http.StatusInternalServerError)
		return
	}

	// конец
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		loger.Logger.Info(fmt.Sprintf("%v", err))
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
		return
	}
	loger.Logger.Info("Success")
}

/*
Подруга дней моих суровых,
Голубка дряхлая моя!
Одна в глуши лесов сосновых
Давно, давно ты ждешь меня.
Ты под окном своей светлицы
Горюешь, будто на часах,
И медлят поминутно спицы
В твоих наморщенных руках.
Глядишь в забытые вороты
На черный отдаленный путь:
Тоска, предчувствия, заботы
Теснят твою всечасно грудь.
То чудится тебе…

*/
