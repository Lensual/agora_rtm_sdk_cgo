package agora

/*
//引入Agora C封装
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtm_sdk_c/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtm_sdk_c -lagora_rtm_sdk_c -lstdc++

//链接AgoraRTM SDK
#cgo CFLAGS: -I${SRCDIR}/../../third_party/agora_rtm_sdk_c/agora_rtm_sdk-prefix/src/agora_rtm_sdk/rtm/sdk/high_level_api/include
#cgo LDFLAGS: -L${SRCDIR}/../../third_party/agora_rtm_sdk_c/agora_rtm_sdk-prefix/src/agora_rtm_sdk/rtm/sdk -lagora_rtm_sdk

#include "C_IAgoraRtmService.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

// #region agora

// #region agora::rtm

type PEER_MESSAGE_STATE C.enum_C_PEER_MESSAGE_STATE

const (
	PEER_MESSAGE_INIT             PEER_MESSAGE_STATE = C.PEER_MESSAGE_INIT
	PEER_MESSAGE_FAILURE          PEER_MESSAGE_STATE = C.PEER_MESSAGE_FAILURE
	PEER_MESSAGE_PEER_UNREACHABLE PEER_MESSAGE_STATE = C.PEER_MESSAGE_PEER_UNREACHABLE
	PEER_MESSAGE_RECEIVED_BY_PEER PEER_MESSAGE_STATE = C.PEER_MESSAGE_RECEIVED_BY_PEER
	PEER_MESSAGE_SENT_TIMEOUT     PEER_MESSAGE_STATE = C.PEER_MESSAGE_SENT_TIMEOUT
)

/**
 * The login error code.
 */
type LOGIN_ERR_CODE C.enum_C_LOGIN_ERR_CODE

const (

	/**
	 * 0: Login succeeds. No error occurs.
	 */
	LOGIN_ERR_OK LOGIN_ERR_CODE = C.LOGIN_ERR_OK

	/**
	 * 1: Login fails for reasons unknown.
	 */
	LOGIN_ERR_UNKNOWN LOGIN_ERR_CODE = C.LOGIN_ERR_UNKNOWN

	/**
	 * 2: The server rejects the login, either because the user has already logged in, or because
	 * the RTM service is not initialized.
	 */
	LOGIN_ERR_REJECTED LOGIN_ERR_CODE = C.LOGIN_ERR_REJECTED // ALREADY LOGIN OR NOT INITIALIZED, SERVER REJECT

	/**
	 * 3: Invalid login arguments.
	 */
	LOGIN_ERR_INVALID_ARGUMENT LOGIN_ERR_CODE = C.LOGIN_ERR_INVALID_ARGUMENT

	/**
	 * 4: The App ID is invalid.
	 */
	LOGIN_ERR_INVALID_APP_ID LOGIN_ERR_CODE = C.LOGIN_ERR_INVALID_APP_ID

	/**
	 * 5: The token is invalid.
	 */
	LOGIN_ERR_INVALID_TOKEN LOGIN_ERR_CODE = C.LOGIN_ERR_INVALID_TOKEN

	/**
	 * 6: The login is rejected because the token has expired.
	 */
	LOGIN_ERR_TOKEN_EXPIRED LOGIN_ERR_CODE = C.LOGIN_ERR_TOKEN_EXPIRED

	/**
	 * 7: Authentication of the RTMP token fails.
	 */
	LOGIN_ERR_NOT_AUTHORIZED LOGIN_ERR_CODE = C.LOGIN_ERR_NOT_AUTHORIZED

	/**
	 * 8: The login times out. The current timeout is set as six seconds.
	 */
	LOGIN_ERR_TIMEOUT LOGIN_ERR_CODE = C.LOGIN_ERR_TIMEOUT
)

/**
 * The logout error code.
 */
type LOGOUT_ERR_CODE C.enum_C_LOGOUT_ERR_CODE

const (

	/**
	 * 0: Logout succeeds. No error occurs.
	 */
	LOGOUT_ERR_OK LOGOUT_ERR_CODE = C.LOGOUT_ERR_OK

	/**
	 * 1: Logout fails.
	 */
	LOGOUT_ERR_REJECTED LOGOUT_ERR_CODE = C.LOGOUT_ERR_REJECTED
)

/**
 * The connection state.
 */
type CONNECTION_STATE C.enum_C_CONNECTION_STATE

const (
	/**
	 * 1: The SDK has logged in the RTM service.
	 */
	CONNECTION_STATE_CONNECTED CONNECTION_STATE = C.CONNECTION_STATE_CONNECTED

	/**
	 * 2: The initial state. The SDK is disconnected from the RTM service.
	 */
	CONNECTION_STATE_DISCONNECTED CONNECTION_STATE = C.CONNECTION_STATE_DISCONNECTED

	/**
	 * 3: The SDK gives up logging in the RTM service, mainly because another instance has logged in the RTM
	 * service with the same user ID.
	 *
	 * Call the logout() method before calling login to log in the RTM service again.
	 */
	CONNECTION_STATE_ABORTED CONNECTION_STATE = C.CONNECTION_STATE_ABORTED
)

/**
 * The state of the channel message.
 */
type CHANNEL_MESSAGE_STATE C.enum_C_CHANNEL_MESSAGE_STATE

const (
	/**
	 * 1: The channel message is received by the server.
	 */
	CHANNEL_MESSAGE_RECEIVED_BY_SERVER CHANNEL_MESSAGE_STATE = C.CHANNEL_MESSAGE_RECEIVED_BY_SERVER
	/**
	 * 3: The SDK has not received a response from the server in five seconds. The current timeout is set as
	 * five seconds.
	 */
	CHANNEL_MESSAGE_SENT_TIMEOUT CHANNEL_MESSAGE_STATE = C.CHANNEL_MESSAGE_SENT_TIMEOUT
)

/**
 * The join channel error.
 */
type JOIN_CHANNEL_ERR C.enum_C_JOIN_CHANNEL_ERR

const (
	/**
	  0: The method call succeeds, or the user joins the channel successfully.
	*/
	JOIN_CHANNEL_ERR_OK JOIN_CHANNEL_ERR = C.JOIN_CHANNEL_ERR_OK

	/**
	  1: Common failure. The user fails to join the channel.
	*/
	JOIN_CHANNEL_ERR_FAILURE JOIN_CHANNEL_ERR = C.JOIN_CHANNEL_ERR_FAILURE

	/**
	  2: **RESERVED FOR FUTURE USE**
	*/
	JOIN_CHANNEL_ERR_REJECTED JOIN_CHANNEL_ERR = C.JOIN_CHANNEL_ERR_REJECTED // Usually occurs when the user is already in the channel

	/**
	  3: The user fails to join the channel because the argument is invalid.
	*/
	JOIN_CHANNEL_ERR_INVALID_ARGUMENT JOIN_CHANNEL_ERR = C.JOIN_CHANNEL_ERR_INVALID_ARGUMENT

	/**
	  4: A timeout occurs when joining the channel. The current timeout is set as five seconds. Possible reasons: The user is in the \ref agora::rtm::CONNECTION_STATE_ABORTED "CONNECTION_STATE_ABORTED" state.
	*/
	JOIN_CHANNEL_TIMEOUT JOIN_CHANNEL_ERR = C.JOIN_CHANNEL_TIMEOUT

	/**
	  5: The number of the RTM channels you are in exceeds the limit of 20.
	*/
	JOIN_CHANNEL_ERR_EXCEED_LIMIT JOIN_CHANNEL_ERR = C.JOIN_CHANNEL_ERR_EXCEED_LIMIT

	/**
	  6: The user is joining or has joined the channel.
	*/
	JOIN_CHANNEL_ERR_ALREADY_JOINED JOIN_CHANNEL_ERR = C.JOIN_CHANNEL_ERR_ALREADY_JOINED

	/**
	  7: The method call frequency exceeds 50 queries every three seconds.
	*/
	JOIN_CHANNEL_ERR_TOO_OFTEN JOIN_CHANNEL_ERR = C.JOIN_CHANNEL_ERR_TOO_OFTEN

	/**
	  8: The frequency of joining the same channel exceeds two times every five seconds.
	*/
	JOIN_CHANNEL_ERR_JOIN_SAME_CHANNEL_TOO_OFTEN JOIN_CHANNEL_ERR = C.JOIN_CHANNEL_ERR_JOIN_SAME_CHANNEL_TOO_OFTEN

	/**
	  101: \ref agora::rtm::IRtmService "IRtmService" is not initialized.
	*/
	JOIN_CHANNEL_ERR_NOT_INITIALIZED JOIN_CHANNEL_ERR = C.JOIN_CHANNEL_ERR_NOT_INITIALIZED

	/**
	  102: The user does not call the \ref agora::rtm::IRtmService::login "login" method, or the method call of \ref agora::rtm::IRtmService::login "login" does not succeed before joining the channel.
	*/
	JOIN_CHANNEL_ERR_USER_NOT_LOGGED_IN JOIN_CHANNEL_ERR = C.JOIN_CHANNEL_ERR_USER_NOT_LOGGED_IN
)

/*
*

	@brief Error codes related to leaving a channel.
*/
type LEAVE_CHANNEL_ERR C.enum_C_LEAVE_CHANNEL_ERR

const (

	/**
	  0: The method call succeeds, or the user leaves the channel successfully.
	*/
	LEAVE_CHANNEL_ERR_OK LEAVE_CHANNEL_ERR = C.LEAVE_CHANNEL_ERR_OK

	/**
	  1: Common failure. The user fails to leave the channel.
	*/
	LEAVE_CHANNEL_ERR_FAILURE LEAVE_CHANNEL_ERR = C.LEAVE_CHANNEL_ERR_OK

	/**
	  2: **RESERVED FOR FUTURE USE**
	*/
	LEAVE_CHANNEL_ERR_REJECTED LEAVE_CHANNEL_ERR = C.LEAVE_CHANNEL_ERR_REJECTED

	/**
	  3: The user is not in the channel.
	*/
	LEAVE_CHANNEL_ERR_NOT_IN_CHANNEL LEAVE_CHANNEL_ERR = C.LEAVE_CHANNEL_ERR_NOT_IN_CHANNEL

	/**
	  101: \ref agora::rtm::IRtmService "IRtmService" is not initialized.
	*/
	LEAVE_CHANNEL_ERR_NOT_INITIALIZED LEAVE_CHANNEL_ERR = C.LEAVE_CHANNEL_ERR_NOT_INITIALIZED

	/**
	  102: The user does not call the \ref agora::rtm::IRtmService::login "login" method, or the method call of \ref agora::rtm::IRtmService::login "login" does not succeed before calling the \ref agora::rtm::IChannel::leave "leave" method.
	*/
	LEAVE_CHANNEL_ERR_USER_NOT_LOGGED_IN LEAVE_CHANNEL_ERR = C.LEAVE_CHANNEL_ERR_USER_NOT_LOGGED_IN
)

/**
 * The reason for a user to leave the channel.
 */
type LEAVE_CHANNEL_REASON C.enum_C_LEAVE_CHANNEL_REASON

const (

	/**
	 * 1: The user quits the channel.
	 */
	LEAVE_CHANNEL_REASON_QUIT LEAVE_CHANNEL_REASON = C.LEAVE_CHANNEL_REASON_QUIT

	/**
	 * 2: The user is kicked off the channel.
	 */
	LEAVE_CHANNEL_REASON_KICKED LEAVE_CHANNEL_REASON = C.LEAVE_CHANNEL_REASON_KICKED
)

/*
*

	@brief Error codes related to sending a channel message.
*/
type CHANNEL_MESSAGE_ERR_CODE C.enum_C_CHANNEL_MESSAGE_ERR_CODE

const (

	/**
	  0: The method call succeeds, or the server receives the channel message.
	*/
	CHANNEL_MESSAGE_ERR_OK CHANNEL_MESSAGE_ERR_CODE = C.CHANNEL_MESSAGE_ERR_OK

	/**
	  1: Common failure. The user fails to send the channel message.
	*/
	CHANNEL_MESSAGE_ERR_FAILURE CHANNEL_MESSAGE_ERR_CODE = C.CHANNEL_MESSAGE_ERR_FAILURE

	/**
	  2: The SDK does not receive a response from the server in 10 seconds. The current timeout is set as 10 seconds. Possible reasons: The user is in the \ref agora::rtm::CONNECTION_STATE_ABORTED "CONNECTION_STATE_ABORTED" state.
	*/
	CHANNEL_MESSAGE_ERR_SENT_TIMEOUT CHANNEL_MESSAGE_ERR_CODE = C.CHANNEL_MESSAGE_ERR_SENT_TIMEOUT

	/**
	  3: The method call frequency exceeds the limit of (RTM SDK for Windows C++) 180 calls every three seconds or (RTM SDK for Linux C++) 1500 calls every three seconds, with channel and peer messages taken together..
	*/
	CHANNEL_MESSAGE_ERR_TOO_OFTEN CHANNEL_MESSAGE_ERR_CODE = C.CHANNEL_MESSAGE_ERR_TOO_OFTEN

	/**
	  4: The message is null or exceeds 32 KB in length.
	*/
	CHANNEL_MESSAGE_ERR_INVALID_MESSAGE CHANNEL_MESSAGE_ERR_CODE = C.CHANNEL_MESSAGE_ERR_INVALID_MESSAGE

	/**
	  101: \ref agora::rtm::IRtmService "IRtmService" is not initialized.
	*/
	CHANNEL_MESSAGE_ERR_NOT_INITIALIZED CHANNEL_MESSAGE_ERR_CODE = C.CHANNEL_MESSAGE_ERR_NOT_INITIALIZED

	/**
	  102: The user does not call the \ref agora::rtm::IRtmService::login "login" method, or the method call of \ref agora::rtm::IRtmService::login "login" does not succeed before sending out a channel message.
	*/
	CHANNEL_MESSAGE_ERR_USER_NOT_LOGGED_IN CHANNEL_MESSAGE_ERR_CODE = C.CHANNEL_MESSAGE_ERR_USER_NOT_LOGGED_IN
)

/**
 * The response code.
 */
type RESPONSE_CODE C.enum_C_RESPONSE_CODE

const (
	/**
	 * 1: The response
	 */
	RESPONSE_CODE_SUCCESS RESPONSE_CODE = C.RESPONSE_CODE_SUCCESS
	RESPONSE_CODE_FAILURE RESPONSE_CODE = C.RESPONSE_CODE_FAILURE
)

/**
 * The message type.
 */
type MESSAGE_TYPE C.enum_C_MESSAGE_TYPE

const (

	/**
	 * 0: The message type is undefined.
	 */
	MESSAGE_TYPE_UNDEFINED MESSAGE_TYPE = C.MESSAGE_TYPE_UNDEFINED

	/**
	 * 1: A text message.
	 */
	MESSAGE_TYPE_TEXT MESSAGE_TYPE = C.MESSAGE_TYPE_TEXT

	/**
	 * 2: A raw message in binary, for example, audio data, or video data.
	 */
	MESSAGE_TYPE_RAW MESSAGE_TYPE = C.MESSAGE_TYPE_BINARY

	/**
	 * 4: A converge message.
	 */
	MESSAGE_TYPE_IMAGE MESSAGE_TYPE = C.MESSAGE_TYPE_CONVERGE
)

type IMessage C.C_IMessage

// #region IMessage

// func CreateMessage() *IMessage {
// 	return (*IMessage)(C.C_IMessage_createMessage())
// }
//TODO

func (this_ *IMessage) Delete() {
	C.C_IMessage_Delete(unsafe.Pointer(this_))
}

/**
 * Gets the ID of the message.
 * @return The ID of the current IMessage instance.
 */
func (this_ *IMessage) GetMessageId() int64 {
	return int64(C.C_IMessage_getMessageId(unsafe.Pointer(this_)))
}

/**
 * Gets the message type.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IMessage) GetMessageType() int {
	return int(C.C_IMessage_getMessageType(unsafe.Pointer(this_)))
}

/**
 * Sets the content of the text message, or the text description of the raw message.
 * @param str The content of the text message, or the text description of the raw message. The maximum length
 * is 32 KB.
 */
func (this_ *IMessage) SetText(str string) {
	cStr := C.CString(str)
	C.C_IMessage_setText(unsafe.Pointer(this_),
		cStr,
	)
	C.free(unsafe.Pointer(cStr))
}

/**
 * Gets the content of the text messsage, or the text description of the raw message.
 * @return
 * - The content of the text message or the text description of the raw message.
 */
func (this_ *IMessage) GetText() string {
	return C.GoString(C.C_IMessage_getText(unsafe.Pointer(this_)))
}

/**
 * Get pointer of custom raw message
 * @return
 * - The content of binary raw message
 */
func (this_ *IMessage) GetRawMessageData() []byte {
	return C.GoBytes(
		unsafe.Pointer(C.C_IMessage_getRawMessageData(unsafe.Pointer(this_))),
		C.int(this_.GetRawMessageLength()),
	)
}

/**
 * Get length of custom raw message
 * @return
 * - The length of binary raw message in bytes
 */
func (this_ *IMessage) GetRawMessageLength() int {
	return int(C.C_IMessage_getRawMessageLength(unsafe.Pointer(this_)))
}

/**
 * Set message type
 */
func (this_ *IMessage) SetMessageType(_type int32) {
	C.C_IMessage_setMessageType(unsafe.Pointer(this_), C.int32_t(_type))
}

/**
 * Set raw binary message
 */
func (this_ *IMessage) SetRawMessage(data []byte, length int) {
	C.C_IMessage_setRawMessage(unsafe.Pointer(this_), (*C.uchar)(C.CBytes(data)), C.int(length))
}

/**
 * Releases the IMessage instance.
 */
func (this_ *IMessage) Release() {
	C.C_IMessage_release(unsafe.Pointer(this_))
}

// #endregion IMessage

/**
 * The IChannelMember class.
 */
type IChannelMember C.C_IChannelMember

// #region IChannelMember

func (this_ *IChannelMember) Delete() {
	C.C_IChannelMember_Delete(unsafe.Pointer(this_))
}

/**
 * Gets the ID of the channel member.
 * @return The ID of the channel member.
 */
func (this_ *IChannelMember) GetMemberId() string {
	return C.GoString(C.C_IChannelMember_getMemberId(unsafe.Pointer(this_)))
}

/**
 * Gets the ID of the channel.
 * @return The ID of the channel.
 */
func (this_ *IChannelMember) GetChannelId() string {
	return C.GoString(C.C_IChannelMember_getChannelId(unsafe.Pointer(this_)))
}

/**
 * Releases the IChannelMember class.
 */
func (this_ *IChannelMember) Release() {
	C.C_IChannelMember_release(unsafe.Pointer(this_))
}

// #endregion IChannelMember

/**
 * The IChannelAttributes class.
 */
type IChannelAttributes C.C_IChannelAttributes

// #region IChannelAttributes

/**
 * Creates an IChannelAttributes instance.
 * @return
 * - An IChannelAttributes instance, if the method call succeeds.
 */
// func CreateChannelAttributes() *IChannelAttributes {
// 	return (*IChannelAttributes)(C.C_IChannelAttributes_createChannelAttributes())
// }
//TODO

/**
 * Adds an attribute to a specified channel.
 *
 * @param key The pointer to the attribute key.
 * @param value The pointer to the attribute value.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IChannelAttributes) AddAttribute(key string, value string) int {
	cKey := C.CString(key)
	cValue := C.CString(value)
	ret := int(C.C_IChannelAttributes_addAttribute(unsafe.Pointer(this_), cKey, cValue))
	C.free(unsafe.Pointer(cKey))
	C.free(unsafe.Pointer(cValue))
	return ret
}

/**
 * Removes an attribute from the channel.
 * @param key The pointer to the attribute key.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IChannelAttributes) RemoveAttribute(key string) int {
	cKey := C.CString(key)
	ret := int(C.C_IChannelAttributes_removeAttribute(unsafe.Pointer(this_), cKey))
	C.free(unsafe.Pointer(cKey))
	return ret
}

/**
 * Gets the size of the attributes.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IChannelAttributes) GetAttributesSize() int {
	return int(C.C_IChannelAttributes_getAttributesSize(unsafe.Pointer(this_)))
}

/**
 * Gets the channel attributes.
 * @param size The size of the channel attributes.
 * @param key The pointer to the key of each channel attribute.
 * @param value The pointer to the value of each channel attribute.
 */
func (this_ *IChannelAttributes) GetAttributes(size int, key []string, value []string) {
	// C.C_IChannelAttributes_getAttributes(unsafe.Pointer(this_),
	// 	size,
	// 	key,
	// 	value,
	// )
	panic("not implemented") //TODO
}

/**
 * Gets the value of a channel attribute using the attribute key.
 * @param key The pointer to the key of the channel attribute that you want to get.
 */
func (this_ *IChannelAttributes) GetAttributeValue(key string) string {
	cKey := C.CString(key)
	ret := C.GoString(C.C_IChannelAttributes_getAttributeValue(unsafe.Pointer(this_), cKey))
	C.free(unsafe.Pointer(cKey))
	return ret
}

/**
 * Releases the IChannelAttributes instance.
 * @param
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IChannelAttributes) Release() int {
	return int(C.C_IChannelAttributes_release(unsafe.Pointer(this_)))
}

// #endregion IChannelAttributes

/**
 * The IChannelEventHandler class.
 */
type IChannelEventHandler C.C_IChannelEventHandler

// #region IChannelEventHandler

func (this_ *IChannelEventHandler) Delete() {
	C.C_IChannelEventHandler_Delete(unsafe.Pointer(this_))
}

/**
 * Occurs when the local user successfully joins a channel.
 */
func (this_ *IChannelEventHandler) OnJoinSuccess() {
	C.C_IChannelEventHandler_onJoinSuccess(unsafe.Pointer(this_))
}

/**
 * Occurs when the local user fails to join a channel.
 * @param errorCode The error code: #JOIN_CHANNEL_ERR.
 */
func (this_ *IChannelEventHandler) OnJoinFailure(errorCode JOIN_CHANNEL_ERR) {
	C.C_IChannelEventHandler_onJoinFailure(unsafe.Pointer(this_), C.enum_C_JOIN_CHANNEL_ERR(errorCode))
}

/**
 * Occurs when the local user leaves a channel.
 * @param errorCode The error code. See #LEAVE_CHANNEL_ERR.
 */
func (this_ *IChannelEventHandler) OnLeave(errorCode LEAVE_CHANNEL_ERR) {
	C.C_IChannelEventHandler_onLeave(unsafe.Pointer(this_), C.enum_C_LEAVE_CHANNEL_ERR(errorCode))
}

/**
 * Occurs when the local user receives a channel message.
 * @param message The pointer to the messsage: IMessage.
 */
func (this_ *IChannelEventHandler) OnMessageReceived(userId string, message *IMessage) {
	cUserId := C.CString(userId)
	C.C_IChannelEventHandler_onMessageReceived(unsafe.Pointer(this_),
		cUserId,
		unsafe.Pointer(message),
	)
	C.free(unsafe.Pointer(cUserId))
}

/**
 * Reports the state of the message sent by the local user.
 * @param messageId ID of the message.
 * @param state The state of the message: #CHANNEL_MESSAGE_STATE.
 */
func (this_ *IChannelEventHandler) OnSendMessageState(messageId int64, state CHANNEL_MESSAGE_STATE) {
	C.C_IChannelEventHandler_onSendMessageState(unsafe.Pointer(this_),
		C.int64_t(messageId),
		C.enum_C_CHANNEL_MESSAGE_STATE(state),
	)
}

/*
*

	Returns the result of the \ref agora::rtm::IChannel::sendMessage "sendMessage" method call.

	 @param messageId The ID of the sent channel message.
	 @param state The error codes. See #CHANNEL_MESSAGE_ERR_CODE.
*/
func (this_ *IChannelEventHandler) OnSendMessageResult(messageId int64, state CHANNEL_MESSAGE_ERR_CODE) {
	C.C_IChannelEventHandler_onSendMessageResult(unsafe.Pointer(this_),
		C.longlong(messageId),
		C.enum_C_CHANNEL_MESSAGE_ERR_CODE(state),
	)
}

/**
 * Occurs when another member joins the channel.
 * @param member The pointer to the member who joins the channel: IChannelMember.
 */
func (this_ *IChannelEventHandler) OnMemberJoined(member *IChannelMember) {
	C.C_IChannelEventHandler_onMemberJoined(unsafe.Pointer(this_),
		unsafe.Pointer(member),
	)
}

/**
 * Occurs when the other member leaves the channel.
 * @param member The pointer to the member who leaves the channel: IChannelMember.
 */
func (this_ *IChannelEventHandler) OnMemberLeft(member *IChannelMember) {
	C.C_IChannelEventHandler_onMemberLeft(unsafe.Pointer(this_),
		unsafe.Pointer(member),
	)
}

/**
 * Reports all the members in the channel.
 * @param members The pointer to each member in the channel: IChannelMember.
 * @param userCount The number of users in the channel.
 */
func (this_ *IChannelEventHandler) OnMembersGotten(members []*IChannelMember, userCount int) {
	C.C_IChannelEventHandler_onMembersGotten(unsafe.Pointer(this_),
		(*unsafe.Pointer)(unsafe.Pointer(unsafe.SliceData(members))),
		C.int(userCount),
	)
}

/**
 * Occurs when the channel attributes are updated.
 * @param attributes The pointer to the current channel attributes: IChannelAttributes.
 */
func (this_ *IChannelEventHandler) OnAttributesUpdated(attributes *IChannelAttributes) {
	C.C_IChannelEventHandler_onAttributesUpdated(unsafe.Pointer(this_),
		unsafe.Pointer(attributes),
	)
}

/**
 * Occurs when the local user calls updateAttributes().
 * @param requestId ID of the current attribute update request.
 * @param resCode The response code: #RESPONSE_CODE.
 */
func (this_ *IChannelEventHandler) OnUpdateAttributesResponse(requestId int64, resCode RESPONSE_CODE) {
	C.C_IChannelEventHandler_onUpdateAttributesResponse(unsafe.Pointer(this_),
		C.int64_t(requestId),
		C.enum_C_RESPONSE_CODE(resCode),
	)
}

/**
 * Occurs when the channel attributes are deleted.
 * @param attributes The pointer to the channel attributes that you want to remove: IChannelAttributes.
 */
func (this_ *IChannelEventHandler) OnAttributesDeleted(attributes *IChannelAttributes) {
	C.C_IChannelEventHandler_onAttributesDeleted(unsafe.Pointer(this_),
		unsafe.Pointer(attributes),
	)
}

/**
 * Occurs when the local user calls deleteAttributes().
 * @param requestId ID of the current attribute delete request.
 * @param resCode The response code: #RESPONSE_CODE.
 */
func (this_ *IChannelEventHandler) OnDeleteAttributesResponse(requestId int64, resCode RESPONSE_CODE) {
	C.C_IChannelEventHandler_onDeleteAttributesResponse(unsafe.Pointer(this_),
		C.int64_t(requestId),
		C.enum_C_RESPONSE_CODE(resCode),
	)
}

// #endregion IChannelEventHandler

/**
 * The IChannel class.
 */
type IChannel C.C_IChannel

// #region IChannel

/**
 * Sets an event handler for IChannel.
 * @param eventHandler The pointer to the event handler of IChannel: IChannelEventHandler.
 */
func (this_ *IChannel) SetEventHandler(eventHandler *IChannelEventHandler) {
	C.C_IChannel_setEventHandler(unsafe.Pointer(this_), unsafe.Pointer(eventHandler))
}

/**
 * Joins the current channel.
 *
 * A successful method call triggers either onJoinSuccess() or onJoinFailure() on the local client, to report
 * the state of joining the channel.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IChannel) Join() int {
	return int(C.C_IChannel_join(unsafe.Pointer(this_)))
}

/**
 * Leaves the current channel.
 *
 * After the local user successfully leaves the channel, the SDK triggers the onLeave() on the local client.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IChannel) Leave() int {
	return int(C.C_IChannel_leave(unsafe.Pointer(this_)))
}

/**
 * Sends a channel message.
 *
 * After you successfully send a channel message, all members in the channel receive the message in the
 * onMessageReceived() callback.
 * @param message The pointer to the channel message that you want to send: IMessage.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IChannel) SendMessage(message *IMessage) int {
	return int(C.C_IChannel_sendMessage(unsafe.Pointer(this_), unsafe.Pointer(message)))
}

/**
 * Updates the channel attributes.
 *
 * A successful method call triggers the onUpdateAttributesResponse() callback on the local client.
 * @param attributes The pointer to the channel attributes that you want to update: IChannelAttributes.
 * @param requestId ID of the current update request.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IChannel) UpdateAttributes(attributes *IChannelAttributes, requestId *int64) int {
	return int(C.C_IChannel_updateAttributes(unsafe.Pointer(this_),
		unsafe.Pointer(attributes),
		(*C.int64_t)(requestId),
	))
}

/**
 * Removes the channel attributes.
 *
 * A successful method call triggers the onDeleteAttributesResponse() callback on the local client.
 * @param attributes The pointer to the channel attributes that you want to remove: IChannelAttributes.
 * @param requestId ID of the current delete request.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IChannel) DeleteAttributes(attributes *IChannelAttributes, requestId *int64) int {
	return int(C.C_IChannel_deleteAttributes(unsafe.Pointer(this_),
		unsafe.Pointer(attributes),
		(*C.int64_t)(requestId),
	))
}

/**
 * Gets the current request ID.
 * @return
 * - The pointer to the request ID, if the method call succeeds.
 * - An empty pointer NULL, if the method call fails.
 */
func (this_ *IChannel) GetId() string {
	return C.GoString(C.C_IChannel_getId(unsafe.Pointer(this_)))
}

// sync_call
/**
 * Releases the IChannel instance.
 *
 * This is a synchronous method call, which means that the SDK reports the result of this method call
 * after the IChannel instance is successfully released.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IChannel) Release() int {
	return int(C.C_IChannel_release(unsafe.Pointer(this_)))
}

// #endregion IChannel

/**
 * The IRtmServiceEventHandler class.
 */
type IRtmServiceEventHandler C.C_IRtmServiceEventHandler

// region IRtmServiceEventHandler
func (this_ *IRtmServiceEventHandler) Delete() {
	C.C_IRtmServiceEventHandler_Delete(unsafe.Pointer(this_))
}

/**
 * Occurs when the user successfully logs in the RTM service.
 */
func (this_ *IRtmServiceEventHandler) OnLoginSuccess() {
	C.C_IRtmServiceEventHandler_onLoginSuccess(unsafe.Pointer(this_))
}

/**
 * Occurs when the user fails to log in the RTM service.
 * @param errorCode The reason for the login failure: #LOGIN_ERR_CODE.
 */
func (this_ *IRtmServiceEventHandler) OnLoginFailure(errorCode LOGIN_ERR_CODE) {
	C.C_IRtmServiceEventHandler_onLoginFailure(unsafe.Pointer(this_),
		C.enum_C_LOGIN_ERR_CODE(errorCode),
	)
}

/**
 * Occurs when the user successfully logs out of the RTM service.
 */
func (this_ *IRtmServiceEventHandler) OnLogout() {
	C.C_IRtmServiceEventHandler_onLogout(unsafe.Pointer(this_))
}

/**
 * Occurs when the connection state of the local user has changed.
 * @param state The current connection state: #CONNECTION_STATE.
 */
func (this_ *IRtmServiceEventHandler) OnConnectionStateChanged(state CONNECTION_STATE) {
	C.C_IRtmServiceEventHandler_onConnectionStateChanged(unsafe.Pointer(this_),
		C.enum_C_CONNECTION_STATE(state),
	)
}

/**
 * Reports the state of sending a message.
 * @param messageId ID of the message.
 * @param state The current state of the message: #PEER_MESSAGE_STATE.
 */
func (this_ *IRtmServiceEventHandler) OnSendMessageState(messageId int64, state PEER_MESSAGE_STATE) {
	C.C_IRtmServiceEventHandler_onSendMessageState(unsafe.Pointer(this_),
		C.int64_t(messageId),
		C.enum_C_PEER_MESSAGE_STATE(state),
	)
}

/**
 * Occurs when the local user receives a message from a remote user.
 * @param peerId ID of the remote user that sends the message.
 * @param message The pointer to the message: IMessage.
 */
// func (this_ *IRtmServiceEventHandler) OnMessageReceivedFromPeer(peerId string, message *IMessage) {
// 	cPeerId := C.CString(peerId)
// 	C.C_IRtmServiceEventHandler_onMessageReceivedFromPeer(unsafe.Pointer(this_),
// 		cPeerId,
// 		unsafe.Pointer(message),
// 	)
// 	C.free(unsafe.Pointer(cPeerId))
// }
//TODO

// #endregion IRtmServiceEventHandler

/**
 * The IRtmService class.
 */
type IRtmService C.C_IRtmService

// #region IRtmService

/**
 * Creates and gets an IRtmService instance.
 * @param appId The pointer to the app ID.
 * @param eventHandler The pointer to the IRtmServiceEventHandler object.
 * @param eventSpace The connection specific ID, used during report to argus.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmService) Initialize(appId string, eventHandler *IRtmServiceEventHandler) int {
	cAppId := C.CString(appId)
	ret := int(C.C_IRtmService_initialize(unsafe.Pointer(this_),
		cAppId,
		unsafe.Pointer(eventHandler),
	))
	C.free(unsafe.Pointer(cAppId))
	return ret
}

/**
 * Releases the IRtmServiceEventHandler object.
 * @param eventHandler The pointer to the IRtmServiceEventHandler object.
 */
func (this_ *IRtmService) UnregisterObserver(eventHandler *IRtmServiceEventHandler) {
	C.C_IRtmService_unregisterObserver(unsafe.Pointer(this_), unsafe.Pointer(eventHandler))
}

/**
 * Releases the IRtmService instance.
 * @param sync Determines whether to report the result of this method call synchronously.
 * - true: Report the result of this method call after the IRtmService instance is released.
 * - false: (Default) Report the result of this method call immediately, even when the IRtmService is not
 * released.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmService) Release(sync bool) {
	C.C_IRtmService_release(unsafe.Pointer(this_), C.bool(sync))
}

/**
 * Logs in the RTM service.
 *
 * @note
 * - If you login with the same user ID from a different instance, your previous login will be kicked.
 * - The call frequency limit of this method is 2 queries per second.
 * @param token The token used to log in the RTM service.
 * @param userId ID of the user logging in the RTM service.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmService) Login(token string, userId string) int {
	cToken := C.CString(token)
	cUserId := C.CString(userId)
	ret := int(C.C_IRtmService_login(unsafe.Pointer(this_),
		cToken,
		cUserId,
	))
	C.free(unsafe.Pointer(cToken))
	C.free(unsafe.Pointer(cUserId))
	return ret
}

/**
 * Logs out of the RTM service.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmService) Logout() int {
	return int(C.C_IRtmService_logout(unsafe.Pointer(this_)))
}

/**
 * Sends a peer message to a specified remote user.
 *
 * @param peerId The pointer to the ID of the remote user.
 * @param message The pointer to message: IMessage.
 * @return
 * - 0: Success.
 * - < 0: Failure.
 */
func (this_ *IRtmService) SendMessageToPeer(peerId string, message *IMessage) int {
	cPeerId := C.CString(peerId)
	ret := int(C.C_IRtmService_sendMessageToPeer(unsafe.Pointer(this_),
		cPeerId,
		unsafe.Pointer(message),
	))
	C.free(unsafe.Pointer(cPeerId))
	return ret
}

/**
 * Creates an RTM channel.
 *
 * @param channelId The unique channel name for an RTM session. Supported character scopes are:
 * - All lowercase English letters: a to z.
 * - All uppercase English letters: A to Z.
 * - All numeric characters: 0 to 9.
 * - The space character.
 * - Punctuation characters and other symbols, including: "!", "#", "$", "%", "&", "(", ")", "+", "-", ":", ";", "<", "=",
 * ".", ">", "?", "@", "[", "]", "^", "_", " {", "}", "|", "~", ","
 *
 * @param eventHandler The pointer to IChannelEventHandler.
 * @return
 * - The pointer to an IChannel instance, if the method call succeeds.
 * - An empty pointer NULL, if the method call fails.
 */
func (this_ *IRtmService) CreateChannel(channelId string, eventHandler *IChannelEventHandler) *IChannel {
	cChannelId := C.CString(channelId)
	ret := (*IChannel)(C.C_IRtmService_createChannel(unsafe.Pointer(this_),
		cChannelId,
		unsafe.Pointer(eventHandler),
	))
	C.free(unsafe.Pointer(cChannelId))
	return ret
}

// #endregion IRtmService

// #endregion agora::rtm

// #endregion agora
