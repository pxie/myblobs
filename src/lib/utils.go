package lib

import (
	"os/exec"
	"log"
	"os"
)

func Add(x, y int64) int64 {
	return x + y
}

func GetUUID() string {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal("fail to generate uuid")
	}
	return string(uuid)
}

func exists(path string) bool {
	_, err :=os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return false
}

func InitLogger(filepath string) {
	f, err := os.OpenFile(filepath, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
}

func LogInfo(requestID, message string) {
	log.Println("request ID: " + requestID + " " + message)
}


