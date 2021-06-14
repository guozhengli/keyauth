// Code generated by protoc-gen-go-http. DO NOT EDIT.

package permission

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				GrpcPath:          "/keyauth.permission.PermissionService/QueryPermission",
				FunctionName:      "QueryPermission",
				Path:              "/namespaces/:id/permissions",
				Method:            "GET",
				Resource:          "permission",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          false,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "perm_admin"},
				Extension:         map[string]string{},
			},
			{
				GrpcPath:          "/keyauth.permission.PermissionService/QueryRoles",
				FunctionName:      "QueryRoles",
				Path:              "/namespaces/:id/roles",
				Method:            "POST",
				Resource:          "role",
				AuthEnable:        true,
				PermissionEnable:  false,
				AuditLog:          false,
				RequiredNamespace: false,
				Labels:            map[string]string{"allow": "perm_admin"},
				Extension:         map[string]string{},
			},
			{
				GrpcPath:          "/keyauth.permission.PermissionService/CheckPermission",
				FunctionName:      "CheckPermission",
				Path:              "/namespaces/:id/permissions/endpoints/:eid",
				Method:            "GET",
				Resource:          "permission",
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
