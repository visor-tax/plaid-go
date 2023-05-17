/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.370.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// AssetReportAddOns A list of add-ons that should be included in the Asset Report.  `fast_assets`: When Fast Assets is requested, Plaid will create two versions of the Asset Report: the Fast Asset Report, which will contain only Identity and Balance information, and the Full Asset Report, which will also contain Transactions information. A `PRODUCT_READY` webhook will be fired for each Asset Report when it is ready, and the `report_type` field will indicate whether the webhook is firing for the `full` or `fast` Asset Report. To retrieve the Fast Asset Report, call `/asset_report/get` with `fast_report` set to `true`. There is no additional charge for using Fast Assets.  `investments`: Request an Asset Report with Investments. This add-on is in closed beta and not generally available.
type AssetReportAddOns string

var _ = fmt.Printf

// List of AssetReportAddOns
const (
	ASSETREPORTADDONS_INVESTMENTS AssetReportAddOns = "investments"
	ASSETREPORTADDONS_FAST_ASSETS AssetReportAddOns = "fast_assets"
)

var allowedAssetReportAddOnsEnumValues = []AssetReportAddOns{
	"investments",
	"fast_assets",
}

func (v *AssetReportAddOns) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := AssetReportAddOns(value)


	*v = enumTypeValue
	return nil
}

// NewAssetReportAddOnsFromValue returns a pointer to a valid AssetReportAddOns
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetReportAddOnsFromValue(v string) (*AssetReportAddOns, error) {
	ev := AssetReportAddOns(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetReportAddOns) IsValid() bool {
	for _, existing := range allowedAssetReportAddOnsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssetReportAddOns value
func (v AssetReportAddOns) Ptr() *AssetReportAddOns {
	return &v
}

type NullableAssetReportAddOns struct {
	value *AssetReportAddOns
	isSet bool
}

func (v NullableAssetReportAddOns) Get() *AssetReportAddOns {
	return v.value
}

func (v *NullableAssetReportAddOns) Set(val *AssetReportAddOns) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportAddOns) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportAddOns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportAddOns(val *AssetReportAddOns) *NullableAssetReportAddOns {
	return &NullableAssetReportAddOns{value: val, isSet: true}
}

func (v NullableAssetReportAddOns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportAddOns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

