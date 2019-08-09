package harsp

import (
	"errors"
	"fmt"
	"net/http"
)

// TODO: add support of more charset.
const (
	MimeText  = "text/plain; charset=utf-8"
	MimeXML   = "text/xml; charset=utf-8"
	MimeHtml  = "text/html; charset=utf-8"
	MimeJson  = "application/json; charset=utf-8"
	MimeJsonP = "text/plain; charset=utf-8"
)

// all the response errCode and msg must be the RetCode struct.
type RetCode struct {
	// return error code
	// use this code to show this request result.
	Code string

	// return error message
	// this is request result message.
	Msg  string
}

var (
	// this is default success code, you also can reset the variable what you want be, but this must be RetCode struct.
	SuccessRet = RetCode{Code: "0", Msg: "Success"}
)

var (
	// when Harsp write the byte to the response, when it failed, this error will return in any response func.
	ErrWriteFailed = errors.New("failed to write message to http response")
)

// default response content padding func.
// if you want define your content padding format, you can set variable with your padding func.
// only JSON|XML|JSONP will use the padding func
var DefaultPadding func(RetCode, interface{}) map[string]interface{} = defaultContentPadding

type TextRspWriter interface {
	// send a text data, which contains text|html.
	// it returns the writer, then you can use the writer continue to write.
	// @param: w http.ResponseWriter, all the data will use the writer write into the response
	// @param: data string, this is the response content
	// @return： w http.ResponseWriter
	Send(http.ResponseWriter) (http.ResponseWriter, error)
}

type ApiRspWriter interface {

	// send data to http response.
	// it returns the writer, then you can use the writer continue to write.
	// @param: w http.ResponseWriter, all the data will use the writer write into the response
	// @param: data interface, this is the response content
	// @return： w http.ResponseWriter
	send(http.ResponseWriter, map[string]interface{}) (http.ResponseWriter, error)

	// when you want to return a success msg to the http client, you can use this func.
	// this func use the default errCode string 0, if you want change the default value,
	// your must to set the package variable of SuccessCode to a new value,
	// you can do it when you start your application, once you do that, it will always affected.
	// @param: w http.ResponseWriter, all the data will use the writer write into the response
	// @param: data interface, this is the response content
	// @return： w http.ResponseWriter
	Success(http.ResponseWriter) (http.ResponseWriter, error)

	// when you want to return a failed msg to the http client, you can use this func.
	// you can define your errCode in your application, when you use this func to send a failed response,
	// you can assign which errCode you wish to return.
	// @param: w http.ResponseWriter, all the data will use the writer write into the response
	// @param: data interface, this is the response content
	// @return： w http.ResponseWriter
	Failed(http.ResponseWriter) (http.ResponseWriter, error)
}

// this func send a error http status code to the client.
// All the http error response will use this write error message.
// all the http error response header of Content-Type wile be set "text/plain; charset=utf-8".
// @param: w http.ResponseWriter
// @param: code int, witch is the http response status code.
// @param: msg string, this is the error http response content of plain
func errStatus(w http.ResponseWriter, code int, msg string) error {
	w.WriteHeader(code)
	_, err := fmt.Fprintln(w, msg)
	if err != nil {
		return err
	}

	return nil
}

// TODO: add the 301 302 response.

// this func return a http response with status code 400
// this means the http client request is Invalid, it's usually point the request param.
func BadRequest(w http.ResponseWriter, msg string) error {
	return errStatus(w, http.StatusBadRequest, msg)
}

// this func return a http response with status code 401
// this means the client is unauthorized, like UnLogin.
func UnAuthorized(w http.ResponseWriter, msg string) error {
	return errStatus(w, http.StatusUnauthorized, msg)
}

// this func return a http response with status code 404.
// witch means the resource what the client want is don't exist.
func NotFound(w http.ResponseWriter, msg string) error {
	return errStatus(w, http.StatusNotFound, msg)
}

// this func return a http response with status code 403.
// which means the client don't has the power to get the resource.
func Forbidden(w http.ResponseWriter, msg string) error {
	return errStatus(w, http.StatusForbidden, msg)
}

// this func return a http response with status code 406
// this means the http request method is not allowed, witch usually used in the application of RESTful.
func MethodNotAllowed(w http.ResponseWriter, msg string) error {
	return errStatus(w, http.StatusMethodNotAllowed, msg)
}

// this func return a http response with status code 408
// this means the http request is timeout
func RequestTimeout(w http.ResponseWriter, msg string) error {
	return errStatus(w, http.StatusRequestTimeout, msg)
}

// this is the default padding func with pad the response data.
// when you want send data with JSON|XML|jsonp, the data will be padding by use this default func.
// you can set the DefaultPadding value, to replace this default padding func.
func defaultContentPadding(rc RetCode, d interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"code": rc.Code,
		"msg":  rc.Msg,
	}

	if nil != d {
		data["data"] = d
	}

	return data
}
