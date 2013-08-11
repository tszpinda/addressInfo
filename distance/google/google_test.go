package google

import (
	"testing"
	"fmt"
	"log"
)

func initTest() {
}

func TestGetDistance(t *testing.T) {
	d, e := GetDistance("EX16 6AB", "EX5 2UX")
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(d)
}
