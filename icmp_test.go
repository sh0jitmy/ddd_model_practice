package tcpip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewICMP(t *testing.T) {
	tests := []struct {
		name string
		want ICMP
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewICMP(), "NewICMP()")
		})
	}
}

func TestICMP_Send(t *testing.T) {
	type fields struct {
		Type           []byte
		Code           []byte
		CheckSum       []byte
		Identification []byte
		SequenceNumber []byte
		Data           []byte
	}
	type args struct {
		ifindex int
		packet  []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ICMP
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			icmp := &ICMP{
				Type:           tt.fields.Type,
				Code:           tt.fields.Code,
				CheckSum:       tt.fields.CheckSum,
				Identification: tt.fields.Identification,
				SequenceNumber: tt.fields.SequenceNumber,
				Data:           tt.fields.Data,
			}
			assert.Equalf(t, tt.want, icmp.Send(tt.args.ifindex, tt.args.packet), "ICMP.Send(%v, %v)", tt.args.ifindex, tt.args.packet)
		})
	}
}
