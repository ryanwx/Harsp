package rsp

import (
    "errors"
    "fmt"
    "net/http"
)

const (
    MimeText = "application/text; charset=utf-8"
    MimeXML = "application/xml; charset=utf-8"
    MimeHtml = "application/html; charset=utf-8"
    MimeJson = "application/json; charset=utf-8"
)

var (
    SuccessCode = "0"
)

var (
    ErrWriteFailed = errors.New("failed to write message to http response")
)

var DefaultPadding func(errCode string, d interface{}) interface{} = defaultContentPadding

type rsp struct {
    data interface{}

    statusCode int

    errCode string

    header http.Header
}

type rspWriter interface {
    Send(http.ResponseWriter, interface{}) (http.ResponseWriter, error)

    Success(http.ResponseWriter, interface{}) (http.ResponseWriter, error)
}

func errStatus(w http.ResponseWriter, code int, msg string) error {
    w.WriteHeader(code)
    _, err := fmt.Fprintln(w, msg)
    if err != nil{
        return err
    }

    return nil
}

func NotFound(w http.ResponseWriter, msg string) error{
    return errStatus(w, http.StatusNotFound, msg)
}

func Forbidden(w http.ResponseWriter, msg string) error{
    return errStatus(w, http.StatusForbidden, msg)
}

func UnAuthorized(w http.ResponseWriter, msg string) error {
    return errStatus(w, http.StatusUnauthorized, msg)
}

func defaultContentPadding(errCode string, d interface{}) interface{} {
    return map[string]interface{}{
        "code": errCode,
        "data": d,
    }
}

