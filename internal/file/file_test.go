package file

import (
	"testing"
)

func TestRead(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "read file",
			args: args{"../../test_data/readfile.txt"},
			want: "read file\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Read(tt.args.path); got != tt.want {
				t.Errorf("Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrite(t *testing.T) {
	path := "../../test_data/write.txt"
	type args struct {
		path string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "write",
			args:    args{path: path, data: "write.txt"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.args.path, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	RemoveFile(path)
}

func TestReadLine(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reade line",
			args: args{path: "../../test_data/readfile.txt"},
			want: "read file",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadLine(tt.args.path); got != tt.want {
				t.Errorf("ReadLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
