package core

// Result defines the structure of a merged uptime result
type Result struct {
  Services map[string]*ServiceResult `json:"services"`
}

// ServiceResult defines the structure of results for a single service
type ServiceResult struct {
  Name string `json:"name"`
  Status int `json:"status"`
  Timestamp int64 `json:"timestamp"`
}
