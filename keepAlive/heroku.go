package keepAlive

import (
	"fmt"
	"os"
)

func Heroku(port string) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Unable to get hostname from os, error: ", err)
	}
	fmt.Println("Heroku keep alive started: " + hostname + ":" + port)

}
