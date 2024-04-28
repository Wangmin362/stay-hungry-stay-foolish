package vnote

import (
	"encoding/json"
	"os"
)

type Meta struct {
	CreatedTime  string    `json:"created_time"`
	Files        []*File   `json:"files"`
	Folders      []*Folder `json:"folders"`
	Id           string    `json:"id"`
	ModifiedTime string    `json:"modified_time"`
	Signature    string    `json:"signature"`
	Version      int       `json:"version"`
}

type File struct {
	AttachmentFolder string   `json:"attachment_folder"`
	CreatedTime      string   `json:"created_time"`
	Id               string   `json:"id"`
	ModifiedTime     string   `json:"modified_time"`
	Name             string   `json:"name"`
	Signature        string   `json:"signature"`
	Tags             []string `json:"tags"`
}

type Folder struct {
	Name string `json:"name"`
}

func UnMarshal(path string) (*Meta, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var meta Meta
	if err = json.Unmarshal(file, &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

func (r *Meta) PersistentMarshal(path string) error {
	marshal, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return err
	}

	if err = os.WriteFile(path, marshal, os.ModePerm); err != nil {
		return err
	}

	return nil
}
