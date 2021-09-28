package api

import (
	"chat_api/errno"
	"chat_api/global"
	"chat_api/proto/gen/go/chat_pb"
	"chat_api/proto/gen/go/userpb"
	"chat_api/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func GetUserFriendList(ctx *gin.Context) {
	userId, ok := ctx.Get("userId")
	if !ok {
		utils.ErrReturn(ctx, http.StatusUnauthorized, errno.ErrUserNotLogin)
		return
	}
	resp, err := global.ChatSrvClient.GetFriendList(ctx, &chat_pb.GetFriendListRequest{UserId: userId.(uint32)})
	if err != nil {
		zap.S().Errorf("Chat服务错误:%s", err.Error())
		utils.ErrReturn(ctx, http.StatusBadRequest, errno.ChatSrvErr)
		return
	}
	// 查询用户名
	var friendIds []uint32
	for _, val := range resp.UserFriendLists {
		friendIds = append(friendIds, val.FriendId)
	}
	userSrvResp, err := global.UserSrvClient.GetUsersInfo(ctx, &userpb.GetUsersInfoReq{Ids: friendIds})
	if err != nil {
		zap.S().Errorf("User服务错误:%s", err.Error())
		utils.ErrReturn(ctx, http.StatusBadRequest, errno.ChatSrvErr)
		return
	}

	group := map[uint32]interface{}{}
	for _, val := range resp.UserFriendLists {
		_, ok := group[val.GroupId]
		var groupList []map[string]interface{}
		groupInfo := map[string]interface{}{}

		if !ok {
			groupInfo["id"] = val.GroupId
			groupInfo["groupname"] = val.GroupName
			groupInfo["list"] = []map[string]interface{}{}
		} else {
			groupInfo = group[val.GroupId].(map[string]interface{})
		}
		groupList = groupInfo["list"].([]map[string]interface{})
		for k, val2 := range userSrvResp.UsersInfo {
			if val2.Id == val.FriendId {
				userInfo := map[string]interface{}{}
				userInfo["id"] = val2.Id
				userInfo["username"] = val2.Mobile
				userInfo["avatar"] = val2.Avatar
				userInfo["sign"] = val2.Sign
				groupList = append(groupList, userInfo)
				userSrvResp.UsersInfo = append(userSrvResp.UsersInfo[:k], userSrvResp.UsersInfo[k+1:]...)
			}
		}
		groupInfo["list"] = groupList
		group[val.GroupId] = groupInfo
	}

	var returnData []interface{}
	for _, val := range group {
		returnData = append(returnData, val)
	}
	utils.OkReturn(ctx, returnData, "")
}
