// Code generated by protoc-gen-go-http. DO NOT EDIT.

package tag

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				Path:         "/keyauth.tag.TagService/CreateTag",
				FunctionName: "CreateTag",
			},
			{
				Path:         "/keyauth.tag.TagService/DescribeTag",
				FunctionName: "DescribeTag",
			},
			{
				Path:         "/keyauth.tag.TagService/DeleteTag",
				FunctionName: "DeleteTag",
			},
			{
				Path:         "/keyauth.tag.TagService/QueryTagKey",
				FunctionName: "QueryTagKey",
			},
			{
				Path:         "/keyauth.tag.TagService/QueryTagValue",
				FunctionName: "QueryTagValue",
			},
		},
	}
	return set
}