package can

import (
	"github.com/DigitalBrainWaveTech/golang-api-simple-secure/auth"
	"testing"
)

func TestDo(t *testing.T) {
	type args struct {
		user       *auth.User
		permission string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid user with correct permission",
			args: args{
				user: &auth.User{
					ID:          "1",
					Email:       "test@test.com",
					Roles:       []string{"admin"},
					Permissions: []string{"read", "write"},
				},
				permission: "read",
			},
			want: true,
		},
		{
			name: "Valid user without correct permission",
			args: args{
				user: &auth.User{
					ID:          "2",
					Email:       "test2@test.com",
					Roles:       []string{"viewer"},
					Permissions: []string{"read"},
				},
				permission: "write",
			},
			want: false,
		},
		{
			name: "Nil user",
			args: args{
				user:       nil,
				permission: "read",
			},
			want: false,
		},
		{
			name: "Valid user with empty permission",
			args: args{
				user: &auth.User{
					ID:          "3",
					Email:       "test3@test.com",
					Roles:       []string{"editor"},
					Permissions: []string{"read", "write"},
				},
				permission: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Do(tt.args.user, tt.args.permission)
			if got != tt.want {
				t.Errorf("Do() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoAny(t *testing.T) {
	type args struct {
		user        *auth.User
		permissions []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid user with one matching permission",
			args: args{
				user: &auth.User{
					ID:          "1",
					Email:       "test@test.com",
					Roles:       []string{"admin"},
					Permissions: []string{"read", "write"},
				},
				permissions: []string{"write"},
			},
			want: true,
		},
		{
			name: "Valid user with multiple matching permissions",
			args: args{
				user: &auth.User{
					ID:          "2",
					Email:       "user2@example.com",
					Roles:       []string{"moderator"},
					Permissions: []string{"delete", "create", "edit"},
				},
				permissions: []string{"edit", "delete"},
			},
			want: true,
		},
		{
			name: "Valid user with no matching permissions",
			args: args{
				user: &auth.User{
					ID:          "3",
					Email:       "test3@test.com",
					Roles:       []string{"editor"},
					Permissions: []string{"read", "write"},
				},
				permissions: []string{"delete"},
			},
			want: false,
		},
		{
			name: "Nil user",
			args: args{
				user:        nil,
				permissions: []string{"read", "write"},
			},
			want: false,
		},
		{
			name: "User with empty permissions list",
			args: args{
				user: &auth.User{
					ID:          "4",
					Email:       "user4@example.com",
					Roles:       []string{"viewer"},
					Permissions: []string{},
				},
				permissions: []string{"read"},
			},
			want: false,
		},
		{
			name: "Empty permissions to check against",
			args: args{
				user: &auth.User{
					ID:          "5",
					Email:       "user5@example.com",
					Roles:       []string{"guest"},
					Permissions: []string{"read"},
				},
				permissions: []string{},
			},
			want: false,
		},
		{
			name: "User with overlapping permissions",
			args: args{
				user: &auth.User{
					ID:          "6",
					Email:       "user6@example.com",
					Roles:       []string{"superuser"},
					Permissions: []string{"manage", "admin", "read"},
				},
				permissions: []string{"read", "write"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DoAny(tt.args.user, tt.args.permissions...); got != tt.want {
				t.Errorf("DoAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoAnyFunc(t *testing.T) {
	type args struct {
		user        *auth.User
		permissions []string
		fn          func()
	}
	tests := []struct {
		name string
		args args
	}{

		{
			name: "Valid user with action executed",
			args: args{
				user: &auth.User{
					ID:          "1",
					Email:       "user1@example.com",
					Roles:       []string{"admin"},
					Permissions: []string{"read", "execute"},
				},
				permissions: []string{"execute"},
				fn: func() {
					t.Log("Action executed successfully")
				},
			},
		},
		{
			name: "Valid user without matching permissions",
			args: args{
				user: &auth.User{
					ID:          "2",
					Email:       "user2@example.com",
					Roles:       []string{"viewer"},
					Permissions: []string{"read"},
				},
				permissions: []string{"write"},
				fn: func() {
					t.Error("This should not be executed")
				},
			},
		},
		{
			name: "Nil user provided",
			args: args{
				user:        nil,
				permissions: []string{"read"},
				fn: func() {
					t.Error("This should not be executed as user is nil")
				},
			},
		},
		{
			name: "Empty permissions list with valid user",
			args: args{
				user: &auth.User{
					ID:          "3",
					Email:       "user3@example.com",
					Roles:       []string{"editor"},
					Permissions: []string{"read"},
				},
				permissions: []string{},
				fn: func() {
					t.Error("This should not be executed as permissions list is empty")
				},
			},
		},
		{
			name: "Valid user with multiple permissions and action executed",
			args: args{
				user: &auth.User{
					ID:          "4",
					Email:       "user4@example.com",
					Roles:       []string{"moderator"},
					Permissions: []string{"read", "write", "delete"},
				},
				permissions: []string{"write", "delete"},
				fn: func() {
					t.Log("Action executed successfully for matching permissions")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoAnyFunc(tt.args.user, tt.args.permissions, tt.args.fn)
		})
	}
}

func TestDoFunc(t *testing.T) {
	type args struct {
		user       *auth.User
		permission string
		fn         func()
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Valid user with matching permission and action executed",
			args: args{
				user: &auth.User{
					ID:          "5",
					Email:       "user5@example.com",
					Roles:       []string{"user"},
					Permissions: []string{"edit", "publish"},
				},
				permission: "publish",
				fn: func() {
					t.Log("Action executed for matching permission")
				},
			},
		},
		{
			name: "Valid user without matching permission",
			args: args{
				user: &auth.User{
					ID:          "6",
					Email:       "user6@example.net",
					Roles:       []string{"contributor"},
					Permissions: []string{"edit"},
				},
				permission: "delete",
				fn: func() {
					t.Error("This should not be executed")
				},
			},
		},
		{
			name: "Nil user provided",
			args: args{
				user:       nil,
				permission: "read",
				fn: func() {
					t.Error("This should not be executed as user is nil")
				},
			},
		},
		{
			name: "Valid user with empty user permissions",
			args: args{
				user: &auth.User{
					ID:          "7",
					Email:       "user7@example.org",
					Roles:       []string{"guest"},
					Permissions: []string{},
				},
				permission: "view",
				fn: func() {
					t.Error("This should not be executed for empty user permissions")
				},
			},
		},
		{
			name: "Valid user with overlapping permissions and action executed",
			args: args{
				user: &auth.User{
					ID:          "8",
					Email:       "user8@example.com",
					Roles:       []string{"manager"},
					Permissions: []string{"approve", "manage", "read"},
				},
				permission: "read",
				fn: func() {
					t.Log("Action executed for matching overlapping permission")
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DoFunc(tt.args.user, tt.args.permission, tt.args.fn)
		})
	}
}
