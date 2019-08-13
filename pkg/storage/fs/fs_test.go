package fs

import (
	"reflect"
	"testing"

	"github.com/rugwirobaker/larissa/pkg/storage"
	"github.com/rugwirobaker/larissa/pkg/types"
	"github.com/spf13/afero"
)

func TestNewBackend(t *testing.T) {
	type args struct {
		rootDir    string
		filesystem afero.Fs
	}
	tests := []struct {
		name    string
		args    args
		want    storage.Backend
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBackend(tt.args.rootDir, tt.args.filesystem)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBackend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBackend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backend_bucketLoc(t *testing.T) {
	type args struct {
		bucket string
	}
	tests := []struct {
		name string
		fs   *backend
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fs.bucketLoc(tt.args.bucket); got != tt.want {
				t.Errorf("backend.bucketLoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backend_fileLoc(t *testing.T) {
	type args struct {
		bucket string
		file   string
	}
	tests := []struct {
		name string
		fs   *backend
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fs.fileLoc(tt.args.bucket, tt.args.file); got != tt.want {
				t.Errorf("backend.fileLoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backend_Put(t *testing.T) {
	type args struct {
		file    string
		bucket  string
		content []byte
	}
	tests := []struct {
		name    string
		fs      *backend
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fs.Put(tt.args.file, tt.args.bucket, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("backend.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backend_Get(t *testing.T) {
	type args struct {
		file   string
		bucket string
	}
	tests := []struct {
		name    string
		fs      *backend
		args    args
		want    *types.Object
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fs.Get(tt.args.file, tt.args.bucket)
			if (err != nil) != tt.wantErr {
				t.Errorf("backend.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("backend.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backend_Del(t *testing.T) {
	type args struct {
		file   string
		bucket string
	}
	tests := []struct {
		name    string
		fs      *backend
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fs.Del(tt.args.file, tt.args.bucket); (err != nil) != tt.wantErr {
				t.Errorf("backend.Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backend_Exists(t *testing.T) {
	type args struct {
		file   string
		bucket string
	}
	tests := []struct {
		name    string
		fs      *backend
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fs.Exists(tt.args.file, tt.args.bucket); (err != nil) != tt.wantErr {
				t.Errorf("backend.Exists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
