package searchhash_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hsmtkk/hybrid_analysis_go/pkg/apikey"
	"github.com/hsmtkk/hybrid_analysis_go/pkg/searchhash"
	"github.com/stretchr/testify/assert"
)

func TestReal(t *testing.T) {
	apiKey, err := apikey.New().LoadAPIKey()
	assert.Nil(t, err, "should be nil")
	hash := "f4d76f4ad2977077b00035901b614d04a1fd5e5dec9d22309279304c8da56865"
	result, err := searchhash.New(apiKey).SearchHash(hash)
	assert.Nil(t, err, "should be nil")
	assert.NotNil(t, result.ThreatLevel, "shout not be nil")
	assert.NotNil(t, result.ThreatScore, "shout not be nil")
	assert.NotEmpty(t, result.Verdict, "shout not be empty")
}

func TestLocal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bs, err := ioutil.ReadFile("./response/example.json")
		assert.Nil(t, err, "should be nil")
		written, err := io.Copy(w, bytes.NewBuffer(bs))
		assert.Nil(t, err, "should be nil")
		assert.Greater(t, written, int64(0), "should be greater than zero")
	}))
	defer ts.Close()

	hash := "f4d76f4ad2977077b00035901b614d04a1fd5e5dec9d22309279304c8da56865"
	result, err := searchhash.NewForTest(ts.Client(), ts.URL, "apiKey").SearchHash(hash)
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, 2, result.ThreatLevel, "should be equal")
	assert.Equal(t, 100, result.ThreatScore, "should be equal")
	assert.Equal(t, "malicious", result.Verdict, "should be equal")
}
