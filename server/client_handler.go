package main

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
		return "cliente conectado"
	default:
		return "Comando no implementado"
	}
}
