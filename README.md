# Planetary Geojson Features Go

## Raw and Unmarshaled geojson features for every region defined by Geofabrik as Go functions

### Installation:

`go get github.com/lambdajack/planetary-geojson-features-go/region`

### Usage:

```go
func main() {
	// Raw geojson
	lux := region.EuropeLuxembourgBytes()
	fmt.Println(string(lux))

	// Unmarshaled geojson
	china, err := region.AsiaChinaUnmarshaled()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(china.BBox)
}
```

#### All functions/regions are listed [here](INDEX.md)

## Acknowledgments

#### Unmarshalling by [Paul Mach](https://github.com/paulmach/go.geojson)

#### Geojson features generated from the .poly files supplied by [Geofabrik](https://download.geofabrik.de/)

## Disclaimer

Not all regions in the world are available (only those supplied by Geofabrik). In line with their disclaimer, the region boundaries do not represent any political statements.
