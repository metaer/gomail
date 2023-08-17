package mail

import (
	"testing"

	"gomail/internal/config"
)

func TestSession_AuthPlain(t *testing.T) {
	cfg := &config.ServerConfig{
		Users: []struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		}{
			{
				Username: "test",
				Password: "test",
			},
		},
	}

	type fields struct {
		auth   bool
		from   string
		rcptTo string
		cfg    *config.ServerConfig
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "negative test",
			fields: fields{
				auth:   false,
				from:   "example@example.com",
				rcptTo: "example@example.com",
				cfg:    cfg,
			},
			args: args{
				username: "username",
				password: "password",
			},
			wantErr: true,
		},
		{
			name: "positive test",
			fields: fields{
				auth:   false,
				from:   "example@example.com",
				rcptTo: "example@example.com",
				cfg:    cfg,
			},
			args: args{
				username: "test",
				password: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				auth:   tt.fields.auth,
				from:   tt.fields.from,
				rcptTo: tt.fields.rcptTo,
				cfg:    tt.fields.cfg,
			}
			if err := s.AuthPlain(tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("AuthPlain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
