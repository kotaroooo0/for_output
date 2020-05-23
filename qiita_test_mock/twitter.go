package qiita_test_mock

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

type ITwitterApiClient interface {
	PostTweet(string, url.Values) (anaconda.Tweet, error)
}

// Golangでよく見るanacondaのTwitterClient
// *anaconda.TwitterApiはPostTweet(string, url.Values) (anaconda.Tweet, error)を実装している
// したがって、TwitterApiClientはITwitterApiClientを満たしている
type TwitterApiClient *anaconda.TwitterApi
