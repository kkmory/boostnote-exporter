package repository

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kkmory/boostnote-exporter/pkg/entity"
)

type LocalRepository struct {
	TargetDirectory string
}

func (l LocalRepository) Export(doc *entity.BoostNoteDoc) error {
	targetPath := fmt.Sprintf("%s/%s", doc.FolderPathname, doc.CreatedAt.Format("200601"))

	if err := createDirectoryIfNotExist(fmt.Sprintf("%s%s", l.TargetDirectory, targetPath)); err != nil {
		return fmt.Errorf("export error: %w", err)
	}

	fPath := fmt.Sprintf("%s%s/%s.md", l.TargetDirectory, targetPath, doc.Title)
	if err := writeFile(fPath, doc.Head.Content); err != nil {
		return fmt.Errorf("export error: %w", err)
	}

	return nil
}

func (l LocalRepository) PutTimeStampFile(folderName string) error {
	dtFilePath := fmt.Sprintf("%s/dt.txt", folderName)
	dt := time.Now().UTC().Format("2006-01-02 15:04:05")

	if err := writeFile(dtFilePath, dt); err != nil {
		return fmt.Errorf("timestamp error: %w", err)
	}

	return nil
}

func createDirectoryIfNotExist(folderName string) error {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.MkdirAll(folderName, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeFile(path, content string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("export error: %w", err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln("[ERROR] export error: failed to close buffer")
		}
	}(f)

	fw := bufio.NewWriter(f)
	_, err = fw.WriteString(content)
	if err != nil {
		return fmt.Errorf("export error: %w", err)
	}

	err = fw.Flush()
	if err != nil {
		return fmt.Errorf("export error: %w", err)
	}

	return nil
}
