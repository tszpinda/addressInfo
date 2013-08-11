package openData

import (
	"testing"
	"fmt"
	"log"
)

func initTest() {
}

func TestGetDistance_fromSingleFile(t *testing.T) {
	d, e := GetDistance("EX16 6AB", "EX5 2UX")
	if e != nil {
		log.Fatal(e)
	}
	if int(d) != 21251 {
		log.Fatal("Invalid distance calculated ", d, " != ", 21251)
	} 
	fmt.Println(d)
}

func TestGetDistance(t *testing.T) {
	d, e := GetDistance("EX16 6AB", "TA22 9RT")
	if e != nil {
		log.Fatal(e)
	}
	if int(d) != 13325 {
		log.Fatal("Invalid distance calculated ", d, " != ", 13325)
	} 
	fmt.Println(d)
}
