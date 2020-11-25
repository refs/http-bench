package service

import "net/http"

// Dummy is a simple, dummy handler that returns a basic json.
func Dummy() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("{'r-u-stoopid': 'yes'}")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}