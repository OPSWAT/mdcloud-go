package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// HashLookupResp by hash
type HashLookupResp struct {
	Success bool `json:"success"`
	Data    struct {
		FileID      string `json:"file_id"`
		DataID      string `json:"data_id"`
		Archived    bool   `json:"archived"`
		ScanResults struct {
			ScanDetails        string    `json:"scan_details"`
			RescanAvailable    bool      `json:"rescan_available"`
			DataID             string    `json:"data_id"`
			ScanAllResultI     int       `json:"scan_all_result_i"`
			StartTime          time.Time `json:"start_time"`
			TotalTime          int       `json:"total_time"`
			TotalAvs           int       `json:"total_avs"`
			TotalDetectedAvs   int       `json:"total_detected_avs"`
			ProgressPercentage int       `json:"progress_percentage"`
			InQueue            int       `json:"in_queue"`
			ScanAllResultA     string    `json:"scan_all_result_a"`
		} `json:"scan_results"`
		FileInfo struct {
			FileSize            int       `json:"file_size"`
			UploadTimestamp     time.Time `json:"upload_timestamp"`
			Md5                 string    `json:"md5"`
			Sha1                string    `json:"sha1"`
			Sha256              string    `json:"sha256"`
			FileTypeCategory    string    `json:"file_type_category"`
			FileTypeDescription string    `json:"file_type_description"`
			FileTypeExtension   string    `json:"file_type_extension"`
			DisplayName         string    `json:"display_name"`
		} `json:"file_info"`
		HashResults struct {
			Wa bool `json:"wa"`
		} `json:"hash_results"`
		TopThreat   int    `json:"top_threat"`
		ShareFile   int    `json:"share_file"`
		RestVersion string `json:"rest_version"`
	} `json:"data"`
}

// HashLookupReq used for rescan body
type HashLookupReq struct {
	Hashes []string `json:"hash"`
}

// HashDetails by file_id
func (api *API) HashDetails(hash string) string {
	req, _ := http.NewRequest("GET", URL+"hash/"+hash, nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}

// HashesDetails by file_ids
func (api *API) HashesDetails(hashes []string) string {
	payload := &HashLookupReq{Hashes: hashes}
	j, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", api.URL+"hash", bytes.NewBuffer(j))
	req.Header.Add("Authorization", "apikey "+api.Token)
	req.Header.Add("content-type", "application/json")
	return FmtResponse(api.Client.Do(req))
}

// HashVulnerabilities by file_ids
func (api *API) HashVulnerabilities(hash string, limit, offset int) string {
	url, _ := url.Parse(URL + "vulnerability/" + hash)
	q := url.Query()
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		q.Set("offset", strconv.Itoa(offset))
	}
	url.RawQuery = q.Encode()
	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}
