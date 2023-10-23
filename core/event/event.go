package event

type Events struct {
	Event string `json:"event"`
	Data  interface{}
}
