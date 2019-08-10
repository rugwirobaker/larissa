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

func Test_backend_bucketLocation(t *testing.T) {
	type fields struct {
		rootDir    string
		filesystem afero.Fs
	}
	type args struct {
		bucket string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &backend{
				rootDir:    tt.fields.rootDir,
				filesystem: tt.fields.filesystem,
			}
			if got := fs.bucketLoc(tt.args.bucket); got != tt.want {
				t.Errorf("backend.bucketLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backend_Put(t *testing.T) {
	type fields struct {
		rootDir    string
		filesystem afero.Fs
	}
	type args struct {
		file    string
		bucket  string
		content []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &backend{
				rootDir:    tt.fields.rootDir,
				filesystem: tt.fields.filesystem,
			}
			if err := fs.Put(tt.args.file, tt.args.bucket, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("backend.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backend_Get(t *testing.T) {
	type fields struct {
		rootDir    string
		filesystem afero.Fs
	}
	type args struct {
		file   string
		bucket string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Object
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &backend{
				rootDir:    tt.fields.rootDir,
				filesystem: tt.fields.filesystem,
			}
			got, err := fs.Get(tt.args.file, tt.args.bucket)
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
	type fields struct {
		rootDir    string
		filesystem afero.Fs
	}
	type args struct {
		file   string
		bucket string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &backend{
				rootDir:    tt.fields.rootDir,
				filesystem: tt.fields.filesystem,
			}
			if err := fs.Del(tt.args.file, tt.args.bucket); (err != nil) != tt.wantErr {
				t.Errorf("backend.Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_backend_Exists(t *testing.T) {
	type fields struct {
		rootDir    string
		filesystem afero.Fs
	}
	type args struct {
		file   string
		bucket string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &backend{
				rootDir:    tt.fields.rootDir,
				filesystem: tt.fields.filesystem,
			}
			if got := fs.Exists(tt.args.file, tt.args.bucket); got != tt.want {
				t.Errorf("backend.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
