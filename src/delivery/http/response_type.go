package http

type response map[string]interface{}

type resultFormat struct {
	Status  int         `json:"status,omitempty"`
	Name    string      `json:"name,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   response    `json:"error,omitempty"`
	Meta    response    `json:"meta,omitempty"`
}

type paginateMeta struct {
}
