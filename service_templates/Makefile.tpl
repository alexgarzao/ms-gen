.PHONY: build doc fmt lint run onlyrun test vet
{{ $service_name := .ServiceName }}
EXECUTABLE={{$service_name}}-server
LOG_FILE=/var/log/${EXECUTABLE}.log
GOFMT=gofmt -w
GODEPS=go get
GOBUILD=go build -v
GOFILES=\
	{{$service_name}}_server/main.go\
	{{$service_name}}_server/service.go\
	{{ range $value := .Paths }}{{$service_name}}_server/{{$value.CodeFilename}}.go\
	{{ end }}{{$service_name}}_server/db.go\
	{{$service_name}}_server/db_models.go


default: build


build: vet
	${GOBUILD} -o ./bin/${EXECUTABLE} ${GOFILES}


doc:
#	TODO


fmt:
	${GOFMT} ${GOFILES}


lint:
	golint ./{{$service_name}}_server


run: build
	./bin/${EXECUTABLE}


onlyrun:
	go run ${GOFILES}


test:
	go test ./{{$service_name}}_server/...


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
	go vet ./{{$service_name}}_server/...
	

#stop:
#	pkill -f ${EXECUTABLE}


#start:
#	-make stop
#	nohup ${EXECUTABLE} >> ${LOG_FILE} 2>&1 </dev/null &


#showprocess:
#	ps aux | grep ${EXECUTABLE}
