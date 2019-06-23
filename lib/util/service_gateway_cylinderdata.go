package util

import (
	log "github.com/cihub/seelog"
	zmq "github.com/zeromq/goczmq"
)

type serviceGatewayCylinderData struct {
	sock  *zmq.Sock
	order chan string
	done  chan struct{}
}

const CylinderDataSharedObjectID = "cylinder"

var gServiceGatewayCylinderData *serviceGatewayCylinderData

func NewLedSudareData() ByteData {
	return NewByteData(CylinderWidth * CylinderHeight * CylinderCount * CylinderColorPlane)
}

func InitSeriveGatewayCylinderData(endpoint string) {

	AddSharedByteData(CylinderDataSharedObjectID, NewLedSudareData())

	gServiceGatewayCylinderData = &serviceGatewayCylinderData{}
	var err error
	gServiceGatewayCylinderData.sock, err = zmq.NewSub(endpoint, "")
	if err != nil {
		panic(err)
	}
	gServiceGatewayCylinderData.sock.Connect(endpoint)
	gServiceGatewayCylinderData.order = make(chan string)
	gServiceGatewayCylinderData.done = make(chan struct{})

	go serviceGatewayCylinderDataWorker(
		gServiceGatewayCylinderData.sock,
		gServiceGatewayCylinderData.order,
		gServiceGatewayCylinderData.done)
}

func rgb565to888(c565 uint32) []byte {
	r := (byte)((c565 & 0xF800) >> 8)
	g := (byte)((c565 & 0x7E0) >> 3)
	b := (byte)((c565 & 0x1F) << 3)
	return []byte{r, g, b}
}

func serviceGatewayCylinderDataWorker(sock *zmq.Sock, c chan string, done chan struct{}) {

	log.Info("serviceGatewayCylinderDataWorker start..")
	defer sock.Destroy()
	defer log.Info("serviceGatewayCylinderDataWorker finish")
	for {
		select {
		case <-c:
			done <- struct{}{}
			return
		default:
			data, _, _ := sock.RecvFrame()
			log.Info("Receive Image from ZeroMQ PUB: ", len(data))

			plane := 0
			if len(data) == 720000 {
				plane = 4
			} else if len(data) == 540000 {
				plane = 3
			} else if len(data) == 360000 {
				plane = 2
			} else {
				log.Warn("Receive Image Size was invalid. ", len(data), data)
				continue
			}

			EditSharedByteData(CylinderDataSharedObjectID,
				func(editable ByteData) error {

					ConcurrentEnum(0, CylinderCount, func(r int) {
						for y := 0; y < CylinderHeight; y++ {
							for x := 0; x < CylinderWidth; x++ {
								idxFrom := (r * CylinderHeight * CylinderWidth * plane) +
									(y * CylinderWidth * plane) + (plane * x)
								idxTo := (r * CylinderHeight * CylinderWidth * CylinderColorPlane) +
									(y * CylinderWidth * CylinderColorPlane) + (CylinderColorPlane * x)

								if plane == 2 {

									tmp := rgb565to888((uint32(data[idxFrom+0]) << 8) + uint32(data[idxFrom+1]))

									for i := 0; i < 3; i++ {
										editable.SetAt(idxTo+i, tmp[i])
									}

								} else {

									for i := 0; i < 3; i++ {
										editable.SetAt(idxTo+i, data[idxFrom+i])
									}

								}
							}
						}
					})

					return nil
				})
		}
	}
}
