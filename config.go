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
	Workloads []string
}

func parseConfig(data []byte) (config configSpec, err error) {
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
