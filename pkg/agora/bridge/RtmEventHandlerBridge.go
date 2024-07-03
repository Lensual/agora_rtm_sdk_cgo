package bridge

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../../third_party/agora_rtm_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../../third_party/agora_rtm_sdk_c -lagora_rtm_sdk_c -lstdc++

//链接AgoraRTM SDK
#cgo CFLAGS: -I${SRCDIR}/../../../third_party/agora_rtm_sdk_c/agora_rtm_sdk-prefix/src/agora_rtm_sdk/rtm/sdk/high_level_api/include
#cgo LDFLAGS: -L${SRCDIR}/../../../third_party/agora_rtm_sdk_c/agora_rtm_sdk-prefix/src/agora_rtm_sdk/rtm/sdk -lagora_rtm_sdk

#include "bridge/C_RtmEventHandlerBridge.h"

void cgo_RtmEventHandlerBridge_onMessageEvent(C_RtmEventHandlerBridge *this_, void *userData,
	struct C_MessageEvent *event);

void cgo_RtmEventHandlerBridge_onPresenceEvent(C_RtmEventHandlerBridge *this_, void *userData,
	struct C_PresenceEvent *event);

void cgo_RtmEventHandlerBridge_onTopicEvent(C_RtmEventHandlerBridge *this_, void *userData,
	struct C_TopicEvent *event);

void cgo_RtmEventHandlerBridge_onLockEvent(C_RtmEventHandlerBridge *this_, void *userData,
	struct C_LockEvent *event);

void cgo_RtmEventHandlerBridge_onStorageEvent(C_RtmEventHandlerBridge *this_, void *userData,
	struct C_StorageEvent *event);

void cgo_RtmEventHandlerBridge_onJoinResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, char *userId, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onLeaveResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, char *userId, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onJoinTopicResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, char *userId, char *topic, char *meta, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onLeaveTopicResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, char *userId, char *topic, char *meta, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onSubscribeTopicResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, char *userId, char *topic, struct C_UserList succeedUsers, struct C_UserList failedUsers, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onConnectionStateChanged(C_RtmEventHandlerBridge *this_, void *userData,
	char *channelName, enum C_RTM_CONNECTION_STATE state, enum C_RTM_CONNECTION_CHANGE_REASON reason);

void cgo_RtmEventHandlerBridge_onTokenPrivilegeWillExpire(C_RtmEventHandlerBridge *this_, void *userData,
	char *channelName);

void cgo_RtmEventHandlerBridge_onSubscribeResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onPublishResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onLoginResult(C_RtmEventHandlerBridge *this_, void *userData,
	enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onSetChannelMetadataResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, enum C_RTM_CHANNEL_TYPE channelType, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onUpdateChannelMetadataResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, enum C_RTM_CHANNEL_TYPE channelType, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onRemoveChannelMetadataResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, enum C_RTM_CHANNEL_TYPE channelType, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onGetChannelMetadataResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, enum C_RTM_CHANNEL_TYPE channelType, C_IMetadata *data, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onSetUserMetadataResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *userId, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onUpdateUserMetadataResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *userId, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onRemoveUserMetadataResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *userId, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onGetUserMetadataResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *userId, C_IMetadata *data, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onSubscribeUserMetadataResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *userId, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onSetLockResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, enum C_RTM_CHANNEL_TYPE channelType, char *lockName, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onRemoveLockResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, enum C_RTM_CHANNEL_TYPE channelType, char *lockName, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onReleaseLockResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, enum C_RTM_CHANNEL_TYPE channelType, char *lockName, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onAcquireLockResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, enum C_RTM_CHANNEL_TYPE channelType, char *lockName, enum C_RTM_ERROR_CODE errorCode, char *errorDetails);

void cgo_RtmEventHandlerBridge_onRevokeLockResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, enum C_RTM_CHANNEL_TYPE channelType, char *lockName, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onGetLocksResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, char *channelName, enum C_RTM_CHANNEL_TYPE channelType, struct C_LockDetail *lockDetailList, size_t count, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onWhoNowResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, struct C_UserState *userStateList, size_t count, char *nextPage, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onGetOnlineUsersResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, struct C_UserState *userStateList, size_t count, char *nextPage, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onWhereNowResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, struct C_ChannelInfo *channels, size_t count, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onGetUserChannelsResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, struct C_ChannelInfo *channels, size_t count, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onPresenceSetStateResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onPresenceRemoveStateResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, enum C_RTM_ERROR_CODE errorCode);

void cgo_RtmEventHandlerBridge_onPresenceGetStateResult(C_RtmEventHandlerBridge *this_, void *userData,
	uint64_t requestId, struct C_UserState *state, enum C_RTM_ERROR_CODE errorCode);

*/
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/Lensual/agora_rtm_sdk_cgo/pkg/agora"
)

type IRtmEventHandlerBridgeHandler interface {
	OnMessageEvent(event *agora.MessageEvent)
	OnPresenceEvent(event *agora.PresenceEvent)
	OnTopicEvent(event *agora.TopicEvent)
	OnLockEvent(event *agora.LockEvent)
	OnStorageEvent(event *agora.StorageEvent)
	OnJoinResult(requestId uint64, channelName string, userId string, errorCode agora.RTM_ERROR_CODE)
	OnLeaveResult(requestId uint64, channelName string, userId string, errorCode agora.RTM_ERROR_CODE)
	OnJoinTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode agora.RTM_ERROR_CODE)
	OnLeaveTopicResult(requestId uint64, channelName string, userId string, topic string, meta string, errorCode agora.RTM_ERROR_CODE)
	OnSubscribeTopicResult(requestId uint64, channelName string, userId string, topic string, succeedUsers agora.UserList, failedUsers agora.UserList, errorCode agora.RTM_ERROR_CODE)
	OnConnectionStateChanged(channelName string, state agora.RTM_CONNECTION_STATE, reason agora.RTM_CONNECTION_CHANGE_REASON)
	OnTokenPrivilegeWillExpire(channelName string)
	OnSubscribeResult(requestId uint64, channelName string, errorCode agora.RTM_ERROR_CODE)
	OnPublishResult(requestId uint64, errorCode agora.RTM_ERROR_CODE)
	OnLoginResult(errorCode agora.RTM_ERROR_CODE)
	OnSetChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, errorCode agora.RTM_ERROR_CODE)
	OnUpdateChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, errorCode agora.RTM_ERROR_CODE)
	OnRemoveChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, errorCode agora.RTM_ERROR_CODE)
	OnGetChannelMetadataResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, data *agora.IMetadata, errorCode agora.RTM_ERROR_CODE)
	OnSetUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE)
	OnUpdateUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE)
	OnRemoveUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE)
	OnGetUserMetadataResult(requestId uint64, userId string, data *agora.IMetadata, errorCode agora.RTM_ERROR_CODE)
	OnSubscribeUserMetadataResult(requestId uint64, userId string, errorCode agora.RTM_ERROR_CODE)
	OnSetLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE)
	OnRemoveLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE)
	OnReleaseLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE)
	OnAcquireLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE, errorDetails string)
	OnRevokeLockResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockName string, errorCode agora.RTM_ERROR_CODE)
	OnGetLocksResult(requestId uint64, channelName string, channelType agora.RTM_CHANNEL_TYPE, lockDetailList *agora.LockDetail, count uint, errorCode agora.RTM_ERROR_CODE)
	OnWhoNowResult(requestId uint64, userStateList *agora.UserState, count uint, nextPage string, errorCode agora.RTM_ERROR_CODE)
	OnGetOnlineUsersResult(requestId uint64, userStateList *agora.UserState, count uint, nextPage string, errorCode agora.RTM_ERROR_CODE)
	OnWhereNowResult(requestId uint64, channels *agora.ChannelInfo, count uint, errorCode agora.RTM_ERROR_CODE)
	OnGetUserChannelsResult(requestId uint64, channels *agora.ChannelInfo, count uint, errorCode agora.RTM_ERROR_CODE)
	OnPresenceSetStateResult(requestId uint64, errorCode agora.RTM_ERROR_CODE)
	OnPresenceRemoveStateResult(requestId uint64, errorCode agora.RTM_ERROR_CODE)
	OnPresenceGetStateResult(requestId uint64, state *agora.UserState, errorCode agora.RTM_ERROR_CODE)
}

type RtmEventHandlerBridge struct {
	handler IRtmEventHandlerBridgeHandler
	cBridge *C.C_RtmEventHandlerBridge
}

func (b *RtmEventHandlerBridge) ToAgoraEventHandler() *agora.IRtmEventHandler {
	return (*agora.IRtmEventHandler)(b.cBridge)
}

func (b *RtmEventHandlerBridge) Delete() {
	C.C_RtmEventHandlerBridge_Delete(unsafe.Pointer(b.cBridge))
	b.handler = nil
	b.cBridge = nil
}

var pinner runtime.Pinner

func NewRtmEventHandlerBridge(handler IRtmEventHandlerBridgeHandler) *RtmEventHandlerBridge {
	b := RtmEventHandlerBridge{}
	userData := unsafe.Pointer(&b)
	b.cBridge = (*C.C_RtmEventHandlerBridge)(C.C_RtmEventHandlerBridge_New(
		C.C_RtmEventHandlerBridge_Callbacks{
			onMessageEvent:                C.C_RtmEventHandlerBridge_onMessageEvent(C.cgo_RtmEventHandlerBridge_onMessageEvent),
			onPresenceEvent:               C.C_RtmEventHandlerBridge_onPresenceEvent(C.cgo_RtmEventHandlerBridge_onPresenceEvent),
			onTopicEvent:                  C.C_RtmEventHandlerBridge_onTopicEvent(C.cgo_RtmEventHandlerBridge_onTopicEvent),
			onLockEvent:                   C.C_RtmEventHandlerBridge_onLockEvent(C.cgo_RtmEventHandlerBridge_onLockEvent),
			onStorageEvent:                C.C_RtmEventHandlerBridge_onStorageEvent(C.cgo_RtmEventHandlerBridge_onStorageEvent),
			onJoinResult:                  C.C_RtmEventHandlerBridge_onJoinResult(C.cgo_RtmEventHandlerBridge_onJoinResult),
			onLeaveResult:                 C.C_RtmEventHandlerBridge_onLeaveResult(C.cgo_RtmEventHandlerBridge_onLeaveResult),
			onJoinTopicResult:             C.C_RtmEventHandlerBridge_onJoinTopicResult(C.cgo_RtmEventHandlerBridge_onJoinTopicResult),
			onLeaveTopicResult:            C.C_RtmEventHandlerBridge_onLeaveTopicResult(C.cgo_RtmEventHandlerBridge_onLeaveTopicResult),
			onSubscribeTopicResult:        C.C_RtmEventHandlerBridge_onSubscribeTopicResult(C.cgo_RtmEventHandlerBridge_onSubscribeTopicResult),
			onConnectionStateChanged:      C.C_RtmEventHandlerBridge_onConnectionStateChanged(C.cgo_RtmEventHandlerBridge_onConnectionStateChanged),
			onTokenPrivilegeWillExpire:    C.C_RtmEventHandlerBridge_onTokenPrivilegeWillExpire(C.cgo_RtmEventHandlerBridge_onTokenPrivilegeWillExpire),
			onSubscribeResult:             C.C_RtmEventHandlerBridge_onSubscribeResult(C.cgo_RtmEventHandlerBridge_onSubscribeResult),
			onPublishResult:               C.C_RtmEventHandlerBridge_onPublishResult(C.cgo_RtmEventHandlerBridge_onPublishResult),
			onLoginResult:                 C.C_RtmEventHandlerBridge_onLoginResult(C.cgo_RtmEventHandlerBridge_onLoginResult),
			onSetChannelMetadataResult:    C.C_RtmEventHandlerBridge_onSetChannelMetadataResult(C.cgo_RtmEventHandlerBridge_onSetChannelMetadataResult),
			onUpdateChannelMetadataResult: C.C_RtmEventHandlerBridge_onUpdateChannelMetadataResult(C.cgo_RtmEventHandlerBridge_onUpdateChannelMetadataResult),
			onRemoveChannelMetadataResult: C.C_RtmEventHandlerBridge_onRemoveChannelMetadataResult(C.cgo_RtmEventHandlerBridge_onRemoveChannelMetadataResult),
			onGetChannelMetadataResult:    C.C_RtmEventHandlerBridge_onGetChannelMetadataResult(C.cgo_RtmEventHandlerBridge_onGetChannelMetadataResult),
			onSetUserMetadataResult:       C.C_RtmEventHandlerBridge_onSetUserMetadataResult(C.cgo_RtmEventHandlerBridge_onSetUserMetadataResult),
			onUpdateUserMetadataResult:    C.C_RtmEventHandlerBridge_onUpdateUserMetadataResult(C.cgo_RtmEventHandlerBridge_onUpdateUserMetadataResult),
			onRemoveUserMetadataResult:    C.C_RtmEventHandlerBridge_onRemoveUserMetadataResult(C.cgo_RtmEventHandlerBridge_onRemoveUserMetadataResult),
			onGetUserMetadataResult:       C.C_RtmEventHandlerBridge_onGetUserMetadataResult(C.cgo_RtmEventHandlerBridge_onGetUserMetadataResult),
			onSubscribeUserMetadataResult: C.C_RtmEventHandlerBridge_onSubscribeUserMetadataResult(C.cgo_RtmEventHandlerBridge_onSubscribeUserMetadataResult),
			onSetLockResult:               C.C_RtmEventHandlerBridge_onSetLockResult(C.cgo_RtmEventHandlerBridge_onSetLockResult),
			onRemoveLockResult:            C.C_RtmEventHandlerBridge_onRemoveLockResult(C.cgo_RtmEventHandlerBridge_onRemoveLockResult),
			onReleaseLockResult:           C.C_RtmEventHandlerBridge_onReleaseLockResult(C.cgo_RtmEventHandlerBridge_onReleaseLockResult),
			onAcquireLockResult:           C.C_RtmEventHandlerBridge_onAcquireLockResult(C.cgo_RtmEventHandlerBridge_onAcquireLockResult),
			onRevokeLockResult:            C.C_RtmEventHandlerBridge_onRevokeLockResult(C.cgo_RtmEventHandlerBridge_onRevokeLockResult),
			onGetLocksResult:              C.C_RtmEventHandlerBridge_onGetLocksResult(C.cgo_RtmEventHandlerBridge_onGetLocksResult),
			onWhoNowResult:                C.C_RtmEventHandlerBridge_onWhoNowResult(C.cgo_RtmEventHandlerBridge_onWhoNowResult),
			onGetOnlineUsersResult:        C.C_RtmEventHandlerBridge_onGetOnlineUsersResult(C.cgo_RtmEventHandlerBridge_onGetOnlineUsersResult),
			onWhereNowResult:              C.C_RtmEventHandlerBridge_onWhereNowResult(C.cgo_RtmEventHandlerBridge_onWhereNowResult),
			onGetUserChannelsResult:       C.C_RtmEventHandlerBridge_onGetUserChannelsResult(C.cgo_RtmEventHandlerBridge_onGetUserChannelsResult),
			onPresenceSetStateResult:      C.C_RtmEventHandlerBridge_onPresenceSetStateResult(C.cgo_RtmEventHandlerBridge_onPresenceSetStateResult),
			onPresenceRemoveStateResult:   C.C_RtmEventHandlerBridge_onPresenceRemoveStateResult(C.cgo_RtmEventHandlerBridge_onPresenceRemoveStateResult),
			onPresenceGetStateResult:      C.C_RtmEventHandlerBridge_onPresenceGetStateResult(C.cgo_RtmEventHandlerBridge_onPresenceGetStateResult),
		},
		userData,
	))
	b.handler = handler
	return &b
}

//export cgo_RtmEventHandlerBridge_onMessageEvent
func cgo_RtmEventHandlerBridge_onMessageEvent(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	event *C.struct_C_MessageEvent) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnMessageEvent((*agora.MessageEvent)(unsafe.Pointer(event)))
}

//export cgo_RtmEventHandlerBridge_onPresenceEvent
func cgo_RtmEventHandlerBridge_onPresenceEvent(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	event *C.struct_C_PresenceEvent) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnPresenceEvent((*agora.PresenceEvent)(unsafe.Pointer(event)))
}

//export cgo_RtmEventHandlerBridge_onTopicEvent
func cgo_RtmEventHandlerBridge_onTopicEvent(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	event *C.struct_C_TopicEvent) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnTopicEvent((*agora.TopicEvent)(unsafe.Pointer(event)))
}

//export cgo_RtmEventHandlerBridge_onLockEvent
func cgo_RtmEventHandlerBridge_onLockEvent(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	event *C.struct_C_LockEvent) {

}

//export cgo_RtmEventHandlerBridge_onStorageEvent
func cgo_RtmEventHandlerBridge_onStorageEvent(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	event *C.struct_C_StorageEvent) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnStorageEvent((*agora.StorageEvent)(unsafe.Pointer(event)))
}

//export cgo_RtmEventHandlerBridge_onJoinResult
func cgo_RtmEventHandlerBridge_onJoinResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, userId *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnJoinResult(
		uint64(requestId),
		C.GoString(channelName),
		C.GoString(userId),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onLeaveResult
func cgo_RtmEventHandlerBridge_onLeaveResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, userId *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnLeaveResult(
		uint64(requestId),
		C.GoString(channelName),
		C.GoString(userId),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onJoinTopicResult
func cgo_RtmEventHandlerBridge_onJoinTopicResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, userId *C.char, topic *C.char, meta *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnJoinTopicResult(
		uint64(requestId),
		C.GoString(channelName),
		C.GoString(userId),
		C.GoString(topic),
		C.GoString(meta),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onLeaveTopicResult
func cgo_RtmEventHandlerBridge_onLeaveTopicResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, userId *C.char, topic *C.char, meta *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnLeaveTopicResult(
		uint64(requestId),
		C.GoString(channelName),
		C.GoString(userId),
		C.GoString(topic),
		C.GoString(meta),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onSubscribeTopicResult
func cgo_RtmEventHandlerBridge_onSubscribeTopicResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, userId *C.char, topic *C.char, succeedUsers C.struct_C_UserList, failedUsers C.struct_C_UserList, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnSubscribeTopicResult(
		uint64(requestId),
		C.GoString(channelName),
		C.GoString(userId),
		C.GoString(topic),
		*(*agora.UserList)(unsafe.Pointer(&succeedUsers)),
		*(*agora.UserList)(unsafe.Pointer(&failedUsers)),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onConnectionStateChanged
func cgo_RtmEventHandlerBridge_onConnectionStateChanged(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	channelName *C.char, state C.enum_C_RTM_CONNECTION_STATE, reason C.enum_C_RTM_CONNECTION_CHANGE_REASON) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnConnectionStateChanged(
		C.GoString(channelName),
		agora.RTM_CONNECTION_STATE(state),
		agora.RTM_CONNECTION_CHANGE_REASON(reason),
	)
}

//export cgo_RtmEventHandlerBridge_onTokenPrivilegeWillExpire
func cgo_RtmEventHandlerBridge_onTokenPrivilegeWillExpire(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	channelName *C.char) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnTokenPrivilegeWillExpire(
		C.GoString(channelName),
	)
}

//export cgo_RtmEventHandlerBridge_onSubscribeResult
func cgo_RtmEventHandlerBridge_onSubscribeResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnSubscribeResult(
		uint64(requestId),
		C.GoString(channelName),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onPublishResult
func cgo_RtmEventHandlerBridge_onPublishResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnPublishResult(
		uint64(requestId),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onLoginResult
func cgo_RtmEventHandlerBridge_onLoginResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnLoginResult(
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onSetChannelMetadataResult
func cgo_RtmEventHandlerBridge_onSetChannelMetadataResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, channelType C.enum_C_RTM_CHANNEL_TYPE, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnSetChannelMetadataResult(
		uint64(requestId),
		C.GoString(channelName),
		agora.RTM_CHANNEL_TYPE(channelType),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onUpdateChannelMetadataResult
func cgo_RtmEventHandlerBridge_onUpdateChannelMetadataResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, channelType C.enum_C_RTM_CHANNEL_TYPE, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnUpdateChannelMetadataResult(
		uint64(requestId),
		C.GoString(channelName),
		agora.RTM_CHANNEL_TYPE(channelType),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onRemoveChannelMetadataResult
func cgo_RtmEventHandlerBridge_onRemoveChannelMetadataResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, channelType C.enum_C_RTM_CHANNEL_TYPE, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnRemoveChannelMetadataResult(
		uint64(requestId),
		C.GoString(channelName),
		agora.RTM_CHANNEL_TYPE(channelType),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onGetChannelMetadataResult
func cgo_RtmEventHandlerBridge_onGetChannelMetadataResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, channelType C.enum_C_RTM_CHANNEL_TYPE, data *C.C_IMetadata, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnGetChannelMetadataResult(
		uint64(requestId),
		C.GoString(channelName),
		agora.RTM_CHANNEL_TYPE(channelType),
		(*agora.IMetadata)(data),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onSetUserMetadataResult
func cgo_RtmEventHandlerBridge_onSetUserMetadataResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, userId *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnSetUserMetadataResult(
		uint64(requestId),
		C.GoString(userId),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onUpdateUserMetadataResult
func cgo_RtmEventHandlerBridge_onUpdateUserMetadataResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, userId *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnUpdateUserMetadataResult(
		uint64(requestId),
		C.GoString(userId),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onRemoveUserMetadataResult
func cgo_RtmEventHandlerBridge_onRemoveUserMetadataResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, userId *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnRemoveUserMetadataResult(
		uint64(requestId),
		C.GoString(userId),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onGetUserMetadataResult
func cgo_RtmEventHandlerBridge_onGetUserMetadataResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, userId *C.char, data *C.C_IMetadata, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnGetUserMetadataResult(
		uint64(requestId),
		C.GoString(userId),
		(*agora.IMetadata)(data),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onSubscribeUserMetadataResult
func cgo_RtmEventHandlerBridge_onSubscribeUserMetadataResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, userId *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnSubscribeUserMetadataResult(
		uint64(requestId),
		C.GoString(userId),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onSetLockResult
func cgo_RtmEventHandlerBridge_onSetLockResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, channelType C.enum_C_RTM_CHANNEL_TYPE, lockName *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnSetLockResult(
		uint64(requestId),
		C.GoString(channelName),
		agora.RTM_CHANNEL_TYPE(channelType),
		C.GoString(lockName),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onRemoveLockResult
func cgo_RtmEventHandlerBridge_onRemoveLockResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, channelType C.enum_C_RTM_CHANNEL_TYPE, lockName *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnRemoveLockResult(
		uint64(requestId),
		C.GoString(channelName),
		agora.RTM_CHANNEL_TYPE(channelType),
		C.GoString(lockName),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onReleaseLockResult
func cgo_RtmEventHandlerBridge_onReleaseLockResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, channelType C.enum_C_RTM_CHANNEL_TYPE, lockName *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnReleaseLockResult(
		uint64(requestId),
		C.GoString(channelName),
		agora.RTM_CHANNEL_TYPE(channelType),
		C.GoString(lockName),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onAcquireLockResult
func cgo_RtmEventHandlerBridge_onAcquireLockResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, channelType C.enum_C_RTM_CHANNEL_TYPE, lockName *C.char, errorCode C.enum_C_RTM_ERROR_CODE, errorDetails *C.char) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnAcquireLockResult(
		uint64(requestId),
		C.GoString(channelName),
		agora.RTM_CHANNEL_TYPE(channelType),
		C.GoString(lockName),
		agora.RTM_ERROR_CODE(errorCode),
		C.GoString(errorDetails),
	)
}

//export cgo_RtmEventHandlerBridge_onRevokeLockResult
func cgo_RtmEventHandlerBridge_onRevokeLockResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, channelType C.enum_C_RTM_CHANNEL_TYPE, lockName *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnRevokeLockResult(
		uint64(requestId),
		C.GoString(channelName),
		agora.RTM_CHANNEL_TYPE(channelType),
		C.GoString(lockName),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onGetLocksResult
func cgo_RtmEventHandlerBridge_onGetLocksResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channelName *C.char, channelType C.enum_C_RTM_CHANNEL_TYPE, lockDetailList *C.struct_C_LockDetail, count C.size_t, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnGetLocksResult(
		uint64(requestId),
		C.GoString(channelName),
		agora.RTM_CHANNEL_TYPE(channelType),
		(*agora.LockDetail)(unsafe.Pointer(lockDetailList)),
		uint(count),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onWhoNowResult
func cgo_RtmEventHandlerBridge_onWhoNowResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, userStateList *C.struct_C_UserState, count C.size_t, nextPage *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnWhoNowResult(
		uint64(requestId),
		(*agora.UserState)(unsafe.Pointer(userStateList)),
		uint(count),
		C.GoString(nextPage),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onGetOnlineUsersResult
func cgo_RtmEventHandlerBridge_onGetOnlineUsersResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, userStateList *C.struct_C_UserState, count C.size_t, nextPage *C.char, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnGetOnlineUsersResult(
		uint64(requestId),
		(*agora.UserState)(unsafe.Pointer(userStateList)),
		uint(count),
		C.GoString(nextPage),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onWhereNowResult
func cgo_RtmEventHandlerBridge_onWhereNowResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channels *C.struct_C_ChannelInfo, count C.size_t, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnWhereNowResult(
		uint64(requestId),
		(*agora.ChannelInfo)(unsafe.Pointer(channels)),
		uint(count),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onGetUserChannelsResult
func cgo_RtmEventHandlerBridge_onGetUserChannelsResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, channels *C.struct_C_ChannelInfo, count C.size_t, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnGetUserChannelsResult(
		uint64(requestId),
		(*agora.ChannelInfo)(unsafe.Pointer(channels)),
		uint(count),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onPresenceSetStateResult
func cgo_RtmEventHandlerBridge_onPresenceSetStateResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnPresenceSetStateResult(
		uint64(requestId),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onPresenceRemoveStateResult
func cgo_RtmEventHandlerBridge_onPresenceRemoveStateResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnPresenceRemoveStateResult(
		uint64(requestId),
		agora.RTM_ERROR_CODE(errorCode),
	)
}

//export cgo_RtmEventHandlerBridge_onPresenceGetStateResult
func cgo_RtmEventHandlerBridge_onPresenceGetStateResult(_ *C.C_RtmEventHandlerBridge, userData unsafe.Pointer,
	requestId C.uint64_t, state *C.struct_C_UserState, errorCode C.enum_C_RTM_ERROR_CODE) {

	if userData == nil {
		return
	}

	bridge := (*RtmEventHandlerBridge)(userData)
	bridge.handler.OnPresenceGetStateResult(
		uint64(requestId),
		(*agora.UserState)(unsafe.Pointer(state)),
		agora.RTM_ERROR_CODE(errorCode),
	)
}
