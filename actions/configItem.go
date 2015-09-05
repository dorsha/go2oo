package actions

import "fmt"

// ConfigItem represents a configuration item in OO (key, value)
type ConfigItem struct {
	Props map[string]interface{}
}

// HandleResponse - print configuration items
func (ci ConfigItem) HandleResponse() {
	for k, v := range ci.Props {
		fmt.Printf("%v\t\t%v\n", k, v)
	}
}
