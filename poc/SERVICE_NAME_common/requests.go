package SERVICE_NAME_common

import (
	"time"
)

type ServiceRequest1 struct {
	FieldName1 string    `json:"fieldName1"`
	FieldName2 string    `json:"fieldName2"`
	FieldName3 time.Time `json:"fieldName3"`
}

//type ServiceRequest2 struct {
//	FieldName1 string `json:"fieldName1"`
//	FieldName2 string `json:"fieldName2"`
//	FieldName3 time.Time `json:"fieldName3"`
//}
//
//...
//
//type ServiceRequestN struct {
//	FieldName1 string `json:"fieldName1"`
//	FieldName2 string `json:"fieldName2"`
//	FieldName3 time.Time `json:"fieldName3"`
//}

//type GET_METHOD_NAMERequest struct {
//	FieldName1 string    `json:"fieldName1"`
//	FieldName2 time.Time `json:"fieldName2"`
//}

type GET_METHOD_NAMEResponse struct {
	FieldName3 uint    `json:"fieldName3"`
	FieldName4 float64 `json:"fieldName4"`
}

type PUT_METHOD_NAMERequest struct {
	FieldName1 string    `json:"fieldName1"`
	FieldName2 time.Time `json:"fieldName2"`
}

type PUT_METHOD_NAMEResponse struct {
	FieldName3 uint    `json:"fieldName3"`
	FieldName4 float64 `json:"fieldName4"`
}

type POST_METHOD_NAMERequest struct {
	FieldName1 string    `json:"fieldName1"`
	FieldName2 time.Time `json:"fieldName2"`
}

type POST_METHOD_NAMEResponse struct {
	FieldName3 uint    `json:"fieldName3"`
	FieldName4 float64 `json:"fieldName4"`
}

type DELETE_METHOD_NAMERequest struct {
	FieldName1 string    `json:"fieldName1"`
	FieldName2 time.Time `json:"fieldName2"`
}

type DELETE_METHOD_NAMEResponse struct {
	FieldName3 uint    `json:"fieldName3"`
	FieldName4 float64 `json:"fieldName4"`
}
