package api

import (
	"context"

	"github.com/goNiki/ReviewService/internal/api/pullrequests"
	"github.com/goNiki/ReviewService/internal/api/teams"
	"github.com/goNiki/ReviewService/internal/api/users"
	revV1 "github.com/goNiki/ReviewService/shared/pkg/openapi/reviewerservice/v1"
)

type CompositeApi struct {
	teams        *teams.Api
	users        *users.Api
	pullrequests *pullrequests.Api
}

func NewCompositeApi(teamsApi *teams.Api, usersApi *users.Api, pullrequestsApi *pullrequests.Api) *CompositeApi {
	return &CompositeApi{
		teams:        teamsApi,
		users:        usersApi,
		pullrequests: pullrequestsApi,
	}
}

func (c *CompositeApi) PullRequestCreatePost(ctx context.Context, req *revV1.PullRequestCreatePostReq) (r revV1.PullRequestCreatePostRes, _ error) {
	return c.pullrequests.PullRequestCreatePost(ctx, req)
}

func (c *CompositeApi) PullRequestMergePost(ctx context.Context, req *revV1.PullRequestMergePostReq) (r revV1.PullRequestMergePostRes, _ error) {
	return c.pullrequests.PullRequestMergePost(ctx, req)
}

func (c *CompositeApi) PullRequestReassignPost(ctx context.Context, req *revV1.PullRequestReassignPostReq) (r revV1.PullRequestReassignPostRes, _ error) {
	return c.pullrequests.PullRequestReassignPost(ctx, req)
}

func (c *CompositeApi) TeamAddPost(ctx context.Context, req *revV1.Team) (r revV1.TeamAddPostRes, _ error) {
	return c.teams.TeamAddPost(ctx, req)
}

func (c *CompositeApi) TeamGetGet(ctx context.Context, params revV1.TeamGetGetParams) (r revV1.TeamGetGetRes, _ error) {
	return c.teams.TeamGetGet(ctx, params)
}

func (c *CompositeApi) UsersGetReviewGet(ctx context.Context, params revV1.UsersGetReviewGetParams) (r revV1.UsersGetReviewGetRes, _ error) {
	return c.users.UsersGetReviewGet(ctx, params)
}

func (c *CompositeApi) UsersSetIsActivePost(ctx context.Context, req *revV1.UsersSetIsActivePostReq) (r revV1.UsersSetIsActivePostRes, _ error) {
	return c.users.UsersSetIsActivePost(ctx, req)
}
