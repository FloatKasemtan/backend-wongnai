package covid

import (
	"github.com/FloatKasemtan/types"
	"github.com/FloatKasemtan/types/response"
	"github.com/FloatKasemtan/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SummaryHandler(c *gin.Context) {
	// initialize province map with -1 key for "N/A" value
	province := map[int64]*types.ProvinceData{
		-1: {
			Name:  "N/A",
			Count: 0,
		},
	}
	// initialize ageGroup map with all possible range
	ageGroup := map[string]int32{
		"N/A":   0,
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
	}
	dataUrl := "http://static.wongnai.com/devinterview/covid-cases.json"

	resp := new(types.CovidResponse)
	//  get data from api, and push in struct
	if err := utils.GetJson[*types.CovidResponse](dataUrl, resp); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error fetching data"})
		return
	}

	for _, data := range resp.Data {
		distributeByProvince(data, province)
		distributeByAge(data, ageGroup)
	}

	// initialize and restructure variable for province response
	resultProvince := map[string]int32{}
	for _, data := range province {
		resultProvince[data.Name] = data.Count
	}

	result := &response.Response{
		Success: true,
		Payload: types.SummaryResponse{
			Province: resultProvince,
			AgeGroup: ageGroup,
		},
	}

	c.IndentedJSON(http.StatusOK, *result)
}

func distributeByProvince(data types.CovidData, province map[int64]*types.ProvinceData) {
	if data.ProvinceId == nil {
		province[-1].Count++
		return
	}
	// check if key is key valid
	if p, ok := province[*data.ProvinceId]; ok {
		p.Count++
	} else {
		province[*data.ProvinceId] = &types.ProvinceData{
			Name:  *data.Province,
			Count: 1,
		}
	}
}

func distributeByAge(data types.CovidData, ageGroup map[string]int32) {
	switch {
	case data.Age == nil:
		ageGroup["N/A"]++
	case *data.Age <= 30:
		ageGroup["0-30"]++
	case *data.Age <= 60:
		ageGroup["31-60"]++
	default:
		ageGroup["61+"]++
	}
}
