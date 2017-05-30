package foo2

import (
	"fmt"
)

type InputServicer interface {
	GetData(url string) ([]byte, error)
}

type Reader struct {
	Servicer InputServicer
}

func (r Reader) GetData(url string) ([]byte, error)  {
	fmt.Println("GetData method from foo2.Reader")
	return []byte("Hello"), nil
}

func (r Reader) Fetch(url string) ([]byte, error) {
	data, err := r.Servicer.GetData(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch the data from url : %v", err)
	}
	return data, nil
}
