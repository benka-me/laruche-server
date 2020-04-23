package db

import (
	"github.com/benka-me/laruche-server/go-pkg/config"
	"github.com/benka-me/laruche/go-pkg/laruche"
	"github.com/benka-me/laruche/go-pkg/laruche/faker"
	"testing"
)

func Test1InsertGetBee(t *testing.T) {
	db := Init(config.Init(true))
	tests := []struct {
		name    string
		bee     *laruche.Bee
		wantErr bool
	}{
		{
			name:    "benka-me/aaaa",
			bee:     faker.AlphaBeesMap["benka-me/aaaa"],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"/insert", func(t *testing.T) {
			got, err := db.InsertBee(tt.bee)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertBee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.bee.GetNamespaceStr() {
				t.Errorf("InsertBee() got = %v, want %v", got, tt.bee.GetNamespaceStr())
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name+"/get", func(t *testing.T) {
			got, err := db.GetBee(laruche.Namespace(tt.name))
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.GetNamespaceStr() != tt.name {
				t.Errorf("GetBee() got = %v, want %v", got.GetNamespaceStr(), tt.name)
			}
		})
	}
}
