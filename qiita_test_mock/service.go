package qiita_test_mock

import "fmt"

// このInterfaceの宣言はなくてもPostHelloWorldは動作します
// ただ、ここもInterfaceとすることでService層をさらに上のUseCase層等へインジェクションすることができます
// そうすると、UseCase層はService層のモックを用いて実装やテストを行うことができます
type Service interface {
	PostHelloWorld(string) (string, error)
}

type ServiceImpl struct {
	TwitterApiClient ITwitterApiClient
}

// ITwitterApiClientを利用して、ツイートする
// 構造体の要素としてDIする(フィールドインジェクション)
func (s ServiceImpl) PostHelloWorld(name string) (string, error) {
	content := fmt.Sprintf("Hello World by %s", name)
	tweet, err := s.TwitterApiClient.PostTweet(content, nil)
	if err != nil {
		return "", err
	}
	return tweet.Text, nil
}

// 上と同様の処理を行う
// 引数としてDIする(あまり見たことはないし非推奨)
func PostHelloWorld(name string, client ITwitterApiClient) (string, error) {
	content := fmt.Sprintf("Hello World by %s", name)
	tweet, err := client.PostTweet(content, nil)
	if err != nil {
		return "", err
	}
	return tweet.Text, nil
}
