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

	"github.com/OPSWAT/mdcloud-go/utils"
	"github.com/google/go-cmp/cmp"
	logger "github.com/sirupsen/logrus"
)

// RescanReq used for rescan body
type RescanReq struct {
	FileIDs []string `json:"file_ids"`
}

// ScanResponse used for scan body
type ScanResponse struct {
	DataID        string   `json:"data_id"`
	Status        string   `json:"status"`
	InQueue       int      `json:"in_queue"`
	QueuePriority string   `json:"queue_priority"`
	Error         ApiError `json:"error"`
}

// EngineResult defines particular engine result details
type EngineResult struct {
	WaitTime    int       `json:"wait_time"`
	ThreatFound string    `json:"threat_found"`
	ScanTime    int       `json:"scan_time"`
	ScanResultI int       `json:"scan_result_i"`
	DefTime     time.Time `json:"def_time"`
}

type FileInfo struct {
	FileSize            int       `json:"file_size"`
	UploadTimestamp     time.Time `json:"upload_timestamp"`
	Md5                 string    `json:"md5"`
	Sha1                string    `json:"sha1"`
	Sha256              string    `json:"sha256"`
	FileTypeCategory    string    `json:"file_type_category"`
	FileTypeDescription string    `json:"file_type_description"`
	FileTypeExtension   string    `json:"file_type_extension"`
	DisplayName         string    `json:"display_name"`
}
type ExtractedFiles struct {
	DataID         string `json:"data_id"`
	FilesInArchive []struct {
		ScanResultI        int    `json:"scan_result_i"`
		ProgressPercentage int    `json:"progress_percentage"`
		FileType           string `json:"file_type"`
		FileSize           int    `json:"file_size"`
		DisplayName        string `json:"display_name"`
		DetectedBy         int    `json:"detected_by"`
		DataID             string `json:"data_id"`
	} `json:"files_in_archive"`
}

// ScanResult used for polling result
type ScanResult struct {
	DataID                  string         `json:"data_id"`
	Archived                bool           `json:"archived"`
	ProcessInfo             ProcessInfo    `json:"process_info"`
	ExtractedFiles          ExtractedFiles `json:"extracted_files"`
	ScanResultHistoryLength int            `json:"scan_result_history_length"`
	Votes                   Votes          `json:"votes"`
	ScanResults             struct {
		ScanDetails        map[string]EngineResult `json:"scan_details"`
		RescanAvailable    bool                    `json:"rescan_available"`
		DataID             string                  `json:"data_id"`
		ScanAllResultI     int                     `json:"scan_all_result_i"`
		StartTime          time.Time               `json:"start_time"`
		TotalTime          int                     `json:"total_time"`
		TotalAvs           int                     `json:"total_avs"`
		TotalDetectedAvs   int                     `json:"total_detected_avs"`
		ProgressPercentage int                     `json:"progress_percentage"`
		InQueue            int                     `json:"in_queue"`
		ScanAllResultA     string                  `json:"scan_all_result_a"`
	} `json:"scan_results"`
	FileInfo    FileInfo `json:"file_info"`
	HashResults struct {
		Wa bool `json:"wa"`
	} `json:"hash_results"`
	TopThreat   int    `json:"top_threat"`
	ShareFile   int    `json:"share_file"`
	RestVersion string `json:"rest_version"`
}

// ScanFile sends to API
func (api *API) ScanFile(path string, headers []string, poll bool) (string, error) {
	url := fmt.Sprintf("%s/file", api.URL)
	file, err := os.Open(path)
	if err != nil {
		logger.Errorln(err)
		return "", err
	}
	defer file.Close()
	req, err := http.NewRequest(http.MethodPost, url, file)
	if err != nil {
		logger.Errorln(err)
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
		logger.Errorln(err)
		return "", err
	}
	var s = new(ScanResponse)
	err = json.NewDecoder(resp.Body).Decode(&s)
	filterHeaders := func(s string) bool { return strings.HasPrefix(s, "X-") }
	rateLimits := utils.FilterMap(resp.Header, filterHeaders)
	api.Limits = rateLimits
	defer resp.Body.Close()
	if err != nil {
		logger.Errorln(err)
		return "", err
	}
	if resp.StatusCode != 200 {
		logger.WithFields(logger.Fields{"status_code": resp.StatusCode, "error_code": s.Error.Code, "error_message": strings.Join(s.Error.Messages, ",")}).Errorln(resp.Status)
		return "", nil
	}
	if !poll {
		logger.WithField("data_id", s.DataID).Infoln("Result data_id")
		if r, e := json.Marshal(s); e == nil {
			return string(r), nil
		}
		logger.Errorln(err)
		return "", err
	}
	var jsonResult string
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	for {
		select {
		case <-ticker.C:
			result := new(ScanResult)
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
			logger.WithFields(logger.Fields{
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
	result := new(HashLookupResp)
	strRes, err := api.HashDetails(hash)
	json.NewDecoder(strings.NewReader(strRes)).Decode(&result)
	if cmp.Equal(result.Error, ApiError{}) {
		logger.WithField("hash", hash).Info("Hash not found sending to scan")
		return api.ScanFile(path, headers, poll)
	}
	return strRes, err
}
