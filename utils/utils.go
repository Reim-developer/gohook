package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/lucasb-eyer/go-colorful"
)

func GetHexColor(colorStr string) (int, error) {

	color, err := colorful.Hex(colorStr)
	if err != nil {
		return 0, err
	}

	r := int(color.R * 255)
	g := int(color.G * 255)
	b := int(color.B * 255)

	return (r<<16 | g<<8 | b), nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func GetTimeNow() string {
	const TimeFormat string = "2006_01_02_15_04"
	var timeNow = time.Now()
	var timeStamp = timeNow.Format(TimeFormat)

	return timeStamp
}

func WriteTo(filePath string, content []byte) error {
	const FilePermission = os.FileMode(0644)

	err := os.WriteFile(filePath, content, FilePermission)
	if err != nil {
		return err
	}

	return nil
}

func IsNonEmpty(str string) bool {
	return strings.TrimSpace(str) != ""
}

func InfoShow(format string, a ...any) {
	fmt.Fprintf(os.Stdout, "[INFO] "+format+"\n", a...)
}

func CriticalShow(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "[CRITICAL] "+format+"\n", a...)
}
