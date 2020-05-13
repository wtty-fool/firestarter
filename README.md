# Firestarter
![CircleCI build status](https://img.shields.io/circleci/build/github/wtty-fool/firestarter)
[![docker build status](https://img.shields.io/docker/cloud/build/wttyf00l/firestarter)](https://hub.docker.com/r/wttyf00l/firestarter)

A simple, containerized Prometheus metrics generator.

> Note: This repository is used mainly for experimenting with Prometheus metrics
reporting. It does not serve any other purpose than educational and should
never be used outside testing/dev/staging setups.

## Adding Firestarter as Prometheus target
The easiest way to do have Firestarter as a target right from the start is to run Prometheus container with `etc-prometheus` mounted as configuration directory:
```
docker run --network host wttyf00l/firestarter
docker run --network host --volume etc-prometheus:/etc/prometheus prom/prometheus
```

Otherwise you may have to edit configuration inside running container and make Prometheus reload it by sending HUP signal.
```
docker run --network host wttyf00l/firestarter
docker run --network host --name my-prometheus prom/prometheus
docker exec -it my-prometheus sh
$ vi /etc/prometheus/prometheus.yml
# add a target, save, exit
$ kill -HUP 1
$ exit
```
