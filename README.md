# Nextcloud InfluxDB importer

I love [OwnTracks](https://owntracks.org/) and store all the collected data in InfluxDB. I also use Nextcloud and I would like to visualize all my tracks in there using [Maps](https://apps.nextcloud.com/apps/maps). 
This app read the stored locations from InfluxDB, converts data from one
day to GPX and stores the resulting files in Nextcloud.

## ToDo

* [x] Read available data from InfluxDB
* [] Check what's already available in Nextcloud
* [] For every missing day read data from InfluxDB
* [] Convert to GPX
* [] Store in Nextcloud

## Notes

Forward InfluxDB port to localhost:

```
kubectl port-forward <pod name> <local port>:8086
```

Run influxdb locally:

```
docker run -it --rm --network host influxdb /bin/bash
influx -precision rfc3339 -port <local port>
```

