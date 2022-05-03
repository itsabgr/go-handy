package handy

import (
	"context"
	"github.com/valyala/fasthttp"
	"io"
)

func HTTP(ctx context.Context, method string, uri *fasthttp.URI, body io.Reader, headers map[string]string) (response []byte, statusCode int, err error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	//
	req.SetURI(uri)
	req.SetBodyStream(body, -1)
	req.Header.SetMethod(method)
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	//
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	//
	if deadline, ok := ctx.Deadline(); ok {
		err = fasthttp.DoDeadline(req, res, deadline)
	} else {
		err = fasthttp.Do(req, res)
	}
	//
	if err != nil {
		return nil, 0, err
	}
	//
	return res.Body(), res.StatusCode(), nil
}
