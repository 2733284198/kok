// Code generated by kok; DO NOT EDIT.
// github.com/RussellLuo/kok

package profilesvc

import (
	"context"
	"net/http"

	httpcodec "github.com/RussellLuo/kok/pkg/codec/httpv2"
	"github.com/RussellLuo/kok/pkg/oasv2"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func NewHTTPRouter(svc Service, codecs httpcodec.Codecs) chi.Router {
	return NewHTTPRouterWithOAS(svc, codecs, nil)
}

func NewHTTPRouterWithOAS(svc Service, codecs httpcodec.Codecs, schema oasv2.Schema) chi.Router {
	r := chi.NewRouter()

	if schema != nil {
		r.Method("GET", "/api", oasv2.Handler(OASv2APIDoc, schema))
	}

	var codec httpcodec.Codec
	var options []kithttp.ServerOption

	codec = codecs.EncodeDecoder("DeleteAddress")
	r.Method(
		"DELETE", "/profiles/{id}/addresses/{addressID}",
		kithttp.NewServer(
			MakeEndpointOfDeleteAddress(svc),
			decodeDeleteAddressRequest(codec),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(options,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("DeleteProfile")
	r.Method(
		"DELETE", "/profiles/{id}",
		kithttp.NewServer(
			MakeEndpointOfDeleteProfile(svc),
			decodeDeleteProfileRequest(codec),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(options,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("GetAddress")
	r.Method(
		"GET", "/profiles/{id}/addresses/{addressID}",
		kithttp.NewServer(
			MakeEndpointOfGetAddress(svc),
			decodeGetAddressRequest(codec),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(options,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("GetAddresses")
	r.Method(
		"GET", "/profiles/{id}/addresses",
		kithttp.NewServer(
			MakeEndpointOfGetAddresses(svc),
			decodeGetAddressesRequest(codec),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(options,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("GetProfile")
	r.Method(
		"GET", "/profiles/{id}",
		kithttp.NewServer(
			MakeEndpointOfGetProfile(svc),
			decodeGetProfileRequest(codec),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(options,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("PatchProfile")
	r.Method(
		"PATCH", "/profiles/{id}",
		kithttp.NewServer(
			MakeEndpointOfPatchProfile(svc),
			decodePatchProfileRequest(codec),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(options,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("PostAddress")
	r.Method(
		"POST", "/profiles/{id}/addresses",
		kithttp.NewServer(
			MakeEndpointOfPostAddress(svc),
			decodePostAddressRequest(codec),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(options,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("PostProfile")
	r.Method(
		"POST", "/profiles",
		kithttp.NewServer(
			MakeEndpointOfPostProfile(svc),
			decodePostProfileRequest(codec),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(options,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	codec = codecs.EncodeDecoder("PutProfile")
	r.Method(
		"PUT", "/profiles/{id}",
		kithttp.NewServer(
			MakeEndpointOfPutProfile(svc),
			decodePutProfileRequest(codec),
			httpcodec.MakeResponseEncoder(codec, 200),
			append(options,
				kithttp.ServerErrorEncoder(httpcodec.MakeErrorEncoder(codec)),
			)...,
		),
	)

	return r
}

func decodeDeleteAddressRequest(codec httpcodec.Codec) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req DeleteAddressRequest

		id := chi.URLParam(r, "id")
		if err := codec.DecodeRequestParam("id", id, &_req.Id); err != nil {
			return nil, err
		}

		addressID := chi.URLParam(r, "addressID")
		if err := codec.DecodeRequestParam("addressID", addressID, &_req.AddressID); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodeDeleteProfileRequest(codec httpcodec.Codec) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req DeleteProfileRequest

		id := chi.URLParam(r, "id")
		if err := codec.DecodeRequestParam("id", id, &_req.Id); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodeGetAddressRequest(codec httpcodec.Codec) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req GetAddressRequest

		id := chi.URLParam(r, "id")
		if err := codec.DecodeRequestParam("id", id, &_req.Id); err != nil {
			return nil, err
		}

		addressID := chi.URLParam(r, "addressID")
		if err := codec.DecodeRequestParam("addressID", addressID, &_req.AddressID); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodeGetAddressesRequest(codec httpcodec.Codec) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req GetAddressesRequest

		id := chi.URLParam(r, "id")
		if err := codec.DecodeRequestParam("id", id, &_req.Id); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodeGetProfileRequest(codec httpcodec.Codec) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req GetProfileRequest

		id := chi.URLParam(r, "id")
		if err := codec.DecodeRequestParam("id", id, &_req.Id); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodePatchProfileRequest(codec httpcodec.Codec) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req PatchProfileRequest

		if err := codec.DecodeRequestBody(r.Body, &_req); err != nil {
			return nil, err
		}

		id := chi.URLParam(r, "id")
		if err := codec.DecodeRequestParam("id", id, &_req.Id); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodePostAddressRequest(codec httpcodec.Codec) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req PostAddressRequest

		if err := codec.DecodeRequestBody(r.Body, &_req); err != nil {
			return nil, err
		}

		id := chi.URLParam(r, "id")
		if err := codec.DecodeRequestParam("id", id, &_req.Id); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodePostProfileRequest(codec httpcodec.Codec) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req PostProfileRequest

		if err := codec.DecodeRequestBody(r.Body, &_req); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}

func decodePutProfileRequest(codec httpcodec.Codec) kithttp.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		var _req PutProfileRequest

		if err := codec.DecodeRequestBody(r.Body, &_req); err != nil {
			return nil, err
		}

		id := chi.URLParam(r, "id")
		if err := codec.DecodeRequestParam("id", id, &_req.Id); err != nil {
			return nil, err
		}

		return &_req, nil
	}
}
