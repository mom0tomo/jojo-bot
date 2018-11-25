package keys

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

// GetTwitterAPI is getting the Twitter API's responce
func GetTwitterAPI() *anaconda.TwitterApi {

	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"))
	return api
}
