package main

import "encoding/json"

func listing1() error {
	b := getMessage()
	var m map[string]any
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	return nil
}

func getMessage() []byte {
	return nil
}
