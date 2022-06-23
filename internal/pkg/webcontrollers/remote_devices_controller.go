package webcontrollers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/rokmonster/ocr/internal/pkg/websocket/remote"
	log "github.com/sirupsen/logrus"
)

func (controller *RemoteDevicesController) getRemoteDevices() map[uuid.UUID]ServerClient {
	return controller.clients
}

// ServerClient - holds basic information about rok-remote instance connected to websocket
type ServerClient struct {
	Name    string
	Address string
	Handler *remote.RemoteServerWS
}

type RemoteDevicesController struct {
	clients      map[uuid.UUID]ServerClient
	upgrader     websocket.Upgrader
	templatesDir string
	tessdataDir  string
}

func NewRemoteDevicesController(templates, tessdata string) *RemoteDevicesController {
	return &RemoteDevicesController{
		clients:      make(map[uuid.UUID]ServerClient),
		templatesDir: templates,
		tessdataDir:  tessdata,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// who care's about CORS?
				// P.s. this is bad idea...
				return true
			},
		},
	}
}

func (controller *RemoteDevicesController) Setup(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		data := gin.H{
			"devices": controller.getRemoteDevices(),
		}

		switch c.NegotiateFormat(gin.MIMEJSON, gin.MIMEHTML) {
		case gin.MIMEHTML:
			c.HTML(http.StatusOK, "devices.html", data)
		case gin.MIMEJSON:
			c.JSON(http.StatusOK, data)
		}
	})

	router.GET("/:id/disconnect", func(ctx *gin.Context) {
		id := uuid.MustParse(ctx.Param("id"))
		if c, ok := controller.clients[id]; ok {
			c.Handler.Disconnect()
		}

		ctx.Redirect(http.StatusFound, "/devices/")
	})

	router.GET("/:id/data", func(ctx *gin.Context) {
		id := uuid.MustParse(ctx.Param("id"))
		if c, ok := controller.clients[id]; ok {
			ctx.JSON(http.StatusOK, c.Handler.GetData())
			return
		}

		ctx.Redirect(http.StatusFound, "/devices/")
	})

	router.GET("/ws", func(c *gin.Context) {
		ws, _ := controller.upgrader.Upgrade(c.Writer, c.Request, nil)

		// don't forget to close the connection & remove client
		defer ws.Close()

		// first message on WS should be our hello
		var deviceInfo struct {
			Serial string `json:"serial"`
		}
		err := ws.ReadJSON(&deviceInfo)

		if err != nil {
			log.Errorf("I don't like this WS Client: %v", err)
			_ = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.ClosePolicyViolation, "I expect you to behave nicely"))
			return
		} else {
			log.Infof("[%v] connected from: %v", deviceInfo.Serial, ws.RemoteAddr())

			sessionId := uuid.New()

			// register handler && start the loop
			handler := remote.NewRemoteServerWS(ws, sessionId.String(), controller.templatesDir, controller.tessdataDir)
			device := ServerClient{
				Address: ws.RemoteAddr().String(),
				Name:    deviceInfo.Serial,
				Handler: handler,
			}

			// put the client into active clients...
			controller.clients[sessionId] = device
			defer delete(controller.clients, sessionId)

			// handle the command / send actions
			handler.Loop()
		}
	})
}
