all: build

APP_NAME = puppet
PKGDIR_TMP = ${TMPDIR}golang

.pre-build:
	mkdir -p build

deps:
	go get -u github.com/golang/dep/...
	go get -u golang.org/x/lint/golint
	dep ensure -vendor-only -v

clean:
	rm -rf build/
	rm -rf ${PKGDIR_TMP}_darwin

build: .pre-build
	GOOS=darwin go build -i -o build/${APP_NAME}.darwin.ext -pkgdir ${PKGDIR_TMP}
	GOOS=windows go build -i -o build/${APP_NAME}.windows.ext.exe -pkgdir ${PKGDIR_TMP}


lastrun:
	rm last_run_report.yaml
	sudo cp /opt/puppetlabs/puppet/cache/state/last_run_report.yaml last_run_report.yaml
	sudo chown graham_gilbert:admin last_run_report.yaml