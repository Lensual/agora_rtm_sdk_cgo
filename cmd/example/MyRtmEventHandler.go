package main

import (
	"fmt"

	"github.com/Lensual/agora_rtm_sdk_cgo/pkg/agora"
)

type MyRtmEventHandler struct{}

func (h *MyRtmEventHandler) OnMessageEvent(event *agora.MessageEvent) {
	fmt.Printf("OnMessageEvent event:%v\n",
		event) //DEBUG
}
func (h *MyRtmEventHandler) OnPresenceEvent(event *agora.PresenceEvent) {
	fmt.Printf("OnPresenceEvent event:%v\n",
		event) //DEBUG
}
func (h *MyRtmEventHandler) OnTopicEvent(event *agora.TopicEvent) {
	fmt.Printf("OnTopicEvent event:%v\n",
		event) //DEBUG
}
func (h *MyRtmEventHandler) OnLockEvent(event *agora.LockEvent) {
	fmt.Printf("OnLockEvent event:%v\n",
		event) //DEBUG
}
func (h *MyRtmEventHandler) OnStorageEvent(event *agora.StorageEvent) {
	fmt.Printf("OnStorageEvent event:%v\n",
		event) //DEBUG
}
func (h *MyRtmEventHandler) OnJoinResult(requestId uint64, channelName string, userId string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnJoinResult requestId:%v channelName:%v userId:%v errorCode:%v\n",
		requestId, channelName, userId, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnLeaveResult(requestId uint64, channelName string, userId string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnLeaveResult requestId:%v channelName:%v userId:%v errorCode:%v\n",
		requestId, channelName, userId, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnJoinTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnJoinTopicResult requestId:%v channelName:%v userId:%v topic:%v meta:%v errorCode:%v\n",
		requestId, channelName, userId, topic, meta, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnLeaveTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnLeaveTopicResult requestId:%v channelName:%v userId:%v topic:%v meta:%v errorCode:%v\n",
		requestId, channelName, userId, topic, meta, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnSubscribeTopicResult(requestId uint64, channelName string, userId string, topic string, succeedUsers agora.UserList, failedUsers agora.UserList, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnSubscribeTopicResult requestId:%v channelName:%v userId:%v topic:%v succeedUsers:%v failedUsers:%v errorCode:%v\n",
		requestId, channelName, userId, topic, succeedUsers, failedUsers, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnConnectionStateChanged(channelName string, state agora.RTM_CONNECTION_STATE, reason agora.RTM_CONNECTION_CHANGE_REASON) {
	fmt.Printf("OnConnectionStateChanged channelName:%v state:%v reason:%v\n",
		channelName, state, reason) //DEBUG
}
func (h *MyRtmEventHandler) OnTokenPrivilegeWillExpire(channelName string) {
	fmt.Printf("OnTokenPrivilegeWillExpire channelName:%v\n",
		channelName) //DEBUG
}
func (h *MyRtmEventHandler) OnSubscribeResult(requestId uint64, channelName string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnSubscribeResult requestId:%v channelName:%v errorCode:%v\n",
		requestId, channelName, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnPublishResult(requestId uint64, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnPublishResult requestId:%v errorCode:%v\n",
		requestId, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnLoginResult(errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnLoginResult errorCode:%v\n",
		errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnSetChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnSetChannelMetadataResult requestId:%v channelName:%v channelType:%v errorCode:%v\n",
		requestId, channelName, channelType, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnUpdateChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnUpdateChannelMetadataResult requestId:%v channelName:%v channelType:%v errorCode:%v\n",
		requestId, channelName, channelType, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnRemoveChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnRemoveChannelMetadataResult requestId:%v channelName:%v channelType:%v errorCode:%v\n",
		requestId, channelName, channelType, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnGetChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, data *agora.IMetadata, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnGetChannelMetadataResult requestId:%v channelName:%v channelType:%v data:%v errorCode:%v\n",
		requestId, channelName, channelType, data, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnSetUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnSetUserMetadataResult requestId:%v userId:%v errorCode:%v\n",
		requestId, userId, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnUpdateUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnUpdateUserMetadataResult requestId:%v userId:%v errorCode:%v\n",
		requestId, userId, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnRemoveUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnRemoveUserMetadataResult requestId:%v userId:%v errorCode:%v\n",
		requestId, userId, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnGetUserMetadataResult(requestId uint64, userId string, data *agora.IMetadata, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnGetUserMetadataResult requestId:%v userId:%v data:%v errorCode:%v\n",
		requestId, userId, data, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnSubscribeUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnSubscribeUserMetadataResult requestId:%v userId:%v errorCode:%v\n",
		requestId, userId, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnSetLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnSetLockResult requestId:%v channelName:%v channelType:%v lockName:%v errorCode:%v\n",
		requestId, channelName, channelType, lockName, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnRemoveLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnRemoveLockResult requestId:%v channelName:%v channelType:%v lockName:%v errorCode:%v\n",
		requestId, channelName, channelType, lockName, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnReleaseLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnReleaseLockResult requestId:%v channelName:%v channelType:%v lockName:%v errorCode:%v\n",
		requestId, channelName, channelType, lockName, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnAcquireLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE, errorDetails string) {
	fmt.Printf("OnAcquireLockResult requestId:%v channelName:%v channelType:%v lockName:%v errorCode:%v\n",
		requestId, channelName, channelType, lockName, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnRevokeLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnRevokeLockResult requestId:%v channelName:%v channelType:%v lockName:%v errorCode:%v\n",
		requestId, channelName, channelType, lockName, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnGetLocksResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockDetailList *agora.LockDetail, count uint, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnGetLocksResult requestId:%v channelName:%v channelType:%v lockDetailList:%v errorCode:%v\n",
		requestId, channelName, channelType, lockDetailList, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnWhoNowResult(requestId uint64, userStateList *agora.UserState, count uint, nextPage string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnWhoNowResult requestId:%v userStateList:%v count:%v nextPage:%v errorCode:%v\n",
		requestId, userStateList, count, nextPage, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnGetOnlineUsersResult(requestId uint64, userStateList *agora.UserState, count uint, nextPage string, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnGetOnlineUsersResult requestId:%v userStateList:%v count:%v nextPage:%v errorCode:%v\n",
		requestId, userStateList, count, nextPage, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnWhereNowResult(requestId uint64, channels *agora.ChannelInfo, count uint, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnWhereNowResult requestId:%v channels:%v count:%v errorCode:%v\n",
		requestId, channels, count, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnGetUserChannelsResult(requestId uint64, channels *agora.ChannelInfo, count uint, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnGetUserChannelsResult requestId:%v channels:%v count:%v errorCode:%v\n",
		requestId, channels, count, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnPresenceSetStateResult(requestId uint64, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnPresenceSetStateResult requestId:%v errorCode:%v\n",
		requestId, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnPresenceRemoveStateResult(requestId uint64, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnPresenceRemoveStateResult requestId:%v errorCode:%v\n",
		requestId, errorCode) //DEBUG
}
func (h *MyRtmEventHandler) OnPresenceGetStateResult(requestId uint64, state *agora.UserState, errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnPresenceGetStateResult requestId:%v state:%v errorCode:%v\n",
		requestId, state, errorCode) //DEBUG
}
