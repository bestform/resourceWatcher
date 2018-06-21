package main

import (
	"crypto/sha256"
	"encoding/gob"
	"io"
)

var db map[string]string

var dbReadWriter io.ReadWriter

func initDb(resource io.ReadWriter) error {
	dbReadWriter = resource
	db = make(map[string]string)
	return readDb()
}

func HasContentForResource(name string) bool {
	_, exists := db[name]
	return exists
}

func SetContentForResource(name string, content string) {
	db[name] = hash(content)
}

func HasContentChanged(name string, content string) bool {
	if _, exists := db[name]; !exists {
		return false
	}

	return db[name] != hash(content)
}

func hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	return string(h.Sum(nil))
}

func persistDb() error {
	e := gob.NewEncoder(dbReadWriter)

	err := e.Encode(db)
	if err != nil {
		return err
	}

	return nil
}

func readDb() error {
	d := gob.NewDecoder(dbReadWriter)

	err := d.Decode(&db)
	if err != nil && err != io.EOF {
		return err
	}

	return nil
}
