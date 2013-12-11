package lib

import (
	"testing"
	"fmt"
	"os/exec"
	"log"
	"os"
	"bufio"
)

func TestAdd(t *testing.T) {
	r := Add(4, 90)
	if r != 94 {
		t.Errorf("Add(4, 90) failed. Got %d, expected 94.", r)
	}
}

func TestGetUUID(t *testing.T) {
	uuid := GetUUID()
	fmt.Println("UUID = ", uuid)
	if uuid == "" {
		t.Errorf("Cannot generate UUID")
	}
}

func TestExists(t *testing.T) {
	path := "/tmp/test.file"
	_, err := exec.Command("touch", path).Output()
	if err != nil {
		log.Fatal("fail to touch one file")
	}

	if exists(path) != true {
		t.Errorf("fail to check file exist.")
	}

	defer os.Remove(path)

}

func TestLogger(t *testing.T) {
	path := "/tmp/test.file"
	InitLogger(path)
	uuid := GetUUID()
	LogInfo(uuid, "log to file")

}

