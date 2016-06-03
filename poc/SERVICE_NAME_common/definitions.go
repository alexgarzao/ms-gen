package SERVICE_NAME_common

import (
	"fmt"
)

type WSServiceError struct {
	Message string `json:"error"`
}

func (e WSServiceError) Error() string {
	return fmt.Sprintf("PROC: %v", e.Message)
}
