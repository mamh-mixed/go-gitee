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

// ?????????????????? GET https://gitee.com/api/v5/users/{username}
// ??????????????????????????? GET https://gitee.com/api/v5/user
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

type UserEditRequest struct {
	Name  *string `json:"name,omitempty"` //
	Blog  *string `json:"blog,omitempty"`
	Weibo *string `json:"weibo,omitempty"`
	Bio   *string `json:"bio,omitempty"`
}

// Edit the authenticated user.
//
//  ??????????????????????????? PATCH https://gitee.com/api/v5/user
func (s *UsersService) Edit(ctx context.Context, user *UserEditRequest) (*User, *Response, error) {
	u := "user"
	req, err := s.client.NewRequest("PATCH", u, user)
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

type SSHKey struct {
	// ??????????????????
	ID        *int64     `json:"id,omitempty"`
	Key       *string    `json:"key,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Title     *string    `json:"title,omitempty"`
	CreatedAt *Timestamp `json:"created_at,omitempty"`
}

func (k SSHKey) String() string {
	return Stringify(k)
}

type KeyCreateRequest struct {
	Key   *string `json:"key"`   // ????????????
	Title *string `json:"title"` // ????????????
}

// ???????????????????????????sshkey??????????????????????????????????????????
// ????????????????????????????????? GET https://gitee.com/api/v5/user/keys
// ????????????????????????????????? GET https://gitee.com/api/v5/users/{username}/keys
func (s *UsersService) ListKeys(ctx context.Context, user string, opts *ListOptions) ([]*SSHKey, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/keys", user)
	} else {
		u = "user/keys"
	}
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var keys []*SSHKey
	resp, err := s.client.Do(ctx, req, &keys)
	if err != nil {
		return nil, resp, err
	}

	return keys, resp, nil

}

// CreateKey adds a public key for the authenticated user.
//
//  ?????????????????? POST https://gitee.com/api/v5/user/keys
func (s *UsersService) CreateKey(ctx context.Context, key *KeyCreateRequest) (*SSHKey, *Response, error) {
	u := "user/keys"

	req, err := s.client.NewRequest("POST", u, key)
	if err != nil {
		return nil, nil, err
	}

	k := new(SSHKey)
	resp, err := s.client.Do(ctx, req, k)
	if err != nil {
		return nil, resp, err
	}

	return k, resp, nil
}

// ??????sshkey???id???????????????
// ?????????????????? GET https://gitee.com/api/v5/user/keys/{id}  id=?????? ID
func (s *UsersService) GetKey(ctx context.Context, id int64) (*SSHKey, *Response, error) {
	var u string
	u = fmt.Sprintf("user/keys/%v", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var keys *SSHKey
	resp, err := s.client.Do(ctx, req, &keys)
	if err != nil {
		return nil, resp, err
	}

	return keys, resp, nil
}

// DeleteKey deletes a public key.
//
//  ?????????????????? DELETE https://gitee.com/api/v5/user/keys/{id}
func (s *UsersService) DeleteKey(ctx context.Context, id int64) (*Response, error) {
	u := fmt.Sprintf("user/keys/%v", id)

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// ?????????????????????????????? GET https://gitee.com/api/v5/users/{username}/followers
// ?????????????????????????????? GET https://gitee.com/api/v5/user/followers ?????????????????? ??????????????? ?????????
func (s *UsersService) ListFollowers(ctx context.Context, user string, opts *ListOptions) ([]*User, *Response, error) {
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

// ??????????????????????????????????????? GET https://gitee.com/api/v5/users/{username}/following
// ???????????????????????????????????? GET https://gitee.com/api/v5/user/following  ???????????????????????????????????????????????????????????????????????????????????????
func (s *UsersService) ListFollowings(ctx context.Context, user string, opts *ListOptions) ([]*User, *Response, error) {
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
	Mode string `url:"mode,omitempty"` // ????????????: project(?????????????????????namepsce)???intrant(????????????namespace)???all(???????????????)?????????(intrant)
}

// ??????????????????????????? Namespace
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

// ??????????????????????????? Namespace GET https://gitee.com/api/v5/user/namespaces
// mode ????????????: project(?????????????????????namepsce)???intrant(????????????namespace)???all(???????????????)?????????(intrant)
func (s *UsersService) ListNamespaces(ctx context.Context, opts *NamespacesOptions) ([]*Namespace, *Response, error) {
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
	Path string `url:"path,omitempty"` // path Namespace path ??????????????????
}

// ??????????????????????????? Namespace GET https://gitee.com/api/v5/user/namespace
// path Namespace path ??????????????????
func (s *UsersService) GetNamespace(ctx context.Context, opts *NamespaceOptions) (*Namespace, *Response, error) {
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

// IsFollowing checks if "user" is following "target". Passing the empty
// string for "user" will check if the authenticated user is following "target".
//
//  ????????????????????????????????????????????? GET https://gitee.com/api/v5/user/following/{username}
//  ?????????????????????????????????????????? GET https://gitee.com/api/v5/users/{username}/following/{target_user}
func (s *UsersService) IsFollowing(ctx context.Context, user, target string) (bool, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/following/%v", user, target)
	} else {
		u = fmt.Sprintf("user/following/%v", target)
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return false, nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	following, err := parseBoolResponse(err)
	return following, resp, err
}

// Follow will cause the authenticated user to follow the specified user.
//
//  ?????????????????? PUT https://gitee.com/api/v5/user/following/{username}
func (s *UsersService) Follow(ctx context.Context, user string) (*Response, error) {
	u := fmt.Sprintf("user/following/%v", user)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Unfollow will cause the authenticated user to unfollow the specified user.
//
//  ???????????????????????? DELETE https://gitee.com/api/v5/user/following/{username}
func (s *UsersService) Unfollow(ctx context.Context, user string) (*Response, error) {
	u := fmt.Sprintf("user/following/%v", user)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
