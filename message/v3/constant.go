package message

type MessageType uint8

const (
	reserved0 MessageType = iota
	CONNECT
	CONNACK
	PUBLISH
	PUBACK
	PUBREC
	PUBREL
	PUBCOMP
	SUBSCRIBE
	SUBACK
	UNSUBSCRIBE
	UNSUBACK
	PINGREQ
	PINGRESP
	DISCONNECT
	AUTH
	reserved16
)

func IsMessageTypeAvailable(v uint8) bool {
	if v > 0 && v < 16 {
		return true
	}
	return false
}

const (
	QoS0 uint8 = iota
	QoS1
	QoS2
)

type ReasonCode uint8

func (r ReasonCode) Byte() byte {
	return byte(r)
}

func IsReasonCodeAvailable(v uint8) bool {
	if (v >= 0x00 && v <= 0x02) ||
		v == 0x04 ||
		(v >= 0x10 && v <= 0x11) ||
		(v >= 0x18 && v <= 0x19) ||
		(v >= 0x80 && v <= 0xA2) {
		return true
	}
	return false
}

const (
	Success                             ReasonCode = 0x00
	NotmalDisconnection                 ReasonCode = 0x00
	GrantedQoS0                         ReasonCode = 0x00
	GrantedQoS1                         ReasonCode = 0x01
	GrantedQoS2                         ReasonCode = 0x02
	DisconnectWithWillMessage           ReasonCode = 0x04
	NoMatchingSubscribers               ReasonCode = 0x10
	NoSubscriptionExisted               ReasonCode = 0x11
	ContinueAuthentication              ReasonCode = 0x18
	ReAuthenticate                      ReasonCode = 0x19
	UnspecifiedError                    ReasonCode = 0x80
	MalformedPacket                     ReasonCode = 0x81
	ProtocolError                       ReasonCode = 0x82
	ImplementationSpecificError         ReasonCode = 0x83
	UnsupportedProtocolVersion          ReasonCode = 0x84
	ClientIdentifierNotValid            ReasonCode = 0x85
	BadUsernameOrPassword               ReasonCode = 0x86
	NotAuthorized                       ReasonCode = 0x87
	ServerUnavailable                   ReasonCode = 0x88
	ServerBusy                          ReasonCode = 0x89
	Banned                              ReasonCode = 0x8A
	ServerShuttingDown                  ReasonCode = 0x8B
	BadAuthenticationMethod             ReasonCode = 0x8C
	KeepAliveTimeout                    ReasonCode = 0x8D
	SessionTakenOver                    ReasonCode = 0x8E
	TopicFilterInvalid                  ReasonCode = 0x8F
	TopicNameInvalid                    ReasonCode = 0x90
	PacketIdentifierInUse               ReasonCode = 0x91
	PacketIdentifierNotFound            ReasonCode = 0x92
	ReceiveMaximumExceeded              ReasonCode = 0x93
	TopicAliasInvalid                   ReasonCode = 0x94
	PacketTooLarge                      ReasonCode = 0x95
	MessageRateTooHigh                  ReasonCode = 0x96
	QuotaExceeded                       ReasonCode = 0x97
	AdministrativeAction                ReasonCode = 0x98
	PayloadFormatInvalid                ReasonCode = 0x99
	RetianlNotSupported                 ReasonCode = 0x9A
	QoSNotSupported                     ReasonCode = 0x9B
	UseAnotherServer                    ReasonCode = 0x9C
	ServerMoved                         ReasonCode = 0x9D
	SharedSubscriptionsNotSupported     ReasonCode = 0x9E
	ConnectionRateExceeded              ReasonCode = 0x9F
	MaximumConnectionTime               ReasonCode = 0xA0
	SubscriptionIdentifiersNotSupported ReasonCode = 0xA1
	WildcardSubscriptionsNotSupported   ReasonCode = 0xA2
)

type PropertyType uint8

func (p PropertyType) Byte() byte {
	return byte(p)
}

func IsPropertyTypeAvailable(v uint8) bool {
	if (v >= 0x01 && v <= 0x03) ||
		(v >= 0x08 && v <= 0x09) ||
		v == 0x0B ||
		(v >= 0x11 && v <= 0x13) ||
		(v >= 0x15 && v <= 0x1A) ||
		v == 0x1C ||
		v == 0x1F ||
		(v >= 0x21 && v <= 0x2A) {
		return true
	}
	return false
}

const (
	PayloadFormatIndecator         PropertyType = 0x01
	MesageExpiryInterval           PropertyType = 0x02
	ContentType                    PropertyType = 0x03
	ResponseTopic                  PropertyType = 0x08
	CorrelationData                PropertyType = 0x09
	SubscriptionIdentifier         PropertyType = 0x0B
	SessionExpirtInterval          PropertyType = 0x11
	AssignedClientIdentifier       PropertyType = 0x12
	ServerkeepAlice                PropertyType = 0x13
	AuthenticationMethod           PropertyType = 0x15
	AuthenticationData             PropertyType = 0x16
	RequestProblemInformation      PropertyType = 0x17
	WillDelayInterval              PropertyType = 0x18
	RequestResponseInformation     PropertyType = 0x19
	ResponseInformation            PropertyType = 0x1A
	ServerReference                PropertyType = 0x1C
	ReasonStringg                  PropertyType = 0x1F
	ReceiveMaximum                 PropertyType = 0x21
	TopicAliasMaximum              PropertyType = 0x22
	TopicAlias                     PropertyType = 0x23
	MaximumQoS                     PropertyType = 0x24
	RetainAvalilable               PropertyType = 0x25
	UserProperty                   PropertyType = 0x26
	MaximumPacketSize              PropertyType = 0x27
	WildcardSubscrived             PropertyType = 0x28
	SubscrptionIdentifierAvailable PropertyType = 0x29
	SharedSubscriptionsAvaliable   PropertyType = 0x2A
)