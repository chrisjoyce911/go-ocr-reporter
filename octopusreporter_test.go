package ocr

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOCR_LoadCustomers(t *testing.T) {

	tests := []struct {
		name    string
		file    string
		want    []Customer
		wantErr bool
	}{
		{
			name:    "ok",
			file:    "LoadCustomers_ok.json",
			wantErr: false,
			want: []Customer{
				{
					ID:                     2,
					WsID:                   "f8e508b7bf16",
					InternalID:             "52",
					Label:                  "Customer_X",
					OwnUse:                 0,
					AfterImport:            "user@ou=%/Customer_X/%\r\nserver@ou=%/Customer_X/%",
					CountUsers:             SpecialInt{5},
					CountServers:           SpecialInt{1},
					CountDirectScanServers: SpecialInt{1},
					Updated:                SpecialUnix{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OCR{}
			file := fmt.Sprintf("test_files/%s", tt.file)
			data, _ := ioutil.ReadFile(file)
			got, err := o.LoadCustomers(data)

			if assert.NoError(t, err) {
				assert.Equal(t, got, tt.want)
			}

		})
	}
}
