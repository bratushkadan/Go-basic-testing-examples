package util

import "fmt"

type Person struct {
	FirstName   string
	LastName    string
	FathersName string
}

func (p *Person) String() string {
	if p.FathersName == "" {
		return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
	}
	return fmt.Sprintf("%s %s %s", p.FirstName, p.FathersName, p.LastName)
}
