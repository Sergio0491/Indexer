package main

import (
	"indexer/controllers"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	if len(os.Args) < 2 {
		log.Fatal("Uso: El segundo argumento debe ser el nombre del directorio dentro de la carpeta files")
	}

	directory := "./files/" + os.Args[1]
	controllers.IndexEmails(directory)
}
