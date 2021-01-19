/*
 * FaaS develop resource
 * Funtion Context model
 * --------------------
 * Copyright (c) 2021 Copyright bespin All Rights Reserved.
 * Author: Moyu Su (Fiathux Su)
 * Date: 2021-01-14
 * Desc:
 */

package model

import (
	"net/http"
	"net/url"
)

type ContextRT struct {
	FnID        string
	ResidentID  string
	ReqestID    string
	ResToken    string
	Method      string
	EventSource string
	Resource    string
	Host        string
	Query       url.Values
	Header      http.Header
}
