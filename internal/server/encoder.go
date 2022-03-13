package server

import (
	"kratos-realworld/internal/errors"
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
)

func errorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-type", "application"+codec.Name())
	if 99 < se.Code && 600 > se.Code {
		w.WriteHeader(se.Code)
	} else {
		w.WriteHeader(500)
	}
	_, _ = w.Write(body)
}
