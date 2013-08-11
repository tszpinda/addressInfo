address-info
============

golang demo - how to use google api and create own rest webservices:

1. returns full address from given postcode ie.: 

http://localhost:8080/ds/address/bs2

2. returns distance (in meters) between two given postcodes, ie.: 

http://localhost:8080/ds/distance/bs2/bs1

to get started (on unix):
```bash
git clone https://github.com/tszpinda/address-info.git
cd address-info
export GOPATH=`pwd`
go run src/main/addressInfo.go

http://localhost:8080/index #to open simple ui
```

API used:

https://developers.google.com/maps/documentation/geocoding/

https://developers.google.com/maps/documentation/directions/

BE AWARE OF:

Note: the Directions API may only be used in conjunction with displaying results on a Google map; using Directions data without displaying a map for which directions data was requested is prohibited. Additionally, calculation of directions generates copyrights and warnings which must be displayed to the user in some fashion. For complete details on allowed usage, consult the Maps API Terms of Service License Restrictions.
