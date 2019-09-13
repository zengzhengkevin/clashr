package adapters

import (
	"encoding/json"
	"errors"
	"github.com/zu1k/gossr/obfs"
	"github.com/zu1k/gossr/protocol"
	"net"
	"strconv"
	"strings"
	"time"

	C "github.com/Dreamacro/clash/constant"

	"github.com/zu1k/gossr"
	"github.com/zu1k/gossr/ssr"
)

type ShadowsocksR struct {
	*Base
	server string
	//ssrquery     *url.URL
	ssrop        ShadowsocksROption
	ObfsData     interface{}
	ProtocolData interface{}
}

type ShadowsocksROption struct {
	Name          string `proxy:"name"`
	Server        string `proxy:"server"`
	Port          int    `proxy:"port"`
	Password      string `proxy:"password"`
	Cipher        string `proxy:"cipher"`
	Protocol      string `proxy:"protocol"`
	ProtocolParam string `proxy:"protocolparam"`
	Obfs          string `proxy:"obfs"`
	ObfsParam     string `proxy:"obfsparam"`
}

func (ssrins *ShadowsocksR) Dial(metadata *C.Metadata) (C.Conn, error) {
	ssrop := ssrins.ssrop
	cipher, err := shadowsocksr.NewStreamCipher(ssrop.Cipher, ssrop.Password)
	if err != nil {
		return nil, err
	}

	dialer := net.Dialer{
		Timeout:   time.Millisecond * 500,
		DualStack: true,
	}
	conn, err := dialer.Dial("tcp", ssrins.server)
	if err != nil {
		return nil, err
	}

	dstcon := shadowsocksr.NewSSTCPConn(conn, cipher)
	if dstcon.Conn == nil || dstcon.RemoteAddr() == nil {
		return nil, errors.New("nil connection")
	}

	rs := strings.Split(dstcon.RemoteAddr().String(), ":")
	port, _ := strconv.Atoi(rs[1])

	dstcon.IObfs = obfs.NewObfs(ssrop.Obfs)
	obfsServerInfo := &ssr.ServerInfoForObfs{
		Host:   rs[0],
		Port:   uint16(port),
		TcpMss: 1460,
		Param:  ssrop.ObfsParam,
	}
	dstcon.IObfs.SetServerInfo(obfsServerInfo)
	dstcon.IProtocol = protocol.NewProtocol(ssrop.Protocol)
	protocolServerInfo := &ssr.ServerInfoForObfs{
		Host:   rs[0],
		Port:   uint16(port),
		TcpMss: 1460,
		Param:  ssrop.ProtocolParam,
	}
	dstcon.IProtocol.SetServerInfo(protocolServerInfo)

	if ssrins.ObfsData == nil {
		ssrins.ObfsData = dstcon.IObfs.GetData()
	}
	dstcon.IObfs.SetData(ssrins.ObfsData)

	if ssrins.ProtocolData == nil {
		ssrins.ProtocolData = dstcon.IProtocol.GetData()
	}
	dstcon.IProtocol.SetData(ssrins.ProtocolData)

	if _, err := dstcon.Write(serializesSocksAddr(metadata)); err != nil {
		dstcon.Close()
		return nil, err
	}
	return newConn(dstcon, ssrins), nil

}

func NewShadowsocksR(ssrop ShadowsocksROption) (*ShadowsocksR, error) {
	server := net.JoinHostPort(ssrop.Server, strconv.Itoa(ssrop.Port))
	return &ShadowsocksR{
		Base: &Base{
			name: ssrop.Name,
			tp:   C.ShadowsocksR,
			udp:  false,
		},
		server: server,
		//ssrquery: u,
		ssrop: ssrop,
	}, nil
}

func (ssr *ShadowsocksR) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"type": ssr.Type().String(),
	})
}

func (ssr *ShadowsocksR) DialUDP(metadata *C.Metadata) (pac C.PacketConn, netaddr net.Addr, err error) {
	return nil, nil, nil
}
