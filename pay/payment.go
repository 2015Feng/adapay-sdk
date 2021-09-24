package adapay

import (
	"strings"

	adapayCore "adapay-sdk/adapay-core"
)

type paymentInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	QueryList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Close(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type Payment struct {
	*Adapay
}

func (p *Payment) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + PAYMENT_CREATE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *Payment) Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + strings.Replace(PAYMENT_QUERY, "{payment_id}", adapayCore.ToString(reqParam["payment_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *Payment) QueryList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + PAYMENT_LIST_QUERY
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *Payment) Close(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + strings.Replace(PAYMENT_CLOSE, "{payment_id}", adapayCore.ToString(reqParam["payment_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, p.HandleConfig(multiMerchConfigId...))
}
