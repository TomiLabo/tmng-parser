package category

// Code category code
type Code string

func (i Code) String() string {
	switch i {
	case Header:
		return "Header"
	case Title:
		return "Title"
	case SubTitle:
		return "SubTitle"
	case ListItem:
		return "ListItem"
	case Empty:
		return "Empty"
	case Error:
		return "Error"
	case PlaneText:
		return "PlaneText"
	default:
		panic("Unexpected category code")
	}
}

const (
	Header    Code = "Header"
	Title     Code = "Title"
	SubTitle  Code = "SubTitle"
	ListItem  Code = "ListItem"
	Empty     Code = "Empty"
	Error     Code = "Error"
	PlaneText Code = "PlaneText"
)
