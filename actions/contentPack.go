package actions

import "fmt"

// ContentPacks represents contnet packs in OO
type ContentPacks []struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// HandleResponse - print content packs
func (cps ContentPacks) HandleResponse() {
	for _, cp := range cps {
		fmt.Printf("%v, %v\n", cp.Name, cp.Version)
	}
}
