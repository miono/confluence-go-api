package goconfluence

import (
	"net/url"
	"strconv"
)

// getContentEndpoint creates the correct api endpoint by given id
func (a *API) getSearchEndpoint() (*url.URL, error) {
	return url.ParseRequestURI(a.endPoint.String() + "/search")
}

// Search querys confluence using CQL
func (a *API) Search(query SearchQuery) (*Search, error) {
	ep, err := a.getSearchEndpoint()
	if err != nil {
		return nil, err
	}
	ep.RawQuery = addSearchQueryParams(query).Encode()
	return a.SendSearchRequest(ep, "GET")
}

// addSearchQueryParams adds the defined query parameters
func addSearchQueryParams(query SearchQuery) *url.Values {

	data := url.Values{}
	if query.CQL != "" {
		data.Set("cql", query.CQL)
	}
	if query.CQLContext != "" {
		data.Set("cqlcontext", query.CQLContext)
	}
	if query.IncludeArchivedSpaces == true {
		data.Set("includeArchivedSpaces", "true")
	}
	if query.Limit != 0 {
		data.Set("limit", strconv.Itoa(query.Limit))
	}
	if query.Start != 0 {
		data.Set("start", strconv.Itoa(query.Start))
	}
	return &data
}
