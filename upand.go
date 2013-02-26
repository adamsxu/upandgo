package upandgo

var (
	Server   *AppServer
	AppPath  string
	HttpPort int
	HttpAddr string
	RunMode  string
	Config   *AppConfig
)

type AppServer struct {
	Id   int
	Addr string
}

type AppConfig struct {
}

func (s *AppServer) Run() {
	addr := fmt.Sprintf("%s:%d", HttpAddr, HttpPort)
	err := http.ListenAndServe(addr, app.Handlers)
	if err != nil {
		Fatal("ListenAndServe: ", err)
	}
}

func NewAppServer() {

}

func init() {
	AppServer = NewAppServer()
}
