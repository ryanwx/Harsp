package harsp

import (
	"encoding/json"
	"net/http"
)

type JSON struct {
	Rc RetCode

	Data interface{}
}

// this func implements the rspWriter func Send.
// this Send a Json data to the client.
func (this JSON) send(w http.ResponseWriter, data map[string]interface{}) (http.ResponseWriter, error) {
	w.Header().Set("Content-Type", MimeJson)
	content, err := json.Marshal(data)
	if err != nil {
		return w, ErrWriteFailed
	}

	_, err = w.Write(content)
	if err != nil {
		return w, ErrWriteFailed
	}

	return w, nil
}

// this implements the rspWriter func Success.
// this Send a Success response to the client with the default success Json data format or
// you set the default success Json data format.
func (this JSON) Success(w http.ResponseWriter) (http.ResponseWriter, error) {
	packedData := DefaultPadding(SuccessRet, this.Data)

	return this.send(w, packedData)
}

// this implements the rspWriter func Failed.
// this func send a json data to the client with the errCode.
func (this JSON) Failed(w http.ResponseWriter) (http.ResponseWriter, error) {
	packedData := DefaultPadding(this.Rc, this.Data)

	return this.send(w, packedData)
}
