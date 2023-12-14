package models

import "time"

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement" `
	DisplayName string    `json:"displayName"`
	EmailId     string    `json:"emailId" gorm:"uniqueIndex"`
	Username    string    `json:"username" gorm:"uniqueIndex"`
	Password    string    `json:"password"`
	AboutMe     string    `json:"aboutMe" gorm:"default:null"`
	Location    string    `json:"location" gorm:"default:null"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	Posts       []Post    `json:"posts" gorm:"foreignKey:UserId"`
}

type Post struct {
	ID               uint      `json:"id" gorm:"primaryKey;autoIncrement" `
	UserId           uint      `json:"userId" gorm:"foreignKey:UserId"`
	User             User      `json:"user" gorm:"foreignKey:UserId"`
	ParentQuestionId uint      `json:"parentQuestionId" gorm:"foreignKey:ParentQuestionId"`
	ParentQuestion   *Post     `json:"parentQuestion" gorm:"foreignKey:ParentQuestionId"`
	PostTypeId       uint      `json:"postTypeId"`
	PostType         PostType  `json:"postType" gorm:"foreignKey:PostTypeId"`
	AcceptedAnswerId uint      `json:"acceptedAnswerId" gorm:"foreignKey:AcceptedAnswerId"`
	AcceptedAnswer   *Post     `json:"acceptedAnswer" gorm:"foreignKey:AcceptedAnswerId"`
	Title            string    `json:"title"`
	Details          string    `json:"details"`
	CreatedAt        time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type PostType struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement" `
	Name string `json:"name"`
}

type Tag struct {
	ID             uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	TagName        string `json:"tagName"`
	TagDescription string `json:"tagDescription"`
}

type PostTag struct {
	ID     uint `json:"id" gorm:"primaryKey;autoIncrement"`
	PostID uint `json:"postId"`
	Post   Post `json:"post" gorm:"foreignKey:PostID"`
	TagID  uint `json:"tagId"`
	Tag    Tag  `json:"tag" gorm:"foreignKey:TagID"`
}

type VoteType struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	VoteType string `json:"voteType"`
}

type Vote struct {
	ID         uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	PostID     uint     `json:"postId"`
	Post       Post     `json:"post" gorm:"foreignKey:PostID"`
	UserID     uint     `json:"userId"`
	User       User     `json:"user" gorm:"foreignKey:UserID"`
	VoteTypeID uint     `json:"voteTypeId"`
	VoteType   VoteType `json:"voteType" gorm:"foreignKey:VoteTypeID"`
}

type Comment struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      uint      `json:"userId"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	PostID      uint      `json:"postId"`
	Post        Post      `json:"post" gorm:"foreignKey:PostID"`
	CommentText string    `json:"commentText"`
	CommentDate time.Time `json:"commentDate" gorm:"autoCreateTime"`
}
