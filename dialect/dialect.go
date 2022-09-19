package dialect

type Name int

func (n Name) String() string {
	return "opengauss"
}

const (
	Invalid Name = iota
	OPENGAUSS
)
