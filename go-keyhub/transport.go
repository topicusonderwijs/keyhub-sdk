package keyhub

import (
	"golang.org/x/oauth2"
	"net/http"
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
