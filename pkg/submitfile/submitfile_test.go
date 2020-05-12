package submitfile_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hsmtkk/hybrid_analysis_go/pkg/apikey"
	"github.com/hsmtkk/hybrid_analysis_go/pkg/submitfile"
	"github.com/stretchr/testify/assert"
)

func TestReal(t *testing.T) {
	apiKey, err := apikey.New().LoadAPIKey()
	assert.Nil(t, err, "should be nil")
	result, err := submitfile.New(apiKey).SubmitFile("./submitfile_test.go", submitfile.WindowsSeven64)
	assert.Nil(t, err, "should be nil")
	fmt.Println(result)
}

func TestLocal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reader, err := os.Open("./response/example.json")
		assert.Nil(t, err, "should be nil")
		written, err := io.Copy(w, reader)
		assert.Nil(t, err, "should be nil")
		assert.Greater(t, written, int64(0), "should be greater than zero")
	}))
	defer ts.Close()

	apiKey, err := apikey.New().LoadAPIKey()
	assert.Nil(t, err, "should be nil")
	result, err := submitfile.NewForTest(ts.Client(), ts.URL, apiKey).SubmitFile("./submitfile_test.go", submitfile.WindowsSeven64)
	assert.Nil(t, err, "should be nil")
	assert.NotEmpty(t, result.JobID, "should not be empty")
	assert.NotEmpty(t, result.SubmissionID, "should not be empty")
	assert.NotEmpty(t, result.SHA256, "should not be empty")
}
