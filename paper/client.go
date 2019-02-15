package paper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type cursor struct {
	Value      string    `json:"value"`
	Expiration time.Time `json:"expiration"`
}

type listPaperDocsArgs struct {
	FilterBy  string `json:"filter_by"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
	Limit     string `json:"limit"`
}

type listPaperDocsResponse struct {
	DocIds  []string `json:"doc_ids"`
	Cursor  cursor   `json:"cursor"`
	HasMore bool     `json:"has_more"`
}

type paperDocExport struct {
	DocId        string `json:"doc_id"`
	ExportFormat string `json:"export_format"`
}

type paperDocExportResult struct {
	Owner    string `json:"owner"`
	Title    string `json:"title"`
	Revision int64  `json:"revision"`
	MimeType string `json:"mime_type"`
}

type client struct {
	AuthToken string
}

func (c client) ListDocuments() (listPaperDocsResponse, error) {
	resp := listPaperDocsResponse{}

	req, err := c.listPaperDocsRequest()
	if err != nil {
		return resp, err
	}

	err = c.sendRpcRequest(req, &resp)
	return resp, err
}

func (c client) DownloadDocument(id string) (paperDocExportResult, []byte, error) {
	resp := paperDocExportResult{}

	req, err := c.downloadPaperDocRequest(id)
	if err != nil {
		return resp, nil, err
	}

	body, err := c.sendContentDownloadRequest(req, &resp)
	return resp, body, err
}

func (c client) listPaperDocsRequest() (*http.Request, error) {
	args := listPaperDocsArgs{FilterBy: "docs_created"}

	bodyBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest("POST", "https://api.dropboxapi.com/2/paper/docs/list", bodyReader)
	if err != nil {
		return nil, err
	}

	authHeader := fmt.Sprintf("Bearer %s", c.AuthToken)
	req.Header.Set("Authorization", authHeader)

	return req, nil
}

func (c client) downloadPaperDocRequest(docId string) (*http.Request, error) {
	args := paperDocExport{
		DocId:        docId,
		ExportFormat: "markdown",
	}

	bodyBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.dropboxapi.com/2/paper/docs/download", nil)
	if err != nil {
		return nil, err
	}

	authHeader := fmt.Sprintf("Bearer %s", c.AuthToken)
	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Dropbox-API-Arg", string(bodyBytes))

	return req, nil
}

func (c client) sendRpcRequest(req *http.Request, v interface{}) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func (c client) sendContentDownloadRequest(req *http.Request, v interface{}) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	result := resp.Header.Get("Dropbox-API-Result")
	err = json.Unmarshal([]byte(result), v)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}
