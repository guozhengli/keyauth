// Code generated by protoc-gen-go-http. DO NOT EDIT.

package role

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				GrpcPath:         "/keyauth.role.RoleService/CreateRole",
				FunctionName:     "CreateRole",
				Path:             "",
				Method:           "",
				Resource:         "role",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "perm_admin", "action": "create"},
			},
			{
				GrpcPath:         "/keyauth.role.RoleService/QueryRole",
				FunctionName:     "QueryRole",
				Path:             "",
				Method:           "",
				Resource:         "role",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "perm_admin", "action": "list"},
			},
			{
				GrpcPath:         "/keyauth.role.RoleService/DescribeRole",
				FunctionName:     "DescribeRole",
				Path:             "",
				Method:           "",
				Resource:         "role",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "perm_admin", "action": "get"},
			},
			{
				GrpcPath:         "/keyauth.role.RoleService/DeleteRole",
				FunctionName:     "DeleteRole",
				Path:             "",
				Method:           "",
				Resource:         "role",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "perm_admin", "action": "delete"},
			},
			{
				GrpcPath:         "/keyauth.role.RoleService/QueryPermission",
				FunctionName:     "QueryPermission",
				Path:             "",
				Method:           "",
				Resource:         "permission",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "perm_admin", "action": "list"},
			},
			{
				GrpcPath:         "/keyauth.role.RoleService/DescribePermission",
				FunctionName:     "DescribePermission",
				Path:             "",
				Method:           "",
				Resource:         "permission",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "perm_admin", "action": "get"},
			},
			{
				GrpcPath:         "/keyauth.role.RoleService/AddPermissionToRole",
				FunctionName:     "AddPermissionToRole",
				Path:             "",
				Method:           "",
				Resource:         "permission",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "perm_admin", "action": "update"},
			},
			{
				GrpcPath:         "/keyauth.role.RoleService/RemovePermissionFromRole",
				FunctionName:     "RemovePermissionFromRole",
				Path:             "",
				Method:           "",
				Resource:         "permission",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "perm_admin", "action": "update"},
			},
			{
				GrpcPath:         "/keyauth.role.RoleService/UpdatePermission",
				FunctionName:     "UpdatePermission",
				Path:             "",
				Method:           "",
				Resource:         "permission",
				AuthEnable:       true,
				PermissionEnable: false,
				AuditLog:         true,
				Labels:           map[string]string{"allow": "perm_admin", "action": "update"},
			},
		},
	}
	return set
}
