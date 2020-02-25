// Code generated by goa v3.0.10, DO NOT EDIT.
//
// tmess HTTP client encoders and decoders
//
// Command:
// $ goa gen terminal-chat/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	tmess "terminal-chat/gen/tmess"
	tmessviews "terminal-chat/gen/tmess/views"

	goahttp "goa.design/goa/v3/http"
)

// BuildLoginRequest instantiates a HTTP request object with method and path
// set to call the "tmess" service "login" endpoint
func (c *Client) BuildLoginRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: LoginTmessPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("tmess", "login", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeLoginRequest returns an encoder for requests sent to the tmess login
// server.
func EncodeLoginRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*tmess.LoginPayload)
		if !ok {
			return goahttp.ErrInvalidType("tmess", "login", "*tmess.LoginPayload", v)
		}
		req.SetBasicAuth(p.User, p.Password)
		return nil
	}
}

// DecodeLoginResponse returns a decoder for responses returned by the tmess
// login endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeLoginResponse may return the following errors:
//	- "unauthorized" (type tmess.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeLoginResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "login", err)
			}
			return body, nil
		case http.StatusUnauthorized:
			var (
				body LoginUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "login", err)
			}
			return nil, NewLoginUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("tmess", "login", resp.StatusCode, string(body))
		}
	}
}

// BuildEchoerRequest instantiates a HTTP request object with method and path
// set to call the "tmess" service "echoer" endpoint
func (c *Client) BuildEchoerRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	scheme := c.scheme
	switch c.scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	}
	u := &url.URL{Scheme: scheme, Host: c.host, Path: EchoerTmessPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("tmess", "echoer", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeEchoerRequest returns an encoder for requests sent to the tmess echoer
// server.
func EncodeEchoerRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*tmess.EchoerPayload)
		if !ok {
			return goahttp.ErrInvalidType("tmess", "echoer", "*tmess.EchoerPayload", v)
		}
		{
			head := p.Token
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeEchoerResponse returns a decoder for responses returned by the tmess
// echoer endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeEchoerResponse may return the following errors:
//	- "invalid-scopes" (type tmess.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type tmess.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeEchoerResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "echoer", err)
			}
			return body, nil
		case http.StatusForbidden:
			var (
				body EchoerInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "echoer", err)
			}
			return nil, NewEchoerInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body EchoerUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "echoer", err)
			}
			return nil, NewEchoerUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("tmess", "echoer", resp.StatusCode, string(body))
		}
	}
}

// BuildListenerRequest instantiates a HTTP request object with method and path
// set to call the "tmess" service "listener" endpoint
func (c *Client) BuildListenerRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	scheme := c.scheme
	switch c.scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	}
	u := &url.URL{Scheme: scheme, Host: c.host, Path: ListenerTmessPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("tmess", "listener", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeListenerRequest returns an encoder for requests sent to the tmess
// listener server.
func EncodeListenerRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*tmess.ListenerPayload)
		if !ok {
			return goahttp.ErrInvalidType("tmess", "listener", "*tmess.ListenerPayload", v)
		}
		{
			head := p.Token
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeListenerResponse returns a decoder for responses returned by the tmess
// listener endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeListenerResponse may return the following errors:
//	- "invalid-scopes" (type tmess.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type tmess.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeListenerResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			return nil, nil
		case http.StatusForbidden:
			var (
				body ListenerInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "listener", err)
			}
			return nil, NewListenerInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body ListenerUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "listener", err)
			}
			return nil, NewListenerUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("tmess", "listener", resp.StatusCode, string(body))
		}
	}
}

// BuildSummaryRequest instantiates a HTTP request object with method and path
// set to call the "tmess" service "summary" endpoint
func (c *Client) BuildSummaryRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	scheme := c.scheme
	switch c.scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	}
	u := &url.URL{Scheme: scheme, Host: c.host, Path: SummaryTmessPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("tmess", "summary", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSummaryRequest returns an encoder for requests sent to the tmess
// summary server.
func EncodeSummaryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*tmess.SummaryPayload)
		if !ok {
			return goahttp.ErrInvalidType("tmess", "summary", "*tmess.SummaryPayload", v)
		}
		{
			head := p.Token
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeSummaryResponse returns a decoder for responses returned by the tmess
// summary endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeSummaryResponse may return the following errors:
//	- "invalid-scopes" (type tmess.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type tmess.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeSummaryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SummaryResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "summary", err)
			}
			p := NewSummaryChatSummaryCollectionOK(body)
			view := "default"
			vres := tmessviews.ChatSummaryCollection{Projected: p, View: view}
			if err = tmessviews.ValidateChatSummaryCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("tmess", "summary", err)
			}
			res := tmess.NewChatSummaryCollection(vres)
			return res, nil
		case http.StatusForbidden:
			var (
				body SummaryInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "summary", err)
			}
			return nil, NewSummaryInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body SummaryUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "summary", err)
			}
			return nil, NewSummaryUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("tmess", "summary", resp.StatusCode, string(body))
		}
	}
}

// BuildSubscribeRequest instantiates a HTTP request object with method and
// path set to call the "tmess" service "subscribe" endpoint
func (c *Client) BuildSubscribeRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	scheme := c.scheme
	switch c.scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	}
	u := &url.URL{Scheme: scheme, Host: c.host, Path: SubscribeTmessPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("tmess", "subscribe", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSubscribeRequest returns an encoder for requests sent to the tmess
// subscribe server.
func EncodeSubscribeRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*tmess.SubscribePayload)
		if !ok {
			return goahttp.ErrInvalidType("tmess", "subscribe", "*tmess.SubscribePayload", v)
		}
		{
			head := p.Token
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeSubscribeResponse returns a decoder for responses returned by the
// tmess subscribe endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeSubscribeResponse may return the following errors:
//	- "invalid-scopes" (type tmess.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type tmess.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeSubscribeResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SubscribeResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "subscribe", err)
			}
			err = ValidateSubscribeResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("tmess", "subscribe", err)
			}
			res := NewSubscribeEventOK(&body)
			return res, nil
		case http.StatusForbidden:
			var (
				body SubscribeInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "subscribe", err)
			}
			return nil, NewSubscribeInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body SubscribeUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "subscribe", err)
			}
			return nil, NewSubscribeUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("tmess", "subscribe", resp.StatusCode, string(body))
		}
	}
}

// BuildHistoryRequest instantiates a HTTP request object with method and path
// set to call the "tmess" service "history" endpoint
func (c *Client) BuildHistoryRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	scheme := c.scheme
	switch c.scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	}
	u := &url.URL{Scheme: scheme, Host: c.host, Path: HistoryTmessPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("tmess", "history", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeHistoryRequest returns an encoder for requests sent to the tmess
// history server.
func EncodeHistoryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*tmess.HistoryPayload)
		if !ok {
			return goahttp.ErrInvalidType("tmess", "history", "*tmess.HistoryPayload", v)
		}
		{
			head := p.Token
			req.Header.Set("Authorization", head)
		}
		values := req.URL.Query()
		if p.View != nil {
			values.Add("view", *p.View)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeHistoryResponse returns a decoder for responses returned by the tmess
// history endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeHistoryResponse may return the following errors:
//	- "invalid-scopes" (type tmess.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type tmess.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeHistoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body HistoryResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "history", err)
			}
			p := NewHistoryChatSummaryOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &tmessviews.ChatSummary{Projected: p, View: view}
			if err = tmessviews.ValidateChatSummary(vres); err != nil {
				return nil, goahttp.ErrValidationError("tmess", "history", err)
			}
			res := tmess.NewChatSummary(vres)
			return res, nil
		case http.StatusForbidden:
			var (
				body HistoryInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "history", err)
			}
			return nil, NewHistoryInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body HistoryUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tmess", "history", err)
			}
			return nil, NewHistoryUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("tmess", "history", resp.StatusCode, string(body))
		}
	}
}