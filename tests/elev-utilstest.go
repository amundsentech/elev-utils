package main

import (
	"fmt"
	"../../elev-utils"
)

const (
	lat	= 45.638347
	lon	= -111.025257
	demdir	= "/data/dem/hdt/"
)

func main() {

        // OPTION #1, use elev-utils package
        z, err := srtm.ElevationFromLatLon(demdir,lat,lon)
        if err != nil {
                fmt.Printf("%s",err.Error())
        }

        fmt.Printf("Your elevation is: %v meters\n",z)
}
