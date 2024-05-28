package main

import (
	"fmt"

	"github.com/Lensual/agora_rtm_sdk_cgo/pkg/agora"
)

type RtmEventHandler struct{}

func (h *RtmEventHandler) OnMessageEvent(event *agora.MessageEvent) {
	fmt.Printf("OnMessageEvent event:%v\n", event) //DEBUG
}
func (h *RtmEventHandler) OnPresenceEvent(event *agora.PresenceEvent) {
	fmt.Printf("OnPresenceEvent event:%v\n", event) //DEBUG
}
func (h *RtmEventHandler) OnTopicEvent(event *agora.TopicEvent) {
	fmt.Printf("OnTopicEvent event:%v\n", event) //DEBUG
}
func (h *RtmEventHandler) OnLockEvent(event *agora.LockEvent) {
	fmt.Printf("OnLockEvent event:%v\n", event) //DEBUG
}
func (h *RtmEventHandler) OnStorageEvent(event *agora.StorageEvent) {
	fmt.Printf("OnStorageEvent event:%v\n", event) //DEBUG
}
func (h *RtmEventHandler) OnJoinResult(requestId uint64, channelName string, userId string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnLeaveResult(requestId uint64, channelName string, userId string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnJoinTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnLeaveTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnSubscribeTopicResult(requestId uint64, channelName string, userId string, topic string, succeedUsers agora.UserList, failedUsers agora.UserList, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnConnectionStateChanged(channelName string, state agora.CONNECTION_STATE, reason agora.RTM_CONNECTION_CHANGE_REASON) {

}
func (h *RtmEventHandler) OnTokenPrivilegeWillExpire(channelName string) {

}
func (h *RtmEventHandler) OnSubscribeResult(requestId uint64, channelName string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnPublishResult(requestId uint64, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnLoginResult(errorCode agora.RTM_ERROR_CODE) {
	fmt.Printf("OnLoginResult errorCode:%v\n", errorCode) //DEBUG
}
func (h *RtmEventHandler) OnSetChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnUpdateChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnRemoveChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnGetChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, data *agora.IMetadata, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnSetUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnUpdateUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnRemoveUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnGetUserMetadataResult(requestId uint64, userId string, data *agora.IMetadata, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnSubscribeUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnSetLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnRemoveLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnReleaseLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnAcquireLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE, errorDetails string) {

}
func (h *RtmEventHandler) OnRevokeLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnGetLocksResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockDetailList *agora.LockDetail, count uint, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnWhoNowResult(requestId uint64, userStateList *agora.UserState, count uint, nextPage string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnGetOnlineUsersResult(requestId uint64, userStateList *agora.UserState, count uint, nextPage string, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnWhereNowResult(requestId uint64, channels *agora.ChannelInfo, count uint, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnGetUserChannelsResult(requestId uint64, channels *agora.ChannelInfo, count uint, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnPresenceSetStateResult(requestId uint64, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnPresenceRemoveStateResult(requestId uint64, errorCode agora.RTM_ERROR_CODE) {

}
func (h *RtmEventHandler) OnPresenceGetStateResult(requestId uint64, state *agora.UserState, errorCode agora.RTM_ERROR_CODE) {

}
