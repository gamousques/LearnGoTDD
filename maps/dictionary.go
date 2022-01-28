package main

import "errors"

var ErrWordNotFound = errors.New("sorry, could not found the word you requested")

type Dictionary map[string]string

func(d Dictionary) Search( searchText string) (string, error) {
	definition, ok := d[searchText]

	if !ok {
		return "", ErrWordNotFound
	}

	return definition,nil
}