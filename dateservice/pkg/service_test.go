package pkg

import (
	"context"
	"testing"
	"time"
)

func Test_dateService_Get(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"today",
			args{context.Background()},
			time.Now().Format("02.01.2006"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := dateService{}
			got, err := da.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dateService_Status(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"ok",
			args{context.Background()},
			"ok",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := dateService{}
			got, err := da.Status(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Status() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Status() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dateService_Validate(t *testing.T) {
	type args struct {
		ctx  context.Context
		date string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			"valid date",
			args{
				context.Background(),
				"31.12.2019",
			},
			true,
			false,
		},
		{
			"invalid date",
			args{
				context.Background(),
				"31.31.2019",
			},
			false,
			true,
		},
		{
			"invalid, USA formatted date",
			args{
				context.Background(),
				"12.31.2019",
			},
			false,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := dateService{}
			got, err := da.Validate(tt.args.ctx, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Validate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
