// Code generated by goa v3.0.4, DO NOT EDIT.
//
// chatter gRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o
// $(GOPATH)/src/goa.design/examples/streaming

package server

import (
	"context"
	"strings"

	chatter "goa.design/examples/streaming/gen/chatter"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/metadata"
)

// EncodeLoginResponse encodes responses from the "chatter" service "login"
// endpoint.
func EncodeLoginResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "login", "string", v)
	}
	resp := NewLoginResponse(result)
	return resp, nil
}

// DecodeLoginRequest decodes requests sent to "chatter" service "login"
// endpoint.
func DecodeLoginRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		user     string
		password string
		err      error
	)
	{
		if vals := md.Get("user"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("user", "metadata"))
		} else {
			user = vals[0]
		}
		if vals := md.Get("password"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("password", "metadata"))
		} else {
			password = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.LoginPayload
	{
		payload = NewLoginPayload(user, password)
	}
	return payload, nil
}

// DecodeEchoerRequest decodes requests sent to "chatter" service "echoer"
// endpoint.
func DecodeEchoerRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.EchoerPayload
	{
		payload = NewEchoerPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// DecodeListenerRequest decodes requests sent to "chatter" service "listener"
// endpoint.
func DecodeListenerRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.ListenerPayload
	{
		payload = NewListenerPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// DecodeSummaryRequest decodes requests sent to "chatter" service "summary"
// endpoint.
func DecodeSummaryRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.SummaryPayload
	{
		payload = NewSummaryPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// DecodeSubscribeRequest decodes requests sent to "chatter" service
// "subscribe" endpoint.
func DecodeSubscribeRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.SubscribePayload
	{
		payload = NewSubscribePayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}

// DecodeHistoryRequest decodes requests sent to "chatter" service "history"
// endpoint.
func DecodeHistoryRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		view  *string
		token string
		err   error
	)
	{
		if vals := md.Get("view"); len(vals) > 0 {
			view = &vals[0]
		}
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	if err != nil {
		return nil, err
	}
	var payload *chatter.HistoryPayload
	{
		payload = NewHistoryPayload(view, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, nil
}
