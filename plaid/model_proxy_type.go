/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.345.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"fmt"
)

// ProxyType An enum indicating whether a network proxy is present and if so what type it is.  `none_detected` indicates the user is not on a detectable proxy network.  `tor` indicates the user was using a Tor browser, which sends encrypted traffic on a decentralized network and is somewhat similar to a VPN (Virtual Private Network).  `vpn` indicates the user is on a VPN (Virtual Private Network)  `web_proxy` indicates the user is on a web proxy server, which may allow them to conceal information such as their IP address or other identifying information.  `public_proxy` indicates the user is on a public web proxy server, which is similar to a web proxy but can be shared by multiple users. This may allow multiple users to appear as if they have the same IP address for instance.
type ProxyType string

var _ = fmt.Printf

// List of ProxyType
const (
	PROXYTYPE_NONE_DETECTED ProxyType = "none_detected"
	PROXYTYPE_TOR ProxyType = "tor"
	PROXYTYPE_VPN ProxyType = "vpn"
	PROXYTYPE_WEB_PROXY ProxyType = "web_proxy"
	PROXYTYPE_PUBLIC_PROXY ProxyType = "public_proxy"
)

var allowedProxyTypeEnumValues = []ProxyType{
	"none_detected",
	"tor",
	"vpn",
	"web_proxy",
	"public_proxy",
}

func (v *ProxyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}

	enumTypeValue := ProxyType(value)


	*v = enumTypeValue
	return nil
}

// NewProxyTypeFromValue returns a pointer to a valid ProxyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProxyTypeFromValue(v string) (*ProxyType, error) {
	ev := ProxyType(v)


	return &ev, nil
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProxyType) IsValid() bool {
	for _, existing := range allowedProxyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProxyType value
func (v ProxyType) Ptr() *ProxyType {
	return &v
}

type NullableProxyType struct {
	value *ProxyType
	isSet bool
}

func (v NullableProxyType) Get() *ProxyType {
	return v.value
}

func (v *NullableProxyType) Set(val *ProxyType) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyType) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyType(val *ProxyType) *NullableProxyType {
	return &NullableProxyType{value: val, isSet: true}
}

func (v NullableProxyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

