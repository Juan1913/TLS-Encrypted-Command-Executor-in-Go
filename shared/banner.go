package shared

import "fmt"

func PrintBanner(role string) {
	if role == "server" {
		fmt.Println(`
	███████╗ █████╗ ████████╗
	██╔════╝██╔══██╗╚══██╔══╝
	███████╗███████║   ██║   
	╚════██║██╔══██║   ██║   
	███████║██║  ██║   ██║   
	╚══════╝╚═╝  ╚═╝   ╚═╝   
	Servidor Avanzado
	`)
	} else if role == "client" {
		fmt.Println(`
	 █████╗  █████╗ ███╗   ██╗
	██╔══██╗██╔══██╗████╗  ██║
	███████║███████║██╔██╗ ██║
	██╔══██║██╔══██║██║╚██╗██║
	██║  ██║██║  ██║██║ ╚████║
	╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝
	Cliente Avanzado
	`)
	}
}
