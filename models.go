package main

// HTTP status code 200 and repository model in data
// swagger:response repoResp
type swaggRepoResp struct {
	// in:body
	Body struct {
		// HTTP status code 200/201
		Code int `json:"code"`
		// Repository Data
		Data string `json:"data"`
	}
}

// Error Forbidden
// swagger:response forbidden
type swaggErrForbidden struct {
	// in:body
	Body struct {
		// HTTP status code 403 -  Forbidden
		Code int `json:"code"`
		// Detailed error message
		Message string `json:"message"`
	}
}

// A validatePaymentParams parameter model.
//
// This is used for operations that want the ID of a user in the path
// swagger:parameters validatePaymentRequest
type validatePaymentParams struct {
	// The intParam1 of user
	//
	// in: formData
	// required: true
	IntParam1 int `json:"intParam1"`

	// The stParam1 of user
	//
	// in: formData
	// required: false
	StrParam1 string `json:"strParam1"`
}
