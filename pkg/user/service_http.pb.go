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
				GrpcPath:          "/keyauth.user.UserService/QueryAccount",
				FunctionName:      "QueryAccount",
				Path:              "",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "list"},
				Extension:         map[string]string{},
			},
			{
				GrpcPath:          "/keyauth.user.UserService/DescribeAccount",
				FunctionName:      "DescribeAccount",
				Path:              "",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "*", "action": "get"},
				Extension:         map[string]string{},
			},
			{
				GrpcPath:          "/keyauth.user.UserService/CreateAccount",
				FunctionName:      "CreateAccount",
				Path:              "",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"action": "create", "allow": "org_admin"},
				Extension:         map[string]string{},
			},
			{
				GrpcPath:          "/keyauth.user.UserService/BlockAccount",
				FunctionName:      "BlockAccount",
				Path:              "",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "update"},
				Extension:         map[string]string{},
			},
			{
				GrpcPath:          "/keyauth.user.UserService/UnBlockAccount",
				FunctionName:      "UnBlockAccount",
				Path:              "",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "update"},
				Extension:         map[string]string{},
			},
			{
				GrpcPath:          "/keyauth.user.UserService/DeleteAccount",
				FunctionName:      "DeleteAccount",
				Path:              "",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"action": "delete", "allow": "org_admin"},
				Extension:         map[string]string{},
			},
			{
				GrpcPath:          "/keyauth.user.UserService/UpdateAccountProfile",
				FunctionName:      "UpdateAccountProfile",
				Path:              "",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "update"},
				Extension:         map[string]string{},
			},
			{
				GrpcPath:          "/keyauth.user.UserService/UpdateAccountPassword",
				FunctionName:      "UpdateAccountPassword",
				Path:              "",
				Method:            "",
				Resource:          "account",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          true,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "org_admin", "action": "update"},
				Extension:         map[string]string{},
			},
			{
				GrpcPath:          "/keyauth.user.UserService/GeneratePassword",
				FunctionName:      "GeneratePassword",
				Path:              "",
				Method:            "",
				Resource:          "",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          false,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "*"},
				Extension:         map[string]string{},
			},
		},
	}
	return set
}
