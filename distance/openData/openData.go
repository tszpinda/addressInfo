package openData

import (
	"encoding/csv"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"errors"
	"fmt"
)

type result struct {
	record []string
	err    error
}

func getFilename(postcode string) string {
	return "/home/tszpinda/dev/code/go/openData/Data/CSV/" + strings.ToLower(postcode)[0:2] + ".csv"
}

func getRecord(postcode string, c chan result) {
	filename := getFilename(postcode)
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		c <- result{nil, errors.New(fmt.Sprintf("Invalid postcode: '%s'", postcode))}
		return
	}
	defer f.Close()
	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			c <- result{nil, err}
			return
		}
		p := strings.Replace(record[0], " ", "", -1)
		if p == postcode {
			//return record, nil
			c <- result{record, nil}
			return
		}
	}
	c <- result{nil, errors.New("Postcode not found: " + postcode)}
}
func getRecordsFromSingleFile(p1, p2 string) ([]string, []string, error) {
	filename := getFilename(p1)
	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)

	count := 0
	var r1 []string
	var r2 []string
	p1found := false
	p2found := false
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}
		count++
		p := strings.Replace(record[0], " ", "", -1)
		if p == p1 {
			r1 = record
			p1found = true
		}
		if p == p2 {
			r2 = record
			p2found = true
		}
		if p1found && p2found {
			break
		}
	}
	
	if !p1found {
		return nil, nil, errors.New("Postcode not found: " + p1)
	}
	
	if !p2found {
		return nil, nil, errors.New("Postcode not found: " + p2)
	}
	
	return r1, r2, nil
}
func getRecords(p1, p2 string) ([]string, []string, error) {
	if getFilename(p1) == getFilename(p2) {
		return getRecordsFromSingleFile(p1, p2)
	}

	//read from two diff files so do it in separate rutines
	c := make(chan result)
	go getRecord(p1, c)
	go getRecord(p2, c)
	r1, r2 := <-c, <-c

	if r1.err != nil {
		return nil, nil, r1.err
	}
	if r2.err != nil {
		return nil, nil, r2.err
	}

	return r1.record, r2.record, nil
}

func GetDistance(p1, p2 string) (float64, error) {
	p1 = strings.ToUpper(strings.Replace(p1, " ", "", -1))
	p2 = strings.ToUpper(strings.Replace(p2, " ", "", -1))

	r1, r2, err := getRecords(p1, p2)
	if err != nil {
		return 0, err
	}

	//simple calculation
	x1, _ := strconv.ParseFloat(r1[2], 64)
	y1, _ := strconv.ParseFloat(r1[3], 64)
	x2, _ := strconv.ParseFloat(r2[2], 64)
	y2, _ := strconv.ParseFloat(r2[3], 64)

	xd := math.Abs(x2 - x1)
	yd := math.Abs(y2 - y1)
	d := math.Sqrt(xd*xd + yd*yd)
	return d, nil

}
