package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	log "github.com/cihub/seelog"
	"github.com/gorilla/websocket"

	"simulator/lib/util"
)

var (
	port         = flag.String("p", ":2345", "http service port")
	logVerbose   = flag.Bool("v", false, "output detailed log.")
	interval     = flag.Int("i", 1000, "advertising interval time(ms)")
	optInputPort = flag.String("r", "127.0.0.1:5563", "Specify IP and port of server main_realsense_serivce.py running.")
)

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
		err := util.EditSharedByteData(util.CylinderDataSharedObjectID,
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
	flag.Parse()

	util.InitSeriveGatewayCylinderData("tcp://" + *optInputPort)

	/// setup and start http server
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})

	log.Infof("Server is running on port: %s\n", *port)
	log.Error(http.ListenAndServe(*port, nil))
}
