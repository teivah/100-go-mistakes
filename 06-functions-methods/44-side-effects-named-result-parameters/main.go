package main

import (
	"context"
	"errors"
)

type loc struct{}

func (l loc) getCoordinates1(ctx context.Context, address string) (
	lat, lng float32, err error) {
	isValid := l.validateAddress(address)
	if !isValid {
		return 0, 0, errors.New("invalid address")
	}

	if ctx.Err() != nil {
		return 0, 0, err
	}

	// Get and return coordinates
	return 0, 0, nil
}

func (l loc) getCoordinates2(ctx context.Context, address string) (
	lat, lng float32, err error) {
	isValid := l.validateAddress(address)
	if !isValid {
		return 0, 0, errors.New("invalid address")
	}

	if err := ctx.Err(); err != nil {
		return 0, 0, err
	}

	// Get and return coordinates
	return 0, 0, nil
}

func (l loc) getCoordinates3(ctx context.Context, address string) (
	lat, lng float32, err error) {
	isValid := l.validateAddress(address)
	if !isValid {
		return 0, 0, errors.New("invalid address")
	}

	if err = ctx.Err(); err != nil {
		return
	}

	// Get and return coordinates
	return 0, 0, nil
}

func (l loc) validateAddress(address string) bool {
	return true
}
