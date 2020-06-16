# Nextcloud InfluxDB importer

I love [OwnTracks](https://owntracks.org/) and store all the collected data in InfluxDB. I also use Nextcloud and I would like to visualize all my tracks in there using [Maps](https://apps.nextcloud.com/apps/maps). 
The app in this project reads the stored locations from InfluxDB, converts data from every single
day to GPX files and stores the resulting files in Nextcloud.

## Configuration

The app can be configured by the following environment variables:

* **INFLUXDB_URI**: URI of the InfluxDB, default "http://localhost:8086"
* **INFLUXDB_DB_NAME**: Name of the InfluxDB database, default "owntracks"
* **INFLUXDB_MEASUREMENT_NAME**: Name of the InfluxDB measurement that keeps the OwnTracks data, default "owntracks"
* **OWNTRACKS_USER**: User whose track data should be extracted, default "holger"
* **NEXTCLOUD_URI**: WebDAV URI of the Nextcloud instance
* **NEXTCLOUD_USER**
* **NEXTCLOUD_PASSWORD**
* **TRACKDIR**: Path on Nextcloud where GPX files should be stored, default "/Tracks/owntracks/"

## Test

```bash
go test ./...
```

## Build 

```bash
go build main.go
```

## Docker

Build docker image with

```bash
make build
```

Push image using

```bash
make push
```

Run docker container using e.g.

```bash
docker run --rm -e "NEXTCLOUD_URI=https://my-nextcloud.de/remote.php/dav/files/holger/" -e "NEXTCLOUD_USER=holger" -e "NEXTCLOUD_PASSWORD=password" windsource/nextcloud-influxdb-tracks-importer
```
