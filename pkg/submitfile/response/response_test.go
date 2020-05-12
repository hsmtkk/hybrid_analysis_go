package response_test

import (
	"io/ioutil"
	"testing"

	"github.com/hsmtkk/hybrid_analysis_go/pkg/submitfile/response"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	js, err := ioutil.ReadFile("./example.json")
	assert.Nil(t, err, "should be nil")
	result, err := response.New().Parse(js)
	assert.Nil(t, err, "should be nil")
	assert.Equal(t, "5eb9e1cf55ea3b4998153586", result.JobID, "should be equal")
	assert.Equal(t, "5eba443ffe581326e15352ca", result.SubmissionID, "should be equal")
	assert.Equal(t, "f4d76f4ad2977077b00035901b614d04a1fd5e5dec9d22309279304c8da56865", result.SHA256, "should be equal")
}
