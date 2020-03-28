package main

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	yaml :=
		`---
jobs:
  - name: 'Trigger Jenkins Tests'
    match:
      events: ['release', 'autorelease']
      services: ['*-dev']
    docker:
      image: curlimages/curl
      command: ['/usr/bin/curl']
      args: ['https://github.com']
`
	cfg, err := parseConfig([]byte(yaml))
	if err != nil {
		t.Errorf("Error while parsing config: %s", err)
	}

	if l := len(cfg.Jobs); l != 1 {
		t.Errorf("len(cfg.Jobs) != 1, got %d", l)
	}

	if l := len(cfg.Jobs[0].Match.Events); l != 2 {
		t.Errorf("len events != 2, got %d", l)
	}

	if cfg.Jobs[0].Match.Events[0] != "release" || cfg.Jobs[0].Match.Events[1] != "autorelease" {
		t.Errorf("specified events were not found.")
	}

	if cfg.Jobs[0].Match.Services[0] != "*-dev" {
		t.Errorf("specified service was not found.")
	}
}
