package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// ProcessInfo post procesing info
type ProcessInfo struct {
	UserAgent          string `json:"user_agent"`
	Result             string `json:"result"`
	ProgressPercentage int    `json:"progress_percentage"`
	Profile            string `json:"profile"`
	PostProcessing     struct {
		CopyMoveDestination  string `json:"copy_move_destination"`
		ConvertedTo          string `json:"converted_to"`
		ConvertedDestination string `json:"converted_destination"`
		ActionsRan           string `json:"actions_ran"`
		ActionsFailed        string `json:"actions_failed"`
	} `json:"post_processing"`
	FileTypeSkippedScan bool   `json:"file_type_skipped_scan"`
	BlockedReason       string `json:"blocked_reason"`
}

// Votes per hash
type Votes struct {
	Up   int `json:"up"`
	Down int `json:"down"`
}

// HashLookupResp by hash
// todo generate from swagger
type HashLookupResp struct {
	Success bool `json:"success"`
	Data    struct {
		ScanResultHistoryLength int            `json:"scan_result_history_length"`
		Votes                   Votes          `json:"votes"`
		FileID                  string         `json:"file_id"`
		DataID                  string         `json:"data_id"`
		Archived                bool           `json:"archived"`
		ProcessInfo             ProcessInfo    `json:"process_info"`
		ExtractedFiles          ExtractedFiles `json:"extracted_files"`
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
	} `json:"data"`
}

// HashLookupReq used for rescan body
type HashLookupReq struct {
	Hashes []string `json:"hash"`
}

// HashDetails by file_id
func (api *API) HashDetails(hash string) (string, error) {
	url := fmt.Sprintf("%s/hash/%s", api.URL, hash)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", api.Authorization)
	return fmtResponse(api.Client.Do(req))
}

// HashesDetails by file_ids
func (api *API) HashesDetails(hashes []string) (string, error) {
	url := fmt.Sprintf("%s/hash", api.URL)
	payload := &HashLookupReq{Hashes: hashes}
	j, _ := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(j))
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", api.Authorization)
	req.Header.Add("content-type", "application/json")
	return fmtResponse(api.Client.Do(req))
}

// HashVulnerabilities by hash
func (api *API) HashVulnerabilities(hash string, limit, offset int) (string, error) {
	url, _ := url.Parse(fmt.Sprintf("%s/vulnerability/%s", api.URL, hash))
	q := url.Query()
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		q.Set("offset", strconv.Itoa(offset))
	}
	url.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", api.Authorization)
	return fmtResponse(api.Client.Do(req))
}
