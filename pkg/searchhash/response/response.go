package response

import (
	"encoding/json"
	"fmt"
)

type Parser interface {
	Parse(json []byte) (Result, error)
}

func New() Parser {
	return &parserImpl{}
}

type Result struct {
	ThreatLevel int
	ThreatScore int
	Verdict     string
}

type parserImpl struct{}

type responseSchema struct {
	ThreatLevel int    `json:"threat_level"`
	ThreatScore int    `json:"threat_score"`
	Verdict     string `json:"verdict"`
}

func (impl *parserImpl) Parse(js []byte) (Result, error) {
	schs := []responseSchema{}
	if err := json.Unmarshal(js, &schs); err != nil {
		return Result{}, err
	}
	if len(schs) == 0 {
		return Result{}, fmt.Errorf("search failure")
	}
	sch := schs[0]
	return Result{ThreatLevel: sch.ThreatLevel, ThreatScore: sch.ThreatScore, Verdict: sch.Verdict}, nil
}
