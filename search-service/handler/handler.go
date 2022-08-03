package handler

import open "github.com/opensearch-project/opensearch-go"

type Handler struct {
	Client *open.Client
}

type Conf struct {
	OpenAddr string `json:"open_addr"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Insecure bool   `json:"insecure"`
}

type openSearchResponse struct {
	Took int64 `json:"took"`
	Hits hits  `json:"hits"`
}

type hits struct {
	Total map[string]interface{} `json:"total"`
	Hits  []hit                  `json:"hits"`
}
type hit struct {
	ID     string                 `json:"_id"`
	Score  float64                `json:"_score"`
	Source map[string]interface{} `json:"_source"`
}

type catIndicesEntry struct {
	Index string `json:"index"`
}

func isValidIndexName(s string) bool {
	if len(s) > 1 {
		return indexNameRegex.MatchString(s)
	}
	return shortIndexNameRegex.MatchString(s)
}
