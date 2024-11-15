package main

import (
	"log"
	"os"
)

// Verificar si un archivo existe (por ejemplo, los certificados TLS)
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Manejar errores de manera centralizada
func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

// Función para validar entradas (por ejemplo, comandos recibidos del cliente)
func isValidCommand(cmd string, allowedCommands []string) bool {
	for _, allowed := range allowedCommands {
		if cmd == allowed {
			return true
		}
	}
	return false
}
