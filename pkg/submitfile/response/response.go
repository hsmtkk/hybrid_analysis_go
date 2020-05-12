package response

import (
	"encoding/json"
)

type Result struct {
	JobID        string
	SubmissionID string
	SHA256       string
}

type Parser interface {
	Parse(js []byte) (Result, error)
}

func New() Parser {
	return &parserImpl{}
}

type parserImpl struct{}

type responseSchema struct {
	JobID         string `json:"job_id"`
	SubmissionID  string `json:"submission_id"`
	EnvironmentID int    `json:"environment_id"`
	SHA256        string `json:"sha256"`
}

func (impl *parserImpl) Parse(js []byte) (Result, error) {
	sch := responseSchema{}
	if err := json.Unmarshal(js, &sch); err != nil {
		return Result{}, err
	}
	return Result{JobID: sch.JobID, SubmissionID: sch.SubmissionID, SHA256: sch.SHA256}, nil
}
