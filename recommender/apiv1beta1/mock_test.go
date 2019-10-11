// Copyright 2019 Google LLC
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

// Code generated by gapic-generator. DO NOT EDIT.

package recommender

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/api/option"
	recommenderpb "google.golang.org/genproto/googleapis/cloud/recommender/v1beta1"

	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockRecommenderServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	recommenderpb.RecommenderServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockRecommenderServer) ListRecommendations(ctx context.Context, req *recommenderpb.ListRecommendationsRequest) (*recommenderpb.ListRecommendationsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*recommenderpb.ListRecommendationsResponse), nil
}

func (s *mockRecommenderServer) GetRecommendation(ctx context.Context, req *recommenderpb.GetRecommendationRequest) (*recommenderpb.Recommendation, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*recommenderpb.Recommendation), nil
}

func (s *mockRecommenderServer) MarkRecommendationClaimed(ctx context.Context, req *recommenderpb.MarkRecommendationClaimedRequest) (*recommenderpb.Recommendation, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*recommenderpb.Recommendation), nil
}

func (s *mockRecommenderServer) MarkRecommendationSucceeded(ctx context.Context, req *recommenderpb.MarkRecommendationSucceededRequest) (*recommenderpb.Recommendation, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*recommenderpb.Recommendation), nil
}

func (s *mockRecommenderServer) MarkRecommendationFailed(ctx context.Context, req *recommenderpb.MarkRecommendationFailedRequest) (*recommenderpb.Recommendation, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*recommenderpb.Recommendation), nil
}

// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
	mockRecommender mockRecommenderServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	serv := grpc.NewServer()
	recommenderpb.RegisterRecommenderServer(serv, &mockRecommender)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clientOpt = option.WithGRPCConn(conn)

	os.Exit(m.Run())
}

func TestRecommenderListRecommendations(t *testing.T) {
	var nextPageToken string = ""
	var recommendationsElement *recommenderpb.Recommendation = &recommenderpb.Recommendation{}
	var recommendations = []*recommenderpb.Recommendation{recommendationsElement}
	var expectedResponse = &recommenderpb.ListRecommendationsResponse{
		NextPageToken:   nextPageToken,
		Recommendations: recommendations,
	}

	mockRecommender.err = nil
	mockRecommender.reqs = nil

	mockRecommender.resps = append(mockRecommender.resps[:0], expectedResponse)

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/recommenders/%s", "[PROJECT]", "[LOCATION]", "[RECOMMENDER]")
	var request = &recommenderpb.ListRecommendationsRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListRecommendations(context.Background(), request).Next()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockRecommender.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	want := (interface{})(expectedResponse.Recommendations[0])
	got := (interface{})(resp)
	var ok bool

	switch want := (want).(type) {
	case proto.Message:
		ok = proto.Equal(want, got.(proto.Message))
	default:
		ok = want == got
	}
	if !ok {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestRecommenderListRecommendationsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockRecommender.err = gstatus.Error(errCode, "test error")

	var formattedParent string = fmt.Sprintf("projects/%s/locations/%s/recommenders/%s", "[PROJECT]", "[LOCATION]", "[RECOMMENDER]")
	var request = &recommenderpb.ListRecommendationsRequest{
		Parent: formattedParent,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.ListRecommendations(context.Background(), request).Next()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestRecommenderGetRecommendation(t *testing.T) {
	var name2 string = "name2-1052831874"
	var description string = "description-1724546052"
	var recommenderSubtype string = "recommenderSubtype-1488504412"
	var etag string = "etag3123477"
	var expectedResponse = &recommenderpb.Recommendation{
		Name:               name2,
		Description:        description,
		RecommenderSubtype: recommenderSubtype,
		Etag:               etag,
	}

	mockRecommender.err = nil
	mockRecommender.reqs = nil

	mockRecommender.resps = append(mockRecommender.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", "[PROJECT]", "[LOCATION]", "[RECOMMENDER]", "[RECOMMENDATION]")
	var request = &recommenderpb.GetRecommendationRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetRecommendation(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockRecommender.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestRecommenderGetRecommendationError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockRecommender.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", "[PROJECT]", "[LOCATION]", "[RECOMMENDER]", "[RECOMMENDATION]")
	var request = &recommenderpb.GetRecommendationRequest{
		Name: formattedName,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.GetRecommendation(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestRecommenderMarkRecommendationClaimed(t *testing.T) {
	var name2 string = "name2-1052831874"
	var description string = "description-1724546052"
	var recommenderSubtype string = "recommenderSubtype-1488504412"
	var etag2 string = "etag2-1293302904"
	var expectedResponse = &recommenderpb.Recommendation{
		Name:               name2,
		Description:        description,
		RecommenderSubtype: recommenderSubtype,
		Etag:               etag2,
	}

	mockRecommender.err = nil
	mockRecommender.reqs = nil

	mockRecommender.resps = append(mockRecommender.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", "[PROJECT]", "[LOCATION]", "[RECOMMENDER]", "[RECOMMENDATION]")
	var etag string = "etag3123477"
	var request = &recommenderpb.MarkRecommendationClaimedRequest{
		Name: formattedName,
		Etag: etag,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.MarkRecommendationClaimed(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockRecommender.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestRecommenderMarkRecommendationClaimedError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockRecommender.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", "[PROJECT]", "[LOCATION]", "[RECOMMENDER]", "[RECOMMENDATION]")
	var etag string = "etag3123477"
	var request = &recommenderpb.MarkRecommendationClaimedRequest{
		Name: formattedName,
		Etag: etag,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.MarkRecommendationClaimed(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestRecommenderMarkRecommendationSucceeded(t *testing.T) {
	var name2 string = "name2-1052831874"
	var description string = "description-1724546052"
	var recommenderSubtype string = "recommenderSubtype-1488504412"
	var etag2 string = "etag2-1293302904"
	var expectedResponse = &recommenderpb.Recommendation{
		Name:               name2,
		Description:        description,
		RecommenderSubtype: recommenderSubtype,
		Etag:               etag2,
	}

	mockRecommender.err = nil
	mockRecommender.reqs = nil

	mockRecommender.resps = append(mockRecommender.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", "[PROJECT]", "[LOCATION]", "[RECOMMENDER]", "[RECOMMENDATION]")
	var etag string = "etag3123477"
	var request = &recommenderpb.MarkRecommendationSucceededRequest{
		Name: formattedName,
		Etag: etag,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.MarkRecommendationSucceeded(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockRecommender.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestRecommenderMarkRecommendationSucceededError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockRecommender.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", "[PROJECT]", "[LOCATION]", "[RECOMMENDER]", "[RECOMMENDATION]")
	var etag string = "etag3123477"
	var request = &recommenderpb.MarkRecommendationSucceededRequest{
		Name: formattedName,
		Etag: etag,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.MarkRecommendationSucceeded(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestRecommenderMarkRecommendationFailed(t *testing.T) {
	var name2 string = "name2-1052831874"
	var description string = "description-1724546052"
	var recommenderSubtype string = "recommenderSubtype-1488504412"
	var etag2 string = "etag2-1293302904"
	var expectedResponse = &recommenderpb.Recommendation{
		Name:               name2,
		Description:        description,
		RecommenderSubtype: recommenderSubtype,
		Etag:               etag2,
	}

	mockRecommender.err = nil
	mockRecommender.reqs = nil

	mockRecommender.resps = append(mockRecommender.resps[:0], expectedResponse)

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", "[PROJECT]", "[LOCATION]", "[RECOMMENDER]", "[RECOMMENDATION]")
	var etag string = "etag3123477"
	var request = &recommenderpb.MarkRecommendationFailedRequest{
		Name: formattedName,
		Etag: etag,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.MarkRecommendationFailed(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockRecommender.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestRecommenderMarkRecommendationFailedError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockRecommender.err = gstatus.Error(errCode, "test error")

	var formattedName string = fmt.Sprintf("projects/%s/locations/%s/recommenders/%s/recommendations/%s", "[PROJECT]", "[LOCATION]", "[RECOMMENDER]", "[RECOMMENDATION]")
	var etag string = "etag3123477"
	var request = &recommenderpb.MarkRecommendationFailedRequest{
		Name: formattedName,
		Etag: etag,
	}

	c, err := NewClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.MarkRecommendationFailed(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}