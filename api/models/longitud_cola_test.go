package models

import (
	"reflect"
	"testing"
	"time"
)

func TestLongitudCola_Validate(t *testing.T) {
	type fields struct {
		Facilidad string
		LongCola  int
		Time      int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Invalid Facility Name Test", fields{"", 10, time.Now().Unix()}, true},
		{"Invalid State Name Test", fields{"ECR Cdo", -10, time.Now().Unix()}, true},
		{"Invalid Time Value Test", fields{"ECR Cdo", 10, time.Now().Unix() + 650}, true},
		{"Invalid Time Value Test", fields{"ECR Cdo", 10, time.Now().Unix() - 650}, true}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := LongitudCola{
				Facilidad: tt.fields.Facilidad,
				LongCola:  &tt.fields.LongCola,
				Time:      tt.fields.Time,
			}
			if err := e.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLongitudCola_Write(t *testing.T) {
	type fields struct {
		Facilidad string
		LongCola  int
		Time      int64
	}
	tests := []struct {
		name    string
		fields  fields
		want    *LongitudCola
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := LongitudCola{
				Facilidad: tt.fields.Facilidad,
				LongCola:  &tt.fields.LongCola,
				Time:      tt.fields.Time,
			}
			got, err := e.Write()
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Write() got = %v, Want %v", got, tt.want)
			}
		})
	}
}
