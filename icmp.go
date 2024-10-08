package tcpip

import (
)

type ICMP struct {
	Type           []byte
	Code           []byte
	CheckSum       []byte
	Identification []byte
	SequenceNumber []byte
	Data           []byte
}

func NewICMP() ICMP {
	// https://www.infraexpert.com/study/tcpip4.html
	icmp := ICMP{
		// ping request
		Type:           []byte{0x08},
		Code:           []byte{0x00},
		CheckSum:       []byte{0x00, 0x00},
		Identification: []byte{0x00, 0x10},
		SequenceNumber: []byte{0x00, 0x01},
		Data:           []byte{0x01, 0x02},
	}

	icmpsum := sumByteArr(toByteArr(icmp))
	icmp.CheckSum = checksum(icmpsum)

	return icmp
}

func (icmp *ICMP) Send(ifindex int, packet []byte) {


}

func parseICMP(packet []byte) ICMP {
	return ICMP{
		Type:           []byte{packet[0]},
		Code:           []byte{packet[1]},
		CheckSum:       []byte{packet[2], packet[3]},
		Identification: []byte{packet[4], packet[5]},
		SequenceNumber: []byte{packet[6], packet[7]},
		Data:           packet[8:],
	}
}
