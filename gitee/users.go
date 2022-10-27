package gitee

import (
	"context"
	"fmt"
)

// UsersService handles communication with the user related
// methods of the gitee API.
//
// gitee API docs:
type UsersService service

// User represents a gitee user.
type User struct {
	ID                *int64     `json:"id,omitempty"`
	Login             *string    `json:"login,omitempty"`
	Name              *string    `json:"name,omitempty"`
	AvatarUrl         *string    `json:"avatar_url,omitempty"`
	Url               *string    `json:"url,omitempty"`
	HtmlUrl           *string    `json:"html_url,omitempty"`
	Remark            *string    `json:"remark,omitempty"`
	FollowersUrl      *string    `json:"followers_url,omitempty"`
	FollowingUrl      *string    `json:"following_url,omitempty"`
	GistsUrl          *string    `json:"gists_url,omitempty"`
	StarredUrl        *string    `json:"starred_url,omitempty"`
	SubscriptionsUrl  *string    `json:"subscriptions_url,omitempty"`
	OrganizationsUrl  *string    `json:"organizations_url,omitempty"`
	ReposUrl          *string    `json:"repos_url,omitempty"`
	EventsUrl         *string    `json:"events_url,omitempty"`
	ReceivedEventsUrl *string    `json:"received_events_url,omitempty"`
	Type              *string    `json:"type,omitempty"`
	SiteAdmin         *bool      `json:"site_admin,omitempty"`
	Blog              *string    `json:"blog,omitempty"`
	Weibo             *string    `json:"weibo,omitempty"`
	Bio               *string    `json:"bio,omitempty"`
	PublicRepos       *int       `json:"public_repos,omitempty"`
	PublicGists       *int       `json:"public_gists,omitempty"`
	Followers         *int       `json:"followers,omitempty"`
	Following         *int       `json:"following,omitempty"`
	Stared            *int       `json:"stared,omitempty"`
	Watched           *int       `json:"watched,omitempty"`
	CreatedAt         *Timestamp `json:"created_at,omitempty"`
	UpdatedAt         *Timestamp `json:"updated_at,omitempty"`
	Email             *string    `json:"email,omitempty"`
}

// UserListOptions specifies optional parameters to the UsersService.ListAll
// method.
type UserListOptions struct {
	// Note: Pagination is powered exclusively by the Since parameter,
	// ListOptions.Page has no effect.
	// ListOptions.PerPage controls an undocumented GitHub API parameter.
	ListOptions
}

func (u User) String() string {
	return Stringify(u)
}

// 获取一个用户 GET https://gitee.com/api/v5/users/{username}
// 获取授权用户的资料 GET https://gitee.com/api/v5/user
func (s *UsersService) Get(ctx context.Context, user string) (*User, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v", user)
	} else {
		u = "user"
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(User)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

type SshKey struct {
	// 获取一个公钥
	ID        *int64     `json:"id,omitempty"`
	Key       *string    `json:"key,omitempty"`
	Url       *string    `json:"url,omitempty"`
	Title     *string    `json:"title,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
}

func (k SshKey) String() string {
	return Stringify(k)
}

// 获取当前授权用户的sshkey，这个能获取多个的，一个列表
// 列出授权用户的所有公钥 GET https://gitee.com/api/v5/user/keys
func (s *UsersService) GetUserKeys(ctx context.Context, opts *ListOptions) ([]*SshKey, *Response, error) {
	u, err := addOptions("user/keys", opts)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var keys []*SshKey
	resp, err := s.client.Do(ctx, req, &keys)
	if err != nil {
		return nil, resp, err
	}

	return keys, resp, nil

}

// 通过sshkey的id来获取公钥
// 获取一个公钥 GET https://gitee.com/api/v5/user/keys/{id}  id=公钥 ID
func (s *UsersService) GetUserKey(ctx context.Context, id int64) (*SshKey, *Response, error) {
	var u string
	u = fmt.Sprintf("user/keys/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var keys *SshKey
	resp, err := s.client.Do(ctx, req, &keys)
	if err != nil {
		return nil, resp, err
	}

	return keys, resp, nil
}

// 列出指定用户的关注者 GET https://gitee.com/api/v5/users/{username}/followers
// 列出授权用户的关注者 GET https://gitee.com/api/v5/user/followers 这个获取的是 我被哪些人 关注了
func (s *UsersService) GetUserFollowers(ctx context.Context, user string, opts *ListOptions) ([]*User, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/followers", user)
	} else {
		u = "user/followers"
	}

	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := s.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil

}

// 列出指定用户正在关注的用户 GET https://gitee.com/api/v5/users/{username}/following
// 列出授权用户正关注的用户 GET https://gitee.com/api/v5/user/following  这是获取的是我关注的哪些人，或者某个账号下面他关注的其他人
func (s *UsersService) GetUserFollowings(ctx context.Context, user string, opts *ListOptions) ([]*User, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/following", user)
	} else {
		u = "user/following"
	}

	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*User
	resp, err := s.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil

}
