package srtm

import	(
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
)


// SrtmTile holds file path and details of a single SRTM file (...which are themselves 'Tiles')
type SrtmTile struct {
	Latitude	int
	Longitude	int
	Name		string
	Dir		string
	Path		string
	SquareSize	int
	Size		int64
}


// getElevation is main handler for a single lat lon input
func ElevationFromLatLon(demdir string, lat, lon float64) (float64, error) {
        srtm, err := getSrtm(demdir, lat, lon)
	if err != nil {
                return 0, err
        }


	elevation, err := srtm.getElevationFromSrtm(lat, lon)
	if err != nil {
                return 0, err
        }

        return elevation, nil
}


// getElevationFromWKT is not implemented yet

// getElevationFromBBOX is not implemented yet

// getSrtm is a specific handler for filling in details of a single SRTM Tile
func getSrtm(demdir string, lat, lon float64) (SrtmTile, error) {
	var srtm SrtmTile

        srtm.Dir = demdir

        srtm.getSrtmFileName(lat, lon)

	fmt.Println(srtm.Path)

        err := srtm.getSquareSize()
	if err != nil {
		return srtm, err
	}

	return srtm, err
}


// getElevationFromSrtm is a specific handler for elevation, if SRTM details are known
func (self *SrtmTile) getElevationFromSrtm(lat, lon float64) (float64, error) {
	row, column := self.getRowAndColumn(lat, lon)

        elevation, err := self.getElevationFromRowAndColumn(row, column)
        if err != nil {
                return 0, err
        }

        return elevation, nil
}


// SRTM compliance prescribes distinct filenames eg. S56W072.hgt 
func (self *SrtmTile) getSrtmFileName(lat, lon float64) {
	y := "S"
	if lat >= 0 {
		y = "N"
	}

	x := "W"
	if lon >= 0 {
		x = "E"
	}

	self.Latitude = int(math.Abs(math.Floor(lat)))
	self.Longitude = int(math.Abs(math.Floor(lon)))

	self.Name = fmt.Sprintf("%s%02d%s%03d.hgt", y, self.Latitude, x, self.Longitude)

	self.Path = filepath.Join(self.Dir, self.Name)
}


// the SquareSize determines the density of integers from the hgt file
// Each 3-arc-second data tile has 1442401 integers representing a 1201×1201 grid
// Each 1-arc-second data tile has 12967201 integers representing a 3601×3601 grid
func (self *SrtmTile) getSquareSize() error {

	// prepare file for observation
	f, err := os.Stat(self.Path)
	if err != nil {
		fmt.Println("Can not find",self.Path,"?  Error: ",err.Error())
		return err
	}

	// get the size
	self.Size = f.Size()

	// get the tile size
	if self.SquareSize <= 0 {
		squareSizeFloat := math.Sqrt(float64(self.Size) / 2.0)

		if squareSizeFloat != float64(self.SquareSize) || self.SquareSize <= 0 {
			return errors.New(fmt.Sprintf("Invalid size for file %s: %d", self.Name, self.Size))
		}

		self.SquareSize = int(squareSizeFloat)
	}

	return nil
}


// getRowAndColumn calculates the lookup []byte in the grid
// NOTE: row and column are int, therefore become FLOOR rounded values
func (self *SrtmTile) getRowAndColumn(lat, lon float64) (int, int) {
	row := int((float64(self.Latitude) + 1.0 - lat) * (float64(self.SquareSize - 1.0)))
	column := int((lon - float64(self.Longitude)) * (float64(self.SquareSize - 1.0)))
	return row, column
}


// find the elevation value associated with the row and column
func (self *SrtmTile) getElevationFromRowAndColumn(row, column int) (float64, error) {
	i := int64(row * self.SquareSize + column)

	// calculate the byte range
	byte1 := i*2
	byte2 := i*2+1

	// open the file for reading
	f, err := os.Open(self.Path)
	if err != nil {
		return 0, err
	}

	// get the results from the byte location
	result1, err := f.Seek(byte1, 0)
	result2, err := f.Seek(byte2, 0)

	// do some magic
	// github.com/tkrajina/go-elevations/blob/master/geoelevations/srtm.go
	final := int(result1)*256 + int(result2)

	if final > 9000 {
		return math.NaN(), err
	}

	return float64(final), nil
}
