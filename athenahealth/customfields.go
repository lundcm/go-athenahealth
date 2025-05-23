package athenahealth

import "context"

type CustomFieldValue struct {
	CustomFieldID    string `json:"customfieldid"`
	CustomFieldValue string `json:"customfieldvalue"`
	OptionID         string `json:"optionid"`
}

type SelectOption struct {
	OptionID    string `json:"optionid"`
	OptionValue string `json:"optionvalue"`
}

type CustomField struct {
	CaseSensitive  bool            `json:"casesensitive"`
	CustomFieldID  string          `json:"customfieldid"`
	DisallowUpdate bool            `json:"disallowupdate"`
	MaxLength      string          `json:"length"`
	Name           string          `json:"name"`
	Ordering       string          `json:"ordering"`
	Searchable     bool            `json:"searchable"`
	SelectList     []*SelectOption `json:"selectlist"`
	Select         bool            `json:"select"`
	Type           string          `json:"type"`
}

// ListCustomFields - List of custom fields for the practice
//
// GET /v1/{practiceid}/customfields
//
// https://docs.athenahealth.com/api/api-ref/custom-fields#Get-practice's-list-of-custom-fields
func (h *HTTPClient) ListCustomFields(ctx context.Context) ([]*CustomField, error) {
	var out []*CustomField

	_, err := h.Get(ctx, "/customfields", nil, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
