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

// A ContextRT contain FaaS event meta-datas
type ContextRT struct {
	FnID      string                 // function ID
	TenantID  string                 // tenant ID
	RequestID string                 // current request ID
	ResToken  string                 // cloud resource token
	Method    string                 // method
	EventName string                 // event source name
	Resource  string                 // event resource path
	Host      string                 // host name of event trigger's
	Query     url.Values             // query data in event URI
	Header    http.Header            // event headers
	Extra     map[string]interface{} // extra values
}
