package main

import (
	"fmt"
	"kode/config"
	"kode/internal/server"
	"kode/loger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Создаю логер и проверяю на ошибку
	if err := loger.NewLogger(); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	// Вытаскиваю конфиг и передаю дальше для создания сервера
	conf := config.NewConfig()
	srv := server.NewServer(conf)

	loger.Logger.Info("Server is working at: http://127.0.0.1:8080")
	fmt.Println("Server is working at: http://127.0.0.1:8080")

	// Создаю два канала для принятия сигнала с клавы и ошибки
	done := make(chan os.Signal, 1)
	waitingForErr := make(chan bool, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			loger.Logger.Info(fmt.Sprintf("%v", err))
			waitingForErr <- true
		}
	}()

	// Жду прихода ошибки или сигнала в заблоченной горутины
	select {
	case s := <-done:
		loger.Logger.Info(fmt.Sprintf("Server is stoped by host: %v", s))
		srv.Close()
		waitingForErr <- true
	case err := <-waitingForErr:
		loger.Logger.Info(fmt.Sprintf("Error occured: %v", err))
	}

}
