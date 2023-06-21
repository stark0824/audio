package routeauth

var apiRouter = []string{
	"/api/usercoll/drama",
	"/api/usercoll/checkDrama",
}

func ApiCheckLoginController(checkUrl string) bool {
	flag := false
	for _, url := range apiRouter {
		if url == checkUrl {
			flag = true
		}
	}
	return flag
}
