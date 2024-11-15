package main

import (
	"fmt"
	"runtime"
)

// Obtener información básica del sistema operativo
func getSystemInfo() string {
	return fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
}
