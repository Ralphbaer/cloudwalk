default: build

package-name=Quake-Log-Parser
service-name=Quake-Log-Parser-service
artifacts-dir=./.artifacts
bin-dir=./.bin

.PHONY: gen
gen:
	go generate ./...

clean_build:
	echo "Cleaning build directory..." && \
	rm -rf $(bin-dir) && \
    mkdir -p $(bin-dir)

build: clean_build
	go build -o $(bin-dir)/$(service-name)

build_linux: clean_build
	GOOS=linux GOARCH=amd64 go build -o $(bin-dir)/$(service-name)

run:
	go run main.go
	
test:
	mkdir -p $(artifacts-dir) && \
	gotestsum --format short-verbose --junitfile $(artifacts-dir)/junit-report.xml \
	./... github.com/Ralphbaer/cloudwalk/common... \
	-coverprofile=$(artifacts-dir)/coverage.out

cover-html:
	go tool cover -html=$(artifacts-dir)/coverage.out -o $(artifacts-dir)/coverage.html

cover: test cover-html
	open $(artifacts-dir)/coverage.html

clean:
	rm -rf $(bin-dir)
	rm -rf $(artifacts-dir)