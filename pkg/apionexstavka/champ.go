package apionexstavka

import (
	"fmt"

	config "github.com/reucot/parser/config/collector"

	"github.com/corpix/uarand"
	"github.com/imroc/req/v3"
)

const urlLineChamp string = "/LineFeed/GetChampsZip"
const urlLiveChamp string = "/LiveFeed/GetChampsZip"

type ChampResponse struct {
	Error     string               `json:"Error,omitempty"`
	ErrorCode int                  `json:"ErrorCode,omitempty"`
	GUID      string               `json:"Guid,omitempty"`
	ID        int                  `json:"Id,omitempty"`
	Success   bool                 `json:"Success,omitempty"`
	Value     []InsteadChampStruct `json:"Value,omitempty"`
}

type InsteadChampStruct struct {
	Ci int    `json:"CI,omitempty"`
	Gc int    `json:"GC,omitempty"`
	L  string `json:"L,omitempty"`
	Le string `json:"LE,omitempty"`
}

type Champ struct {
	c       *req.Client
	sportID int
}

func NewChamp(sportID int) *Champ {
	c := req.C()

	return &Champ{
		c:       c,
		sportID: sportID,
	}
}

func (c *Champ) GetChamps(isLive bool) ([]InsteadChampStruct, error) {
	c.c.SetUserAgent(uarand.GetRandom())

	urlChamp := urlLineChamp
	if isLive {
		urlChamp = urlLiveChamp
	}

	cr := &ChampResponse{}

	resp, err := c.c.R().
		SetHeader("Accept", "application/json").
		SetResult(cr).
		Get(fmt.Sprintf("https://%s%s?sport=%d&virtualSports=true&groupChamps=true",
			config.Get().Domains.OneXStavka,
			urlChamp,
			c.sportID))

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() || !cr.Success {
		return nil, fmt.Errorf("error response: status code %d", resp.StatusCode)
	}

	return cr.Value, nil
}
