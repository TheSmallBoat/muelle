package muelle

const (
	ServiceTwin      = "Muelle-Twin"
	ServicePublish   = "Muelle-Pub"
	ServiceSubscribe = "Muelle-Sub"

	ActionHeader    = "HA"
	TopicHeader     = "HT"
	QosHeader       = "HQ"
	BodyTitleHeader = "HBT"

	ResponseActionHeader = "HRA"
	ResponseStatusHeader = "HRS"

	ActionTwin        = "AT"
	ActionPublish     = "AP"
	ActionSubscribe   = "AS"
	ActionUnSubscribe = "AUS"

	ResponseTwinHeader      = "RT"
	ResponsePublishHeader   = "RP"
	ResponseSubscribeHeader = "RS"

	ResponseSuccess = "s"
	ResponseFailure = "f"

	//BodyTitle
	ServiceError = "service-error"
	ActionError  = "action-error"
	ProcessTime  = "process-time"
	PayLoad      = "PayLoad"
)

var MuNode = &Node{}
