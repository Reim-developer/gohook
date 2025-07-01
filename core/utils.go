package core

import (
	"fmt"
	"os"
	"strings"

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

func IsNonEmpty(str string) bool {
	return strings.TrimSpace(str) != ""
}

func InfoShow(format string, a ...any) {
	fmt.Fprintf(os.Stdout, "[INFO] "+format+"\n", a...)
}

func CriticalShow(format string, a ...any) {
	fmt.Fprintf(os.Stderr, "[CRITICAL] "+format+"\n", a...)
}
