package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
	"rat-advanced/shared"
)

func handleClientCommand(command string) string {
	allowedCommands := []string{"ping", "ls", "whoami"}
	if !isValidCommand(command, allowedCommands) {
		return "Comando no válido"
	}

	// Procesar comando válido
	switch command {
	case "ping":
		return "pong"
	case "ls":
		return "Listado de archivos no disponible aún"
	case "whoami":
		return "Eres el cliente conectado"
	default:
		return "Comando no implementado"
	}
}

func isValidCommand(cmd string, allowedCommands []string) bool {
	for _, allowed := range allowedCommands {
		if cmd == allowed {
			return true
		}
	}
	return false
}

func handleCommands(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		command := scanner.Text()
		log.Printf("Comando recibido: %s", command)

		// Ejecutar el comando y enviar la respuesta
		response := handleClientCommand(command)
		_, err := conn.Write([]byte(response + "\n"))
		if err != nil {
			log.Printf("Error enviando respuesta: %v", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error leyendo desde conexión: %v", err)
	}
}

func main() {
	// Mostrar banner
	shared.PrintBanner("cliente")

	// Conectar al servidor
	config := &tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", "127.0.0.1:8443", config)
	if err != nil {
		log.Fatalf("Error conectando al servidor: %v", err)
	}
	defer conn.Close()

	log.Println("Conectado al servidor")

	// Manejar comandos
	handleCommands(conn)
}
