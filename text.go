package harsp

import (
	"net/http"
)

type TEXT struct {
	Data string
}

// this func implements the rspWriter func Send.
// this Send a Json data to the client.
func (this TEXT) Send(w http.ResponseWriter) (http.ResponseWriter, error) {
	w.Header().Set("Content-Type", MimeText)

	_, err := w.Write([]byte(this.Data))
	if err != nil {
		return w, ErrWriteFailed
	}

	return w, nil
}
