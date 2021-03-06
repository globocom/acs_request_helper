/*
 * This file is part of Graylog Plugin Oauth.
 *
 * Graylog Plugin Oauth is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * Graylog Plugin Oauth is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Foobar. If not, see <https://www.gnu.org/licenses/>
 */

package acsrequesthelper

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"net/url"
	"sort"
	"strings"
)

type fields = map[string]string

type RequestObject struct {
	host             string
	port             int
	apiKey           string
	secretKey        string
	params           fields
	encodedSignature string
}

func BuildRequestObject(host string, port int, apiKey string, secretKey string, params fields) RequestObject {
	ro := RequestObject{
		host:      host,
		port:      port,
		apiKey:    apiKey,
		secretKey: secretKey,
		params:    params,
	}

	ro.encodedSignature = ro.buildSignature()
	return ro
}

func (ro RequestObject) BuildQueryString() string {
	params := copyMap(ro.params)
	params["apiKey"] = ro.apiKey
	params["signature"] = ro.encodedSignature

	return toSortedString(params)
}

func (ro RequestObject) buildSignature() string {
	params := copyMap(ro.params)
	params["apiKey"] = ro.apiKey

	escaped := escapeAndToLowerCase(params)
	sorted := toSortedString(escaped)
	return strings.Replace(enc(sorted, ro.secretKey), "_", "%2F", -1)
}

func enc(signature string, secretKey string) string {
	h := hmac.New(sha1.New, []byte(secretKey))
	io.WriteString(h, signature)
	return url.QueryEscape(base64.URLEncoding.EncodeToString(h.Sum(nil)))
}

func toSortedString(params fields) string {
	var tokens []string
	for k, v := range params {
		tokens = append(tokens, k+"="+v)
	}
	sort.Strings(tokens)
	return strings.Join(tokens, "&")
}

func escapeAndToLowerCase(params fields) fields {
	var newFields = fields{}
	for k, v := range params {
		lk := strings.ToLower(k)
		lv := strings.ToLower(v)
		lev := url.QueryEscape(lv)
		rlev := strings.Replace(lev, "+", "%20", -1)

		newFields[lk] = rlev
	}

	return newFields
}

func copyMap(in map[string]string) map[string]string {
	var out = map[string]string{}
	for k, v := range in {
		out[k] = v
	}
	return out
}
