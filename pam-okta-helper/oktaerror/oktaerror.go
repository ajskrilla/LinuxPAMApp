package oktaerror

// maybe move to the util folder?
type OktaError struct {
	ErrorCode    string `json:"errorCode"`
	ErrorSummary string `json:"errorSummary"`
	ErrorLink    string `json:"errorLink,omitempty"`
	ErrorId      string `json:"errorId,omitempty"`
}