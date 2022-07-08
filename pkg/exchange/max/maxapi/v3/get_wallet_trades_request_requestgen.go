// Code generated by "requestgen -method GET -url /api/v3/wallet/:walletType/trades -type GetWalletTradesRequest -responseType []Trade"; DO NOT EDIT.

package v3

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wanewang/bbgo/pkg/exchange/max/maxapi"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

func (g *GetWalletTradesRequest) Market(market string) *GetWalletTradesRequest {
	g.market = market
	return g
}

func (g *GetWalletTradesRequest) From(from uint64) *GetWalletTradesRequest {
	g.from = &from
	return g
}

func (g *GetWalletTradesRequest) StartTime(startTime time.Time) *GetWalletTradesRequest {
	g.startTime = &startTime
	return g
}

func (g *GetWalletTradesRequest) EndTime(endTime time.Time) *GetWalletTradesRequest {
	g.endTime = &endTime
	return g
}

func (g *GetWalletTradesRequest) Limit(limit uint64) *GetWalletTradesRequest {
	g.limit = &limit
	return g
}

func (g *GetWalletTradesRequest) WalletType(walletType max.WalletType) *GetWalletTradesRequest {
	g.walletType = walletType
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetWalletTradesRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetWalletTradesRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check market field -> json key market
	market := g.market

	// TEMPLATE check-required
	if len(market) == 0 {
		return nil, fmt.Errorf("market is required, empty string given")
	}
	// END TEMPLATE check-required

	// assign parameter of market
	params["market"] = market
	// check from field -> json key from_id
	if g.from != nil {
		from := *g.from

		// assign parameter of from
		params["from_id"] = from
	} else {
	}
	// check startTime field -> json key start_time
	if g.startTime != nil {
		startTime := *g.startTime

		// assign parameter of startTime
		// convert time.Time to milliseconds time stamp
		params["start_time"] = strconv.FormatInt(startTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check endTime field -> json key end_time
	if g.endTime != nil {
		endTime := *g.endTime

		// assign parameter of endTime
		// convert time.Time to milliseconds time stamp
		params["end_time"] = strconv.FormatInt(endTime.UnixNano()/int64(time.Millisecond), 10)
	} else {
	}
	// check limit field -> json key limit
	if g.limit != nil {
		limit := *g.limit

		// assign parameter of limit
		params["limit"] = limit
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetWalletTradesRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if g.isVarSlice(_v) {
			g.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetWalletTradesRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetWalletTradesRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check walletType field -> json key walletType
	walletType := g.walletType

	// TEMPLATE check-required
	if len(walletType) == 0 {
		return nil, fmt.Errorf("walletType is required, empty string given")
	}
	// END TEMPLATE check-required

	// assign parameter of walletType
	params["walletType"] = walletType

	return params, nil
}

func (g *GetWalletTradesRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (g *GetWalletTradesRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (g *GetWalletTradesRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (g *GetWalletTradesRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (g *GetWalletTradesRequest) Do(ctx context.Context) ([]max.Trade, error) {

	// empty params for GET operation
	var params interface{}
	query, err := g.GetParametersQuery()
	if err != nil {
		return nil, err
	}

	apiURL := "/api/v3/wallet/:walletType/trades"
	slugs, err := g.GetSlugsMap()
	if err != nil {
		return nil, err
	}

	apiURL = g.applySlugsToUrl(apiURL, slugs)

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse []max.Trade
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return apiResponse, nil
}
