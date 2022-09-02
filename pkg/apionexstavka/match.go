package apionexstavka

import (
	"fmt"

	config "github.com/reucot/parser/config/collector"

	"github.com/corpix/uarand"
	"github.com/imroc/req/v3"
)

const (
	urlLineMatches string = "/LineFeed/Get1x2_VZip"
	urlLiveMatches string = "/LiveFeed/Get1x2_VZip"
	urlLineMatch   string = "/LineFeed/GetGameZip"
	urlLiveMatch   string = "/LiveFeed/GetGameZip"
)

type Match struct {
	c       *req.Client
	sportID int
}

func NewMatch(sportID int) *Match {
	return &Match{
		c:       req.C(),
		sportID: sportID,
	}
}

type MatchesResponse struct {
	Error     string                 `json:"Error,omitempty"`
	ErrorCode int                    `json:"ErrorCode,omitempty"`
	GUID      string                 `json:"Guid,omitempty"`
	ID        int                    `json:"Id,omitempty"`
	Success   bool                   `json:"Success,omitempty"`
	Value     []InsteadMatchesStruct `json:"Value,omitempty"`
}

type InsteadMatchesStruct struct {
	I int `json:"I,omitempty"` // ID - матча
	// Ci  int    `json:"CI,omitempty"`  // ID - матча
	Coi int    `json:"COI,omitempty"` // ID - турнира
	L   string `json:"L,omitempty"`   // Наименование турнира RU
	LE  string `json:"LE,omitempty"`  // Наименование турнира EN
	O1  string `json:"O1,omitempty"`  // Наименование RU 1 участника
	O1E string `json:"O1E,omitempty"` // Наименование EN 1 участника
	O2  string `json:"O2,omitempty"`  // Наименование RU 2 участника
	O2E string `json:"O2E,omitempty"` // Наименование EN 2 участника
	S   int    `json:"S,omitempty"`   // Дата матча
	Sc  struct {
		Fs struct {
			S1 int `json:"S1,omitempty"`
			S2 int `json:"S2,omitempty"`
		} `json:"FS,omitempty"` // Общий счет (может не быть если 0 : 0) смотреть по PS
		Ps []struct {
			Key   int      `json:"Key,omitempty"`   // Раунд
			Value struct{} `json:"Value,omitempty"` // Значение
		} `json:"PS,omitempty"` // Голы по раундам
		Ss struct {
			S1 string `json:"S1,omitempty"`
			S2 string `json:"S2,omitempty"`
		} `json:"SS,omitempty"` // Счет 15 : 15 для Тенниса
		Tr int `json:"TR,omitempty"`
		Ts int `json:"TS,omitempty"` //Пройденное время в секундах (футбол)
	} `json:"SC,omitempty"` // Информация о текущем событии
}

func (mr *Match) GetMatches(isLive bool) ([]InsteadMatchesStruct, error) {
	mr.c.SetUserAgent(uarand.GetRandom())

	urlMatches := urlLineMatches
	if isLive {
		urlMatches = urlLiveMatches
	}

	cr := &MatchesResponse{}

	resp, err := mr.c.R().
		SetHeader("Accept", "application/json").
		SetResult(cr).
		Get(fmt.Sprintf("https://%s%s?sports=%d&count=50&mode=4",
			config.Get().Domains.OneXStavka,
			urlMatches,
			mr.sportID))

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() || !cr.Success {
		return nil, fmt.Errorf("error response: status code %d", resp.StatusCode)
	}

	return cr.Value, nil
}

type MatchResponse struct {
	Error     string             `json:"Error,omitempty"`
	ErrorCode int                `json:"ErrorCode,omitempty"`
	GUID      string             `json:"Guid,omitempty"`
	ID        int                `json:"Id,omitempty"`
	Success   bool               `json:"Success,omitempty"`
	Value     InsteadMatchStruct `json:"Value,omitempty"`
}

type InsteadMatchStruct struct {
	I   int `json:"I,omitempty"`   // ID - матча
	Coi int `json:"COI,omitempty"` // ID - чемпионата
	Ge  []struct {
		E [][]struct {
			C float64 `json:"C,omitempty"`
			G int     `json:"G,omitempty"`
			T int     `json:"T,omitempty"`
		} `json:"E,omitempty"`
		G int `json:"G,omitempty"`
	} `json:"GE,omitempty"` // Основные ставки
	L   string `json:"L,omitempty"`  // Наименование чемпионата RU
	Le  string `json:"LE,omitempty"` // Наименование чемпионата EN
	Li  int    `json:"LI,omitempty"` // ID чемпионата
	Mec []struct {
		Ec int    `json:"EC,omitempty"`
		Mt int    `json:"MT,omitempty"`
		N  string `json:"N,omitempty"`
	} `json:"MEC,omitempty"` // Связь заголовков и ставок
	O1    string   `json:"O1,omitempty"`    // Наименование RU 1 участника
	O1E   string   `json:"O1E,omitempty"`   // Наименование EN 1 участника
	O1Img []string `json:"O1IMG,omitempty"` // Ссылка на картинку 1 участника
	O2    string   `json:"O2,omitempty"`    // Наименование RU 2 участника
	O2E   string   `json:"O2E,omitempty"`   // Наименование EN 2 участника
	O2Img []string `json:"O2IMG,omitempty"` // Ссылка на картинку 2 участника
	Se    string   `json:"SE,omitempty"`    // Наименование категории En
	Si    int      `json:"SI,omitempty"`    // Категория ID
	Sn    string   `json:"SN,omitempty"`    // Наименование категории
	Sg    []struct {
		Ec  int `json:"EC,omitempty"`
		Egc int `json:"EGC,omitempty"`
		Ge  []struct {
			E [][]struct {
				C float64 `json:"C,omitempty"`
				G int     `json:"G,omitempty"`
				T int     `json:"T,omitempty"`
			} `json:"E,omitempty"`
			G int `json:"G,omitempty"`
		} `json:"GE,omitempty"`
		I   int `json:"I,omitempty"`
		Mec []struct {
			Ec int    `json:"EC,omitempty"`
			Mt int    `json:"MT,omitempty"`
			N  string `json:"N,omitempty"`
		} `json:"MEC,omitempty"`
		Mg int    `json:"MG,omitempty"`
		N  int    `json:"N,omitempty"`
		P  int    `json:"P,omitempty"`
		Pn string `json:"PN,omitempty"`
		Si int    `json:"SI,omitempty"`
		R  int    `json:"R,omitempty"`
	} `json:"SG,omitempty"` // Доп ставки на тайм
}

func (mr *Match) GetMatch(gameID int, isLive bool) (InsteadMatchStruct, error) {
	mr.c.SetUserAgent(uarand.GetRandom())

	urlMatch := urlLineMatch
	if isLive {
		urlMatch = urlLiveMatch
	}

	cr := &MatchResponse{}

	resp, err := mr.c.R().
		SetHeader("Accept", "application/json").
		SetResult(cr).
		Get(fmt.Sprintf("https://%s%s?id=%d&lng=ru&GroupEvents=true&allEventsGroupSubGames=true&countevents=250",
			config.Get().Domains.OneXStavka,
			urlMatch,
			gameID))

	if err != nil {
		return InsteadMatchStruct{}, err
	}

	if !resp.IsSuccess() {
		return InsteadMatchStruct{}, fmt.Errorf("error response: status code %d", resp.StatusCode)
	}

	if !cr.Success {
		fmt.Println(cr)
		return InsteadMatchStruct{}, fmt.Errorf("error response: %s", cr.Error)
	}

	return cr.Value, nil
}
