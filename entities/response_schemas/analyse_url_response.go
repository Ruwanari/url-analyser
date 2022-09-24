package response_schemas

type AnalyseUrlResponse struct {
	Version               string   `json:"version"`
	Title                 string   `json:"title"`
	Headers               []Header `json:"headers"`
	InternalLinks         int      `json:"internalLinks"`
	ExternalLinks         int      `json:"externalLinks"`
	InaccessibleLinkCount int      `json:"inaccessibleLinkCount"`
	LoginFormPresent      bool     `json:"loginFormPresent"`
}

type Header struct {
	HeadingType string `json:"headingType"`
	Count       int    `json:"count"`
}
