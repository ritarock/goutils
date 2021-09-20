package date

import (
	"testing"
	"time"
)

func TestToYYYYMMDD(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "to YYYYMMDD",
			args: args{t: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)},
			want: "20210101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToYYYYMMDD(tt.args.t); got != tt.want {
				t.Errorf("ToYYYYMMDD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToHHMM(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "to HHMM",
			args: args{t: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)},
			want: "0000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToHHMM(tt.args.t); got != tt.want {
				t.Errorf("ToHHMM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSpan(t *testing.T) {
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "timespan",
			args: args{
				from: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local),
				to:   time.Date(2021, 1, 1, 1, 0, 0, 0, time.Local),
			},
			want: time.Duration(3600) * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeSpan(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("TimeSpan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToYYYY_MM_DD(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "to YYYY-MM-DD",
			args: args{t: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)},
			want: "2021-01-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToYYYY_MM_DD(tt.args.t); got != tt.want {
				t.Errorf("ToYYYY_MM_DD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToYYYY_MM_DD_HH_MM_SS(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "to YYYY-MM-DD HH:MM:SS",
			args: args{t: time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)},
			want: "2021-01-01 00:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToYYYY_MM_DD_HH_MM_SS(tt.args.t); got != tt.want {
				t.Errorf("ToYYYY_MM_DD_HH_MM_SS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBeginningOfMonth(t *testing.T) {
	now := time.Date(2021, 9, 1, 0, 0, 0, 0, time.Local)
	nowTime = func() time.Time {
		return now
	}()
	tests := []struct {
		name string
		want string
	}{
		{
			name: "return 20210901",
			want: "20210901",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBeginningOfMonth(); got != tt.want {
				t.Errorf("GetBeginningOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEndOfMonth(t *testing.T) {
	now := time.Date(2021, 9, 1, 0, 0, 0, 0, time.Local)
	nowTime = func() time.Time {
		return now
	}()
	tests := []struct {
		name string
		want string
	}{
		{
			name: "return 20210901",
			want: "20210930",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEndOfMonth(); got != tt.want {
				t.Errorf("GetEndOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}
