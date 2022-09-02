package apionexstavka

import (
	"fmt"

	config "github.com/reucot/parser/config/collector"

	"github.com/corpix/uarand"
	"github.com/imroc/req/v3"
)

type CategoryResponse struct {
	Error     string                  `json:"Error,omitempty"`
	ErrorCode int                     `json:"ErrorCode,omitempty"`
	GUID      string                  `json:"Guid,omitempty"`
	ID        int                     `json:"Id,omitempty"`
	Success   bool                    `json:"Success,omitempty"`
	Value     []InsteadCategoryStruct `json:"Value,omitempty"`
}

type InsteadCategoryStruct struct {
	E string `json:"E,omitempty"`
	I int    `json:"I,omitempty"`
	N string `json:"N,omitempty"`
}

const urlLineCategory = "/LineFeed/GetSportsShortZip"

type Category struct {
	c *req.Client
}

func NewCategory() *Category {
	c := req.C()

	return &Category{
		c: c,
	}
}

func (c *Category) GetLineCategories() ([]InsteadCategoryStruct, error) {
	c.c.SetUserAgent(uarand.GetRandom())

	cr := &CategoryResponse{}

	resp, err := c.c.R().
		SetHeader("Accept", "application/json").
		SetResult(cr).
		Get(fmt.Sprintf("https://%s%s", config.Get().Domains.OneXStavka, urlLineCategory))

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() || !cr.Success {
		return nil, fmt.Errorf("error response: status code %d", resp.StatusCode)
	}

	return cr.Value, nil
}
