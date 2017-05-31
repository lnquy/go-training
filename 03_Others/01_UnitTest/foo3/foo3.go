package foo3

import (
	"fmt"
)

type InputServicer interface {
	GetData(url string) ([]byte, error)
	GetID(url string) ([]byte, error)
	DeleteData(url string) ([]byte, error)
	UpdateData(url string) ([]byte, error)
}

// In case interface has too many functions and we don't want to implement all those functions,
// We can make interface becomes the anonymous (embedded) field of struct.
// So we can override the only function(s) we need.
// CAREFUL: May cause runtime panic when calling a function which is not implemented yet.
type Reader struct {
	InputServicer // Embedded interface
}

func (r Reader) GetData(url string) ([]byte, error)  {
	fmt.Println("GetData method from foo3.Reader")
	return []byte("Hello"), nil
}

func (r Reader) Fetch(url string) ([]byte, error) {
	data, err := r.GetData(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch the data from url : %v", err)
	}
	return data, nil
}
