package filecache

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"math/rand"
	"time"
)

var testFile = "/tmp/test.txt"

func removeFile() {
	_, err := os.OpenFile(testFile, os.O_RDONLY, 0666)
	if err == nil {
		err = os.Remove(testFile)
	}
}

func initTest() {
	removeFile()
	//we testing single instance so need to clear cache before every test method
	cache = nil
	filename = testFile
}

func TestCacheDistance(t *testing.T) {
	initTest()
	CacheDistance("postcode-11", "postcode-12", 14)
	CacheDistance("postcode-21", "postcode-22", 24)

	f, err := os.OpenFile(testFile, os.O_RDONLY, 0666)
	if os.IsNotExist(err) {
		t.Fatal("Test file does not exists")
	}
	scanner := bufio.NewScanner(bufio.NewReader(f))
	i := 0
	for scanner.Scan() {
		i++

		if i == 1 && scanner.Text() != "postcode-11,postcode-12,14" {
			t.Fatal("1st line has invalid format")
		}
		if i == 2 && scanner.Text() != "postcode-21,postcode-22,24" {
			t.Fatal("2nd line has invalid format")
		}
		if i == 3 {
			t.Fatal("To many entires found")
		}
	}
	if i == 0 {
		t.Fatal("No cached entries found")
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}

func TestGetDistance(t *testing.T) {
	initTest()
	CacheDistance("postcode-11", "postcode-12", 14)
	CacheDistance("postcode-21", "postcode-22", 24)

	dist, ok := GetDistance("postcode-11", "postcode-12")
	if !ok {
		t.Fatal("dist 1 not found")
	}
	dist, ok = GetDistance("postcode-21", "postcode-22")
	if !ok {
		t.Fatal("dist 2 not found")
	}
	dist, ok = GetDistance("postcode-22", "postcode-21")
	if !ok {
		t.Fatal("reverted dist 2 not found")
	}
	if dist.Meters != 24 {
		t.Fatal("invalid distance")
	}
	dist, ok = GetDistance("postcode-22", "postcode-22")
	if ok {
		t.Fatal("expected not entry")
	}
}

func TestCacheDistanceWhenCacheAlreadyLoaded(t *testing.T) {
	initTest()
	CacheDistance("postcode-11", "postcode-12", 14)
	//this will load from cache
	dist, ok := GetDistance("postcode-11", "postcode-12")
	fmt.Println(dist)
	if !ok {
		t.Fatal("dist 1 not found")
	}

	//this should add to file and cache
	CacheDistance("postcode-21", "postcode-22", 24)
	//this should read from cache
	_, ok = GetDistance("postcode-11", "postcode-12")
	if !ok {
		t.Fatal("dist 1 not found (2nd try)")
	}

	//this should read from cache
	dist, ok = GetDistance("postcode-21", "postcode-22")
	fmt.Println(dist)
	if !ok {
		t.Fatal("dist 2 not found")
	}
}

func readAndCache(p string, c chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		
		p1 := fmt.Sprintf("%s:1:%v", p, i)
		p2 := fmt.Sprintf("%s:2:%v", p, i)
		GetDistance(p1, p2)
		CacheDistance(p1, p2, float64(i))
		GetDistance(p1, p2)
	}
	c <- fmt.Sprintf("%s-readCache-done", p)
}
func read(p string, c chan string) {
	for i := 0; i < 50; i++ {
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		
		p1 := fmt.Sprintf("%s:1:%v", p, 1)
		p2 := fmt.Sprintf("%s:2:%v", p, 1)
		GetDistance(p1, p2)
	}
	c <- fmt.Sprintf("%v",p)
}

func TestConcurrentUse(t *testing.T) {
	initTest()

	c := make(chan string)
	go readAndCache("p1", c)
	go readAndCache("p2", c)
	go readAndCache("p3", c)
	go readAndCache("p4", c)
	go readAndCache("p5", c)
	go readAndCache("p6", c)
	go readAndCache("p7", c)
	go readAndCache("p8", c)
	go readAndCache("p9", c)
	go readAndCache("p10", c)

	for i := 0; i < 50;  i++ {
		go read(fmt.Sprintf("p%v", i), c)
	}
	
	//wait for all 50 + 10 routines to finish
	for i := 0; i < 60;  i++ {
		<-c
		//fmt.Println(g)
	}
	fmt.Println("\n")
	if len(cache) != 200 {
		t.Fatal("Expected 10 rutines to write 10 x 2 each")
	}	
}
