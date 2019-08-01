package packs

import (
	"os/exec"
	"strings"
)

// GetGem returns all globally installed gems.
//
// This is tricky, because multiple versions of
// the same gem can be installed.
// We're currently just adding all of them to
// the versions field. That'll probably change.
func GetGem() ([]Package, error) {
	var packs []Package

	cmdName := "gem"
	cmdArgs := []string{"list"}

	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return packs, err
	}

	lines := strings.Split(string(cmdOut), "\n")
	for _, line := range lines {
		dat := strings.SplitAfterN(line, " ", 2)
		if len(line) > 3 {
			pack := Package{
				Name:    dat[0],
				Version: dat[1],
			}
			packs = append(packs, pack)
		}
	}
	return packs, nil
}
