package custom_errors

import "fmt"

type NotSetError struct {
	ConfigName string
	Global     bool
}

func (e *NotSetError) Error() string {
	var configType string

	if e.Global {
		configType = "global"
	} else {
		configType = "local"
	}

	return fmt.Sprintf("no %s %s set", configType, e.ConfigName)
}
