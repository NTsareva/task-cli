package models

import "fmt"

func (t Task) Validate() error {
	if !t.Status.IsValid() {
		return fmt.Errorf("invalid task status: %s", t.Status)
	}
	return nil
}
