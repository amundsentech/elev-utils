# SRTM: Elevation Utilities
This package contains several utilities that take a lat / lon value and returns the associated elevation *in meters_* anywhere in the world.

**SRTM** is designed specifically as an interface between local SRTM data and a lat/lon query, reducing the need to rely on shelling out to some other software (eg. shelling a call to a gdal utility `gdallocationinfo -valonly /path/to/dem.vrt -geoloc lat lon`).

The primary use of this utility is to provide a fast service for making multiple repeat elevation queries at scale, eg. finding elevation values for coordinates forming point/line/polygon geometries... or millions of them.

An example use of this code is to first establish one or more srtms' details by calling `getSrtm` for each lat/lon int (a single srtm is good for about 20-30m squared), then fetching elevation for one/more series of coordinates.

This code converts the lat lon to a byte value within the file, and seeks for just that byte value, therefore does not require loading an SRTM into memory for each request.

For more information on SRTM, please visit `https://wiki.openstreetmap.org/wiki/SRTM`.

Currently, the author is hosting a worldwide set of srtm data on a private google storage bucket located at `gs://data.map.life/raw/dem`, which may be made available on a per-request basis.


### TBD
If the srtm (hgt) file does not exist locally, build in a method to download it.  This would allow a client to maintain a limited set of srtm data, without downloading the entire world (~ 500 gb).


### Notes
Several functions in the **SRTM** package is adapted from `https://github.com/tkrajina/go-elevations`, whose repo is designed as a standalone http client service that utilizes srtm data available online... useful for maintaining a small footprint but does not work at scale.

This package requires that 1-arc-second or 3-arc-second SRTM data are already available in a directory accessible by the code that relies on this library.  If the data is not present, the response will be an error of "srtm not available, please download" or similar.
