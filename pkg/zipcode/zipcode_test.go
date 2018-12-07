package zipcode

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
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

func randSuffix() string {
	return fmt.Sprintf("-%04d", rand.Int31n(10000))
}

func TestResetInsertProd(t *testing.T) {
	tz := make([]string, 0, len(testZips))
	copy(tz, testZips)
	tz[len(tz)/3] = tz[len(tz)/3] + randSuffix()
	testResetInsert(t, tz)
}

func TestResetInsertTest(t *testing.T) {
	testResetInsert(t, testZips)
}

func testResetInsert(t *testing.T, tz []string) {
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
	err = c.insertOrder(tz...)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCities(t *testing.T) {
	t.SkipNow()
	s, err := New()
	if err != nil {
		t.Fatal(err.Error())
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			t.Log(err.Error())
		}
	}()
	defer s.Close()
	time.Sleep(time.Second)
	c := newClient(testHost)
	cities, err := c.orderCountByCity()
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(cities) == 0 {
		t.Fatal("No cities returned: please test with data")
	}
	if strings.Contains(cities[0], "could not get city") {
		t.Fatal(cities[0])
	}
	t.Logf("Orders in %d cities recorded!", len(cities))
}
