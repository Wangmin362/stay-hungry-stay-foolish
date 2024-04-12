package main

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
