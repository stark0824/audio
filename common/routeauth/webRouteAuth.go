package routeauth

var webRouter = []string{
	"/web/usercoll/drama",
	"/web/usercoll/checkDrama",
}

func WebCheckLoginController(checkUrl string) bool {
	flag := false
	for _, url := range webRouter {
		if url == checkUrl {
			flag = true
		}
	}
	return flag
}
