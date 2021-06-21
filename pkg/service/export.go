package service

import (
	"fmt"
	"time"

	"github.com/kkmory/boostnote-exporter/pkg/entity"
)

type ExportService struct {
	Boostnote BoostNoteRepository
	Local     LocalRepository
}

func (s ExportService) Export(folder, dist, dt string) (int64, error) {
	pv, err := time.Parse("2006-01-02 15:04:05", dt)
	config := &entity.AppConfig{
		FolderName:          folder,
		BasePath:            dist,
		PreviousProcessedAt: pv,
	}

	var count int64
	f, err := s.Boostnote.GetFolderByName(config.FolderName)
	if err != nil {
		return 0, err
	}

	for _, d := range f.ChildDocsIDs {
		doc, err := s.Boostnote.ImportDocByID(d)
		if err != nil {
			return 0, err
		}
		err = s.Local.Export(doc)
		count++
	}

	if err := s.Local.PutTimeStampFile(fmt.Sprintf("%s/%s", config.BasePath, config.FolderName)); err != nil {
		return 0, err
	}

	return count, nil
}
