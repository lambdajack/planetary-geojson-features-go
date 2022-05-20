# Planetary Geojson Features Go

## Raw or Unmarshaled geojson features for every region defined by Geofabrik

## Usage:

`go get github.com/lambdajack/planetary-geojson-features-go/region`

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

### All functions/regions are listed [here](INDEX.md)
