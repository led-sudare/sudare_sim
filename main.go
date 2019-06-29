package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github.com/cihub/seelog"
	"github.com/gorilla/websocket"

	"simulator/lib"
	"simulator/lib/util"
)

type Configs struct {
	Port       int    `json:"port"`
	LogVorbose bool   `json:"logVorbose"`
	ZmqTarget  string `json:"zmqTarget"`
}

func NewConfigs() Configs {
	return Configs{
		Port:       2345,
		LogVorbose: false,
		ZmqTarget:  "0.0.0.0:5563",
	}
}

var upgrader = websocket.Upgrader{}

func setupLoggingFramwork(isLogVorbose bool) {

	var logLevel string
	if isLogVorbose {
		logLevel = "trace"
	} else {
		logLevel = "info"
	}

	logConfig := fmt.Sprintf("<seelog type=\"sync\" minlevel=\"%s\">"+
		"<outputs><console/></outputs>"+
		"</seelog>", logLevel)

	fmt.Printf("set seelog config: %s", logConfig)

	logger, _ := log.LoggerFromConfigAsBytes([]byte(logConfig))
	log.ReplaceLogger(logger)

}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("Upgrade error: ", err)
		return
	}
	defer ws.Close()

	clock := time.NewTicker(time.Duration(100 * time.Millisecond))
	defer clock.Stop()

	for {
		<-clock.C
		err := util.EditSharedByteData(lib.CylinderDataSharedObjectID,
			func(editable util.ByteData) error {
				log.Info("Send WebSocket Data: ", len(editable.GetBytes()))
				ws.SetWriteDeadline(time.Now().Add(200 * time.Millisecond))
				if err := ws.WriteMessage(websocket.BinaryMessage, editable.GetBytes()); err != nil {
					log.Debug("ws.WriteMessage", err)
					http.Error(w, "Internal Error"+err.Error(), http.StatusInternalServerError)
					return err
				}
				return nil
			})
		if err != nil {
			log.Error(err)
			break
		}
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path == "/" {
		http.ServeFile(w, r, "./www/index.html")
	} else {
		http.ServeFile(w, r, "./www/"+r.URL.Path)
	}
}

func main() {

	configs := NewConfigs()
	util.ReadConfig(&configs)

	var (
		port         = flag.Int("p", configs.Port, "http service port")
		optInputPort = flag.String("r", configs.ZmqTarget, "Specify IP and port of server zeromq PUB running.")
	)

	flag.Parse()

	lib.InitSeriveGatewayCylinderData("tcp://" + *optInputPort)

	/// setup and start http server
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})

	log.Infof("Server is running on port: %v\n", *port)
	log.Error(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
