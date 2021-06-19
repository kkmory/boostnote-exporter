package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	folder := flag.String("folder", "", "folder name for export [required]")
	dist := flag.String("dist", "./", "target directory for export [optional]")
	dt := flag.String("dt", "19700101", "export pages if updated_at > dt [optional]")
	flag.Parse()

	if len(*folder) == 0 {
		log.Fatalf("[ERROR] faild to run action. please set --folder option.")
	}

	log.Println("[INFO] starting export action...")

	log.Println(*folder)
	log.Println(*dist)
	log.Println(*dt)

	log.Println("[INFO] successfully finished")

	os.Exit(0)
}
