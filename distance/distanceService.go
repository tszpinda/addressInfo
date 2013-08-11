package distance

import (
	"code.google.com/p/gorest"
	"fmt"
	"strings"
	gd "github.com/tszpinda/addressInfo/distance/google"
	od "github.com/tszpinda/addressInfo/distance/openData"
	model "github.com/tszpinda/addressInfo/model"
	filecache "github.com/tszpinda/addressInfo/distance/filecache"
)

type DistanceService struct {
	gorest.RestService `root:"/ds/" consumes:"application/json" produces:"application/json"`
	getDistance        gorest.EndPoint `method:"GET" path:"/distance/{postcode1:string}/{postcode2:string}/{lookupType:int}" output:"Distance"`
}

func (serv DistanceService) GetDistance(postcode1, postcode2 string, lookupType int) model.Distance {
	fmt.Println("incoming GetDistance request: ", postcode1, postcode2, lookupType)
	
	cacheDisabled := true
	inCache := false
	var d model.Distance;
	if !cacheDisabled {
		d, inCache = filecache.GetDistance(postcode1, postcode2)
	}
	m := float64(0)
	var err error 
	if inCache && !cacheDisabled {
		m = d.Meters
	}else {
		if lookupType == 1 {
			m, err = gd.GetDistance(postcode1, postcode2)
		}else {
			m, err = od.GetDistance(postcode1, postcode2)
		}
	}
	
	if err != nil {
		if strings.ContainsAny(err.Error(), "Unable to find distance") {		
			serv.ResponseBuilder().SetResponseCode(404).WriteAndOveride([]byte(err.Error()))
		}else{
			serv.ResponseBuilder().SetResponseCode(500).WriteAndOveride([]byte(err.Error()))
		}
	} else {
		filecache.CacheDistance(postcode1, postcode2, m)
	}
	return model.Distance{postcode1, postcode2, m}
}
