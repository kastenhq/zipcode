package zipcode

import (
	"testing"
)

const testHost = "localhost:8000"

var testZips = []string{
	"96146",
	"96148",
	"96133",
	"96160",
	"96161",
	"96162",
	"96134",
	"96135",
	"96137",
	"90201",
	"90202",
	"90270",
	"90209",
	"90210",
	"90211",
	"90212",
	"90213",
	"90220",
	"90221",
	"90222",
	"90223",
	"90224",
	"90230",
	"90231",
	"90232",
	"90233",
}

func TestResetInsert(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Fatal(err.Error())
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			t.Error(err.Error())
		}
	}()
	defer s.Close()
	c := newClient(testHost)
	err = c.resetOrders()
	if err != nil {
		t.Fatal(err.Error())
	}
	err = c.insertOrder(testZips...)
	if err != nil {
		t.Fatal(err.Error())
	}
	return
	for _, zip := range testZips {
		err := c.insertOrder(zip)
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}