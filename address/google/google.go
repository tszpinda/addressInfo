package google

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"net/url"
	"errors"
	model "github.com/tszpinda/addressInfo/model"
	filecache "github.com/tszpinda/addressInfo/address/filecache"
)


type PostCodeResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float32
				Lng float32
			}
		}
	}
}

type AddressResponse struct {
	Results []struct {
		Address_components []struct {
			Long_name string
			Types     []string
		}
	}
}


func getGeometry(postcode string) (lat, lng float32, err error) {
	postcode = url.QueryEscape(postcode)
	urlGeometry := "http://maps.googleapis.com/maps/api/geocode/json?address=" + postcode + "&sensor=false"
	res, err := http.Get(urlGeometry)
	if err != nil {
		return 0, 0, err
	}
	rawResponse, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return 0, 0, err
	}
	
	postCodeJson := PostCodeResponse{}
	if err := json.Unmarshal(rawResponse, &postCodeJson); err == nil {
		if len(postCodeJson.Results) == 0 {
			return 0, 0, errors.New(fmt.Sprintf("Unable to find postcode: '%s'", postcode))
		}
		
		lat = postCodeJson.Results[0].Geometry.Location.Lat
		lng = postCodeJson.Results[0].Geometry.Location.Lng
	} else {
		return 0, 0, err
	}
	return
}

func contains(sl []string, text string) bool {
	for _, v := range sl {
		if strings.EqualFold(v, text) {
			return true
		}
	}
	return false
}

//uses 2 google api for more precise address 
//first one to get latitude and longitude
//the from coordinates get address
func GetAddress(postcode string) (model.Address, error) {
	cacheDisabled := true
	if !cacheDisabled {
		a, f := filecache.GetAddress(postcode)
		if f {
			return a, nil
		}
	}
	
	lat, lng, err := getGeometry(postcode)
	if err != nil {
		address := model.Address{}
		return address, err
	}
	address, err := getAddress(lat, lng)
	if err == nil {
		filecache.CacheAddress(address)
	}
	return address, err
}
func getAddress(lat, lng float32) (model.Address, error) {
	urlTemplate := "http://maps.googleapis.com/maps/api/geocode/json?latlng=%g,%g&sensor=false"
	geocodeUrl := fmt.Sprintf(urlTemplate, lat, lng)
	address := model.Address{}

	res, err := http.Get(geocodeUrl)
	if err != nil {
		return address, err
	}
	rawResponse, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return address, err
	}

	addressResponse := AddressResponse{}
	if err := json.Unmarshal(rawResponse, &addressResponse); err == nil {
		//get first one as its most accurate
		if len(addressResponse.Results) == 0 {
			return address, errors.New(fmt.Sprintf("Unable to find address for a latitude '%s' and longtitude '%s'postcode: '%s'", lat, lng))
		}		
		
		addressSlice := addressResponse.Results[0]
		addressData := addressSlice.Address_components
		for i := 0; i < len(addressData); i++ {
			addressElem := addressData[i]
			//fmt.Printf("p[%v] == %v\n", i, addressElem.Long_name)
			if contains(addressElem.Types, "street_number") {
				address.HouseNumber = addressElem.Long_name
			} else if contains(addressElem.Types, "route") {
				address.Street = addressElem.Long_name
			} else if contains(addressElem.Types, "postal_town") {
				address.Town = addressElem.Long_name
			} else if contains(addressElem.Types, "postal_code") {
				address.Postcode = addressElem.Long_name
			} else if contains(addressElem.Types, "country") {
				address.Country = addressElem.Long_name
			} else if contains(addressElem.Types, "administrative_area_level_2") {
				address.County = addressElem.Long_name
			}
		}
	} else {
		return address, err
	}

	return address, nil
}
