package packs

import (
	"os/exec"
	"strings"
)

// GetCargo returns all packages globally with Cargo
func GetCargo() ([]Package, error) {
	var packs []Package

	cmdName := "cargo"
	cmdArgs := []string{"install", "--list"}

	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return packs, err
	}

	lines := strings.Split(string(cmdOut), "\n")
	for _, line := range lines {
		if len(line) > 0 && line[len(line)-1:] == ":" {
			fields := strings.Fields(line[:len(line)-1])
			pack := Package{
				Name:    fields[0],
				Version: fields[1],
			}
			packs = append(packs, pack)
		}
	}

	return packs, nil
}
