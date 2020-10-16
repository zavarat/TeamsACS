package common

import (
	"testing"
)

func TestToCamelCase(t *testing.T) {
	t.Log(ToCamelCase("user_name"))
}

func TestToSnakeCase(t *testing.T) {
	t.Log(ToSnakeCase("UserName"))
}

func TestToCamelCase1(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"user_name", args{"user_name"}, "UserName"},
		{"username", args{"username"}, "Username"},
		{"id", args{"id"}, "Id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.args.str); got != tt.want {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSnakeCase1(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ID", args{"ID"}, "id"},
		{"UserName", args{"UserName"}, "user_name"},
		{"UserNAME", args{"UserNAME"}, "user_name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSnakeCase(tt.args.str); got != tt.want {
				t.Errorf("ToSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestUUID(t *testing.T) {
	t.Log(UUID())
}

