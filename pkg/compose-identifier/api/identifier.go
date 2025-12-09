package api

type Identifier interface {
	Value() uint32
	String() string
}
