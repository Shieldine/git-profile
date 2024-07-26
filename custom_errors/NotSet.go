package custom_errors

import "fmt"

type NotSetError struct {
	ConfigName string
}

func (e *NotSetError) Error() string {
	return fmt.Sprintf("no local %s set", e.ConfigName)
}
