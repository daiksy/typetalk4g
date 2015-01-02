package typetalk4g

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	EndpointUrl     = "https://typetalk.in/api/v1/topics/{topicId}?typetalkToken={apiKey}"
	TopicIdReplacer = "{topicId}"
	ApiKeyReplacer  = "{apiKey}"
)

func buildUrl(topicId int, token string) string {
	r := strings.NewReplacer(TopicIdReplacer, strconv.Itoa(topicId), ApiKeyReplacer, token)
	return r.Replace(EndpointUrl)
}

type Client struct {
	topicId      int
	token        string
	MessagesSize int
}

func NewClient(topicId int, token string) *Client {
	client := &Client{
		topicId,
		token,
		20,
	}
	return client
}

func (client Client) Post(message string) (resp *http.Response, err error) {
	values := url.Values{}
	values.Add("message", message)

	resp, err = http.PostForm(buildUrl(client.topicId, client.token), values)

	return
}

func (client Client) GetTopicMessages() (messages Posts, err error) {
	resp, err := http.Get(buildUrl(client.topicId, client.token) + "&count=" + strconv.Itoa(client.MessagesSize))

	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Posts{}, err
		}
		err = json.Unmarshal(body, &messages)
	}
	defer resp.Body.Close()

	return
}
