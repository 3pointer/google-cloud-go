// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package apiv1

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	adminpb "google.golang.org/genproto/googleapis/firestore/admin/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newFirestoreAdminClientHook clientHook

// FirestoreAdminCallOptions contains the retry settings for each method of FirestoreAdminClient.
type FirestoreAdminCallOptions struct {
	CreateIndex     []gax.CallOption
	ListIndexes     []gax.CallOption
	GetIndex        []gax.CallOption
	DeleteIndex     []gax.CallOption
	GetField        []gax.CallOption
	UpdateField     []gax.CallOption
	ListFields      []gax.CallOption
	ExportDocuments []gax.CallOption
	ImportDocuments []gax.CallOption
}

func defaultFirestoreAdminGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("firestore.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("firestore.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://firestore.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultFirestoreAdminCallOptions() *FirestoreAdminCallOptions {
	return &FirestoreAdminCallOptions{
		CreateIndex: []gax.CallOption{},
		ListIndexes: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Internal,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetIndex: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Internal,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteIndex: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Internal,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetField: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Internal,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateField: []gax.CallOption{},
		ListFields: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Internal,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ExportDocuments: []gax.CallOption{},
		ImportDocuments: []gax.CallOption{},
	}
}

// internalFirestoreAdminClient is an interface that defines the methods availaible from Cloud Firestore API.
type internalFirestoreAdminClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateIndex(context.Context, *adminpb.CreateIndexRequest, ...gax.CallOption) (*CreateIndexOperation, error)
	CreateIndexOperation(name string) *CreateIndexOperation
	ListIndexes(context.Context, *adminpb.ListIndexesRequest, ...gax.CallOption) *IndexIterator
	GetIndex(context.Context, *adminpb.GetIndexRequest, ...gax.CallOption) (*adminpb.Index, error)
	DeleteIndex(context.Context, *adminpb.DeleteIndexRequest, ...gax.CallOption) error
	GetField(context.Context, *adminpb.GetFieldRequest, ...gax.CallOption) (*adminpb.Field, error)
	UpdateField(context.Context, *adminpb.UpdateFieldRequest, ...gax.CallOption) (*UpdateFieldOperation, error)
	UpdateFieldOperation(name string) *UpdateFieldOperation
	ListFields(context.Context, *adminpb.ListFieldsRequest, ...gax.CallOption) *FieldIterator
	ExportDocuments(context.Context, *adminpb.ExportDocumentsRequest, ...gax.CallOption) (*ExportDocumentsOperation, error)
	ExportDocumentsOperation(name string) *ExportDocumentsOperation
	ImportDocuments(context.Context, *adminpb.ImportDocumentsRequest, ...gax.CallOption) (*ImportDocumentsOperation, error)
	ImportDocumentsOperation(name string) *ImportDocumentsOperation
}

// FirestoreAdminClient is a client for interacting with Cloud Firestore API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Operations are created by service FirestoreAdmin, but are accessed via
// service google.longrunning.Operations.
type FirestoreAdminClient struct {
	// The internal transport-dependent client.
	internalClient internalFirestoreAdminClient

	// The call options for this service.
	CallOptions *FirestoreAdminCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FirestoreAdminClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FirestoreAdminClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *FirestoreAdminClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateIndex creates a composite index. This returns a google.longrunning.Operation
// which may be used to track the status of the creation. The metadata for
// the operation will be the type IndexOperationMetadata.
func (c *FirestoreAdminClient) CreateIndex(ctx context.Context, req *adminpb.CreateIndexRequest, opts ...gax.CallOption) (*CreateIndexOperation, error) {
	return c.internalClient.CreateIndex(ctx, req, opts...)
}

// CreateIndexOperation returns a new CreateIndexOperation from a given name.
// The name must be that of a previously created CreateIndexOperation, possibly from a different process.
func (c *FirestoreAdminClient) CreateIndexOperation(name string) *CreateIndexOperation {
	return c.internalClient.CreateIndexOperation(name)
}

// ListIndexes lists composite indexes.
func (c *FirestoreAdminClient) ListIndexes(ctx context.Context, req *adminpb.ListIndexesRequest, opts ...gax.CallOption) *IndexIterator {
	return c.internalClient.ListIndexes(ctx, req, opts...)
}

// GetIndex gets a composite index.
func (c *FirestoreAdminClient) GetIndex(ctx context.Context, req *adminpb.GetIndexRequest, opts ...gax.CallOption) (*adminpb.Index, error) {
	return c.internalClient.GetIndex(ctx, req, opts...)
}

// DeleteIndex deletes a composite index.
func (c *FirestoreAdminClient) DeleteIndex(ctx context.Context, req *adminpb.DeleteIndexRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteIndex(ctx, req, opts...)
}

// GetField gets the metadata and configuration for a Field.
func (c *FirestoreAdminClient) GetField(ctx context.Context, req *adminpb.GetFieldRequest, opts ...gax.CallOption) (*adminpb.Field, error) {
	return c.internalClient.GetField(ctx, req, opts...)
}

// UpdateField updates a field configuration. Currently, field updates apply only to
// single field index configuration. However, calls to
// FirestoreAdmin.UpdateField should provide a field mask to avoid
// changing any configuration that the caller isn’t aware of. The field mask
// should be specified as: { paths: "index_config" }.
//
// This call returns a google.longrunning.Operation which may be used to
// track the status of the field update. The metadata for
// the operation will be the type FieldOperationMetadata.
//
// To configure the default field settings for the database, use
// the special Field with resource name:
// projects/{project_id}/databases/{database_id}/collectionGroups/__default__/fields/*.
func (c *FirestoreAdminClient) UpdateField(ctx context.Context, req *adminpb.UpdateFieldRequest, opts ...gax.CallOption) (*UpdateFieldOperation, error) {
	return c.internalClient.UpdateField(ctx, req, opts...)
}

// UpdateFieldOperation returns a new UpdateFieldOperation from a given name.
// The name must be that of a previously created UpdateFieldOperation, possibly from a different process.
func (c *FirestoreAdminClient) UpdateFieldOperation(name string) *UpdateFieldOperation {
	return c.internalClient.UpdateFieldOperation(name)
}

// ListFields lists the field configuration and metadata for this database.
//
// Currently, FirestoreAdmin.ListFields only supports listing fields
// that have been explicitly overridden. To issue this query, call
// FirestoreAdmin.ListFields with the filter set to
// indexConfig.usesAncestorConfig:false.
func (c *FirestoreAdminClient) ListFields(ctx context.Context, req *adminpb.ListFieldsRequest, opts ...gax.CallOption) *FieldIterator {
	return c.internalClient.ListFields(ctx, req, opts...)
}

// ExportDocuments exports a copy of all or a subset of documents from Google Cloud Firestore
// to another storage system, such as Google Cloud Storage. Recent updates to
// documents may not be reflected in the export. The export occurs in the
// background and its progress can be monitored and managed via the
// Operation resource that is created. The output of an export may only be
// used once the associated operation is done. If an export operation is
// cancelled before completion it may leave partial data behind in Google
// Cloud Storage.
func (c *FirestoreAdminClient) ExportDocuments(ctx context.Context, req *adminpb.ExportDocumentsRequest, opts ...gax.CallOption) (*ExportDocumentsOperation, error) {
	return c.internalClient.ExportDocuments(ctx, req, opts...)
}

// ExportDocumentsOperation returns a new ExportDocumentsOperation from a given name.
// The name must be that of a previously created ExportDocumentsOperation, possibly from a different process.
func (c *FirestoreAdminClient) ExportDocumentsOperation(name string) *ExportDocumentsOperation {
	return c.internalClient.ExportDocumentsOperation(name)
}

// ImportDocuments imports documents into Google Cloud Firestore. Existing documents with the
// same name are overwritten. The import occurs in the background and its
// progress can be monitored and managed via the Operation resource that is
// created. If an ImportDocuments operation is cancelled, it is possible
// that a subset of the data has already been imported to Cloud Firestore.
func (c *FirestoreAdminClient) ImportDocuments(ctx context.Context, req *adminpb.ImportDocumentsRequest, opts ...gax.CallOption) (*ImportDocumentsOperation, error) {
	return c.internalClient.ImportDocuments(ctx, req, opts...)
}

// ImportDocumentsOperation returns a new ImportDocumentsOperation from a given name.
// The name must be that of a previously created ImportDocumentsOperation, possibly from a different process.
func (c *FirestoreAdminClient) ImportDocumentsOperation(name string) *ImportDocumentsOperation {
	return c.internalClient.ImportDocumentsOperation(name)
}

// firestoreAdminGRPCClient is a client for interacting with Cloud Firestore API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type firestoreAdminGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing FirestoreAdminClient
	CallOptions **FirestoreAdminCallOptions

	// The gRPC API client.
	firestoreAdminClient adminpb.FirestoreAdminClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewFirestoreAdminClient creates a new firestore admin client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Operations are created by service FirestoreAdmin, but are accessed via
// service google.longrunning.Operations.
func NewFirestoreAdminClient(ctx context.Context, opts ...option.ClientOption) (*FirestoreAdminClient, error) {
	clientOpts := defaultFirestoreAdminGRPCClientOptions()
	if newFirestoreAdminClientHook != nil {
		hookOpts, err := newFirestoreAdminClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := FirestoreAdminClient{CallOptions: defaultFirestoreAdminCallOptions()}

	c := &firestoreAdminGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		firestoreAdminClient: adminpb.NewFirestoreAdminClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *firestoreAdminGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *firestoreAdminGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *firestoreAdminGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *firestoreAdminGRPCClient) CreateIndex(ctx context.Context, req *adminpb.CreateIndexRequest, opts ...gax.CallOption) (*CreateIndexOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateIndex[0:len((*c.CallOptions).CreateIndex):len((*c.CallOptions).CreateIndex)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firestoreAdminClient.CreateIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *firestoreAdminGRPCClient) ListIndexes(ctx context.Context, req *adminpb.ListIndexesRequest, opts ...gax.CallOption) *IndexIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListIndexes[0:len((*c.CallOptions).ListIndexes):len((*c.CallOptions).ListIndexes)], opts...)
	it := &IndexIterator{}
	req = proto.Clone(req).(*adminpb.ListIndexesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*adminpb.Index, string, error) {
		var resp *adminpb.ListIndexesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.firestoreAdminClient.ListIndexes(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetIndexes(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

func (c *firestoreAdminGRPCClient) GetIndex(ctx context.Context, req *adminpb.GetIndexRequest, opts ...gax.CallOption) (*adminpb.Index, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIndex[0:len((*c.CallOptions).GetIndex):len((*c.CallOptions).GetIndex)], opts...)
	var resp *adminpb.Index
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firestoreAdminClient.GetIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firestoreAdminGRPCClient) DeleteIndex(ctx context.Context, req *adminpb.DeleteIndexRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteIndex[0:len((*c.CallOptions).DeleteIndex):len((*c.CallOptions).DeleteIndex)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.firestoreAdminClient.DeleteIndex(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *firestoreAdminGRPCClient) GetField(ctx context.Context, req *adminpb.GetFieldRequest, opts ...gax.CallOption) (*adminpb.Field, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetField[0:len((*c.CallOptions).GetField):len((*c.CallOptions).GetField)], opts...)
	var resp *adminpb.Field
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firestoreAdminClient.GetField(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *firestoreAdminGRPCClient) UpdateField(ctx context.Context, req *adminpb.UpdateFieldRequest, opts ...gax.CallOption) (*UpdateFieldOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "field.name", url.QueryEscape(req.GetField().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateField[0:len((*c.CallOptions).UpdateField):len((*c.CallOptions).UpdateField)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firestoreAdminClient.UpdateField(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateFieldOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *firestoreAdminGRPCClient) ListFields(ctx context.Context, req *adminpb.ListFieldsRequest, opts ...gax.CallOption) *FieldIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListFields[0:len((*c.CallOptions).ListFields):len((*c.CallOptions).ListFields)], opts...)
	it := &FieldIterator{}
	req = proto.Clone(req).(*adminpb.ListFieldsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*adminpb.Field, string, error) {
		var resp *adminpb.ListFieldsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.firestoreAdminClient.ListFields(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetFields(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

func (c *firestoreAdminGRPCClient) ExportDocuments(ctx context.Context, req *adminpb.ExportDocumentsRequest, opts ...gax.CallOption) (*ExportDocumentsOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ExportDocuments[0:len((*c.CallOptions).ExportDocuments):len((*c.CallOptions).ExportDocuments)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firestoreAdminClient.ExportDocuments(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ExportDocumentsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *firestoreAdminGRPCClient) ImportDocuments(ctx context.Context, req *adminpb.ImportDocumentsRequest, opts ...gax.CallOption) (*ImportDocumentsOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ImportDocuments[0:len((*c.CallOptions).ImportDocuments):len((*c.CallOptions).ImportDocuments)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.firestoreAdminClient.ImportDocuments(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportDocumentsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateIndexOperation manages a long-running operation from CreateIndex.
type CreateIndexOperation struct {
	lro *longrunning.Operation
}

// CreateIndexOperation returns a new CreateIndexOperation from a given name.
// The name must be that of a previously created CreateIndexOperation, possibly from a different process.
func (c *firestoreAdminGRPCClient) CreateIndexOperation(name string) *CreateIndexOperation {
	return &CreateIndexOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateIndexOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*adminpb.Index, error) {
	var resp adminpb.Index
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateIndexOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*adminpb.Index, error) {
	var resp adminpb.Index
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateIndexOperation) Metadata() (*adminpb.IndexOperationMetadata, error) {
	var meta adminpb.IndexOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateIndexOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateIndexOperation) Name() string {
	return op.lro.Name()
}

// ExportDocumentsOperation manages a long-running operation from ExportDocuments.
type ExportDocumentsOperation struct {
	lro *longrunning.Operation
}

// ExportDocumentsOperation returns a new ExportDocumentsOperation from a given name.
// The name must be that of a previously created ExportDocumentsOperation, possibly from a different process.
func (c *firestoreAdminGRPCClient) ExportDocumentsOperation(name string) *ExportDocumentsOperation {
	return &ExportDocumentsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportDocumentsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*adminpb.ExportDocumentsResponse, error) {
	var resp adminpb.ExportDocumentsResponse
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ExportDocumentsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*adminpb.ExportDocumentsResponse, error) {
	var resp adminpb.ExportDocumentsResponse
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ExportDocumentsOperation) Metadata() (*adminpb.ExportDocumentsMetadata, error) {
	var meta adminpb.ExportDocumentsMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportDocumentsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportDocumentsOperation) Name() string {
	return op.lro.Name()
}

// ImportDocumentsOperation manages a long-running operation from ImportDocuments.
type ImportDocumentsOperation struct {
	lro *longrunning.Operation
}

// ImportDocumentsOperation returns a new ImportDocumentsOperation from a given name.
// The name must be that of a previously created ImportDocumentsOperation, possibly from a different process.
func (c *firestoreAdminGRPCClient) ImportDocumentsOperation(name string) *ImportDocumentsOperation {
	return &ImportDocumentsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ImportDocumentsOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *ImportDocumentsOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *ImportDocumentsOperation) Metadata() (*adminpb.ImportDocumentsMetadata, error) {
	var meta adminpb.ImportDocumentsMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ImportDocumentsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ImportDocumentsOperation) Name() string {
	return op.lro.Name()
}

// UpdateFieldOperation manages a long-running operation from UpdateField.
type UpdateFieldOperation struct {
	lro *longrunning.Operation
}

// UpdateFieldOperation returns a new UpdateFieldOperation from a given name.
// The name must be that of a previously created UpdateFieldOperation, possibly from a different process.
func (c *firestoreAdminGRPCClient) UpdateFieldOperation(name string) *UpdateFieldOperation {
	return &UpdateFieldOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateFieldOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*adminpb.Field, error) {
	var resp adminpb.Field
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateFieldOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*adminpb.Field, error) {
	var resp adminpb.Field
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateFieldOperation) Metadata() (*adminpb.FieldOperationMetadata, error) {
	var meta adminpb.FieldOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateFieldOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateFieldOperation) Name() string {
	return op.lro.Name()
}

// FieldIterator manages a stream of *adminpb.Field.
type FieldIterator struct {
	items    []*adminpb.Field
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*adminpb.Field, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *FieldIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *FieldIterator) Next() (*adminpb.Field, error) {
	var item *adminpb.Field
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *FieldIterator) bufLen() int {
	return len(it.items)
}

func (it *FieldIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// IndexIterator manages a stream of *adminpb.Index.
type IndexIterator struct {
	items    []*adminpb.Index
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*adminpb.Index, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *IndexIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *IndexIterator) Next() (*adminpb.Index, error) {
	var item *adminpb.Index
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *IndexIterator) bufLen() int {
	return len(it.items)
}

func (it *IndexIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
