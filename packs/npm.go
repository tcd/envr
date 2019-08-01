package packs

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// GetNPM returns all globally installed NPM Packages.
//
// npm ls -g is painfully slow.
// We may be able to do this faster manually by finding
// the global npm installs folder and just reading the
// project.json files in the folders there.
func GetNPM() ([]Package, error) {
	cmdOut, err := npmCmd()
	if err != nil {
		return []Package{}, err
	}

	emptyInterface := makeGenericInterface(cmdOut)
	packs := parseObjectOfKeys(emptyInterface)

	return packs, nil
}

func npmCmd() ([]byte, error) {
	cmdName := "npm"
	cmdArgs := []string{"ls", "-g", "--depth", "0", "--json"}

	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		return []byte{}, err
	}

	return cmdOut, nil
}

// Unmarshal JSON bytes using a generic interface
func makeGenericInterface(bytes []byte) interface{} {
	var x interface{}
	err := json.Unmarshal(bytes, &x)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	return x
}

// Iterate through objects inside a JSON object and wrangle them into shape
func parseObjectOfKeys(x interface{}) []Package {
	var packs []Package

	topLevelObject := x.(map[string]interface{})

	for _, nestedMap := range topLevelObject {

		for k, v := range nestedMap.(map[string]interface{}) {
			switch jsonObj := v.(type) {
			case interface{}:
				p := Package{Name: k}
				for itemKey, itemValue := range jsonObj.(map[string]interface{}) {
					switch itemKey {
					case "version":
						switch itemValue := itemValue.(type) {
						case string:
							p.Version = itemValue
						default:
							break
						}
					case "from":
						switch itemValue := itemValue.(type) {
						case string:
							p.From = itemValue
						default:
							break
						}
					case "resolved":
						switch itemValue := itemValue.(type) {
						case string:
							p.Resolved = itemValue
						default:
							break
						}
					default:
						break
					}
				}
				packs = append(packs, p)
			default:
				break
			}
		}
	}
	return packs
}
