package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/OPSWAT/mdcloud-go/pkg/utils"
	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
)

// RescanReq used for rescan body
type RescanReq struct {
	FileIDs []string `json:"file_ids"`
}

type ScanResp struct {
	DataID      string `json:"data_id"`
	ScanResults struct {
		ProgressPercentage int `json:"progress_percentage"`
	} `json:"scan_results"`
	Error ApiError `json:"error"`
}

// ScanFile sends to API
func (api *API) ScanFile(path string, headers []string, poll bool) (string, error) {
	url := fmt.Sprintf("%s/file", api.URL)
	file, err := os.Open(path)
	if err != nil {
		logrus.Errorln(err)
		return "", err
	}
	defer file.Close()
	req, err := http.NewRequest(http.MethodPost, url, file)
	if err != nil {
		logrus.Errorln(err)
		return "", err
	}
	req.Header.Add("apikey", api.Token)
	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("x-filename", filepath.Base(path))
	if len(headers) > 0 {
		for _, v := range headers {
			h := strings.Split(v, "=")
			req.Header.Add(h[0], h[1])
		}
	}
	resp, err := api.Client.Do(req)
	if err != nil {
		logrus.Errorln(err)
		return "", err
	}
	defer resp.Body.Close()
	var s = new(ScanResp)
	err = json.NewDecoder(resp.Body).Decode(&s)
	filterHeaders := func(s string) bool { return strings.HasPrefix(s, "X-") }
	rateLimits := utils.FilterMap(resp.Header, filterHeaders)
	api.Limits = rateLimits
	if err != nil {
		logrus.Errorln(err)
		return "", err
	}
	if resp.StatusCode != 200 {
		logrus.WithFields(logrus.Fields{"status_code": resp.StatusCode, "error_code": s.Error.Code, "error_message": strings.Join(s.Error.Messages, ",")}).Errorln(resp.Status)
		return "", nil
	}
	if !poll {
		logrus.WithField("data_id", s.DataID).Infoln("Result data_id")
		if r, e := json.Marshal(s); e == nil {
			return string(r), nil
		}
		logrus.Errorln(err)
		return "", err
	}
	var jsonResult string
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	for {
		select {
		case <-ticker.C:
			result := new(ScanResp)
			resDataID, err := api.ResultsByDataID(s.DataID)
			if err != nil {
				return "", errors.New("Failed to get results for: " + resDataID)
			}
			json.NewDecoder(strings.NewReader(resDataID)).Decode(&result)
			if result.ScanResults.ProgressPercentage == 100 {
				response, _ := json.Marshal(result)
				jsonResult = string(response)
				go func() { done <- true }()
			}
			logrus.WithFields(logrus.Fields{
				"data_id":  s.DataID,
				"progress": result.ScanResults.ProgressPercentage,
			}).Info("Scan progress")
		case <-done:
			close(done)
			return jsonResult, nil
		}
	}
}

// ResultsByDataID by data_id
func (api *API) ResultsByDataID(dataID string) (string, error) {
	url := fmt.Sprintf("%s/file/%s", api.URL, dataID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("apikey", api.Token)
	return fmtResponse(api.Client.Do(req))
}

// RescanFile by file_id
func (api *API) RescanFile(fileID string) (string, error) {
	url := fmt.Sprintf("%s/file/%s/rescan", api.URL, fileID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("apikey", api.Token)
	return fmtResponse(api.Client.Do(req))
}

// RescanFiles by file_ids
func (api *API) RescanFiles(fileIDs []string) (string, error) {
	url := fmt.Sprintf("%s/file/rescan", api.URL)
	payload := &RescanReq{FileIDs: fileIDs}
	j, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	if err != nil {
		return "", err
	}
	req.Header.Add("apikey", api.Token)
	req.Header.Add("content-type", "application/json")
	return fmtResponse(api.Client.Do(req))
}

// GetSanitizedLink Retrieve the download link for a sanitized file
func (api *API) GetSanitizedLink(fileID string) (string, error) {
	url := fmt.Sprintf("%s/file/%s/sanitizedLink", api.URL, fileID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("apikey", api.Token)
	return fmtResponse(api.Client.Do(req))
}

// FindOrScan file by hash
func (api *API) FindOrScan(path, hash string, headers []string, lookup, poll bool) (string, error) {
	if !lookup {
		return api.ScanFile(path, headers, poll)
	}
	result := new(ScanResp)
	strRes, err := api.HashDetails(hash)
	json.NewDecoder(strings.NewReader(strRes)).Decode(&result)
	if cmp.Equal(result.Error, ApiError{}) {
		logrus.WithField("hash", hash).Info("Hash not found sending to scan")
		return api.ScanFile(path, headers, poll)
	}
	return strRes, err
}
