package system

import (
	"russ/system/util"
)

var HttpMethod = map[string]string{
	"GET":     "GET",
	"POST":    "POST",
	"PUT":     "PUT",
	"DELETE":  "DELETE",
	"PATCH":   "PATCH",
	"OPTIONS": "OPTIONS",
	"HEAD":    "HEAD",
	"TRACE":   "TRACE",
	"CONNECT": "CONNECT",
}

type Router struct {
	tree map[string]*Tree
}

func NewRouter() *Router {
	return &Router{
		tree: make(map[string]*Tree),
	}
}

func (r *Router) Add(pattern string, ctrlInfo *ControllerInfo) {
	segment := util.PathSplit(pattern)
	for _, hm := range HttpMethod {
		r.AddBase(hm, segment, ctrlInfo)
	}
}

func (r *Router) AddBase(hm string, segment []string, ctrlInfo *ControllerInfo) {
	_, ok := r.tree[hm]
	if !ok {
		r.tree[hm] = NewTree()
	}
	r.tree[hm].Add(segment, ctrlInfo)
}

func (r *Router) Get(hm, pattern string) *ControllerInfo {
	tree, ok := r.tree[hm]
	if !ok {
		return nil
	}
	segment := util.PathSplit(pattern)
	cntlInfo := tree.Get(segment)
	if cntlInfo == nil {
		return nil
	} else {
		return cntlInfo.(*ControllerInfo)
	}
}
