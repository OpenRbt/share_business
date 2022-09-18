package firebase_service

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"testing"
)

func TestFirebaseService_VerifyClaims(t *testing.T) {
	type fields struct {
		app  *firebase.App
		auth *auth.Client
	}
	type args struct {
		user           FirebaseProfile
		requiredClaims []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "No error case",
			fields: fields{
				app:  nil,
				auth: nil,
			},
			args: args{
				user: FirebaseProfile{
					CustomClaims: map[string]interface{}{"admin": "admin", "wash_owner": "wash_owner"},
				},
				requiredClaims: []string{"admin"},
			},
			wantErr: false,
		},
		{
			name: "Error case",
			fields: fields{
				app:  nil,
				auth: nil,
			},
			args: args{
				user: FirebaseProfile{
					CustomClaims: map[string]interface{}{"admin": "admin", "wash_owner": "wash_owner"},
				},
				requiredClaims: []string{"wash_engineer"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &FirebaseService{
				app:  tt.fields.app,
				auth: tt.fields.auth,
			}
			if err := svc.VerifyClaims(tt.args.user, tt.args.requiredClaims...); (err != nil) != tt.wantErr {
				t.Errorf("VerifyClaims() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
