package contract

const IDKey = "k:id"

type IDService interface {
	NewID() string
}
