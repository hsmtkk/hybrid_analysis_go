package response_test

import (
	"io/ioutil"
	"testing"

	"github.com/hsmtkk/hybrid_analysis_go/pkg/searchhash/response"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	js, err := ioutil.ReadFile("./example.json")
	assert.Nil(t, err, "should be nil")
	result, err := response.New().Parse(js)
	assert.Equal(t, 2, result.ThreatLevel, "should be equal")
	assert.Equal(t, 100, result.ThreatScore, "should be equal")
	assert.Equal(t, "malicious", result.Verdict, "should be equal")
}
