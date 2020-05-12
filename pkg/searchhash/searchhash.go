package searchhash

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/hsmtkk/hybrid_analysis_go/pkg"
	"github.com/hsmtkk/hybrid_analysis_go/pkg/searchhash/response"
)

type HashSearcher interface {
	SearchHash(hash string) (response.Result, error)
}

func New(apiKey string) HashSearcher {
	url := pkg.BaseURL + "/search/hash"
	return &hashSearcherImpl{client: http.DefaultClient, url: url, apiKey: apiKey}
}

func NewForTest(client *http.Client, url, apiKey string) HashSearcher {
	return &hashSearcherImpl{client: client, url: url, apiKey: apiKey}
}

type hashSearcherImpl struct {
	client *http.Client
	url    string
	apiKey string
}

func (rcv *hashSearcherImpl) SearchHash(hash string) (response.Result, error) {
	kvs := url.Values{}
	kvs.Set("hash", hash)
	req, err := http.NewRequest(http.MethodPost, rcv.url, strings.NewReader(kvs.Encode()))
	if err != nil {
		return response.Result{}, err
	}
	req.Header.Set(pkg.HeaderAPIKEY, rcv.apiKey)
	req.Header.Set(pkg.HeaderUserAgent, pkg.DefaultUserAgent)
	req.Header.Set(pkg.HeaderContentType, pkg.AppEncoded)
	req.Header.Set(pkg.HeaderAccept, pkg.AppJSON)
	resp, err := rcv.client.Do(req)
	if err != nil {
		return response.Result{}, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response.Result{}, err
	}
	result, err := response.New().Parse(respBody)
	if err != nil {
		return response.Result{}, err
	}
	return result, nil
}
