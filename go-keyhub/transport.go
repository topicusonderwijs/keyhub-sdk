/* Licensed to the Apache Software Foundation (ASF) under one or more
   contributor license agreements.  See the NOTICE file distributed with
   this work for additional information regarding copyright ownership.
   The ASF licenses this file to You under the Apache License, Version 2.0
   (the "License"); you may not use this file except in compliance with
   the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License. */

package keyhub

import (
	"net/http"

	"golang.org/x/oauth2"
)

type Transport struct {
	Base http.RoundTripper
}

// Based on oauth2.Transport.RoundTrip()
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.Base.(*oauth2.Transport).Source.Token()
	if err != nil {
		return nil, err
	}

	req2 := cloneRequest(req) // per RoundTripper contract
	req2.Header.Add("topicus-Vault-session", token.Extra("vaultSession").(string))
	res, err := t.Base.RoundTrip(req2)
	return res, err
}

// See oauth2.Transport.cloneRequest()
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
