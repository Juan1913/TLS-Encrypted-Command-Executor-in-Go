package main

import (
	"bufio"
	"fmt"
	"net"
	"rat-advanced/shared"
)

// Manejar comandos y respuestas
func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	// Leer información inicial del cliente
	if scanner.Scan() {
		clientInfo := scanner.Text()
		fmt.Printf("Cliente conectado: %s\n", clientInfo)
	}

	// Interactuar con el cliente
	for scanner.Scan() {
		// Leer mensaje del cliente
		input := scanner.Text()
		message := shared.Deserialize(input)

		// Procesar comando y enviar respuesta
		if message.Command == "ping" {
			response := &shared.Message{
				Command: "pong",
				Payload: "¡Conexión exitosa!",
			}
			conn.Write([]byte(response.Serialize() + "\n"))
		}
	}
}
