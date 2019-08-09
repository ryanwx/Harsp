package harsp

import (
	"encoding/json"
	"net/http"
)

type JSONP struct {
	Rc RetCode

	Data interface{}

	// when you use jsonp data response, you must be set a callback func.
	Callback string
}

// send data to http response.
// it returns the writer, then you can use the writer continue to write.
// @param: w http.ResponseWriter, all the data will use the writer write into the response
// @param: data interface, this is the response content
// @return： w http.ResponseWriter
func (this JSONP) send(w http.ResponseWriter, data map[string]interface{}) (http.ResponseWriter, error) {
	w.Header().Set("Content-Type", MimeJsonP)
	content, err := json.Marshal(data)
	if err != nil {
		return w, ErrWriteFailed
	}

	content = append([]byte(this.Callback+"("), content...)
	content = append(content, []byte(")")...)

	_, err = w.Write(content)
	if err != nil {
		return w, ErrWriteFailed
	}

	return w, nil
}

// when you want to return a success msg to the http client, you can use this func.
// this func use the default errCode string 0, if you want change the default value,
// your must to set the package variable of SuccessCode to a new value,
// you can do it when you start your application, once you do that, it will always affected.
// @param: w http.ResponseWriter, all the data will use the writer write into the response
// @param: data interface, this is the response content
// @return： w http.ResponseWriter
func (this JSONP) Success(w http.ResponseWriter) (http.ResponseWriter, error) {
	packedData := DefaultPadding(SuccessRet, this.Data)

	return this.send(w, packedData)
}

// when you want to return a failed msg to the http client, you can use this func.
// you can define your errCode in your application, when you use this func to send a failed response,
// you can assign which errCode you wish to return.
// @param: w http.ResponseWriter, all the data will use the writer write into the response
// @param: data interface, this is the response content
// @return： w http.ResponseWriter
func (this JSONP) Failed(w http.ResponseWriter) (http.ResponseWriter, error) {
	packedData := DefaultPadding(this.Rc, this.Data)

	return this.send(w, packedData)
}
