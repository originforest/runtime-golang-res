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
	ContentLength int64         // event content length
	ContentType   string        // content MIME type
	Body          io.ReadCloser // content IO
	Result        EventResult   // object who allow process result directly
}

// A Result is a custom serializable object that give response buffer to event
type Result interface {
	Bytes() []byte
}

// the ResultBytes adapt a byte slice to allow it use as a Result
type ResultBytes []byte

// Bytes implement Result interface
func (r ResultBytes) Bytes() []byte {
	return []byte(r)
}

// EvtHandler define standard FaaS event handler
type EvtHandler func(*EventRT, *ContextRT) (Result, error)
