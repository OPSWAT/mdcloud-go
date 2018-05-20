package api

import (
	"bytes"
	"encoding/json"
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
func (api *API) ScanFile(path string, headers []string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	req, err := http.NewRequest("POST", api.URL+"file", file)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "apikey "+api.Token)
	req.Header.Add("Content-Type", "binary/octet-stream")
	req.Header.Add("x-filename", filepath.Base(path))
	if len(headers) > 0 {
		for _, hdr := range headers {
			header := strings.Split("=", hdr)
			req.Header.Add(header[0], header[1])
		}
	}
	resp, err := api.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var s = new(ScanResponse)
	err = json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		log.Fatal(err)
	}

	var jsonResult string
	for range time.Tick(500 * time.Millisecond) {
		result := new(DataIDResponse)
		json.NewDecoder(strings.NewReader(api.ResultsByDataID(s.Data.DataID))).Decode(&result)
		if result.Data.ScanResults.ProgressPercentage == 100 {
			response, _ := json.Marshal(result)
			jsonResult = string(response)
			break
		}
		log.Println("progress: ", result.Data.ScanResults.ProgressPercentage)
	}
	return jsonResult
}

// ResultsByDataID by data_id
func (api *API) ResultsByDataID(dataID string) string {
	req, _ := http.NewRequest("GET", URL+"file/"+dataID, nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}

// RescanFile by file_id
func (api *API) RescanFile(fileID string) string {
	req, _ := http.NewRequest("GET", URL+"file/"+fileID+"/rescan", nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}

// RescanFiles by file_ids
func (api *API) RescanFiles(fileIDs []string) string {
	payload := &RescanReq{FileIDs: fileIDs}
	j, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", api.URL+"file/rescan", bytes.NewBuffer(j))
	req.Header.Add("Authorization", "apikey "+api.Token)
	req.Header.Add("content-type", "application/json")
	return FmtResponse(api.Client.Do(req))
}

// GetSanitizedLink Retrieve the download link for a sanitized file
func (api *API) GetSanitizedLink(fileID string) string {
	req, _ := http.NewRequest("GET", URL+"file/"+fileID+"/sanitizedLink", nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}
