package export_ws

import (
	"github.com/lw000/gocommon/utils"
	"net/http"
	"sync"

	"github.com/olahol/melody"
	log "github.com/sirupsen/logrus"
)

type WsServer struct {
	m *melody.Melody
}

var (
	wsClients sync.Map
)

var (
	GateServe     *WsServer
	gateServeOnce sync.Once
)

func NewGatewayServer() *WsServer {
	GateServe = &WsServer{
		m: melody.New(),
	}
	GateServe.m.Config.MaxMessageSize = 1024
	GateServe.m.Config.MessageBufferSize = 512
	GateServe.HandleConnect(onConnectHandler)
	GateServe.HandleMessageBinary(onBinaryMessageHandler)
	GateServe.HandleDisconnect(onDisconnectHandler)
	GateServe.HandleError(onErrorHandler)
	return GateServe
}

func StartWsService() {
	gateServeOnce.Do(func() {
		GateServe = NewGatewayServer()
	})
}

func (g *WsServer) HandleRequest(w http.ResponseWriter, r *http.Request) error {
	return g.m.HandleRequestWithKeys(w, r, nil)
}

func (g *WsServer) HandleConnect(fn func(*melody.Session)) {
	g.m.HandleConnect(fn)
}

func (g *WsServer) HandleDisconnect(fn func(*melody.Session)) {
	g.m.HandleDisconnect(fn)
}

func (g *WsServer) HandleMessageBinary(fn func(*melody.Session, []byte)) {
	g.m.HandleMessageBinary(fn)
}

func (g *WsServer) HandleError(fn func(*melody.Session, error)) {
	g.m.HandleError(fn)
}

func onConnectHandler(s *melody.Session) {
	sessionId := tyutils.HashCode(tyutils.UUID())
	s.Set("sessionId", sessionId)
	wsClients.Store(sessionId, s)
	log.Infof("客户端连接, sessionId:%d", sessionId)
}

func onErrorHandler(s *melody.Session, e error) {
	sessionId := getSessionId(s)
	log.Infof("客户端错误, sessionId:%d, err:%s", sessionId, e.Error())
}

func onBinaryMessageHandler(s *melody.Session, msg []byte) {
	var err error
	value, ok := s.Get("sessionId")
	if !ok {
		err = s.CloseWithMsg([]byte("error"))
		if err != nil {
			log.Error(err)
		}
		return
	}
	sessionId := value.(uint32)
	log.Info(sessionId)

	log.Info(string(msg))

}

func onDisconnectHandler(s *melody.Session) {
	sessionId := getSessionId(s)
	wsClients.Delete(sessionId)
	log.Infof("客户端断开, sessionId:%d", sessionId)
}

func getSessionId(s *melody.Session) uint32 {
	val, ok := s.Get("sessionId")
	if !ok {
		log.Info("error")
		return 0
	}
	sessionId := val.(uint32)
	return sessionId
}
