package employee

type Gender string // TODO: Consider restricted values e.g. "non-binary", "female", "male"

type Employee struct {
	Gender Gender `json:"gender"`
	Id     int    `json:"id"`
}
