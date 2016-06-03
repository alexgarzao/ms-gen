package main

import (
	"fmt"

	"net/http"

	"github.com/alexgarzao/ms-gen/poc/SERVICE_NAME_common"

	"github.com/ant0ine/go-json-rest/rest"
)

func (s *Service) DELETE_METHOD_NAME(w rest.ResponseWriter, r *rest.Request) {
	// par := r.PathParam("PARAMETER_NAME")
	_ = r.PathParam("PARAMETER_NAME")

	request := SERVICE_NAME_common.DELETE_METHOD_NAMERequest{}
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

	response := SERVICE_NAME_common.DELETE_METHOD_NAMEResponse{}

	w.WriteJson(response)
	w.WriteHeader(http.StatusNoContent) // 204
}
