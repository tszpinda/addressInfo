package filecache

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
	model "github.com/tszpinda/addressInfo/model"
)

var filename = os.TempDir() + "/cache-address.txt" //"/tmp/cache-address.txt"
var cache map[string]model.Address
var cacheLock sync.RWMutex

func CacheAddress(address model.Address) (err error) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	if cache == nil {
		loadCache()
	}
	cache[address.Postcode] = address

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	line := fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", address.Country, 
															 address.County, 
															 address.Town, 
															 address.Postcode, 
															 address.Street, 
															 address.HouseNumber)
	_, err = f.WriteString(line)
	if err != nil {
		return err
	}
	

	return
}

func GetAddress(postcode string) (dist model.Address, found bool) {
	if cache == nil {
		cacheLock.Lock()
		defer cacheLock.Unlock()
		if cache == nil {
			loadCache()
		}
	}

	dist, ok := cache[postcode]
	return dist, ok
}

func loadCache() {
	cache = make(map[string]model.Address)
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
		var addr model.Address = model.Address{}
		
		addr.Country = record[0] 
		addr.County = record[1]
		addr.Town = record[2]
		addr.Postcode = record[3]
		addr.Street = record[4]
		addr.HouseNumber = record[5]
		cache[addr.Postcode] = addr
	}
	fmt.Println("Address Cache loaded")
}
