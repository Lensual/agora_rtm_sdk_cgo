package agora

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../../agora_rtm_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../../agora_rtm_sdk_c -lagora_rtm_sdk_c -lstdc++

//链接AgoraRTM SDK
#cgo CFLAGS: -I${SRCDIR}/../../../agora_rtm_sdk/libs/include
#cgo LDFLAGS: -L${SRCDIR}/../../../agora_rtm_sdk/libs -lagora_rtm_sdk

#include "C_IAgoraRtmStorage.h"
#include <stdlib.h>

*/
import "C"
import "unsafe"

// #region agora
// #region agora::rtm

/**
 * Metadata options.
 */
type MetadataOptions C.struct_C_MetadataOptions

// #region MetadataOptions

/**
 * Indicates whether or not to notify server update the modify timestamp of metadata
 */
func (this_ *MetadataOptions) GetRecordTs() bool {
	return bool(this_.recordTs)
}

/**
 * Indicates whether or not to notify server update the modify timestamp of metadata
 */
func (this_ *MetadataOptions) SetRecordTs(recordTs bool) {
	this_.recordTs = C.bool(recordTs)
}

/**
 * Indicates whether or not to notify server update the modify user id of metadata
 */
func (this_ *MetadataOptions) GetRecordUserId() bool {
	return bool(this_.recordUserId)
}

/**
 * Indicates whether or not to notify server update the modify user id of metadata
 */
func (this_ *MetadataOptions) SetRecordUserId(recordUserId bool) {
	this_.recordUserId = C.bool(recordUserId)
}

func NewMetadataOptions() *MetadataOptions {
	return (*MetadataOptions)(C.C_MetadataOptions_New())
}
func (this_ *MetadataOptions) Delete() {
	C.C_MetadataOptions_Delete((*C.struct_C_MetadataOptions)(this_))
}

// #endregion MetadataOptions

type MetadataItem C.struct_C_MetadataItem

// #region MetadataItem

/**
 * The key of the metadata item.
 */
func (this_ *MetadataItem) GetKey() string {
	return C.GoString(this_.key)
}

/**
 * The key of the metadata item.
 */
func (this_ *MetadataItem) SetKey(key string) {
	this_.key = C.CString(key)
}

/**
 * The value of the metadata item.
 */
func (this_ *MetadataItem) GetValue() string {
	return C.GoString(this_.value)
}

/**
 * The value of the metadata item.
 */
func (this_ *MetadataItem) SetValue(value string) {
	this_.value = C.CString(value)
}

/**
 * The User ID of the user who makes the latest update to the metadata item.
 */
func (this_ *MetadataItem) GetAquthorUserId() string {
	return C.GoString(this_.authorUserId)
}

/**
 * The User ID of the user who makes the latest update to the metadata item.
 */
func (this_ *MetadataItem) SetAuthorUserId(authorUserId string) {
	this_.authorUserId = C.CString(authorUserId)
}

/**
 * The revision of the metadata item.
 */
func (this_ *MetadataItem) GetRevision() int64 {
	return int64(this_.revision)
}

/**
 * The revision of the metadata item.
 */
func (this_ *MetadataItem) SetRevision(revision int64) {
	this_.revision = C.int64_t(revision)
}

/**
 * The Timestamp when the metadata item was last updated.
 */
func (this_ *MetadataItem) GetUpdateTs() int64 {
	return int64(this_.updateTs)
}

/**
 * The Timestamp when the metadata item was last updated.
 */
func (this_ *MetadataItem) SetUpdateTs(updateTs int64) {
	this_.updateTs = C.int64_t(updateTs)
}

func NewMetadataItem() *MetadataItem {
	return (*MetadataItem)(C.C_MetadataItem_New())
}
func (this_ *MetadataItem) Delete() {
	C.C_MetadataItem_Delete((*C.struct_C_MetadataItem)(this_))
}

// #endregion MetadataItem

type IMetadata C.C_IMetadata

// #region IMetadata

/**
 * Set the major revision of metadata.
 *
 * @param [in] revision The major revision of the metadata.
 */
func (this_ *IMetadata) SetMajorRevision(revision int64) {
	C.C_IMetadata_setMajorRevision(unsafe.Pointer(this_), C.int64_t(revision))
}

/**
 * Get the major revision of metadata.
 *
 * @return the major revision of metadata.
 */
func (this_ *IMetadata) GetMajorRevision() int64 {
	return int64(C.C_IMetadata_getMajorRevision(unsafe.Pointer(this_)))
}

/**
 * Add or revise a metadataItem to current metadata.
 */
func (this_ *IMetadata) SetMetadataItem(item *MetadataItem) {
	C.C_IMetadata_setMetadataItem(unsafe.Pointer(this_), (*C.struct_C_MetadataItem)(item))
}

/**
 * Get the metadataItem array of current metadata.
 *
 * @param [out] items The address of the metadataItem array.
 * @param [out] size The size the metadataItem array.
 */
func (this_ *IMetadata) GetMetadataItems(item **MetadataItem, size *uint) {
	C.C_IMetadata_getMetadataItems(unsafe.Pointer(this_), (**C.struct_C_MetadataItem)(unsafe.Pointer(item)), (*C.size_t)(unsafe.Pointer(size)))
}

/**
 * Clear the metadataItem array & reset major revision
 */
func (this_ *IMetadata) ClearMetadata() {
	C.C_IMetadata_clearMetadata(unsafe.Pointer(this_))
}

/**
 * Release the metadata instance.
 */
func (this_ *IMetadata) Release() {
	C.C_IMetadata_release(unsafe.Pointer(this_))
}

// #endregion IMetadata

type IRtmStorage C.C_IRtmStorage

// #region IRtmStorage

/** Creates the metadata object and returns the pointer.
 * @return Pointer of the metadata object.
 */
func (this_ *IRtmStorage) CreateMetadata() *IMetadata {
	return (*IMetadata)(C.C_IRtmStorage_createMetadata(unsafe.Pointer(this_)))
}

/**
 * Set the metadata of a specified channel.
 *
 * @param [in] channelName The name of the channel.
 * @param [in] channelType Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE.
 * @param [in] data Metadata data.
 * @param [in] options The options of operate metadata.
 * @param [in] lock lock for operate channel metadata.
 * @param [out] requestId The unique ID of this request.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmStorage) SetChannelMetadata(channelName string, channelType RTM_CHANNEL_TYPE, data *IMetadata, options *MetadataOptions, lockName string, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	cLockName := C.CString(lockName)
	ret := int(C.C_IRtmStorage_setChannelMetadata(unsafe.Pointer(this_),
		cChannelName,
		C.enum_C_RTM_CHANNEL_TYPE(channelType),
		unsafe.Pointer(data),
		(*C.struct_C_MetadataOptions)(options),
		cLockName,
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cChannelName))
	C.free(unsafe.Pointer(cLockName))
	return ret
}

/**
 * Update the metadata of a specified channel.
 *
 * @param [in] channelName The channel Name of the specified channel.
 * @param [in] channelType Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE.
 * @param [in] data Metadata data.
 * @param [in] options The options of operate metadata.
 * @param [in] lock lock for operate channel metadata.
 * @param [out] requestId The unique ID of this request.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmStorage) UpdateChannelMetadata(channelName string, channelType RTM_CHANNEL_TYPE, data *IMetadata, options *MetadataOptions, lockName string, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	cLockName := C.CString(lockName)
	ret := int(C.C_IRtmStorage_updateChannelMetadata(unsafe.Pointer(this_),
		cChannelName,
		C.enum_C_RTM_CHANNEL_TYPE(channelType),
		unsafe.Pointer(data),
		(*C.struct_C_MetadataOptions)(options),
		cLockName,
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cChannelName))
	C.free(unsafe.Pointer(cLockName))
	return ret
}

/**
 * Remove the metadata of a specified channel.
 *
 * @param [in] channelName The channel Name of the specified channel.
 * @param [in] channelType Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE.
 * @param [in] data Metadata data.
 * @param [in] options The options of operate metadata.
 * @param [in] lock lock for operate channel metadata.
 * @param [out] requestId The unique ID of this request.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmStorage) RemoveChannelMetadata(channelName string, channelType RTM_CHANNEL_TYPE, data *IMetadata, options *MetadataOptions, lockName string, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	cLockName := C.CString(lockName)
	ret := int(C.C_IRtmStorage_removeChannelMetadata(unsafe.Pointer(this_),
		cChannelName,
		C.enum_C_RTM_CHANNEL_TYPE(channelType),
		unsafe.Pointer(data),
		(*C.struct_C_MetadataOptions)(options),
		cLockName,
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cChannelName))
	C.free(unsafe.Pointer(cLockName))
	return ret
}

/**
 * Get the metadata of a specified channel.
 *
 * @param [in] channelName The channel Name of the specified channel.
 * @param [in] channelType Which channel type, RTM_CHANNEL_TYPE_STREAM or RTM_CHANNEL_TYPE_MESSAGE.
 * @param requestId The unique ID of this request.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmStorage) GetChannelMetadata(channelName string, channelType RTM_CHANNEL_TYPE, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	ret := int(C.C_IRtmStorage_getChannelMetadata(unsafe.Pointer(this_),
		cChannelName,
		C.enum_C_RTM_CHANNEL_TYPE(channelType),
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cChannelName))
	return ret
}

/**
 * Set the metadata of a specified user.
 *
 * @param [in] userId The user ID of the specified user.
 * @param [in] data Metadata data.
 * @param [in] options The options of operate metadata.
 * @param [out] requestId The unique ID of this request.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmStorage) SetUserMetadata(userId string, data *IMetadata, options *MetadataOptions, requestId *uint64) int {
	cUserId := C.CString(userId)
	ret := int(C.C_IRtmStorage_setUserMetadata(unsafe.Pointer(this_),
		cUserId,
		unsafe.Pointer(data),
		(*C.struct_C_MetadataOptions)(options),
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

/**
 * Update the metadata of a specified user.
 *
 * @param [in] userId The user ID of the specified user.
 * @param [in] data Metadata data.
 * @param [in] options The options of operate metadata.
 * @param [out] requestId The unique ID of this request.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmStorage) UpdateUserMetadata(userId string, data *IMetadata, options *MetadataOptions, requestId *uint64) int {
	cUserId := C.CString(userId)
	ret := int(C.C_IRtmStorage_updateUserMetadata(unsafe.Pointer(this_),
		cUserId,
		unsafe.Pointer(data),
		(*C.struct_C_MetadataOptions)(options),
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

/**
 * Remove the metadata of a specified user.
 *
 * @param [in] userId The user ID of the specified user.
 * @param [in] data Metadata data.
 * @param [in] options The options of operate metadata.
 * @param [out] requestId The unique ID of this request.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmStorage) RemoveUserMetadata(userId string, data *IMetadata, options *MetadataOptions, requestId *uint64) int {
	cUserId := C.CString(userId)
	ret := int(C.C_IRtmStorage_removeUserMetadata(unsafe.Pointer(this_),
		cUserId,
		unsafe.Pointer(data),
		(*C.struct_C_MetadataOptions)(options),
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

/**
 * Get the metadata of a specified user.
 *
 * @param [in] userId The user ID of the specified user.
 * @param [out] requestId The unique ID of this request.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmStorage) GetUserMetadata(userId string, requestId *uint64) int {
	cUserId := C.CString(userId)
	ret := int(C.C_IRtmStorage_getUserMetadata(unsafe.Pointer(this_),
		cUserId,
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

/**
 * Subscribe the metadata update event of a specified user.
 *
 * @param [in] userId The user ID of the specified user.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmStorage) SubscribeUserMetadata(userId string, requestId *uint64) int {
	cUserId := C.CString(userId)
	ret := int(C.C_IRtmStorage_subscribeUserMetadata(unsafe.Pointer(this_),
		cUserId,
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

/**
 * unsubscribe the metadata update event of a specified user.
 *
 * @param [in] userId The user ID of the specified user.
 *
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmStorage) UnsubscribeUserMetadata(userId string) int {
	cUserId := C.CString(userId)
	ret := int(C.C_IRtmStorage_unsubscribeUserMetadata(unsafe.Pointer(this_),
		cUserId,
	))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

// #endregion IRtmStorage

// #endregion agora::rtm

// #endregion agora
