# SRTM: Elevation Utilities
This package contains several utilities that take a lat / lon value and returns the associated elevation *in meters_* anywhere in the world.

Most of these functions are designed specifically as an interface between SRTM data and a lat/lon query, reducing the need to rely on some other non-golang software (eg. shelling a call to gdal with `gdallocationinfo -valonly /path/to/dem.vrt -geoloc lat lon`).

Currently, the primary function is `srtm.ElevationFromLatLong`, which performs as you'd expect.  Future support for `srtm.ElevationFromWKT` and `srtm.ElevationFromBBOX` is forthcoming.

The primary use case for this utility is to provide a very fast service for making multiple repeat elevation queries at scale, eg. finding elevation values for coordinates forming point/line/polygon geometries... or millions of them.

An example use of this code is to first establish one or more SRTMs' details by calling `getSrtm` for a series of lat/lons that cover the area of investigation (a single 1-arc-second SRTM tile is good for about 20-30m square), then fetching elevation for *n* coordinates.

In application, this code converts the lat lon to a byte lookup value within the file; and seeks for just that byte value; thus does not require loading any SRTM into memory.  This way the process is fast and lightweight.

For more information on SRTM, please visit `https://wiki.openstreetmap.org/wiki/SRTM`.

Currently, the author maintains a private worldwide set of SRTM data on a private google storage bucket.  This set may be made available on a per-inquiry basis.


### TBD
If any given SRTM (.hgt) file does not exist locally, construct a method to fetch it.  This functionality would allow a user to maintain a limited set of SRTM data without downloading the entire world (~ 500 gb).


### Notes
1) Several functions in the **SRTM** package were adapted from `https://github.com/tkrajina/go-elevations`, whose repo is designed as a standalone http client service that utilizes SRTM data available online... reading and storing the SRTM tile into memory... useful for maintaining a small footprint but is not applicable for coordinate lookups worldwide @ scale.

2) This package requires that 1-arc-second or 3-arc-second SRTM data are already available in a directory accessible by the code that calls this package.  If the data is not present, the response will be an error of "SRTM not available, please download" or similar.

3) No efforts are planned to keep this repo backwards-compatible.
