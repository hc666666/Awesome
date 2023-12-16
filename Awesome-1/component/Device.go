package component

import (
	"Awesome/component/model"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"sync"
	"time"
)

// 网口对象
type NetFlow struct {
	device     *pcap.Interface
	snapLen    int32
	sampleTime time.Duration

	handler    *pcap.Handle
	ch_packets chan gopacket.Packet
	ch_len     int32
	Flows      map[gopacket.Flow]model.FlowMetaInfo
	Summon     uint64
}

func MonitorFactory(snapLen int32, sampleTime time.Duration) ([]*NetFlow, error) {
	devs, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal("FindAllDevs", err)
		return nil, err
	}
	res := []*NetFlow{}
	for i := range devs {
		log.Print(devs[i].Addresses)
		res = append(res, &NetFlow{
			device:     &devs[i],
			snapLen:    snapLen,
			sampleTime: sampleTime,
		})
	}
	return res, nil
}

// todo 一个网络层链接的通道缓存多少合适
func (n *NetFlow) newNetMonitor() {
	//监听网口
	handle, err := pcap.OpenLive(n.device.Name, n.snapLen, false, n.sampleTime)
	defer handle.Close()
	if err != nil {
		log.Fatal("fail in OpenLive")
		return
	}
	//DecodeFragment Fragment contains all
	n.ch_packets = make(chan gopacket.Packet, 65535)
	n.Flows = make(map[gopacket.Flow]model.FlowMetaInfo)
	packetSource := gopacket.NewPacketSource(handle, gopacket.DecodeFragment)
	//packetSource.DecodeOptions.NoCopy = true;
	for packet := range packetSource.Packets() {
		n.ch_packets <- packet
		n.Summon++
	}
}

var once sync.Once

func (n *NetFlow) GetFlow(flow gopacket.Flow) model.FlowMetaInfo {
	if n.Flows[flow] == nil {
		once.Do(func() {
			metaFlow := model.MetaFlow{
				Src:       flow.Src().String(),
				Dst:       flow.Dst().String(),
				In_total:  0,
				Out_total: 0,
				In_Udp:    0,
				Out_udp:   0,
				In_tcp:    0,
				Out_tcp:   0,
				In_ICMP:   0,
				Out_ICMP:  0,
				Status:    "",
			}
			n.Flows[flow] = &metaFlow
		})
	}
	return n.Flows[flow]
}
