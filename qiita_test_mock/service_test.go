package qiita_test_mock

import (
	"net/url"
	"testing"

	"github.com/ChimeraCoder/anaconda"
	"github.com/google/go-cmp/cmp"
)

type TwitterApiClientMock struct{}

// ITwitterApiClientを満たすようにPostTweetを実装する
func (m TwitterApiClientMock) PostTweet(content string, values url.Values) (anaconda.Tweet, error) {
	return anaconda.Tweet{Text: content}, nil
}

func TestPostHelloWorld(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{
			input:  "Bob",
			output: "Hello World by Bob",
		},
		{
			input:  "Alice",
			output: "Hello World by Alice",
		},
		{
			input:  "Mike",
			output: "Hello World by Mike",
		},
	}
	// TwitterApiClientMockをインジェクションすることで、anacondaのAPIクライアントを用いない
	// したがって、実際にツイートすることなくテストを実行することができる
	serviceImpl := ServiceImpl{TwitterApiClient: TwitterApiClientMock{}}
	for _, tt := range cases {
		tweetContent, err := serviceImpl.PostHelloWorld(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(tweetContent, tt.output); diff != "" {
			t.Errorf("Diff: (-got +want)\n%s", diff)
		}
		tweetContent, err = PostHelloWorld(tt.input, TwitterApiClientMock{})
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(tweetContent, tt.output); diff != "" {
			t.Errorf("Diff: (-got +want)\n%s", diff)
		}
	}
}
