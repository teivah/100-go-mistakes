package main

import (
	"fmt"
	"log"
)

type Route struct{}

func GetRoute1(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates1(srcLat, srcLng)
	if err != nil {
		log.Println("failed to validate source coordinates")
		return Route{}, err
	}

	err = validateCoordinates1(dstLat, dstLng)
	if err != nil {
		log.Println("failed to validate target coordinates")
		return Route{}, err
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates1(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		log.Printf("invalid latitude: %f", lat)
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		log.Printf("invalid longitude: %f", lng)
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}

func GetRoute2(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{}, err
	}

	err = validateCoordinates2(dstLat, dstLng)
	if err != nil {
		return Route{}, err
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates2(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}

func GetRoute3(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate source coordinates: %w", err)
	}

	err = validateCoordinates2(dstLat, dstLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate target coordinates: %w", err)
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func getRoute(lat, lng, lat2, lng2 float32) (Route, error) {
	return Route{}, nil
}
