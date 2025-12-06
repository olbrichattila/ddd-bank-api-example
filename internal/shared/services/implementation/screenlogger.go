// This logger is only for example, logging to the screen, stdout
package implementation

import (
	"eaglebank/internal/shared/services"
	"fmt"
)

func NewScreenLogger() services.Logger {
	return &logger{}
}

type logger struct {
}

func (l *logger) Debug(msg string) {
	fmt.Printf("DEBUG: %s\n", msg)
}

func (l *logger) Error(msg string) {
	fmt.Printf("ERROR: %s\n", msg)
}

func (l *logger) Info(msg string) {
	fmt.Printf("INFO: %s\n", msg)
}
