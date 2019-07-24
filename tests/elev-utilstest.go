package main

import (
	"fmt"
	"../../elev-utils"
)

const (
	// NW Quadrant (Bozeman) z = ~1530m
/*	lat	= 45.638347
	lon	= -111.025257
*/
	// SE Quadrant, (Maputo Africa) z = ~280m
	lat	= -25.9692
	lon	= 32.5732

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
