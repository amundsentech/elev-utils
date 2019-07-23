import (
	"gihub.com/lumin8/elev-utils"
)

const (
	lon	= 45.676998
	lat	= -111.042931
)

func main {

        // OPTION #1, use elev-utils package
        z, err := srtm.ElevationFromLatLon(demdir,lat,lon)
        if err != nil {
                log.Printf("%s",err.Error())
        }

        log.Printf("Your elevation is: %v meters",z)
}
