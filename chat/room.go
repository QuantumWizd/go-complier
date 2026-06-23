package chat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type User struct {
	ID   string
	Name string
	Conn net.Conn
}

type Room struct {
	mu      sync.RWMutex
	users   map[string]*User
	history *History
	hostIP  string
}

func NewRoom(hostIP string) *Room {
	return &Room{
		users:   make(map[string]*User),
		history: NewHistory(),
		hostIP:  hostIP,
	}
}

func (r *Room) AddUser(user *User) {
	r.mu.Lock()
	r.users[user.ID] = user
	r.mu.Unlock()

	// send WHO list to new user
	r.sendWho(user)

	// send history to new user
	r.sendHistory(user)

	// broadcast join to everyone else
	r.Broadcast(Message{
		Type: TypeSys,
		From: user.Name,
		Body: "joined",
	}, user.ID)

	PrintSystem(fmt.Sprintf("%s has joined the room", user.Name))
}

func (r *Room) RemoveUser(userID string) {
	r.mu.Lock()
	user, exists := r.users[userID]
	if !exists {
		r.mu.Unlock()
		return
	}
	delete(r.users, userID)
	r.mu.Unlock()

	r.Broadcast(Message{
		Type: TypeSys,
		From: user.Name,
		Body: "left",
	}, userID)

	PrintSystem(fmt.Sprintf("%s has left the room", user.Name))
}

func (r *Room) Broadcast(msg Message, excludeID string) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	line := Format(msg)
	for id, user := range r.users {
		if id == excludeID {
			continue
		}
		user.Conn.Write([]byte(line))
	}
}

func (r *Room) HandleUser(conn net.Conn, userID string) {
	reader := bufio.NewReader(conn)

	// first message must be SYS:name:joined
	line, err := reader.ReadString('\n')
	if err != nil {
		conn.Close()
		return
	}

	msg, err := Parse(line)
	if err != nil || msg.Type != TypeSys || msg.Body != "joined" {
		conn.Close()
		return
	}

	user := &User{
		ID:   userID,
		Name: msg.From,
		Conn: conn,
	}

	r.AddUser(user)

	defer func() {
		r.RemoveUser(userID)
		conn.Close()
	}()

	// main read loop
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		msg, err := Parse(line)
		if err != nil {
			continue
		}

		switch msg.Type {
		case TypeMsg:
			r.history.Add(msg)
			PrintMessage(msg)
			r.Broadcast(msg, userID)

		case TypeSys:
			if msg.Body == "left" {
				return
			}
		}
	}
}

func (r *Room) sendWho(target *User) {
	r.mu.RLock()
	names := make([]string, 0, len(r.users))
	for _, u := range r.users {
		names = append(names, u.Name)
	}
	r.mu.RUnlock()

	msg := Message{
		Type: TypeWho,
		From: "HOST",
		Body: strings.Join(names, ","),
	}
	target.Conn.Write([]byte(Format(msg)))
}

func (r *Room) sendHistory(target *User) {
	messages := r.history.GetAll()
	for _, m := range messages {
		hist := Message{
			Type: TypeHist,
			From: m.From,
			Body: m.Body,
		}
		target.Conn.Write([]byte(Format(hist)))
	}
}

func (r *Room) UserCount() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.users)
}