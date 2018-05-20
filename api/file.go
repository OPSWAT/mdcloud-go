package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// RescanReq used for rescan body
type RescanReq struct {
	FileIDs []string `json:"file_ids"`
}

// ScanFile sends to API
func (api *API) ScanFile(path string) {
	fmt.Printf(path)
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
	req.Header.Add("x-sample-sharing", "1")
	return FmtResponse(api.Client.Do(req))
}
