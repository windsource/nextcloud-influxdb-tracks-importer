.PHONY: build push

IMAGE = windsource/nextcloud-influxdb-tracks-importer
VERSION = 1.0.0

build:
	docker build -t $(IMAGE):$(VERSION) -t $(IMAGE):latest .

push: build
	docker push $(IMAGE):$(VERSION)
	docker push $(IMAGE):latest