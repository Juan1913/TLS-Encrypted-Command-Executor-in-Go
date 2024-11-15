package shared

import (
	"fmt"
	"strings"
)

// Mensaje genérico
type Message struct {
	Command string // El comando enviado al cliente
	Payload string // Respuesta o datos adicionales
}

// Serializar un mensaje a string
func (m *Message) Serialize() string {
	return fmt.Sprintf("%s|%s", m.Command, m.Payload)
}

// Deserializar un string a mensaje
func Deserialize(data string) *Message {
	parts := strings.SplitN(data, "|", 2)
	if len(parts) < 2 {
		return &Message{Command: data, Payload: ""}
	}
	return &Message{Command: parts[0], Payload: parts[1]}
}
