package env

import "os"

//get api-key of api
func GetNewsTokenAPI() string {
	return os.Getenv("NEWS_TOKEN_API")
}

func GetBaseURL() string {
	return os.Getenv("BASE_URL")
}
