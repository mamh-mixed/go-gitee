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

// OrganizationsService provides access to the organization related functions
// in the gitee API.
type OrganizationsService service

type Organization struct {
	ID          *int64  `json:"id,omitempty"`
	Login       *string `json:"login,omitempty"`
	Name        *string `json:"name,omitempty"`
	URL         *string `json:"url,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
	ReposURL    *string `json:"repos_url,omitempty"`
	EventsURL   *string `json:"events_url,omitempty"`
	MembersURL  *string `json:"members_url,omitempty"`
	Description *string `json:"description,omitempty"`
	FollowCount *int64  `json:"follow_count,omitempty"`
}

func (o Organization) String() string {
	return Stringify(o)
}

type OrganizationListOptions struct {
	Admin *bool `url:"admin,omitempty"` // 只列出授权用户管理的组织
	ListOptions
}

// List the organizations for a user. Passing the empty string will list
// organizations for the authenticated user.
//
// 列出授权用户所属的组织 GET https://gitee.com/api/v5/user/orgs
// 列出用户所属的组织    GET https://gitee.com/api/v5/users/{username}/orgs
func (s *OrganizationsService) List(ctx context.Context, user string, opts *OrganizationListOptions) ([]*Organization, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("users/%v/orgs", user)
	} else {
		u = "user/orgs"
	}
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var orgs []*Organization
	resp, err := s.client.Do(ctx, req, &orgs)
	if err != nil {
		return nil, resp, err
	}

	return orgs, resp, nil
}

type Membership struct {
	URL             *string       `json:"url,omitempty"`
	Active          *bool         `json:"active,omitempty"`
	Remark          *string       `json:"remark,omitempty"`
	Role            *string       `json:"role,omitempty"`
	OrganizationURL *string       `json:"organization_url,omitempty"`
	Organization    *Organization `json:"organization,omitempty"`
	User            *User         `json:"user,omitempty"`
}

func (o Membership) String() string {
	return Stringify(o)
}

type MembershipListOptions struct {
	Active *bool `url:"active,omitempty"` // 根据成员是否已激活进行筛选资料，缺省返回所有资料
	ListOptions
}

// 列出授权用户在所属组织的成员资料 GET https://gitee.com/api/v5/user/memberships/orgs
func (s *OrganizationsService) ListOrgMemberships(ctx context.Context, opts *MembershipListOptions) ([]*Membership, *Response, error) {
	u := "user/memberships/orgs"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var memberships []*Membership
	resp, err := s.client.Do(ctx, req, &memberships)
	if err != nil {
		return nil, resp, err
	}

	return memberships, resp, nil
}

// 获取授权用户在一个组织的成员资料 GET https://gitee.com/api/v5/user/memberships/orgs/{org}
// 获取授权用户所属组织的一个成员 GET https://gitee.com/api/v5/orgs/{org}/memberships/{username}
func (s *OrganizationsService) GetOrgMembership(ctx context.Context, user, org string) (*Membership, *Response, error) {
	var u string
	if user != "" {
		u = fmt.Sprintf("orgs/%v/memberships/%v", org, user)
	} else {
		u = fmt.Sprintf("user/memberships/orgs/%v", org)
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	membership := new(Membership)
	resp, err := s.client.Do(ctx, req, membership)
	if err != nil {
		return nil, resp, err
	}

	return membership, resp, nil
}

type MembershipEditRequest struct {
	Remark *string `json:"remark,omitempty"` //在组织中的备注信息
}

// 更新授权用户在一个组织的成员资料 PATCH https://gitee.com/api/v5/user/memberships/orgs/{org}
func (s *OrganizationsService) EditOrgMembership(ctx context.Context, org string, membership *MembershipEditRequest) (*Membership, *Response, error) {
	var u string

	u = fmt.Sprintf("user/memberships/orgs/%v", org)

	req, err := s.client.NewRequest("PATCH", u, membership)
	if err != nil {
		return nil, nil, err
	}

	m := new(Membership)
	resp, err := s.client.Do(ctx, req, m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}
