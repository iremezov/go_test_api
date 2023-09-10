package helpers

import (
	"log"
	"os"
)

func Log_msg(method string, msg string) int {
	file, err := os.OpenFile("logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.Println(method + " " + msg)
	return 1
}
