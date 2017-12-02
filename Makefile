BINARY=goboilerplatd
TESTDIRS=`go list ./... | grep -v /vendor/ `
BASEDIR=$(shell pwd)

default:
	rm -rf build/
	mkdir build
	CGO_ENABLED=1 GOOS=linux go build -ldflags "-linkmode external -extldflags -static" -o build/$(BINARY)

test:
	GO_BOILERPLATE_CONFIGPATH=$(BASEDIR) GO_BOILERPLATE_CONFIGFILE=config_test go test --cover -v $(TESTDIRS)


clean:
	rm -rf build/
