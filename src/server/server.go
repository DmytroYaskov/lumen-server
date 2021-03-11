package server

import (
	"log"
	"net/http"
	"strings"

	"lumen-server/client"
	"lumen-server/remotedevice"

	"github.com/gorilla/websocket"
)

type Server struct {
	upgrader websocket.Upgrader

	Devices map[int64]*remotedevice.Device
	Clients map[int64]*client.Client

	// tasks chan task
}

func NewServer() *Server {
	upg := websocket.Upgrader{}

	upg.CheckOrigin = func(r *http.Request) bool { return true }

	newServer := new(Server)

	newServer.upgrader = upg
	newServer.Devices = make(map[int64]*remotedevice.Device)
	newServer.Clients = make(map[int64]*client.Client)

	return newServer
}

// type task struct {
// 	obj     *interface{}
// 	message string
// }

func (s Server) ConnectioHandler(writer http.ResponseWriter, request *http.Request) {
	conn, err := s.upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	// Client identifing
	// To do:
	//  - identifyin by json with additional information about device
	mt, msg, err := conn.ReadMessage()
	if err != nil {
		log.Print("read:", err)
		return
	}

	if mt != websocket.TextMessage {
		return
	}

	log.Print("Received identifier: ", string(msg))

	//parse message
	connType, identifier := parseMessage(string(msg))
	// connType, _ := parseMessage(string(msg))

	if connType {

		err = conn.WriteMessage(websocket.TextMessage, []byte("recv"))
		if err != nil {
			log.Fatal("Responce msg: ", err)
		}

		log.Print("Connected new remote device")
		ID, newDevice := identifyDevice(identifier)
		newDevice.Connection = conn
		if newDevice != nil {
			s.addDevice(ID, newDevice)
			go newDevice.RemoteDeviceService()
		} else {
			log.Print("Can`t identify device!")
		}
	} else {
		client.ClientService(conn)
		log.Print("Connected new remote device")
	}
}

func parseMessage(message string) (connectionType bool, identyfier string) {
	data := strings.Split(message, ":")
	switch data[0] {
	case "rd":
		{
			connectionType = true
		}
	case "client":
		{
			connectionType = false
		}
	}
	identyfier = data[1]
	return
}

func (s Server) addDevice(ID int64, device *remotedevice.Device) {
	_, exists := s.Devices[ID]
	if !exists {
		s.Devices[ID] = device
	} else {
		log.Printf("Device with id %d is exists", ID)
	}
}

func identifyDevice(uniqueName string) (ID int64, newDevice *remotedevice.Device) {
	// Check device and get device information from database
	if uniqueName == "test" || uniqueName == "42" {
		ID = 1
		newDevice = remotedevice.CreateDefaultDevice()
	} else {
		newDevice = nil
	}
	return
}

func (s Server) remodeDevice(ID int64) {
	_, exists := s.Devices[ID]
	if exists {
		delete(s.Devices, ID)
	} else {
		log.Printf("Device with id %d isn't exists", ID)
	}
}

func (s Server) Superviser() {

}

func (s Server) addTask(Obj *interface{}, message string) {

}
