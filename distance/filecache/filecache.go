package filecache

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	model "github.com/tszpinda/addressInfo/model"
)

var filename = os.TempDir() + "/cache-distance.txt"
var cache map[string]model.Distance
var cacheLock sync.RWMutex

func CacheDistance(p1, p2 string, meters float64) (err error) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	if cache == nil {
		loadCache()
	}
	cache[p1+p2] = model.Distance{p1, p2, meters}
	cache[p2+p1] = model.Distance{p1, p2, meters}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	line := fmt.Sprintf("%v,%v,%v\n", p1, p2, meters)
	_, err = f.WriteString(line)
	if err != nil {
		return err
	}
	

	return
}

func GetDistance(postcode1, postcode2 string) (dist model.Distance, found bool) {
	if cache == nil {
		cacheLock.Lock()
		defer cacheLock.Unlock()
		if cache == nil {
			loadCache()
		}
	}

	dist, ok := cache[postcode1+postcode2]
	return dist, ok
}

func loadCache() {
	cache = make(map[string]model.Distance)
	fmt.Println("Loading cache")

	f, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	fmt.Println(filename)
	if err != nil {
		return
	}
	defer f.Close()
	reader := csv.NewReader(f)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Loading: ", record)
		p1 := record[0]
		p2 := record[1]
		m, _ := strconv.ParseFloat(record[2], 64)
		cache[p1+p2] = model.Distance{p1, p2, m}
		cache[p2+p1] = model.Distance{p2, p1, m}
	}
	fmt.Println("Cache loaded")
}
