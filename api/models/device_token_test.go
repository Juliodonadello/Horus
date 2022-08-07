package models

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		{"Token exists in cache ", args{"3d862ca45c9f1038031adfe4b6b08c5b"}, true, "ECR Cdo"},
		{"Token does not exists in cache ", args{"11111111111111111111111111111111"}, false, ""},
		{"Malformed Token", args{"111"}, false, ""},
		{"Empty Token", args{""}, false, ""},
	}
	var testToken DeviceToken
	testToken.Create("ECR Cdo")
	TokenCache["3d862ca45c9f1038031adfe4b6b08c5b"] = &testToken
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CheckToken(tt.args.token)
			if got != tt.want {
				t.Errorf("CheckToken() got = %v, Want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CheckToken() got1 = %v, Want %v", got1, tt.want1)
			}
		})
	}
}

func TestDeviceToken_Create(t *testing.T) {
	tokenLength = 16
	type args struct {
		device string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Correct device name", args{"ECR Cdo"}, false},
		{"Empty device name", args{""}, true},
		{"Long device name", args{"weasdzxcqweasdzxcqweasdzxcqweasdzxcqweasdzxcqweasdzxcqweasdzxc"}, true},
		{"whitespace device name", args{"                 "}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var testToken DeviceToken
			err := testToken.Create(tt.args.device)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create(device) error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeviceToken_Delete(t *testing.T) {
	// Pre setup
	tokenLength = 16
	var testToken DeviceToken
	err := testToken.Create("Generador")
	if err != nil {
		panic(err)
	}
	type fields struct {
		Token *DeviceToken
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Correct delete token", fields{&testToken}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.fields.Token
			if err := e.Delete(); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeviceToken_Validate(t *testing.T) {
	tokenLength = 16
	var correctToken DeviceToken
	err := correctToken.Create("Red Mat Pers")
	if err != nil {
		panic(err)
	}
	type fields struct {
		Device    string
		Token     string
		CreatedAt time.Time
		ExpiresAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Correct token", fields{correctToken.Device, correctToken.Token, correctToken.CreatedAt, correctToken.ExpiresAt}, false},
		{"No device string", fields{"", correctToken.Token, correctToken.CreatedAt, correctToken.ExpiresAt}, true},
		{"Created in the future", fields{correctToken.Device, correctToken.Token, correctToken.CreatedAt.AddDate(50, 0, 0), correctToken.ExpiresAt}, true},
		{"invalid token", fields{correctToken.Device, "", correctToken.CreatedAt, correctToken.ExpiresAt}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := DeviceToken{
				Device:    tt.fields.Device,
				Token:     tt.fields.Token,
				CreatedAt: tt.fields.CreatedAt,
				ExpiresAt: tt.fields.ExpiresAt,
			}
			if err := e.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeviceToken_Write(t *testing.T) {
	tokenLength = 16
	var correctToken DeviceToken
	err := correctToken.Create("Red Mat Pers")
	if err != nil {
		panic(err)
	}
	type fields struct {
		Device    string
		Token     string
		CreatedAt time.Time
		ExpiresAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Write token", fields{correctToken.Device, correctToken.Token, correctToken.CreatedAt, correctToken.ExpiresAt}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := DeviceToken{
				Device:    tt.fields.Device,
				Token:     tt.fields.Token,
				CreatedAt: tt.fields.CreatedAt,
				ExpiresAt: tt.fields.ExpiresAt,
			}
			if err := e.Write(); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	err = correctToken.Delete()
	if err != nil {
		panic(err)
	}
}

func TestFindByDevice(t *testing.T) {
	clearCaches()
	tokenLength = 16
	var correctToken DeviceToken
	err := correctToken.Create("Red Mat Pers")
	if err != nil {
		panic(err)
	}
	err = correctToken.Write()
	if err != nil {
		panic(err)
	}
	type args struct {
		device string
	}
	tests := []struct {
		name string
		args args
		want *DeviceToken
	}{
		{"device exists", args{"Red Mat Pers"}, &correctToken},
		{"device does not exists", args{"Cachirulo"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindByDevice(tt.args.device)
			// Existe un bug con Go que hace que, si bien el tiempo se guarda correctamente en las struct, en la struct Got
			// Se llama al metodo string de time.Time, pero en tt.Want, a pesar de estar exportada y su contenido tambien,
			// no lo hace, imprimiendo el contenido raw del puntero. Por alguna causa esto tambien hace fallar al DeepEqual
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByDevice() got = %v, Want %v", got, tt.want)
			}
		})
	}
	err = correctToken.Delete()
	if err != nil {
		panic(err)
	}
}

func TestFindByToken(t *testing.T) {
	clearCaches()
	tokenLength = 16
	var correctToken DeviceToken
	err := correctToken.Create("Red Mat Pers")
	if err != nil {
		panic(err)
	}
	err = correctToken.Write()
	if err != nil {
		panic(err)
	}
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		Want *DeviceToken
	}{
		{"Token exists", args{correctToken.Token}, &correctToken},
		{"Token does not exists", args{"3d862ca45c9f1038031adfe4b6b08c5b"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindByToken(tt.args.token)
			// Existe un bug con Go que hace que, si bien el tiempo se guarda correctamente en las struct, en la struct Got
			// Se llama al metodo string de time.Time, pero en tt.Want, a pesar de estar exportada y su contenido tambien,
			// no lo hace, imprimiendo el contenido raw del puntero. Por alguna causa esto tambien hace fallar al DeepEqual
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("FindByToken() got = %v, Want %v", got, tt.Want)
			}
		})
	}
	err = correctToken.Delete()
	if err != nil {
		panic(err)
	}
}

func TestRecoverTokens(t *testing.T) {
	clearCaches()
	tokenLength = 16
	var token1 DeviceToken
	var token2 DeviceToken
	var token3 DeviceToken
	err := token1.Create("Red Mat Pers")
	if err != nil {
		panic(err)
	}
	err = token1.Write()
	if err != nil {
		panic(err)
	}
	err = token2.Create("Red Cdo Op")
	if err != nil {
		panic(err)
	}
	err = token2.Write()
	if err != nil {
		panic(err)
	}
	err = token3.Create("Red Alarm")
	if err != nil {
		panic(err)
	}
	err = token3.Write()
	if err != nil {
		panic(err)
	}
	clearCaches()
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{"Recover tokens from DB", 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RecoverTokens()
			if (err != nil) != tt.wantErr {
				t.Errorf("RecoverTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RecoverTokens() got = %v, Want %v", got, tt.want)
			}
		})
	}
	err = token1.Delete()
	if err != nil {
		panic(err)
	}
	err = token2.Delete()
	if err != nil {
		panic(err)
	}
	err = token3.Delete()
	if err != nil {
		panic(err)
	}
}

func Test_generateSecureToken(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Correct int lenght", args{16}, 32},
		{"Excesive int lenght", args{64}, 0},
		{"Negative int lenght", args{-16}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateSecureToken(tt.args.length); len(got) != tt.want {
				t.Errorf("generateSecureToken() = %v, Want %v", got, tt.want)
			}
		})
	}
}

func clearCaches() {
	for k := range DeviceCache {
		delete(DeviceCache, k)
	}
	for k := range TokenCache {
		delete(TokenCache, k)
	}
}
