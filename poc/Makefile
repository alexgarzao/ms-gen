.PHONY: build_server build_bdd doc fmt lint run onlyrun test bdd vet

SERVER_EXECUTABLE=SERVICE_NAME-server
TEST_EXECUTABLE=SERVICE_NAME-test

LOG_FILE=/var/log/${SERVER_EXECUTABLE}.log
GOFMT=gofmt -w
GODEPS=go get
GOBUILD=go build -v

GOSERVERFILES=\
	SERVICE_NAME_server/main.go\
	SERVICE_NAME_server/db.go\
	SERVICE_NAME_server/db_models.go\
	SERVICE_NAME_server/service.go\
	SERVICE_NAME_server/request_validation.go\
	SERVICE_NAME_server/get_method.go\
	SERVICE_NAME_server/put_method.go\
	SERVICE_NAME_server/post_method.go\
	SERVICE_NAME_server/delete_method.go

GOTESTFILES=\
	SERVICE_NAME_test/tests.go\
	SERVICE_NAME_test/test_requests.go\
	SERVICE_NAME_test/test_behaviour_1.go

default: build_server


build_server: vet
	${GOBUILD} -o ./bin/${SERVER_EXECUTABLE} ${GOSERVERFILES}

build_bdd: vet
	${GOBUILD} -o ./bin/${TEST_EXECUTABLE} ${GOTESTFILES}


doc:
#	TODO


fmt:
	${GOFMT} ${GOSERVERFILES} ${GOTESTFILES}


lint:
	golint ./SERVICE_NAME_server


run: build_server
	./bin/${SERVER_EXECUTABLE}


onlyrun:
	go run ${GOSERVERFILES}


test:
	go test ./SERVICE_NAME_server/...


bdd: build_bdd
	./bin/${TEST_EXECUTABLE}


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
