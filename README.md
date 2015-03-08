# docker with golang

* simple sample scripts for myself

## check environment

* Ubuntu 14.04 

## Setup

* Install golang(docker host)
```
$ sudo apt-get install golang
```

* Install Docker
```
$ sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 36A1D7869245C8950F966E92D8576A8BA88D21E9
$ sudo sh -c "echo deb https://get.docker.com/ubuntu docker main > /etc/apt/sources.list.d/docker.list"
$ sudo apt-get update
$ sudo apt-get install lxc-docker
```

* Install fig
```
$ curl -L https://github.com/docker/fig/releases/download/1.0.1/fig-`uname -s`-`uname -m` > ./fig
$ chmod +x ./fig
$ sudo mv /usr/local/bin/fig
```

* build sample go program
```
$ export GOPATH=/home/<your username>/golib
$ golang-go get
$ golang-go build
```

* try it with fig background mode
```
$ fig up -d
```

## Usage

* Using the browser, please go to this URL
    * http://&lt;docker host&gt;:8080/
