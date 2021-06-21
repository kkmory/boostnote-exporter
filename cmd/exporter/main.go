package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/kkmory/boostnote-exporter/pkg/repository"
	"github.com/kkmory/boostnote-exporter/pkg/service"
)

func main() {
	folder := flag.String("folder", "", "folder name for export [required]")
	dist := flag.String("dist", ".", "target directory for export [optional]")
	token := flag.String("token", "", "target directory for export [optional]")
	dt := flag.String("dt", "1970-01-01 00:00:00", "export pages if updated_at > dt [optional]")
	flag.Parse()

	if len(*folder) == 0 || len(*token) == 0 {
		log.Fatalf("[ERROR] faild to run action. please set --folder option.")
	}

	log.Println("[INFO] starting export action...")

	httpClient := new(http.Client)
	defer httpClient.CloseIdleConnections()

	boostNoteRepository := repository.BoostnoteRepository{Client: httpClient, Token: *token}
	localRepository := repository.LocalRepository{TargetDirectory: *dist}

	s := service.ExportService{
		Boostnote: boostNoteRepository,
		Local:     localRepository,
	}

	count, err := s.Export(*folder, *dist, *dt)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}

	log.Printf("[INFO] %d records successfully exported\n", count)

	os.Exit(0)
}
