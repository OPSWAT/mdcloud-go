package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// RescanReq used for rescan body
type RescanReq struct {
	FileIDs []string `json:"file_ids"`
}

// ScanResponse used for scan body
type ScanResponse struct {
	Success bool `json:"success"`
	Data    struct {
		DataID        string `json:"data_id"`
		Status        string `json:"status"`
		InQueue       int    `json:"in_queue"`
		QueuePriority string `json:"queue_priority"`
	} `json:"data"`
}

// DataIDResponse used for polling result
type DataIDResponse struct {
	Success bool `json:"success"`
	Data    struct {
		DataID      string `json:"data_id"`
		Archived    bool   `json:"archived"`
		ScanResults struct {
			ScanDetails struct {
			} `json:"scan_details"`
			RescanAvailable    bool   `json:"rescan_available"`
			DataID             string `json:"data_id"`
			ScanAllResultI     int    `json:"scan_all_result_i"`
			StartTime          string `json:"start_time"`
			TotalTime          int    `json:"total_time"`
			TotalAvs           int    `json:"total_avs"`
			TotalDetectedAvs   int    `json:"total_detected_avs"`
			ProgressPercentage int    `json:"progress_percentage"`
			InQueue            int    `json:"in_queue"`
			ScanAllResultA     string `json:"scan_all_result_a"`
		} `json:"scan_results"`
		FileInfo struct {
			FileSize            int    `json:"file_size"`
			UploadTimestamp     string `json:"upload_timestamp"`
			Md5                 string `json:"md5"`
			Sha1                string `json:"sha1"`
			Sha256              string `json:"sha256"`
			FileTypeCategory    string `json:"file_type_category"`
			FileTypeDescription string `json:"file_type_description"`
			FileTypeExtension   string `json:"file_type_extension"`
			DisplayName         string `json:"display_name"`
		} `json:"file_info"`
		TopThreat   int    `json:"top_threat"`
		ShareFile   int    `json:"share_file"`
		RestVersion string `json:"rest_version"`
	} `json:"data"`
}

// ScanFile sends to API
func (api *API) ScanFile(path string, headers []string) (string, error) {
	url := fmt.Sprintf("%s/file", api.URL)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	req, err := http.NewRequest("POST", url, file)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "apikey "+api.Token)
	req.Header.Add("Content-Type", "binary/octet-stream")
	req.Header.Add("x-filename", filepath.Base(path))
	if len(headers) > 0 {
		for _, v := range headers {
			h := strings.Split(v, "=")
			req.Header.Add(h[0], h[1])
		}
	}
	resp, err := api.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	var s = new(ScanResponse)
	err = json.NewDecoder(resp.Body).Decode(&s)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	if s.Success == false {
		log.Fatal(resp.Status)
	}

	var jsonResult string
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	for {
		select {
		case <-ticker.C:
			result := new(DataIDResponse)
			resDataID, err := api.ResultsByDataID(s.Data.DataID)
			if err != nil {
				return "", errors.New("failed to get results for: " + resDataID)
			}
			json.NewDecoder(strings.NewReader(resDataID)).Decode(&result)
			if result.Data.ScanResults.ProgressPercentage == 100 {
				response, _ := json.Marshal(result)
				jsonResult = string(response)
				go func() { done <- true }()
			}
			log.Printf("progress for %s: %d \n", s.Data.DataID, result.Data.ScanResults.ProgressPercentage)
		case <-done:
			close(done)
			return jsonResult, nil
		}
	}
}

// ResultsByDataID by data_id
func (api *API) ResultsByDataID(dataID string) (string, error) {
	url := fmt.Sprintf("%s/file/%s", api.URL, dataID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return fmtResponse(api.Client.Do(req))
}

// RescanFile by file_id
func (api *API) RescanFile(fileID string) (string, error) {
	url := fmt.Sprintf("%s/file/%s/rescan", api.URL, fileID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return fmtResponse(api.Client.Do(req))
}

// RescanFiles by file_ids
func (api *API) RescanFiles(fileIDs []string) (string, error) {
	url := fmt.Sprintf("%s/file/rescan", api.URL)
	payload := &RescanReq{FileIDs: fileIDs}
	j, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(j))
	req.Header.Add("Authorization", "apikey "+api.Token)
	req.Header.Add("content-type", "application/json")
	return fmtResponse(api.Client.Do(req))
}

// GetSanitizedLink Retrieve the download link for a sanitized file
func (api *API) GetSanitizedLink(fileID string) (string, error) {
	url := fmt.Sprintf("%s/file/%s/sanitizedLink", api.URL, fileID)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return fmtResponse(api.Client.Do(req))
}

// FindOrScan file by hash
func (api *API) FindOrScan(path, hash string, headers []string) (string, error) {
	result := new(HashLookupResp)
	strRes, _ := api.HashDetails(hash)
	json.NewDecoder(strings.NewReader(strRes)).Decode(&result)
	if result.Success == false {
		log.Println("Hash not found sending to scan")
		return api.ScanFile(path, headers)
	}
	return strRes, nil
}
