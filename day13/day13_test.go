package day13

import "testing"

func TestPacketPrint(t *testing.T) {
	inexps := []*Packet{
		&Packet{nil, 4},
		&Packet{nil, 2},
	}
	items := []*Packet{
		&Packet{nil, 1},
		&Packet{inexps, -1},
		&Packet{nil, 3},
	}

	e := &Packet{items, -1}

	res := ""
	e.Print(&res)

	expected := "[1,[4,2],3]"
	if res != expected {
		t.Errorf("String %s != %s", res, expected)
	}
}

func TestConvertToPacket(t *testing.T) {
	packetStr := "[1,[4,2],3]"
	packet, _ := convertToPacket(packetStr)

	res := ""
	packet.Print(&res)

	expected := "[1,[4,2],3]"
	if res != packetStr {
		t.Errorf("String %s != %s", res, expected)
	}
}
