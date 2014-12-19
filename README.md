SUCH MAGIC, MUCH CASSANDRA WOW
===

## What is cmagic?

A cassandra object mapper using gocql under the hood.

## Why is the readme empty?

The library has just undergone substantial changes.


### Example: cassa geoindex

The following thing:

```
CREATE TABLE geo_index (
  geohash text,
  realm text,
  unique_id int,
  update_time timestamp,
  lat float,
  lon float,
  PRIMARY KEY (geohash, realm, unique_id)
)
```

Would be represented in Go like:

```go
type GeoIndex struct {
	Geohash string
	Realm string
	UniqueId int
	UpdateTime time.Time
	Lat float64
	Long float64
}
keys := Keys{
	PartitionKeys: []string{"Geohash", "Realm", "UniqueId"},
}
geoTable := keyspace.Table("GeoIndex", GeoIndex{}, Keys)
```

Then

```go
queryString := fmt.Sprintf("INSERT INTO %v (geohash, realm, unique_id, update_time, lat, lon) VALUES ('%s', '%s', %d, dateof(now()), %f, %f) USING TTL %d;", table, geoHashes["centre"], realm, uniqueId, latitude, longitude, ttl)

// Equals to

g := GeoIndex{
	GeoHash: "ff8989x",
	Realm: "London",
	UniqueId: 42,
	Lat: 0.1,
	Long: 0.2
}
// Note: TTL is missing, Insert needs no selection? There are problems with this...
geoTable.Insert(g)
```

Querying:

```go
queryString := fmt.Sprintf("SELECT geohash, realm, unique_id, lat, lon, update_time FROM %v WHERE geohash IN (%v) AND realm = '%v';", table, geoHashesInList, realm)

// Equals to
geoHashesInList := []string{"absdsd3", "fddff833f", "hsbrh3g4h3", "j3hg43h4g3hg4"}
rows, err := geoTable.Select(geoHashesInList, "London").Query().Read()
```

