.PHONY: build_server build_bdd doc fmt lint run onlyrun test bdd vet
{{ $service_name := .ServiceName }}
SERVER_EXECUTABLE={{$service_name}}-server
BDD_EXECUTABLE={{$service_name}}-test

LOG_FILE=/var/log/${SERVER_EXECUTABLE}.log
GOFMT=gofmt -w
GODEPS=go get
GOBUILD=go build -v

GOSERVERFILES=\
	{{$service_name}}_server/main.go\
	{{$service_name}}_server/service.go\
	{{$service_name}}_server/request_validation.go\
	{{ range $value := .Methods }}{{$service_name}}_server/{{$value.CodeFilename}}\
	{{ end }}{{$service_name}}_server/db.go\
	{{$service_name}}_server/db_models.go

GOTESTFILES=\
	{{$service_name}}_test/tests.go\
	{{ range $value := .Methods }}{{$service_name}}_test/test_{{$value.CodeFilename}}\
	{{ end }}{{$service_name}}_test/test_requests.go

GOCOMMONFILES=\
	{{$service_name}}_common/definitions.go\
	{{$service_name}}_common/requests.go\
	{{$service_name}}_common/utils.go


default: build


build: build_server build_bdd


build_server: vet
	${GOBUILD} -o ./bin/${SERVER_EXECUTABLE} ${GOSERVERFILES}


build_bdd:
	${GOBUILD} -o ./bin/${BDD_EXECUTABLE} ${GOTESTFILES}


doc:
#	TODO


fmt:
	${GOFMT} ${GOSERVERFILES} ${GOTESTFILES} ${GOCOMMONFILES}


lint:
	golint ./{{$service_name}}_server ./{{$service_name}}_common


run: build_server
	./bin/${SERVER_EXECUTABLE}


onlyrun:
	go run ${GOSERVERFILES}


test:
	go test ./{{$service_name}}_server/...


bdd: build_bdd
	./bin/${BDD_EXECUTABLE}


install:
	go install


deps:
	${GODEPS} github.com/jinzhu/gorm
	${GODEPS} github.com/lib/pq
	${GODEPS} github.com/spf13/viper
	${GODEPS} github.com/ant0ine/go-json-rest/rest
	${GODEPS} github.com/golang/lint/golint


update:
	git pull
	make deps
	make install


vet:
	go vet ./{{$service_name}}_server/... ./{{$service_name}}_common/...
	

#stop:
#	pkill -f ${EXECUTABLE}


#start:
#	-make stop
#	nohup ${EXECUTABLE} >> ${LOG_FILE} 2>&1 </dev/null &


#showprocess:
#	ps aux | grep ${EXECUTABLE}
