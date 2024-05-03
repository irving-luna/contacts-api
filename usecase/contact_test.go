package usecase

import (
	"accelone-api/mocks"
	"accelone-api/model"
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestContact_Get(t *testing.T) {
	type args struct {
		ctx       context.Context
		contactID string
	}
	var(
		expected = model.Contact{
			ID: "aaaa",
		}
	)
	tests := []struct {
		name    string
		args    args
		want    *model.Contact
		wantErr bool
	}{
		{
			name: "Positive scenario - existing contact",
			args: args{
				ctx:       context.Background(),
				contactID: "aaaa",
			},
			want: &expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			contactRepo := mocks.ContactRepository{}
			contactRepo.On("Get", mock.Anything, mock.Anything).Return(
				tt.want, nil,
			)

			u := &Contact{
				contactRepo: &contactRepo,
			}
			got, err := u.Get(tt.args.ctx, tt.args.contactID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Contact.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Contact.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
