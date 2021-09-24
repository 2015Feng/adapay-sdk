package merchant

import (
	adapayCore "../adapay-core"
)

type merProfileInterface interface {
	MerProfilePicture(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)

	MerProfileForAudit(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error)
}

type MerProfile struct {
	*Merchant
}

func (e *MerProfile) MerProfilePicture(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + MER_PROFILE_PICTURE

	return adapayCore.UploadAdaPay(reqUrl, reqParam, "file", e.HandleConfig(multiMerchConfigId...))
}

func (e *MerProfile) MerProfileForAudit(reqParam map[string]interface{}, multiMerchConfigId ...string) (map[string]interface{}, *adapayCore.ApiError, error) {
	reqUrl := BASE_URL + MER_PROFILE_FOR_AUDIT
	return adapayCore.RequestAdaPay(reqUrl, adapayCore.POST, reqParam, e.HandleConfig(multiMerchConfigId...))
}
