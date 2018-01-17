package gotalk

import (
  "golang.org/x/net/websocket"
  "github.com/rsms/gotalk/js"
  "fmt"
  "net/http"
  "strings"
  "io"
  "strconv"
  "time"
)

type WebSocketServer struct {
  Limits
  Handlers *Handlers
  OnAccept SockHandler

  // Template value for accepted sockets. Defaults to 0 (no automatic heartbeats)
  HeartbeatInterval time.Duration

  // Template value for accepted sockets. Defaults to nil
  OnHeartbeat func(load int, t time.Time)

  Server   *websocket.Server
}

const gotalkJSSuffix = "/gotalk.js"
const gotalkJSMapSuffix = "/gotalk.js.map"

func (s *WebSocketServer) serveResource(w http.ResponseWriter, r *http.Request, f func()) {
  // serve javascript library
  w.Header()["Cache-Control"] = []string{"public, max-age=300"}
  etag := "\"" + gotalkjs.BrowserLibSHA1Base64 + r.URL.Path + "\""
  w.Header()["ETag"] = []string{etag}
  reqETag := r.Header["If-None-Match"]

  if len(reqETag) != 0 && reqETag[0] == etag {
    w.WriteHeader(http.StatusNotModified)
  } else {
    f()
  }
}

func (s *WebSocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if strings.HasSuffix(r.URL.Path, gotalkJSSuffix) {
    s.serveResource(w, r, func () {
      w.Header()["Content-Type"] = []string{"text/javascript"}
      wsPath := r.URL.Path[:len(r.URL.Path)-len(gotalkJSSuffix)+1]
      serveURL := "this.gotalkResponderAt={ws:'"+wsPath+"'};"
      sizeStr := strconv.FormatInt(int64(len(serveURL) + len(gotalkjs.BrowserLibString)), 10)
      w.Header()["Content-Length"] = []string{sizeStr}
      w.WriteHeader(http.StatusOK)
      // Note: w conforms to interface { WriteString(string)(int,error) }
      io.WriteString(w, serveURL)
      io.WriteString(w, gotalkjs.BrowserLibString)
    })
  } else if strings.HasSuffix(r.URL.Path, gotalkJSMapSuffix) {
    s.serveResource(w, r, func () {
      w.Header()["Content-Type"] = []string{"application/json"}
      w.Header()["Content-Length"] = []string{strconv.FormatInt(
        int64(len(gotalkjs.BrowserLibSourceMapString)),
        10,
      )}
      w.WriteHeader(http.StatusOK)
      io.WriteString(w, gotalkjs.BrowserLibSourceMapString)
    })
  } else {
    // upgrade request connection to web socket protocol
    s.Server.ServeHTTP(w, r)
  }
}

// Handler that can be used with the http package
func WebSocketHandler() *WebSocketServer {
  server := &WebSocketServer{
    Limits:   DefaultLimits,
    Handlers: DefaultHandlers,
  }

  handler := func(ws *websocket.Conn) {
    s := NewSock(server.Handlers)
    ws.PayloadType = websocket.BinaryFrame; // websocket.TextFrame;
    s.Adopt(ws)
    if err := s.Handshake(); err != nil {
      s.Close()
    } else {
      if server.OnAccept != nil {
        server.OnAccept(s)
      }
      s.HeartbeatInterval = server.HeartbeatInterval
      s.OnHeartbeat = server.OnHeartbeat
      s.Read(server.Limits)
    }
  }

  server.Server = &websocket.Server{Handler: handler, Handshake: checkOrigin}

  return server
}

func checkOrigin(config *websocket.Config, req *http.Request) (err error) {
  config.Origin, err = websocket.Origin(config, req)
  if err == nil && config.Origin == nil {
    return fmt.Errorf("null origin")
  }
  return err
}
