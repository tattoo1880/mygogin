package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"mygogin/config"
	"net/http"
	"sync"
)

var (
	upgrader       = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }} // 允许跨域
	ClientConnMap  = make(map[string]*websocket.Conn)
	ClientConnLock = sync.RWMutex{}
)

func WebSocketHandler(ctx *gin.Context) {
	userId := ctx.DefaultQuery("user_id", "")
	if userId == "" {
		ctx.JSON(400, gin.H{"error": "user_id 必须"})
		return
	}

	wsConn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {

		//Logger.Error("WebSocket 升级失败", zap.Error(err))
		config.Logger.Error("WebSocket 升级失败", zap.Error(err))
		return
	}

	// 保存连接
	ClientConnLock.Lock()
	ClientConnMap[userId] = wsConn
	ClientConnLock.Unlock()

	//Logger.Info("WebSocket 连接成功", zap.String("user_id", userId))
	config.Logger.Info("WebSocket 连接成功", zap.String("user_id", userId))
	// 启动 RabbitMQ 消费协程
	go ConsumeService(userId)

	// 监听客户端是否主动断开
	for {
		_, _, err := wsConn.ReadMessage()
		if err != nil {
			//Logger.Info("WebSocket 断开", zap.String("user_id", userId))
			config.Logger.Error("WebSocket 断开", zap.Error(err))
			break
		}
	}

	// 清理断开的连接
	ClientConnLock.Lock()
	delete(ClientConnMap, userId)
	ClientConnLock.Unlock()
	//Logger.Info("连接清理完成", zap.String("user_id", userId))
	config.Logger.Info("连接清理完成", zap.String("user_id", userId))
}
