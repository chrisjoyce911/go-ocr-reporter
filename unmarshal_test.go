package ocr

import (
	"testing"
	"time"
)

func TestSpecialDate_UnmarshalJSON(t *testing.T) {

	// Examples
	// "2020-05-14 18:00:06"
	// "2013-06-19 00:00:00"
	// null

	type fields struct {
		Time time.Time
	}
	type args struct {
		input []byte
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
			sd := &SpecialDate{
				Time: tt.fields.Time,
			}
			if err := sd.UnmarshalJSON(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("SpecialDate.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpecialBool_UnmarshalJSON(t *testing.T) {

	// Examples
	// "No"
	// "0"
	// 1
	// "Yes"

	type fields struct {
		bool bool
	}
	type args struct {
		input []byte
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
			sb := &SpecialBool{
				bool: tt.fields.bool,
			}
			if err := sb.UnmarshalJSON(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("SpecialBool.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSpecialUnix_UnmarshalJSON(t *testing.T) {

	// Example
	// 1589479206
	// 1589452085
	// 1589496250
	// 1589496250

	type fields struct {
		Time time.Time
	}
	type args struct {
		input []byte
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
			st := &SpecialUnix{
				Time: tt.fields.Time,
			}
			if err := st.UnmarshalJSON(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("SpecialUnix.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
