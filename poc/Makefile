.PHONY: build doc fmt lint run onlyrun test vet

EXECUTABLE=SERVICE_NAME-server
LOG_FILE=/var/log/${EXECUTABLE}.log
GOFMT=gofmt -w
GODEPS=go get
GOBUILD=go build -v
GOFILES=\
	SERVICE_NAME_server/main.go\
	SERVICE_NAME_server/db.go\
	SERVICE_NAME_server/db_models.go\
	SERVICE_NAME_server/service.go\
	SERVICE_NAME_server/get_method.go\
	SERVICE_NAME_server/put_method.go\
	SERVICE_NAME_server/post_method.go\
	SERVICE_NAME_server/delete_method.go


default: build


build: vet
	${GOBUILD} -o ./bin/${EXECUTABLE} ${GOFILES}


doc:
#	TODO


fmt:
	${GOFMT} ${GOFILES}


lint:
	golint ./SERVICE_NAME_server


run: build
	./bin/${EXECUTABLE}


onlyrun:
	go run ${GOFILES}


test:
	go test ./SERVICE_NAME_server/...


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
	go vet ./SERVICE_NAME_server/...
	

#stop:
#	pkill -f ${EXECUTABLE}


#start:
#	-make stop
#	nohup ${EXECUTABLE} >> ${LOG_FILE} 2>&1 </dev/null &


#showprocess:
#	ps aux | grep ${EXECUTABLE}
