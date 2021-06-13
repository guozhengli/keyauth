// Code generated by protoc-gen-go-http. DO NOT EDIT.

package user

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				GrpcPath:         "/keyauth.user.UserService/QueryAccount",
				FunctionName:     "QueryAccount",
				Path:             "",
				Method:           "",
				Resource:         "account",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "org_admin", "action": "list"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/DescribeAccount",
				FunctionName:     "DescribeAccount",
				Path:             "",
				Method:           "",
				Resource:         "account",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "*", "action": "get"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/CreateAccount",
				FunctionName:     "CreateAccount",
				Path:             "",
				Method:           "",
				Resource:         "account",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "org_admin", "action": "create"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/BlockAccount",
				FunctionName:     "BlockAccount",
				Path:             "",
				Method:           "",
				Resource:         "account",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "org_admin", "action": "update"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/UnBlockAccount",
				FunctionName:     "UnBlockAccount",
				Path:             "",
				Method:           "",
				Resource:         "account",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "org_admin", "action": "update"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/DeleteAccount",
				FunctionName:     "DeleteAccount",
				Path:             "",
				Method:           "",
				Resource:         "account",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "org_admin", "action": "delete"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/UpdateAccountProfile",
				FunctionName:     "UpdateAccountProfile",
				Path:             "",
				Method:           "",
				Resource:         "account",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"action": "update", "allow": "org_admin"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/UpdateAccountPassword",
				FunctionName:     "UpdateAccountPassword",
				Path:             "",
				Method:           "",
				Resource:         "account",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "org_admin", "action": "update"},
			},
			{
				GrpcPath:         "/keyauth.user.UserService/GeneratePassword",
				FunctionName:     "GeneratePassword",
				Path:             "",
				Method:           "",
				Resource:         "",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         false,
				Labels:           map[string]string{"allow": "*"},
			},
		},
	}
	return set
}
