package store

import (
	"encoding/json"
	"log"
	"os"
)

type IStore interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type FileStore struct {
	FileName string
}

func NewFileStore(fileName string) IStore {
	return &FileStore{FileName: fileName}
}

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, fileData, 0o644)
}

func (fs *FileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(file, &data); err != nil {
		return err
	}
	if &data == nil {
		log.Println("not product registered")
	}

	return json.Unmarshal(file, &data)
}
