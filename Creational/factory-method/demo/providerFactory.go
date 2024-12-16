package main

import "errors"

func getProvider(providerType string) (iProvider, error) {
	if providerType == "mailjet" {
		return newMailjet(), nil
	}

	if providerType == "ses" {
		return newSes(), nil
	}

	return nil, errors.New("provider type not recognized")
}
