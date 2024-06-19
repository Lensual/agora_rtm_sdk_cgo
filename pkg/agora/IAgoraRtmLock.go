package agora

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtm_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtm_sdk_c -lagora_rtm_sdk_c -lstdc++

//链接AgoraRTM SDK
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtm_sdk_c/agora_rtm_sdk-prefix/src/agora_rtm_sdk/rtm/sdk/high_level_api/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtm_sdk_c/agora_rtm_sdk-prefix/src/agora_rtm_sdk/rtm/sdk -lagora_rtm_sdk

#include "C_IAgoraRtmLock.h"
#include <stdlib.h>

*/
import "C"
import "unsafe"

// #region agora

// #region agora::rtm

/**
 * The IRtmLock class.
 *
 * This class provides the rtm lock methods that can be invoked by your app.
 */
type IRtmLock C.C_IRtmLock

// #region IRtmLock

/**
 * sets a lock
 *
 * @param [in] channelName The name of the channel.
 * @param [in] channelType The type of the channel.
 * @param [in] lockName The name of the lock.
 * @param [in] ttl The lock ttl.
 * @param [out] requestId The related request id of this operation.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmLock) SetLock(channelName string, channelType RTM_CHANNEL_TYPE, lockName string, ttl uint32, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	cLockName := C.CString(lockName)
	ret := int(C.C_IRtmLock_setLock(unsafe.Pointer(this_),
		cChannelName,
		C.enum_C_RTM_CHANNEL_TYPE(channelType),
		cLockName,
		C.uint32_t(ttl),
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cChannelName))
	C.free(unsafe.Pointer(cLockName))
	return ret
}

/**
 * gets locks in the channel
 *
 * @param [in] channelName The name of the channel.
 * @param [in] channelType The type of the channel.
 * @param [out] requestId The related request id of this operation.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmLock) GetLocks(channelName string, channelType RTM_CHANNEL_TYPE, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	ret := int(C.C_IRtmLock_getLocks(unsafe.Pointer(this_),
		cChannelName,
		C.enum_C_RTM_CHANNEL_TYPE(channelType),
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cChannelName))
	return ret
}

/**
 * removes a lock
 *
 * @param [in] channelName The name of the channel.
 * @param [in] channelType The type of the channel.
 * @param [in] lockName The name of the lock.
 * @param [out] requestId The related request id of this operation.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmLock) RemoveLock(channelName string, channelType RTM_CHANNEL_TYPE, lockName string, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	cLockName := C.CString(lockName)
	ret := int(C.C_IRtmLock_removeLock(unsafe.Pointer(this_),
		cChannelName,
		C.enum_C_RTM_CHANNEL_TYPE(channelType),
		cLockName,
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cChannelName))
	C.free(unsafe.Pointer(cLockName))
	return ret
}

/**
 * acquires a lock
 *
 * @param [in] channelName The name of the channel.
 * @param [in] channelType The type of the channel.
 * @param [in] lockName The name of the lock.
 * @param [in] retry Whether to automatically retry when acquires lock failed
 * @param [out] requestId The related request id of this operation.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmLock) AcquireLock(channelName string, channelType RTM_CHANNEL_TYPE, lockName string, retry bool, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	cLockName := C.CString(lockName)
	ret := int(C.C_IRtmLock_acquireLock(unsafe.Pointer(this_),
		cChannelName,
		C.enum_C_RTM_CHANNEL_TYPE(channelType),
		cLockName,
		C.bool(retry),
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cChannelName))
	C.free(unsafe.Pointer(cLockName))
	return ret
}

/**
 * releases a lock
 *
 * @param [in] channelName The name of the channel.
 * @param [in] channelType The type of the channel.
 * @param [in] lockName The name of the lock.
 * @param [out] requestId The related request id of this operation.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmLock) ReleaseLock(channelName string, channelType RTM_CHANNEL_TYPE, lockName string, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	cLockName := C.CString(lockName)
	ret := int(C.C_IRtmLock_releaseLock(unsafe.Pointer(this_),
		cChannelName,
		C.enum_C_RTM_CHANNEL_TYPE(channelType),
		cLockName,
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cChannelName))
	C.free(unsafe.Pointer(cLockName))
	return ret
}

/**
 * disables a lock
 *
 * @param [in] channelName The name of the channel.
 * @param [in] channelType The type of the channel.
 * @param [in] lockName The name of the lock.
 * @param [in] owner The lock owner.
 * @param [out] requestId The related request id of this operation.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmLock) RevokeLock(channelName string, channelType RTM_CHANNEL_TYPE, lockName string, owner string, requestId *uint64) int {
	cChannelName := C.CString(channelName)
	cLockName := C.CString(lockName)
	cOwner := C.CString(owner)
	ret := int(C.C_IRtmLock_revokeLock(unsafe.Pointer(this_),
		cChannelName,
		C.enum_C_RTM_CHANNEL_TYPE(channelType),
		cLockName,
		cOwner,
		(*C.uint64_t)(requestId),
	))
	C.free(unsafe.Pointer(cChannelName))
	C.free(unsafe.Pointer(cLockName))
	C.free(unsafe.Pointer(cOwner))
	return ret
}

// #endregion IRtmLock

// #endregion agora::rtm

// #endregion agora
