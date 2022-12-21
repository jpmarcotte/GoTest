package employee

type Gender string // TODO: Implement limited values

type Employee struct {
	Gender Gender `json:"gender"`
	Id     int    `json:"id"`
}
