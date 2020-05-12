package submitfile

import "fmt"

type Environment int

const (
	Undef = iota
	Linux
	Android
	WindowsSeven32
	WindowsSeven64
)

func (e Environment) String() (string, error) {
	switch e {
	case Linux:
		return "300", nil
	case Android:
		return "200", nil
	case WindowsSeven32:
		return "100", nil
	case WindowsSeven64:
		return "120", nil
	default:
		return "", fmt.Errorf("unknown environment")
	}
}
