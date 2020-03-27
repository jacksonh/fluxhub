package main

import (
	"gopkg.in/yaml.v2"
)

type configSpec struct {
	Jobs []jobSpec
}

type jobSpec struct {
	Name   string
	Match  matchSpec
	Docker dockerSpec
}

type dockerSpec struct {
	Image   string
	Command []string
	Args    []string
}

type matchSpec struct {
	Events    []string
	Service   string
	Namespace string
}

func parseConfig(data []byte) (*configSpec, error) {
	config := configSpec{}

	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
