# This is the makefile for Ubuntu distro (Tested on 13.10). Sorry for other users :(
MONGO_INSTALL = sudo apt-get install mongodb
BZR_INSTALL = sudo apt-get install bzr
MONGO_GO_DRIVER_INSTALL = go get labix.org/v2/mgo


all:
	$(MONGO_INSTALL)
	$(BZR_INSTALL)
	$(MONGO_GO_DRIVER_INSTALL)
test:
	go test
