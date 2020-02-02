package server

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

func banner() string {
	return `
                                   @@@
                                  // \\
                                 / \_/ \
                                / ./_(-.\
                               ( /      \\
                                )|| @/ // )
                               /  @>@<@/ (
                               '-.;)@\ \.-'
                                  |    |
                                  |    |
                                  |    |
                                  |    ` + "`" + `'.
                                  |        ` + "`" + `'-.
                                  /           '-..
                                  |          . .. '-.__
                                  |.' .' .'.'__;.--'--.-'
                                   ` + "`" + `"-----'"` + "`" + `

 .--.   .--. .-./` + "`" + `) ,---.   .--.  .-_'''-.    ______         ,-----.    ,---.    ,---.
 |  | _/  /  \ .-.')|    \  |  | '_( )_   \  |    _ ` + "`" + `''.   .'  .-,  '.  |    \  /    |
 | (` + "`" + `' ) /   / ` + "`" + `-' \|  ,  \ |  ||(_ o _)|  ' | _ | ) _  \ / ,-.|  \ _ \ |  ,  \/  ,  |
 |(_ ()_)     ` + "`" + `-'` + "`" + `"` + "`" + `|  |\_ \|  |. (_,_)/___| |( ''_'  ) |;  \  '_ /  | :|  |\_   /|  |
 | (_,_)   __ .---. |  _( )_\  ||  |  .-----.| . (_) ` + "`" + `. ||  _` + "`" + `,/ \ _/  ||  _( )_/ |  |
 |  |\ \  |  ||   | | (_ o _)  |'  \  '-   .'|(_    ._) ': (  '\_/ \   ;| (_ o _) |  |
 |  | \ ` + "`" + `'   /|   | |  (_,_)\  | \  ` + "`" + `-'` + "`" + `   | |  (_.\.' /  \ ` + "`" + `"/  \  ) / |  (_,_)  |  |
 |  |  \    / |   | |  |    |  |  \        / |       .'    '. \_/` + "`" + `` + "`" + `".'  |  |      |  |
 ` + "`" + `--'   ` + "`" + `'-'  '---' '--'    '--'   ` + "`" + `'-...-'  '-----'` + "`" + `        '-----'    '--'      '--'
`
}

type Server struct {
	*echo.Echo
	Dir  string
	Port int
}

func NewServer(port int, dir string) *Server {
	e := echo.New()
	s := &Server{e, dir, port}

	e.GET("/princess/random", s.randomPrincess)
	e.GET("/princess/:id", s.showPrincess)
	e.POST("/princess", s.savePrincess)

	e.HideBanner = true

	return s
}

func (s *Server) Start() {
	fmt.Println(banner())
	s.Logger.Fatal(s.Echo.Start(":" + strconv.Itoa(s.Port)))
}
