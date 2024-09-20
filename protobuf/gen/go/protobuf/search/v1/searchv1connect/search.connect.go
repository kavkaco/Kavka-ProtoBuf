// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: protobuf/search/v1/search.proto

package searchv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/kavkaco/Kavka-Core/protobuf/gen/go/protobuf/search/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SearchServiceName is the fully-qualified name of the SearchService service.
	SearchServiceName = "search.v1.SearchService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SearchServiceSearchProcedure is the fully-qualified name of the SearchService's Search RPC.
	SearchServiceSearchProcedure = "/search.v1.SearchService/Search"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	searchServiceServiceDescriptor      = v1.File_protobuf_search_v1_search_proto.Services().ByName("SearchService")
	searchServiceSearchMethodDescriptor = searchServiceServiceDescriptor.Methods().ByName("Search")
)

// SearchServiceClient is a client for the search.v1.SearchService service.
type SearchServiceClient interface {
	Search(context.Context, *connect.Request[v1.SearchRequest]) (*connect.Response[v1.SearchResponse], error)
}

// NewSearchServiceClient constructs a client for the search.v1.SearchService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSearchServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SearchServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &searchServiceClient{
		search: connect.NewClient[v1.SearchRequest, v1.SearchResponse](
			httpClient,
			baseURL+SearchServiceSearchProcedure,
			connect.WithSchema(searchServiceSearchMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// searchServiceClient implements SearchServiceClient.
type searchServiceClient struct {
	search *connect.Client[v1.SearchRequest, v1.SearchResponse]
}

// Search calls search.v1.SearchService.Search.
func (c *searchServiceClient) Search(ctx context.Context, req *connect.Request[v1.SearchRequest]) (*connect.Response[v1.SearchResponse], error) {
	return c.search.CallUnary(ctx, req)
}

// SearchServiceHandler is an implementation of the search.v1.SearchService service.
type SearchServiceHandler interface {
	Search(context.Context, *connect.Request[v1.SearchRequest]) (*connect.Response[v1.SearchResponse], error)
}

// NewSearchServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSearchServiceHandler(svc SearchServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	searchServiceSearchHandler := connect.NewUnaryHandler(
		SearchServiceSearchProcedure,
		svc.Search,
		connect.WithSchema(searchServiceSearchMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/search.v1.SearchService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SearchServiceSearchProcedure:
			searchServiceSearchHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSearchServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSearchServiceHandler struct{}

func (UnimplementedSearchServiceHandler) Search(context.Context, *connect.Request[v1.SearchRequest]) (*connect.Response[v1.SearchResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("search.v1.SearchService.Search is not implemented"))
}
