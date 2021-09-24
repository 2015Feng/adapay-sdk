package adapay

import (
	"strings"

	adapayCore "adapay-sdk/adapay-core"
)

type paymentReverseInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	QueryList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type PaymentReverse struct {
	*Adapay
}

func (p *PaymentReverse) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + PAYMENT_REVERSE
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *PaymentReverse) Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + strings.Replace(PAYMENT_QUERY_REVERSE, "{reverse_id}", adapayCore.ToString(reqParam["reverse_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *PaymentReverse) QueryList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + PAYMENT_QUERY_REVERSE_LIST
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, p.HandleConfig(multiMerchConfigId...))
}
