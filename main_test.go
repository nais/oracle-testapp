package main

import "testing"

func TestParseOracleURL(t *testing.T) {
	tests := []struct {
		name        string
		oracleURL   string
		username    string
		password    string
		wantConnStr string
		wantErr     bool
	}{
		{
			name:        "simple password",
			oracleURL:   "jdbc:oracle:thin:@myhost:1521/mydb",
			username:    "user",
			password:    "pass",
			wantConnStr: "oracle://user:pass@myhost:1521/mydb",
		},
		{
			name:        "password with special chars",
			oracleURL:   "jdbc:oracle:thin:@myhost:1521/mydb",
			username:    "user",
			password:    "p@ss:word/123",
			wantConnStr: "oracle://user:p%40ss%3Aword%2F123@myhost:1521/mydb",
		},
		{
			name:        "scan hostname with dash",
			oracleURL:   "jdbc:oracle:thin:@dmv07-scan.adeo.no:1521/ngoraq0",
			username:    "nais",
			password:    "bDdvr",
			wantConnStr: "oracle://nais:bDdvr@dmv07-scan.adeo.no:1521/ngoraq0",
		},
		{
			name:      "missing prefix",
			oracleURL: "oracle:thin:@myhost:1521/mydb",
			username:  "user",
			password:  "pass",
			wantErr:   true,
		},
		{
			name:      "missing database",
			oracleURL: "jdbc:oracle:thin:@myhost:1521",
			username:  "user",
			password:  "pass",
			wantErr:   true,
		},
		{
			name:      "missing port",
			oracleURL: "jdbc:oracle:thin:@myhost/mydb",
			username:  "user",
			password:  "pass",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOracleURL(tt.oracleURL, tt.username, tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOracleURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.wantConnStr {
				t.Errorf("parseOracleURL() = %q, want %q", got, tt.wantConnStr)
			}
		})
	}
}
