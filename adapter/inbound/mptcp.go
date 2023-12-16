package inbound

import "net"

const multipathTCPAvailable = true

func setMultiPathTCP(listenConfig *net.ListenConfig, open bool) {
	listenConfig.SetMultipathTCP(open)
}
