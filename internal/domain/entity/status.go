package entity

type Status struct {
	Name string
}

func NewStatus(status string) *Status {
	return &Status{
		Name: status,
	}
}
