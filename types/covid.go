package types

type CovidResponse struct {
	Data []CovidData `json:"Data"`
}

type CovidData struct {
	ConfirmDate    *string `json:"ConfirmDate"`
	No             *int64  `json:"No"`
	Age            *int32  `json:"Age"`
	Gender         *string `json:"Gender"`
	GenerEn        *string `json:"GenerEn"`
	Nation         *string `json:"Nation"`
	NationEn       *string `json:"NationEn"`
	Province       *string `json:"Province"`
	ProvinceId     *int64  `json:"ProvinceId"`
	District       *string `json:"District"`
	ProvinceEn     *string `json:"ProvinceEn"`
	StatQuarantine *int16  `json:"StatQuarantine"`
}

type ProvinceData struct {
	Name  string
	Count int32
}

type SummaryResponse struct {
	Province map[string]int32
	AgeGroup map[string]int32
}
