package handler

import (
	"bytes"
	"comm/auth"
	"comm/errors"
	"comm/logger"
	"context"
	"encoding/json"
	"io/ioutil"
	"proto/search"
	"regexp"

	"github.com/google/uuid"
	openapi "github.com/opensearch-project/opensearch-go/opensearchapi"
	"google.golang.org/protobuf/types/known/structpb"
)

var (
	indexNameRegex      = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9]$`)
	shortIndexNameRegex = regexp.MustCompile(`[a-zA-Z0-9]`)
)

func (h *Handler) CreateIndex(ctx context.Context, request *search.CreateIndexRequest, response *search.CreateIndexResponse) error {
	method := "search.CreateIndex"

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do CreateIndex", acc.Name)
	}

	if !isValidIndexName(request.Index) {
		return errors.BadRequest(method, "Index name should contain only alphanumerics and hyphens")
	}
	req := openapi.IndicesCreateRequest{
		Index: request.Index,
		Body:  nil, // TODO populate with fields and their types
	}
	rsp, err := req.Do(ctx, h.Client)
	if err != nil {
		logger.Errorf(ctx, "Error creating index %s", err)
		return errors.InternalServerError(method, "Error creating index")
	}
	defer rsp.Body.Close()
	if rsp.IsError() {
		logger.Errorf(ctx, "Error creating index %s", rsp.String())
		return errors.InternalServerError(method, "Error creating index")
	}
	return nil
}

func (h *Handler) Index(ctx context.Context, request *search.IndexRequest, response *search.IndexResponse) error {
	method := "search.Index"

	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Index", acc.Name)
	}

	if request.Data == nil {
		return errors.BadRequest(method, "Missing data")
	}
	if len(request.Id) == 0 {
		request.Id = uuid.New().String()
	}
	if len(request.Index) == 0 {
		return errors.BadRequest(method, "Missing index")
	}
	if !isValidIndexName(request.Index) {
		return errors.BadRequest(method, "Index name should contain only alphanumerics and hyphens")
	}

	b, err := request.Data.MarshalJSON()
	if err != nil {
		return errors.BadRequest(method, "Error processing document")
	}
	req := openapi.CreateRequest{
		Index:      request.Index,
		DocumentID: request.Id,
		Body:       bytes.NewBuffer(b),
	}
	rsp, err := req.Do(ctx, h.Client)
	if err != nil {
		logger.Errorf(ctx, "Error indexing doc %s", err)
		return errors.InternalServerError(method, "Error indexing document")
	}
	defer rsp.Body.Close()
	if rsp.IsError() {
		logger.Errorf(ctx, "Error indexing doc %s", rsp.String())
		return errors.InternalServerError(method, "Error indexing document")
	}
	response.Record = &search.Record{
		Id:   req.DocumentID,
		Data: request.Data,
	}

	return nil
}

func (h *Handler) Delete(ctx context.Context, request *search.DeleteRequest, response *search.DeleteResponse) error {
	method := "search.Delete"
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Delete", acc.Name)
	}

	if len(request.Index) == 0 {
		return errors.BadRequest(method, "Missing index param")
	}
	req := openapi.DeleteRequest{
		Index:      request.Index,
		DocumentID: request.Id,
	}
	rsp, err := req.Do(ctx, h.Client)
	if err != nil {
		logger.Errorf(ctx, "Error deleting doc %s", err)
		return errors.InternalServerError(method, "Error deleting document")
	}
	defer rsp.Body.Close()
	if rsp.IsError() {
		logger.Errorf(ctx, "Error deleting doc %s", rsp.String())
		return errors.InternalServerError(method, "Error deleting document")
	}
	return nil
}

func (h *Handler) Search(ctx context.Context, request *search.SearchRequest, response *search.SearchResponse) error {
	method := "search.Search"
	if len(request.Index) == 0 {
		return errors.BadRequest(method, "Missing index param")
	}

	// Search models to support https://opensearch.org/docs/latest/opensearch/ux/
	// - Simple query
	// - Autocomplete (prefix) queries
	// - pagination
	// - Sorting
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do Search", acc.Name)
	}

	// TODO fuzzy
	if len(request.Query) == 0 {
		return errors.BadRequest(method, "Missing query param")
	}

	qs, err := parseQueryString(request.Query)
	if err != nil {
		logger.Errorf(ctx, "Error parsing string %s %s", request.Query, err)
		return errors.BadRequest(method, "%s", err)
	}
	b, _ := qs.MarshalJSON()
	req := openapi.SearchRequest{
		Index: []string{request.Index},
		Body:  bytes.NewBuffer(b),
	}
	rsp, err := req.Do(ctx, h.Client)
	if err != nil {
		logger.Errorf(ctx, "Error searching index %s", err)
		return errors.InternalServerError(method, "Error searching documents")
	}
	defer rsp.Body.Close()
	if rsp.IsError() {
		if rsp.StatusCode == 404 { // index not found
			return errors.NotFound(method, "Index not found")
		}
		logger.Errorf(ctx, "Error searching index %s", rsp.String())
		return errors.InternalServerError(method, "Error searching documents")
	}
	b, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		logger.Errorf(ctx, "Error searching index %s", rsp.String())
		return errors.InternalServerError(method, "Error searching documents")
	}
	var os openSearchResponse
	if err := json.Unmarshal(b, &os); err != nil {
		logger.Errorf(ctx, "Error unmarshalling doc %s", err)
		return errors.InternalServerError(method, "Error searching documents")
	}
	logger.Infof(ctx, "%s", string(b))
	for _, v := range os.Hits.Hits {
		vs, err := structpb.NewStruct(v.Source)
		if err != nil {
			logger.Errorf(ctx, "Error unmarshalling doc %s", err)
			return errors.InternalServerError(method, "Error searching documents")
		}
		response.Records = append(response.Records, &search.Record{
			Id:   v.ID,
			Data: vs,
		})
	}
	return nil
}

func (h *Handler) DeleteIndex(ctx context.Context, request *search.DeleteIndexRequest, response *search.DeleteIndexResponse) error {
	method := "search.DeleteIndex"
	acc, ok := auth.FromContext(ctx)
	if ok {
		logger.Infof(ctx, "%v Do DeleteIndex", acc.Name)
	}

	if len(request.Index) == 0 {
		return errors.BadRequest(method, "Missing index param")
	}
	return h.deleteIndices(ctx, []string{request.Index}, method)
}

func (h *Handler) deleteIndices(ctx context.Context, indices []string, method string) error {
	req := openapi.IndicesDeleteRequest{
		Index: indices,
	}
	rsp, err := req.Do(ctx, h.Client)
	if err != nil {
		logger.Errorf(ctx, "Error deleting index %s", err)
		return errors.InternalServerError(method, "Error deleting index")
	}
	defer rsp.Body.Close()
	if rsp.IsError() {
		logger.Errorf(ctx, "Error deleting index %s", rsp.String())
		return errors.InternalServerError(method, "Error deleting index")
	}
	logger.Infof(ctx, "Deleted indices: %v", indices)
	return nil

}
