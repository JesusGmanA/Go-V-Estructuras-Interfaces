package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"./multimedia"
)

const OPC_AGREGAR_IMAGEN = 1
const OPC_AGREGAR_AUDIO = 2
const OPC_AGREGAR_VIDEO = 3
const MOSTRAR = 4
const EXIT = 5

func main() {
	var opc = 0
	cm := multimedia.ContenidoMultimedia{Multimedias: []multimedia.Multimedia{}}
	scanner := bufio.NewReader(os.Stdin)
	for opc != EXIT {
		limpiarPantalla()
		opc = obtenerOpcMenu(&cm, scanner)
	}

	fmt.Print("Hasta luego!")
}

func limpiarPantalla() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func obtenerOpcMenu(cm *multimedia.ContenidoMultimedia, s *bufio.Reader) int {
	var option int
	fmt.Println("Escribe la accion a realizar:")
	fmt.Println("1. Agregar imagen nueva")
	fmt.Println("2. Agregar nuevo audio")
	fmt.Println("3. Agregar un video nuevo")
	fmt.Println("4. Mostrar contenido multimedia")
	fmt.Println("5. Salir")
	fmt.Print("Opcion: ")
	fmt.Scanln(&option)

	switch option {
	case OPC_AGREGAR_IMAGEN:
		var titulo, formato string
		var canales int64
		fmt.Print("Dame el titulo de la imagen: ")
		titulo, _ = s.ReadString('\n')              //Para poder leer lineas con espacios
		titulo = strings.TrimSuffix(titulo, "\r\n") //No del todo necesario, solo para el formato de mostrar
		fmt.Print("Dame el formato de la imagen: ")
		fmt.Scanln(&formato)
		fmt.Print("Dame los canales de la imagen: ")
		fmt.Scanf("%d", &canales)
		img := multimedia.Imagen{Titulo: titulo, Formato: formato, Canales: canales}
		cm.AgregarMultimedia(&img)

	case OPC_AGREGAR_AUDIO:
		var titulo, formato string
		var duracion int64
		fmt.Print("Dame el titulo del audio: ")
		titulo, _ = s.ReadString('\n')              //Para poder leer lineas con espacios
		titulo = strings.TrimSuffix(titulo, "\r\n") //No del todo necesario, solo para el formato de mostrar
		fmt.Print("Dame el formato del audio: ")
		fmt.Scan(&formato)
		fmt.Print("Dame la duracion del audio en segundos: ")
		fmt.Scan(&duracion)
		a := multimedia.Audio{Titulo: titulo, Formato: formato, Duracion: duracion}
		cm.AgregarMultimedia(&a)

	case OPC_AGREGAR_VIDEO:
		var titulo, formato string
		var frames int64
		fmt.Print("Dame el titulo del video: ")
		titulo, _ = s.ReadString('\n')              //Para poder leer lineas con espacios
		titulo = strings.TrimSuffix(titulo, "\r\n") //No del todo necesario, solo para el formato de mostrar
		fmt.Print("Dame el formato del video: ")
		fmt.Scan(&formato)
		fmt.Print("Dame los frames que tiene el video: ")
		fmt.Scan(&frames)
		v := multimedia.Video{Titulo: titulo, Formato: formato, Frames: frames}
		cm.AgregarMultimedia(&v)

	case MOSTRAR:
		fmt.Print(cm.Mostrar())

	case EXIT:
	default:
		fmt.Println("Opcion inexistente intente de nuevo!")
	}

	if option != EXIT {
		fmt.Print("Presiona 'Enter' para continuar...")
		fmt.Scanln() //El primero se come el "Enter" atorado en el buffer cuando se lee algo desde consola
		fmt.Scanln()
	}
	return option
}
