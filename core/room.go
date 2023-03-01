package core

import (
	"encoding/json"
	"fmt"
	"github.com/Akilan1999/remotegameplay/config"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"os"
	"time"
)

type Room struct {
	ID             string `json:"id"`
	Sessions       map[string]*StreamSession
	CallerConn     *websocket.Conn
	BarrierSession *Barrier
}

type StreamSession struct {
	ID                  string
	Offer               string
	Answer              string
	CallerIceCandidates []string
	CalleeIceCandidates []string
	CallerConn          *websocket.Conn
	CalleeConn          *websocket.Conn
}

var roomMap = make(map[string]*Room)

func GetRoom(id string) *Room {
	return roomMap[id]
}

func NewRoom(callerConn *websocket.Conn) *Room {
	room := Room{
		ID:         newRoomID(),
		Sessions:   make(map[string]*StreamSession),
		CallerConn: callerConn,
	}
	roomMap[room.ID] = &room

	err := room.writeToFile()
	if err != nil {
		println(err)
	}

	return &room
}

func newRoomID() string {
	id := GetRandomName(0)
	for GetRoom(id) != nil {
		id = GetRandomName(0)
	}
	return id
}

func RemoveRoom(id string) {
	roomMap[id] = nil
	// Writes nil to the file
	roomMap[id].writeToFile()
}

func (room *Room) GetSession(id string) *StreamSession {
	return room.Sessions[id]
}

func (room *Room) NewSession(calleeConn *websocket.Conn) *StreamSession {
	session := StreamSession{
		ID:                  room.newSessionID(),
		CallerIceCandidates: []string{},
		CalleeIceCandidates: []string{},
		CallerConn:          room.CallerConn,
		CalleeConn:          calleeConn,
	}

	room.Sessions[session.ID] = &session

	return &session
}

func (room *Room) newSessionID() string {
	id := fmt.Sprintf("%s$%s", room.ID, GetRandomName(0))
	for GetRoom(id) != nil {
		id = fmt.Sprintf("%s$%s", room.ID, GetRandomName(0))
	}
	return id
}

// Write to rooms.json file
func (room *Room) writeToFile() error {
	file, err := json.MarshalIndent(room, "", " ")
	if err != nil {
		return err
	}

	// Get Path from config
	config, err := config.ConfigInit()
	if err != nil {
		return err
	}

	// Write room struct to json file
	err = ioutil.WriteFile(config.Rooms, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

// ReadRoomsFile Reads rooms file and return struct room id
func ReadRoomsFile() (*Room, error) {
	time.Sleep(3 * time.Second)
	config, err := config.ConfigInit()
	if err != nil {
		return nil, err
	}
	jsonFile, err := os.Open(config.Rooms)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var rooms *Room

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &rooms)

	return rooms, nil
}
