package times

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tt, err := time.Parse("2006-01-02T15:04:05-0700", "2018-09-14T10:05:00+0200")

	if err != nil {
		panic(err)
	}

	tests := []struct {
		name    string
		value   string
		want    time.Time
		wantErr bool
	}{
		{"t1", "2018-09-14T10:05:00+0200", tt, false},
		{"t2", "2018-09-14T10:05:00+02:00", tt, false},
		{"z", "2018-09-14T08:05:00Z", tt, false},
		{"fail", "2018-09", time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.value)

			if (err != nil) && !tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !got.Equal(tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustParse_Success(t *testing.T) {
	ts := "2018-09-14T10:05:00+0200"
	tt, err := time.Parse("2006-01-02T15:04:05-0700", ts)

	if err != nil {
		panic(err)
	}

	got := MustParse(ts)

	if !got.Equal(tt) {
		t.Errorf("MustParse() - %v, want %v", got, tt)
	}
}

func TestMustParse_Failure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("panic expected")
		}
	}()

	MustParse("invalid")
}
