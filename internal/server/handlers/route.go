package handlers

import "github.com/gorilla/mux"

func GetRoutes() *mux.Router {
	r := mux.NewRouter()

	// проверка сервера на готовность
	r.HandleFunc("/api/ping", pingServer).Methods("GET")

	// регистрация и аунтификация
	r.HandleFunc("/api/user/register", newUser).Methods("POST")
	r.HandleFunc("/api/user/login", getUser).Methods("GET")

	// получить пост и пост  по ID
	r.HandleFunc("/api/user/notes", newNote).Methods("POST")
	r.HandleFunc("/api/user/notes/{noteId}", getNote).Methods("GET")

	// получить все посты
	r.HandleFunc("/api/user/notes", getNotes).Methods("GET")

	return r
}
