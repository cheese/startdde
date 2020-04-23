package utils

import (
	x "github.com/linuxdeepin/go-x11-client"
	"github.com/linuxdeepin/go-x11-client/ext/randr"
	libutils "pkg.deepin.io/lib/utils"
)

func GetOutputEDID(conn *x.Conn, output randr.Output) ([]byte, error) {
	atomEDID, err := conn.GetAtom("EDID")
	if err != nil {
		return nil, err
	}

	reply, err := randr.GetOutputProperty(conn, output,
		atomEDID, x.AtomInteger,
		0, 32, false, false).Reply(conn)
	if err != nil {
		return nil, err
	}
	return reply.Value, nil
}

func GetEDIDChecksum(edid []byte) string {
	if len(edid) < 128 {
		return ""
	}

	// 最后俩字节两边获取到的可能会不同
	id, _ := libutils.SumStrMd5(string(edid[:126]))
	return id
}
