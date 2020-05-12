package apikey

import (
	"fmt"
	"os"
)

const HybridAnalysisAPIKEY = "HYBRID_ANALYSIS_API_KEY"

type APIKeyLoader interface {
	LoadAPIKey() (string, error)
}

func New() APIKeyLoader {
	return &apiKeyLoaderImpl{}
}

type apiKeyLoaderImpl struct{}

func (imp *apiKeyLoaderImpl) LoadAPIKey() (string, error) {
	val := os.Getenv(HybridAnalysisAPIKEY)
	if val == "" {
		return "", fmt.Errorf("environment variable %s is not defined", HybridAnalysisAPIKEY)
	}
	return val, nil
}
