package util

import gonanoid "github.com/matoous/go-nanoid/v2"

func GenerateID() string {
	id, _ := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", 10)
	return id
}
