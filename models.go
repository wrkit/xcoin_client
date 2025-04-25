package xcoin_client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// CommandType представляет тип команды бота
type CommandType string

// Доступные команды бота
const (
	CommandUnset    CommandType = "Unset"
	CommandStopWait CommandType = "StopWait"
	CommandStop     CommandType = "Stop"
	CommandRunned   CommandType = "Runned"
)

// VersionResponse представляет ответ от эндпоинта версии
type VersionResponse struct {
	BotID      string     `json:"bot_id"`
	BotAddress string     `json:"bot_address"`
	PairCount  int        `json:"pair_count"`
	Pairs      []PairInfo `json:"pairs"`
}

// PairInfo представляет информацию о торговой паре
type PairInfo struct {
	KeyWork int    `json:"key_work"`
	Pair    string `json:"pair"`
}

// KeyValueResponse представляет ответ с настройками пары
type KeyValueResponse struct {
	Val []KeyValue `json:"val"`
}

// KeyValue представляет пару ключ-значение
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Order представляет информацию о торговом ордере
type Order struct {
	ClientOrderId  string  `json:"ClientOrderId"`
	Comission1     float64 `json:"Comission1"`
	Comission2     float64 `json:"Comission2"`
	CreateTime     string  `json:"CreateTime"`
	DataUpdate     string  `json:"DataUpdate"`
	Level          int     `json:"Level"`
	Mode           int     `json:"Mode"`
	OrderId        int     `json:"OrderId"`
	Price          float64 `json:"Price"`
	Quantity       float64 `json:"Quantity"`
	QuantityFilled float64 `json:"QuantityFilled"`
	Side           int     `json:"Side"`
	Status         int     `json:"Status"`
	Symbol         string  `json:"Symbol"`
	Type           int     `json:"Type"`
	KeySetka       int     `json:"keySetka"`
	MyManual       int     `json:"myManual"`
}

type PostBotCmdResponse struct {
	Val []KeyValue `json:"Val"`
}

type Color struct {
	A int `json:"A"`
	R int `json:"R"`
	G int `json:"G"`
	B int `json:"B"`
}

type LotSizeFilter struct {
	MaxQuantity float64 `json:"MaxQuantity"`
	MinQuantity float64 `json:"MinQuantity"`
	StepSize    float64 `json:"StepSize"`
}

type MinNotionalFilter struct {
	MinNotional float64 `json:"MinNotional"`
}

type PriceFilter struct {
	MaxPrice float64 `json:"MaxPrice"`
	MinPrice float64 `json:"MinPrice"`
	TickSize float64 `json:"TickSize"`
}

type SVal struct {
	BaseAsset         string            `json:"BaseAsset"`
	QuoteAsset        string            `json:"QuoteAsset"`
	Symbol            string            `json:"Symbol"`
	LotSizeFilter     LotSizeFilter     `json:"LotSizeFilter"`
	MinNotionalFilter MinNotionalFilter `json:"MinNotionalFilter"`
	PriceFilter       PriceFilter       `json:"PriceFilter"`
}

type Pair struct {
	Birga         string  `json:"Birga"`
	CurSale       float64 `json:"CurSale"`
	CurUsredn     float64 `json:"CurUsredn"`
	Depo          int     `json:"Depo"`
	Description   string  `json:"Description"`
	FreeMoneta    int     `json:"FreeMoneta"`
	InZakup       float64 `json:"InZakup"`
	Invest        float64 `json:"Invest"`
	KeyWork       int     `json:"KeyWork"`
	KolFill       int     `json:"KolFill"`
	KolFillColor  Color   `json:"KolFillColor"`
	KolOtkup      int     `json:"KolOtkup"`
	KolOtkupColor Color   `json:"KolOtkupColor"`
	KolZakup      int     `json:"KolZakup"`
	KolZakupColor Color   `json:"KolZakupColor"`
	Moneta        float64 `json:"Moneta"`
	Price         float64 `json:"Price"`
	Rasst         string  `json:"Rasst"`
	RasstColor    Color   `json:"RasstColor"`
	RasstNum      float64 `json:"RasstNum"`
	Status        string  `json:"Status"`
	StatusInt     int     `json:"StatusInt"`
	Symbol        string  `json:"Symbol"`
	UpMoneta      string  `json:"UpMoneta"`
	Val1Color     Color   `json:"Val1Color"`
	Val2Color     Color   `json:"Val2Color"`
	AllDay        int     `json:"allDay"`
	Setting       Setting `json:"Setting"`
}

type Setting struct {
	Birga                  int        `json:"Birga"`
	DownOrderCurrentWorkId int        `json:"DownOrderCurrentWorkId"`
	DownOrderKol           int        `json:"DownOrderKol"`
	DownOrderPercent       int        `json:"DownOrderPercent"`
	KolOrder               int        `json:"KolOrder"`
	PaintColor             Color      `json:"PaintColor"`
	Primechanie            string     `json:"Primechanie"`
	SVal                   SVal       `json:"SVal"`
	Val1                   string     `json:"Val1"`
	Val2                   string     `json:"Val2"`
	AddDescr               string     `json:"addDescr"`
	AftBreak               int        `json:"aftBreak"`
	AftBreakRunNext        bool       `json:"aftBreakRunNext"`
	BezEnterEnabled        bool       `json:"bezEnterEnabled"`
	BezEnterTimeInterval   string     `json:"bezEnterTimeInterval"`
	BezKolKline            int        `json:"bezKolKline"`
	BezKolOff              int        `json:"bezKolOff"`
	BezOrderKontrol        string     `json:"bezOrderKontrol"`
	BezOtstupKline         int        `json:"bezOtstupKline"`
	BezStartOn             bool       `json:"bezStartOn"`
	CurrentOrderManual     string     `json:"currentOrderManual"`
	DepositOrder           int        `json:"depositOrder"`
	DepositOrderVal        string     `json:"depositOrderVal"`
	FinDelay               int        `json:"finDelay"`
	FinJobKorrect          int        `json:"finJobKorrect"`
	FinKeyWork             int        `json:"finKeyWork"`
	FinLevel               int        `json:"finLevel"`
	FinOtstup              int        `json:"finOtstup"`
	FinVariant             int        `json:"finVariant"`
	FirstStep              float64    `json:"firstStep"`
	KeyImport              int        `json:"keyImport"`
	LimitDeposit           int        `json:"limitDeposit"`
	LimitDepositReserv     int        `json:"limitDepositReserv"`
	Martingale             int        `json:"martingale"`
	MartingaleManual       string     `json:"martingaleManual"`
	MinPlusStep            int        `json:"minPlusStep"`
	Mnogitel               int        `json:"mnogitel"`
	OffGlobalWarning       bool       `json:"offGlobalWarning"`
	OrderStep              float64    `json:"orderStep"`
	OrderStepManual        string     `json:"orderStepManual"`
	PlusStep               float64    `json:"plusStep"`
	PlusStepManual         string     `json:"plusStepManual"`
	PriceAutoCountKlines   int        `json:"priceAutoCountKlines"`
	PriceAutoEnabled       bool       `json:"priceAutoEnabled"`
	PriceAutoProcentStop   float64    `json:"priceAutoProcentStop"`
	PriceAutoTimeInterval  string     `json:"priceAutoTimeInterval"`
	PriceAutoWait          int        `json:"priceAutoWait"`
	PriceStop              int        `json:"priceStop"`
	PriceStop2             int        `json:"priceStop2"`
	Profit                 float64    `json:"profit"`
	ProfitManual           string     `json:"profitManual"`
	Reload                 float64    `json:"reload"`
	RestartAuto            bool       `json:"restartAuto"`
	SmsOrder               int        `json:"smsOrder"`
	SmsPost                string     `json:"smsPost"`
	SmsText                string     `json:"smsText"`
	SmsWarningHours        int        `json:"smsWarningHours"`
	SmsWarningNum          int        `json:"smsWarningNum"`
	StopAuto               bool       `json:"stopAuto"`
	Strateg                int        `json:"strateg"`
	TimeActivate           DotNetTime `json:"timeActivate"`
	TimeDeActivate         DotNetTime `json:"timeDeActivate"`
	Trailing               float64    `json:"trailing"`
	TrailingProc           int        `json:"trailingProc"`
	TrailingProcUp         int        `json:"trailingProcUp"`
}

type DotNetTime time.Time

func (t *DotNetTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || len(data) == 0 {
		*t = DotNetTime(time.Time{})
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	startIndex := strings.Index(s, "(")
	plusIndex := strings.Index(s, "+")
	endIndex := strings.Index(s, ")")
	ts := s[startIndex+1 : plusIndex]

	zone := s[plusIndex+1 : endIndex]

	tsNumber, err := strconv.Atoi(ts)
	if err != nil {
		return err
	}
	zoneNumber, err := strconv.Atoi(zone)
	if err != nil {
		zoneNumber = 0
	}

	timeParse := time.Unix(int64(tsNumber/1000), 0).UTC()
	timeResult := timeParse.In(time.FixedZone("", zoneNumber/100*60*60))
	*t = DotNetTime(timeResult)

	return nil
}
func (t *DotNetTime) MarshalJSON() ([]byte, error) {
	//"\/Date(1647952719369+0200)\/"

	tm := time.Time(*t)
	_, offset := tm.Zone()

	if offset != 0 {
		offset = (offset / 60 / 60) * 100
	}

	offsetStr := fmt.Sprintf("%d", offset)

	if len(offsetStr) < 4 {
		pref := make([]string, 4-len(offsetStr)+1)
		offsetStr = strings.Join(pref, "0") + offsetStr
	}

	if offset < 0 {
		offsetStr = "-" + offsetStr
	} else {
		offsetStr = "+" + offsetStr
	}

	format := `\/Date(%d%s)\/`
	dateResult := fmt.Sprintf(format, tm.Unix()*1000, offsetStr)

	data, err := json.Marshal(dateResult)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (t *DotNetTime) String() string {
	return time.Time(*t).String()
}
func (t *DotNetTime) ToJsonString() string {

	tRaw := strings.Replace(t.String(), " #", "", -1)
	d, err := time.Parse("2006-01-02 15:04:05 -0700 -0700", tRaw)
	if err != nil {
		return fmt.Sprintf("%s", err)
	}
	d.UTC()
	return fmt.Sprintf("%s", d.Format(time.RFC3339))
}
func (t *DotNetTime) Time() time.Time {
	return time.Time(*t)
}
