package main

import (
	"context"
	"errors"
)

type locator struct{}

func (l locator) getCoordinates1(ctx context.Context, address string) (
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

func (l locator) getCoordinates2(ctx context.Context, address string) (
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

func (l locator) validateAddress(address string) bool {
	return true
}
