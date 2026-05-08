package schemas

type PostType string

const (
	PostTypeText  PostType = "TEXT"
	PostTypeMedia PostType = "MEDIA"
)

type UserMicropostVoteResponse struct {
	UserVoted bool `json:"user_voted"`
}
