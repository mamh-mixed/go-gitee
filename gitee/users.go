//Copyright magesfc bright.ma
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

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

type BasicUser struct {
	ID                *int64  `json:"id,omitempty"`
	Login             *string `json:"login,omitempty"`
	Name              *string `json:"name,omitempty"`
	AvatarURL         *string `json:"avatar_url,omitempty"`
	URL               *string `json:"url,omitempty"`
	HTMLURL           *string `json:"html_url,omitempty"`
	Remark            *string `json:"remark,omitempty"`
	FollowersURL      *string `json:"followers_url,omitempty"`
	FollowingURL      *string `json:"following_url,omitempty"`
	GistsURL          *string `json:"gists_url,omitempty"`
	StarredURL        *string `json:"starred_url,omitempty"`
	SubscriptionsURL  *string `json:"subscriptions_url,omitempty"`
	OrganizationsURL  *string `json:"organizations_url,omitempty"`
	ReposURL          *string `json:"repos_url,omitempty"`
	EventsURL         *string `json:"events_url,omitempty"`
	ReceivedEventsURL *string `json:"received_events_url,omitempty"`
	Type              *string `json:"type,omitempty"`
}

// User represents a gitee user.
type User struct {
	*BasicUser
	SiteAdmin   *bool      `json:"site_admin,omitempty"`
	Blog        *string    `json:"blog,omitempty"`
	Weibo       *string    `json:"weibo,omitempty"`
	Bio         *string    `json:"bio,omitempty"`
	PublicRepos *int       `json:"public_repos,omitempty"`
	PublicGists *int       `json:"public_gists,omitempty"`
	Followers   *int       `json:"followers,omitempty"`
	Following   *int       `json:"following,omitempty"`
	Stared      *int       `json:"stared,omitempty"`
	Watched     *int       `json:"watched,omitempty"`
	CreatedAt   *Timestamp `json:"created_at,omitempty"`
	UpdatedAt   *Timestamp `json:"updated_at,omitempty"`
	Email       *string    `json:"email,omitempty"`
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
	URL       *string    `json:"url,omitempty"`
	Title     *string    `json:"title,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
}

func (k SshKey) String() string {
	return Stringify(k)
}

// 获取当前授权用户的sshkey，这个能获取多个的，一个列表
// 列出授权用户的所有公钥 GET https://gitee.com/api/v5/user/keys
func (s *UsersService) ListSshKeys(ctx context.Context, opts *ListOptions) ([]*SshKey, *Response, error) {
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
func (s *UsersService) GetSshKey(ctx context.Context, id int64) (*SshKey, *Response, error) {
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

type NamespacesOptions struct {
	Mode string `url:"mode,omitempty"` // 参与方式: project(所有参与仓库的namepsce)、intrant(所加入的namespace)、all(包含前两者)，默认(intrant)
}

// 获取授权用户的一个 Namespace
type Namespace struct {
	ID      *int64     `json:"id,omitempty"`
	Type    *string    `json:"type,omitempty"`
	Name    *string    `json:"name,omitempty"`
	Path    *string    `json:"path,omitempty"`
	HTMLURL *string    `json:"html_url,omitempty"`
	Parent  *Namespace `json:"parent,omitempty"`
}

func (n Namespace) String() string {
	return Stringify(n)
}

// 列出授权用户所有的 Namespace GET https://gitee.com/api/v5/user/namespaces
// mode 参与方式: project(所有参与仓库的namepsce)、intrant(所加入的namespace)、all(包含前两者)，默认(intrant)
func (s *UsersService) GetUserNamespaces(ctx context.Context, opts *NamespacesOptions) ([]*Namespace, *Response, error) {
	u, err := addOptions("user/namespaces", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var namespaces []*Namespace
	resp, err := s.client.Do(ctx, req, &namespaces)
	if err != nil {
		return nil, resp, err
	}

	return namespaces, resp, nil

}

type NamespaceOptions struct {
	Path string `url:"path,omitempty"` // path Namespace path 需要一个参数
}

// 获取授权用户的一个 Namespace GET https://gitee.com/api/v5/user/namespace
// path Namespace path 需要一个参数
func (s *UsersService) GetUserNamespace(ctx context.Context, opts *NamespaceOptions) (*Namespace, *Response, error) {
	u, err := addOptions("user/namespace", opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	namespace := new(Namespace)
	resp, err := s.client.Do(ctx, req, namespace)
	if err != nil {
		return nil, resp, err
	}

	return namespace, resp, nil

}
