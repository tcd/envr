package packs

import (
	"encoding/json"
	"os/exec"
)

// GetPip3 returns all packages globally installed with pip3.
func GetPip3() ([]Package, error) {
	var packs []Package

	cmdName := "pip3"
	cmdArgs := []string{"list", "--form", "json"}

	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return packs, err
	}

	json.Unmarshal(cmdOut, &packs)

	return packs, nil
}

// GetPip2 returns all packages globally installed with pip2.
func GetPip2() ([]Package, error) {
	var packs []Package

	cmdName := "pip2"
	cmdArgs := []string{"list", "--form", "json"}

	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return packs, err
	}

	json.Unmarshal(cmdOut, &packs)

	return packs, nil
}
