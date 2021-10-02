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
	"github.com/iancoleman/orderedmap"
	"sort"
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

	usersInfo := make(map[uint32]interface{},0)
	for _, val := range userSrvResp.UsersInfo {
		usersInfo[val.Id] = val
	}

	groupList := orderedmap.New()

	for _, v1 := range resp.UserGroup {
		groupInfo := make(map[string]interface{},0)
		groupInfo["id"] = v1.Id
		groupInfo["groupname"] = v1.Name
		groupInfo["list"] = make([]map[string]interface{},0)
		i := 0
		for k2,v2 := range resp.UserFriendLists {
			if v1.Id == v2.GroupId {
				userFiendInfo := make(map[string]interface{},0)
				userFiendInfo["id"] = v2.FriendId
				userFiendInfo["username"] = usersInfo[v2.FriendId].(*userpb.GetUserInfoResponse).Mobile
				userFiendInfo["avatar"] = usersInfo[v2.FriendId].(*userpb.GetUserInfoResponse).Avatar
				userFiendInfo["sign"] = usersInfo[v2.FriendId].(*userpb.GetUserInfoResponse).Sign
				groupInfo["list"] = append(groupInfo["list"].([]map[string]interface{}),userFiendInfo)
				resp.UserFriendLists = append(resp.UserFriendLists[:k2 - i],resp.UserFriendLists[k2 - i:]...)
				i++
			}
		}
		groupList.Set(string(int32(v1.Id)),groupInfo)
	}
	groupList.SortKeys(sort.Strings)

	returnData := make([]interface{},0)
	for _, k := range groupList.Keys() {
		v, _ := groupList.Get(k)
		returnData = append(returnData,v)
	}

	utils.OkReturn(ctx, returnData, "")
}
