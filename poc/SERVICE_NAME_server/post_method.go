package main

import (
	"fmt"
	"net/http"

	"github.com/alexgarzao/ms-gen/poc/SERVICE_NAME_common"

	"github.com/ant0ine/go-json-rest/rest"
)

func (s *Service) POST_METHOD_NAME(w rest.ResponseWriter, r *rest.Request) {
	request := SERVICE_NAME_common.POST_METHOD_NAMERequest{}
	if err := r.DecodeJsonPayload(&request); err != nil {
		rest.Error(
			w,
			fmt.Sprintf("When decoding json: %s", err.Error()),
			http.StatusInternalServerError,
		)
		return
	}

	// Verify if all necessary data are present in request.
	// TODO
	//	if XXX {
	//		rest.Error(
	//			w,
	//			"Some data not found: XXX",
	//			http.StatusBadRequest,
	//		)
	//		return
	//	}

	response := SERVICE_NAME_common.POST_METHOD_NAMEResponse{}

	// Business rules here :-)

	w.WriteJson(&response)
}
