package service

import "github.com/kkmory/boostnote-exporter/pkg/entity"

type BoostNoteRepository interface {
	GetFolderByName(folderName string) (*entity.BoostNoteFolder, error)
	ImportDocByID(docID string) (*entity.BoostNoteDoc, error)
}

type LocalRepository interface {
	Export(doc *entity.BoostNoteDoc) error
	PutTimeStampFile(folderName string) error
}
