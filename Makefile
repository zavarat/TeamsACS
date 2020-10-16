BUILD_VERSION   := latest
BUILD_TIME      := $(shell date "+%F %T")
BUILD_NAME      := teamsacs
RELEASE_VERSION := v1.0.1
SOURCE          := main.go
RELEASE_DIR     := ./release
COMMIT_SHA1     := $(shell git show -s --format=%H )
COMMIT_DATE     := $(shell git show -s --format=%cD )
COMMIT_USER     := $(shell git show -s --format=%ce )
COMMIT_SUBJECT     := $(shell git show -s --format=%s )

clean:
	rm -f teamsacs

gen:
	go generate

build:
	go generate
	CGO_ENABLED=0 go build -a -ldflags \
	'\
	-X "main.BuildVersion=${BUILD_VERSION}"\
	-X "main.ReleaseVersion=${RELEASE_VERSION}"\
	-X "main.BuildTime=${BUILD_TIME}"\
	-X "main.BuildName=${BUILD_NAME}"\
	-X "main.CommitID=${COMMIT_SHA1}"\
	-X "main.CommitDate=${COMMIT_DATE}"\
	-X "main.CommitUser=${COMMIT_USER}"\
	-X "main.CommitSubject=${COMMIT_SUBJECT}"\
	-s -w -extldflags "-static"\
	' \
    -o ${BUILD_NAME} ${SOURCE}

build-linux:
	go generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags \
	'\
	-X "main.BuildVersion=${BUILD_VERSION}"\
	-X "main.ReleaseVersion=${RELEASE_VERSION}"\
	-X "main.BuildTime=${BUILD_TIME}"\
	-X "main.BuildName=${BUILD_NAME}"\
	-X "main.CommitID=${COMMIT_SHA1}"\
	-X "main.CommitDate=${COMMIT_DATE}"\
	-X "main.CommitUser=${COMMIT_USER}"\
	-X "main.CommitSubject=${COMMIT_SUBJECT}"\
	-s -w -extldflags "-static"\
	' \
    -o ${RELEASE_DIR}/${BUILD_NAME} ${SOURCE}

upx:
	upx ${RELEASE_DIR}/${BUILD_NAME}

rpccert:
	echo "make rpc cert"
	openssl req -new -nodes -x509 -out /var/teamsacs/private/bssrpc.tls.crt -keyout /var/teamsacs/private/teamsacs-rpc.tls.key -days 36500 -subj \
	"/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=rpc.teamsacs.net/emailAddress=cert@teamsacs.net"

webcert:
	echo "make web cert"
	openssl req -new -nodes -x509 -out /var/teamsacs/private/bssweb.tls.crt -keyout /var/teamsacs/private/teamsacs-web.tls.key -days 36500 -subj \
	"/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=sweb.teamsacs.net/emailAddress=cert@teamsacs.net"

ci:
	@read -p "type commit message: " cimsg; \
	git ci -am "$(shell date "+%F %T") $${cimsg}"

push:
	@read -p "type commit message: " cimsg; \
	git ci -am "$(shell date "+%F %T") $${cimsg}"
	git push origin master

.PHONY: clean build rpccert webcert


