package handlers

import (
	"net/http"

	"github.com/justinpjose/cushon-assignment/internal/logging"
)

// WriteRsp writes the given string or []byte to the response body
func WriteRsp[M []byte | string](w http.ResponseWriter, log logging.Logger, msg M) {
	var err error

	switch v := any(msg).(type) {
	case string:
		_, err = w.Write([]byte(v))
	case []byte:
		_, err = w.Write(v)
	}

	if err != nil {
		log.Errorf("failed to write response: %v", err)
	}
}
