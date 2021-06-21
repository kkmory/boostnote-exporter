package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kkmory/boostnote-exporter/pkg/entity"
)

type BoostnoteRepository struct {
	Client *http.Client
	Token  string
}

const folderEndpoint = "https://boostnote.io/api/folders"
const docEndpoint = "https://boostnote.io/api/docs"

type folderResponse struct {
	Folders []*entity.BoostNoteFolder `json:"folders"`
}

type docResponse struct {
	Doc *entity.BoostNoteDoc `json:"doc"`
}

func (b BoostnoteRepository) GetFolderByName(folderName string) (*entity.BoostNoteFolder, error) {
	req, _ := http.NewRequest("GET", folderEndpoint, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", b.Token))

	resp, err := b.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln("boost note api error: failed to close buffer")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("boostnote api error: %v", err)
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("boostnote api error: %v", err)
	}

	var folders *folderResponse
	if err := json.Unmarshal(res, &folders); err != nil {
		return nil, fmt.Errorf("boostnote api error: %v", err)
	}

	for _, f := range folders.Folders {
		if f.Name == folderName {
			return f, nil
		}
	}

	return nil, fmt.Errorf("boostnote api error: cannot find specified folder")
}

func (b BoostnoteRepository) ImportDocByID(docID string) (*entity.BoostNoteDoc, error) {
	url := fmt.Sprintf("%s/%s", docEndpoint, docID)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", b.Token))

	resp, err := b.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln("[ERROR] boost note api error: failed to close buffer")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("boostnote api error: %v", err)
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("boostnote api error: %v", err)
	}

	var doc *docResponse
	if err := json.Unmarshal(res, &doc); err != nil {
		return nil, fmt.Errorf("boostnote api error: %v", err)
	}

	return doc.Doc, nil
}
