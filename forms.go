package main

import (
	"errors"
	"strconv"
)

func validateArrayIndexFromForm(aiStr string, alen int) (int, error) {
	if aiStr == "" {
		return -1, errors.New("GroupId is required")
	}
	ai, err := strconv.Atoi(aiStr)
	if err != nil {
		return -1, errors.New("Index is required")
	}
	if ai < 0 || ai >= alen {
		return -1, errors.New("Index out of Bounds")
	}
	return ai, nil
}
