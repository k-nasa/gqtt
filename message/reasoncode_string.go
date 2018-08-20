// Code generated by "stringer -type=ReasonCode"; DO NOT EDIT.

package message

import "strconv"

const (
	_ReasonCode_name_0 = "SuccessGrantedQoS1GrantedQoS2"
	_ReasonCode_name_1 = "DisconnectWithWillMessage"
	_ReasonCode_name_2 = "NoMatchingSubscribersNoSubscriptionExisted"
	_ReasonCode_name_3 = "ContinueAuthenticationReAuthenticate"
	_ReasonCode_name_4 = "UnspecifiedErrorMalformedPacketProtocolErrorImplementationSpecificErrorUnsupportedProtocolVersionClientIdentifierNotValidBadUsernameOrPasswordNotAuthorizedServerUnavailableServerBusyBannedServerShuttingDownBadAuthenticationMethodKeepAliveTimeoutSessionTakenOverTopicFilterInvalidTopicNameInvalidPacketIdentifierInUsePacketIdentifierNotFoundReceiveMaximumExceededTopicAliasInvalidPacketTooLargeMessageRateTooHighQuotaExceededAdministrativeActionPayloadFormatInvalidRetianlNotSupportedQoSNotSupportedUseAnotherServerServerMovedSharedSubscriptionsNotSupportedConnectionRateExceededMaximumConnectionTimeSubscriptionIdentifiersNotSupportedWildcardSubscriptionsNotSupported"
)

var (
	_ReasonCode_index_0 = [...]uint8{0, 7, 18, 29}
	_ReasonCode_index_2 = [...]uint8{0, 21, 42}
	_ReasonCode_index_3 = [...]uint8{0, 22, 36}
	_ReasonCode_index_4 = [...]uint16{0, 16, 31, 44, 71, 97, 121, 142, 155, 172, 182, 188, 206, 229, 245, 261, 279, 295, 316, 340, 362, 379, 393, 411, 424, 444, 464, 483, 498, 514, 525, 556, 578, 599, 634, 667}
)

func (i ReasonCode) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _ReasonCode_name_0[_ReasonCode_index_0[i]:_ReasonCode_index_0[i+1]]
	case i == 4:
		return _ReasonCode_name_1
	case 16 <= i && i <= 17:
		i -= 16
		return _ReasonCode_name_2[_ReasonCode_index_2[i]:_ReasonCode_index_2[i+1]]
	case 24 <= i && i <= 25:
		i -= 24
		return _ReasonCode_name_3[_ReasonCode_index_3[i]:_ReasonCode_index_3[i+1]]
	case 128 <= i && i <= 162:
		i -= 128
		return _ReasonCode_name_4[_ReasonCode_index_4[i]:_ReasonCode_index_4[i+1]]
	default:
		return "ReasonCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}