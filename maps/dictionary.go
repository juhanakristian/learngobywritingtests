package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func Search(dictionary Dictionary, word string) (string, error) {
	definition, ok := dictionary[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
