package harsp

import (
	"net/http"
)

type HTML struct {
	Data string
}

// this func implements the rspWriter func Send.
// this Send a Json data to the client.
func (this HTML) Send(w http.ResponseWriter) (http.ResponseWriter, error) {
	w.Header().Set("Content-Type", MimeHtml)

	_, err := w.Write([]byte(this.Data))
	if err != nil {
		return w, ErrWriteFailed
	}

	return w, nil
}
