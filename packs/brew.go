package packs

import (
	"os/exec"
	"strings"
)

// GetBrew returns all software manually installed by Homebrew (not dependencies)
func GetBrew() ([]Package, error) {
	var packs []Package

	versions, _ := getBrewVersions()

	cmdName := "brew"
	cmdArgs := []string{"leaves"}

	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return packs, err
	}

	lines := strings.Split(string(cmdOut), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(line) > 3 {
			name := strings.TrimSpace(fields[0])
			pack := Package{
				Name:    name,
				Version: versions[name],
			}
			packs = append(packs, pack)
		}
	}
	return packs, nil
}

// getBrewVersions returns a map of key-value strings where the
// key is the package name and the value is the version
func getBrewVersions() (map[string]string, error) {
	packs := make(map[string]string)

	cmdName := "brew"
	cmdArgs := []string{"list", "--versions"}

	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return packs, err
	}

	lines := strings.Split(string(cmdOut), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		// TODO: Handle multiple versions
		if len(line) > 3 {
			packs[strings.TrimSpace(fields[0])] = strings.TrimSpace(fields[1])
		}
	}
	return packs, nil
}

// GetBrewCask returns all applications currently installed by Homebrew Cask.
func GetBrewCask() ([]Package, error) {
	var packs []Package

	cmdName := "brew"
	cmdArgs := []string{"cask", "list", "--versions"}

	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return packs, err
	}

	packLines := strings.Split(string(cmdOut), "\n")
	for _, line := range packLines {
		fields := strings.Fields(line)
		if len(line) > 3 {
			pack := Package{
				Name:    strings.TrimSpace(fields[0]),
				Version: strings.TrimSpace(fields[1]),
			}
			packs = append(packs, pack)
		}
	}
	return packs, nil
}
