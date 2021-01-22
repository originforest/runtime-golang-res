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

// EventResult define a way to write result to FaaS event
type EventResult struct {
	Header     http.Header // header
	StatusCode int         // result HTTP code
	Body       io.Writer   // Body data IO
}

// EventRT define a FaaS runtime event object
type EventRT struct {
	ContentLength uint64        // event content length
	ContentType   string        // content MIME type
	Body          io.ReadCloser // content IO
	Result        EventResult   // object who allow process result directly
}
