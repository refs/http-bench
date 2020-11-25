package service

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkDummy(b *testing.B) {
	server := httptest.NewServer(Dummy())
	defer server.Close()

	for i:=0; i< b.N; i++ {
		res, err := http.Get(server.URL)
		if err != nil {
			log.Fatal(err)
		}
		res.Body.Close()
		assert.Equal(b, 200, res.StatusCode)
	}
}