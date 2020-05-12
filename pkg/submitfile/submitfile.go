package submitfile

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"

	"github.com/hsmtkk/hybrid_analysis_go/pkg"
	"github.com/hsmtkk/hybrid_analysis_go/pkg/submitfile/response"
)

type FileSubmitter interface {
	SubmitFile(filePath string, env Environment) (response.Result, error)
}

func New(apiKey string) FileSubmitter {
	url := pkg.BaseURL + "/submit/file"
	return &fileSubmitterImpl{client: http.DefaultClient, url: url, apiKey: apiKey}
}

func NewForTest(client *http.Client, url, apiKey string) FileSubmitter {
	return &fileSubmitterImpl{client: client, url: url, apiKey: apiKey}
}

type fileSubmitterImpl struct {
	client *http.Client
	url    string
	apiKey string
}

func (rcv *fileSubmitterImpl) SubmitFile(filePath string, env Environment) (response.Result, error) {
	envStr, err := env.String()
	if err != nil {
		return response.Result{}, err
	}
	body, contentType, err := makeFormData(filePath, envStr)
	if err != nil {
		return response.Result{}, err
	}
	req, err := http.NewRequest(http.MethodPost, rcv.url, bytes.NewBuffer(body))
	if err != nil {
		return response.Result{}, err
	}
	req.Header.Set(pkg.HeaderAPIKEY, rcv.apiKey)
	req.Header.Set(pkg.HeaderUserAgent, pkg.DefaultUserAgent)
	req.Header.Set(pkg.HeaderContentType, contentType)
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

func makeFormData(filePath string, envID string) ([]byte, string, error) {
	var buf bytes.Buffer
	mimeWriter := multipart.NewWriter(&buf)

	// file part
	fileName := path.Base(filePath)
	reader, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	fileWriter, err := mimeWriter.CreateFormFile("file", fileName)
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(fileWriter, reader)
	if err != nil {
		return nil, "", err
	}

	// environment_id part
	envWriter, err := mimeWriter.CreateFormField("environment_id")
	if err != nil {
		return nil, "", err
	}
	_, err = envWriter.Write([]byte(envID))
	if err != nil {
		return nil, "", err
	}

	mimeWriter.Close()
	return buf.Bytes(), mimeWriter.FormDataContentType(), nil
}
