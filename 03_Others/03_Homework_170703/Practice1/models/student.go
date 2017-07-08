package models

type (
	Name struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	Student struct {
		Id      uint          `json:"id"`
		Name    `json:"name"` // Embedded struct
		Class   string        `json:"class"`
		School  string        `json:"school"`
		Message string        `json:"message"`
	}
)

func (s *Student) ChangeFirstName(fn string) {
	s.FirstName = fn // Student can directly call the fields of embedded Name struct rather than s.Name.FirstName
}

func (s *Student) ChangeMessage(msg string) {
	s.Message = msg
}
