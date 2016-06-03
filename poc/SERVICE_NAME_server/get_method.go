package main

import (
	"github.com/alexgarzao/ms-gen/poc/SERVICE_NAME_common"

	"github.com/ant0ine/go-json-rest/rest"
)

func (s *Service) GET_METHOD_NAME(w rest.ResponseWriter, r *rest.Request) {
	// par := r.PathParam("PARAMETER_NAME")
	_ = r.PathParam("PARAMETER_NAME")

	response := SERVICE_NAME_common.GET_METHOD_NAMEResponse{}

	// Business rules here :-)

	w.WriteJson(response)
}
