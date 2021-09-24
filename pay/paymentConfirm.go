package adapay

import (
	"strings"

	adapayCore "adapay-sdk/adapay-core"
)

type paymentConfirmInterface interface {
	Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	QueryList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type PaymentConfirm struct {
	*Adapay
}

func (p *PaymentConfirm) Create(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + PAYMENT_CONFIRM
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *PaymentConfirm) Query(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + strings.Replace(PAYMENT_QUERY_CONFIRM, "{payment_confirm_id}", adapayCore.ToString(reqParam["payment_confirm_id"]), -1)
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, p.HandleConfig(multiMerchConfigId...))
}

func (p *PaymentConfirm) QueryList(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + PAYMENT_QUERY_CONFIRM_LIST
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.GET, reqParam, p.HandleConfig(multiMerchConfigId...))
}
