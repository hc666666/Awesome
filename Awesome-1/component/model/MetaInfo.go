package model

import "fmt"

type FlowMetaInfo interface {
	Detail() string
	Refresh() bool
}
type MetaFlow struct {
	Src       string
	Dst       string
	In_total  uint64
	Out_total uint64
	In_Udp    int64
	Out_udp   int64
	In_tcp    int64
	Out_tcp   int64
	In_ICMP   int64
	Out_ICMP  int64
	Status    string
}

func (m *MetaFlow) Intotal() uint64 {
	return 0
}
func (m *MetaFlow) OutTotal() uint64 {
	return 0
}
func (m *MetaFlow) Detail() string {
	str := fmt.Sprintf("Packet:"+
		"[src:%v | dst:%v]\n"+
		"[In_Packet:%v | Out_Packet:%v]\n"+
		"[Status                  %v]",
		m.Src, m.Dst, m.In_total, m.Out_total, m.Status)
	return str
}
func (m *MetaFlow) Refresh() bool {
	return true
}
