package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Bibles struct {
	Bibles []Bible `json:"data"`
}

type Bible struct {
	ID                string        `json:"id"`
	DBlID             string        `json:"dblid"`
	Abbreviation      string        `json:"abbreviation"`
	AbbreviationLocal string        `json:"abbreviationLocal"`
	Copyright         string        `json:"copyright"`
	Language          Language      `json:"language"`
	Countries         []Countries   `json:"countries"`
	Name              string        `json:"name"`
	NameLocal         string        `json:"nameLocal"`
	Description       string        `json:"description"`
	DescriptionLocal  string        `json:"descriptionLocal"`
	Info              string        `json:"info"`
	Type              string        `json:"type"`
	UpdatedAt         string        `json:"updatedAt"`
	RelatedDbl        string        `json:"relatedDbl"`
	AudioBibles       []AudioBibles `json:"audioBibles"`
}

type Language struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	NameLocal         string `json:"nameLocal"`
	Script            string `json:"script"`
	ScriptDescription string `json:"scriptDescription"`
}

type Countries struct {
	Countries []Country `json:"countries"`
}

type Country struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	NameLocal string `json:"nameLocal"`
}

type AudioBibles struct {
	audioBibles []AudioBible `json:"audioBibles"`
}

type AudioBible struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	NameLocal        string `json:"nameLocal"`
	Description      string `json:"description"`
	DescriptionLocal string `json:"descriptionLocal"`
}

var authorize = "57e13daca0fee16ab8ae2bdbbca58b2d"
var FetchUrl = "https://api.scripture.api.bible/v1/"

func GetBibles() (Bibles, error) {
	var bibles Bibles
	url := fmt.Sprintf("%sbibles", FetchUrl)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("api-key", authorize)
	res, err := client.Do(req)
	if err != nil {
		return bibles, err
	}

	defer res.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	bufbytes := buf.Bytes()
	err = json.Unmarshal(bufbytes, &bibles)
	if err != nil {
		return bibles, err
	}

	return bibles, nil

}

func GetBooks(ID string) ([]string, error) {
	var books []string

	url :=
}
