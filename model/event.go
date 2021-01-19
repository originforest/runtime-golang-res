/*
 * FaaS develop resource
 * Funtion Event model
 * --------------------
 * Copyright (c) 2021 Copyright bespin All Rights Reserved.
 * Author: Moyu Su (Fiathux Su)
 * Date: 2021-01-14
 * Desc:
 */

package model

import (
	"io"
	"net/http"
)

//
type EventResult struct {
	Header     http.Header
	StatusCode int
	Body       io.Writer
}

type EventRT struct {
	ContentLength uint64
	ContentType   string
	Body          io.ReadCloser
	Result        EventResult
}
