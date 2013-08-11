package google

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
	"errors"
)

//used to Unmarshal response from google, requires public access of the fields (so needs to be upercase) 
type DistanceResponse struct {
	Routes []struct {
		Legs []struct {
			Distance struct {
				Text  string
				Value float64
			}
		}
	}
}
func GetDistance(p1, p2 string) (meters float64, err error) {
	p1 = url.QueryEscape(p1)
	p2 = url.QueryEscape(p2)

	urlTemplate := "http://maps.googleapis.com/maps/api/directions/json?origin=%s&destination=%s&sensor=false"
	driveTimeUrl := fmt.Sprintf(urlTemplate, p1, p2)

	res, err := http.Get(driveTimeUrl)
	if err != nil {
		return 0, err
	}
	rawResponse, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return 0, err
	}

	distanceResponse := DistanceResponse{}
	if err := json.Unmarshal(rawResponse, &distanceResponse); err == nil {
		if len(distanceResponse.Routes) == 0 {
			return 0, errors.New(fmt.Sprintf("Unable to find distance between '%s' and '%s'", p1, p2))
		}
		
		distance := distanceResponse.Routes[0].Legs[0].Distance
		meters = distance.Value
	} else {
		return 0, err
	}
	return
}