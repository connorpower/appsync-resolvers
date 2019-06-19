package resolvers

import "encoding/json"

type Context struct {
	Arguments json.RawMessage  `json:"arguments"`
	Source    json.RawMessage  `json:"source"`
	Identity  *json.RawMessage `json:"identity"`
}

type Invocation struct {
	Resolve string  `json:"resolve"`
	Context Context `json:"context"`
}

func (in Invocation) isRoot() bool {
	return in.Context.Source == nil || string(in.Context.Source) == "null"
}

func (in Invocation) payload() json.RawMessage {
	if in.isRoot() {
		return in.Context.Arguments
	}

	return in.Context.Source
}

func (in Invocation) identity() *json.RawMessage {
	return in.Context.Identity
}
