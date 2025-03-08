package nietzsche

type StartResponse struct {
	Server           string `json:"server"`
	Task             string `json:"task"`
	RemainingCredits int    `json:"remaining_credits"`
}

type UploadParams struct {
	Task string `json:"task"`
	URL  string `json:"url"`
}

type UploadResponse struct {
	ServerFileName string `json:"server_file_name"`
}

type ProcessResponse struct {
	ServerFileName string `json:"server_file_name"`
}
