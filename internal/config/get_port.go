package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	defaultPort = "8080"
	minPort     = 1024
	maxPort     = 65535
)

func GetPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080", nil
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return "", err
	}
	if p < minPort || p > maxPort {
		return "", errors.New("port out of range")
	}

	return port, nil
}
