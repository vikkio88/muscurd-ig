package enums

type AppState uint8

const (
	Login AppState = iota
	List
	Details
	Quit
)

func (a AppState) String() string {
	switch a {
	case Login:
		return "LOGIN"
	case List:
		return "LIST"
	case Details:
		return "DETAILS"
	case Quit:
		return "QUIT"
	}

	return "INVALID_STATE"
}
