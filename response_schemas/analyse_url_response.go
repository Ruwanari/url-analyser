package response_schemas

type AnalyseUrlResponse struct {
	Protocol          string `json:"protocol"`
	Title             string `json:"title"`
	Headers           Header `json:"headers"`
	InternalLinks     int    `json:"internalLinks"`
	ExternalLinks     int    `json:"externalLinks"`
	InaccessibleLinks bool   `json:"inaccessibleLinks"`
	LoginFormPresent  bool   `json:"loginFormPresent"`
}

type Header struct {
	H1 int `json:"h1"`
	H2 int `json:"h2"`
	H3 int `json:"h3"`
	H4 int `json:"h4"`
	H5 int `json:"h5"`
	H6 int `json:"h6"`
}
