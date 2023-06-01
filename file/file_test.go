package file

import (
	"testing"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    string
		wantErr bool
	}{
		{
			name:    "exists read file",
			args:    "testdata/readfile.txt",
			want:    "readfile\n",
			wantErr: false,
		},
		{
			name:    "not exists read file",
			args:    "testdata/readfile.tx",
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		got, err := Read(tc.args)
		if (err != nil) != tc.wantErr {
			t.Errorf("error: %v, wantErr: %v", err, tc.wantErr)
			return
		}
		if got != tc.want {
			t.Errorf("got: %v, want: %v", got, tc.want)
		}
	}
}

func TestReadLine(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    string
		wantErr bool
	}{
		{
			name:    "exists read file",
			args:    "testdata/readfile.txt",
			want:    "readfile",
			wantErr: false,
		},
		{
			name:    "not exists read file",
			args:    "testdata/readfile.tx",
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		got, err := ReadLine(tc.args)
		if (err != nil) != tc.wantErr {
			t.Errorf("error: %v, wantErr: %v", err, tc.wantErr)
			return
		}
		if got != tc.want {
			t.Errorf("got: %v, want: %v", got, tc.want)
		}
	}
}

func TestWrite(t *testing.T) {
	gotErr := Write("testdata/writefile.txt", "write file\n")
	wantErr := false

	if (gotErr != nil) != wantErr {
		t.Errorf("gotErr: %v, wantErr: %v", gotErr, wantErr)
	}
}
