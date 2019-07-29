package templates

func Url_for(url string) string {
	switch url {
	case "index":
		return "/idex.html"
	case "login":
		return "/login.html"
	case "loginForm":
		return "/login.html"
	case "register":
		return "/resigter.html"
	default:
		return ""
	}
}
