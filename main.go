package main

import (
	"encoding/json"
	"fmt"
	"lobe-ext/pkg/calls"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/api/gateway", gatewayHandler(APIHandlers{
		"currentTime":  wrap(calls.HandleCurrentTime),
		"currentState": wrap(calls.HandleCurrentState),
	}))
	http.HandleFunc("/manifest-dev.json", func(w http.ResponseWriter, r *http.Request) {
		setCORSHeaders(w)
		http.ServeFile(w, r, "manifest-dev.json")
	})

	// 启动服务器
	port := ":3400"
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("API Gateway available at: http://localhost:3400/api/gateway")
	fmt.Println("Manifest available at: http://localhost:3400/manifest-dev.json")
	log.Fatal(http.ListenAndServe(port, nil))
}

// APIRequest 表示API网关请求的结构
type APIRequest struct {
	ID         string      `json:"id"`
	Type       string      `json:"type"`
	APIName    string      `json:"apiName"`
	Arguments  string      `json:"arguments"`
	Identifier string      `json:"identifier"`
	Manifest   interface{} `json:"manifest"`
}

// APIResponse 表示API响应的结构
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type APIHandler = func(APIRequest) (interface{}, error)
type APIHandlers = map[string]APIHandler

// 泛型包装函数
func wrap[T any, R any](handler func(T) (R, error)) APIHandler {
	return func(req APIRequest) (interface{}, error) {
		// 解析参数
		var args T
		if err := json.Unmarshal([]byte(req.Arguments), &args); err != nil {
			return nil, fmt.Errorf("failed to parse arguments: %v", err)
		}

		// 调用实际的处理函数
		return handler(args)
	}
}

// 设置CORS头
func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "x-lobe-chat-auth,x-lobe-plugin-settings,x-lobe-trace,Content-Type")
	w.Header().Set("Vary", "Access-Control-Request-Headers, Accept-Encoding")
}

// 处理API网关请求 - 使用类型别名简化参数
func gatewayHandler(funcs APIHandlers) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setCORSHeaders(w)
		// 处理CORS预检请求
		if r.Method == "OPTIONS" {
			w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 只处理POST请求
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 解析请求体
		var apiReq APIRequest
		if err := json.NewDecoder(r.Body).Decode(&apiReq); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		log.Printf("req: %+v", apiReq)
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Header().Set("Date", time.Now().UTC().Format(http.TimeFormat))

		// 根据apiName调用对应的函数
		handler, ok := funcs[apiReq.APIName]
		if !ok {
			err := fmt.Sprintf("Unknown API: %s", apiReq.APIName)
			response := APIResponse{
				Success: false,
				Error:   err,
			}
			log.Printf("error: %v", err)
			json.NewEncoder(w).Encode(response)
			return
		}

		// 调用处理函数
		res, err := handler(apiReq)
		if err != nil {
			response := APIResponse{
				Success: false,
				Error:   err.Error(),
			}
			log.Printf("error: %v", err)
			json.NewEncoder(w).Encode(response)
			return
		}
		log.Printf("res %+v", res)

		response := APIResponse{
			Success: true,
			Data:    res,
		}
		json.NewEncoder(w).Encode(response)
	}
}
