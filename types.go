package typetalk4g

import "time"

type Team struct {
	Id        int       `json:id`
	Name      string    `json:name`
	ImageUrl  string    `json:imageUrl`
	CreatedAt time.Time `json:createdAt`
	UpdatedAt time.Time `json:updatedAt`
}

type Topic struct {
	Id           int       `json:id`
	Name         string    `json:name`
	Suggestion   string    `json:suggestion`
	LastPostedAt time.Time `json:lastPostedAt`
	CreatedAt    time.Time `json:createdAt`
	UpdatedAt    time.Time `json:updatedAt`
}

type Bookmark struct {
	PostId    int       `json:postId`
	UpdatedAt time.Time `json:updatedAt`
}

type Account struct {
	Id         int       `json:id`
	Name       string    `json:name`
	FullName   string    `fullName`
	Suggestion string    `json:suggestion`
	ImageUrl   string    `json:imageUrl`
	CreatedAt  time.Time `json:createdAt`
	UpdatedAt  time.Time `json:updatedAt`
}

type Post struct {
	Id      int     `json:id`
	TopicId int     `json:topicId`
	Message string  `json:message`
	Account Account `json:account`
}

type Posts struct {
	Team     Team     `json:team`
	Topic    Topic    `json:tipic`
	Bookmark Bookmark `json:bookmark`
	Posts    []Post   `json:posts`
}
