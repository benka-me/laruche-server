package db

import (
	"github.com/benka-me/laruche-server/go-pkg/config"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"testing"
)

func Test2SceneInsertGetHives(t *testing.T) {
	db := Init(config.Init(true))
	tests := []struct {
		name    string
		hive    *laruche.Hive
		want    string
		wantErr bool
	}{
		{
			name: "benka-me/example-a",
			hive: &laruche.Hive{
				Name:   "example-a",
				Repo:   "test/benka-me/example-a",
				Author: "benka-me",
				Public: false,
				Deps:   nil,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.InsertHive(tt.hive)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertHive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.hive.GetNamespaceStr() {
				t.Errorf("InsertHive() got = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := db.GetHive(laruche.Namespace(tt.name))
			if got.GetNamespaceStr() != tt.name {
				t.Errorf("GetHive() = %v, want %v", got.GetNamespaceStr(), tt.name)
			}
		})
	}
}
