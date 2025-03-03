package logger

import (
	"fmt"
	"os"
	"time"
)

func Log(Case int, str string, errInput error) {
	file, err := os.OpenFile("./log/server.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o644)
	if err != nil {
		fmt.Println("Error :Log section : Error oppening file", err)
		return
	}
	logMessage := ""
	Time := time.Now().Format("2006-01-02 15:04:05")
	// Info msgs
	if Case == 1 {
		logMessage = fmt.Sprintf("Info: %s %s", Time, str)
		// Err msgs
	} else if Case == 2 {
		msg := fmt.Sprintf("%s\n", errInput)
		logMessage = fmt.Sprintf("Error: %s %s", Time, msg)
	}
	_, err = file.WriteString(logMessage)
	if err != nil {
		fmt.Println("Error :Log section : Error writing to file", err)
		return
	}
	defer file.Close()
}
