// WARNING: This file has automatically been generated on Sat, 04 Nov 2017 07:47:37 MSK.
// By https://git.io/c-for-go. DO NOT EDIT.

package wpaclient

const (
	// WpaCtrlReq as defined in wpactrl/wpa_ctrl.h:19
	WpaCtrlReq = "CTRL-REQ-"
	// WpaCtrlRsp as defined in wpactrl/wpa_ctrl.h:22
	WpaCtrlRsp = "CTRL-RSP-"
	// WpaEventConnected as defined in wpactrl/wpa_ctrl.h:26
	WpaEventConnected = "CTRL-EVENT-CONNECTED "
	// WpaEventDisconnected as defined in wpactrl/wpa_ctrl.h:28
	WpaEventDisconnected = "CTRL-EVENT-DISCONNECTED "
	// WpaEventAssocReject as defined in wpactrl/wpa_ctrl.h:30
	WpaEventAssocReject = "CTRL-EVENT-ASSOC-REJECT "
	// WpaEventAuthReject as defined in wpactrl/wpa_ctrl.h:32
	WpaEventAuthReject = "CTRL-EVENT-AUTH-REJECT "
	// WpaEventTerminating as defined in wpactrl/wpa_ctrl.h:34
	WpaEventTerminating = "CTRL-EVENT-TERMINATING "
	// WpaEventPasswordChanged as defined in wpactrl/wpa_ctrl.h:36
	WpaEventPasswordChanged = "CTRL-EVENT-PASSWORD-CHANGED "
	// WpaEventEapNotification as defined in wpactrl/wpa_ctrl.h:38
	WpaEventEapNotification = "CTRL-EVENT-EAP-NOTIFICATION "
	// WpaEventEapStarted as defined in wpactrl/wpa_ctrl.h:40
	WpaEventEapStarted = "CTRL-EVENT-EAP-STARTED "
	// WpaEventEapProposedMethod as defined in wpactrl/wpa_ctrl.h:42
	WpaEventEapProposedMethod = "CTRL-EVENT-EAP-PROPOSED-METHOD "
	// WpaEventEapMethod as defined in wpactrl/wpa_ctrl.h:44
	WpaEventEapMethod = "CTRL-EVENT-EAP-METHOD "
	// WpaEventEapPeerCert as defined in wpactrl/wpa_ctrl.h:46
	WpaEventEapPeerCert = "CTRL-EVENT-EAP-PEER-CERT "
	// WpaEventEapPeerAlt as defined in wpactrl/wpa_ctrl.h:48
	WpaEventEapPeerAlt = "CTRL-EVENT-EAP-PEER-ALT "
	// WpaEventEapTLSCertError as defined in wpactrl/wpa_ctrl.h:50
	WpaEventEapTLSCertError = "CTRL-EVENT-EAP-TLS-CERT-ERROR "
	// WpaEventEapStatus as defined in wpactrl/wpa_ctrl.h:52
	WpaEventEapStatus = "CTRL-EVENT-EAP-STATUS "
	// WpaEventEapRetransmit as defined in wpactrl/wpa_ctrl.h:54
	WpaEventEapRetransmit = "CTRL-EVENT-EAP-RETRANSMIT "
	// WpaEventEapRetransmit2 as defined in wpactrl/wpa_ctrl.h:55
	WpaEventEapRetransmit2 = "CTRL-EVENT-EAP-RETRANSMIT2 "
	// WpaEventEapSuccess as defined in wpactrl/wpa_ctrl.h:57
	WpaEventEapSuccess = "CTRL-EVENT-EAP-SUCCESS "
	// WpaEventEapSuccess2 as defined in wpactrl/wpa_ctrl.h:58
	WpaEventEapSuccess2 = "CTRL-EVENT-EAP-SUCCESS2 "
	// WpaEventEapFailure as defined in wpactrl/wpa_ctrl.h:60
	WpaEventEapFailure = "CTRL-EVENT-EAP-FAILURE "
	// WpaEventEapFailure2 as defined in wpactrl/wpa_ctrl.h:61
	WpaEventEapFailure2 = "CTRL-EVENT-EAP-FAILURE2 "
	// WpaEventEapTimeoutFailure as defined in wpactrl/wpa_ctrl.h:63
	WpaEventEapTimeoutFailure = "CTRL-EVENT-EAP-TIMEOUT-FAILURE "
	// WpaEventEapTimeoutFailure2 as defined in wpactrl/wpa_ctrl.h:64
	WpaEventEapTimeoutFailure2 = "CTRL-EVENT-EAP-TIMEOUT-FAILURE2 "
	// WpaEventTempDisabled as defined in wpactrl/wpa_ctrl.h:66
	WpaEventTempDisabled = "CTRL-EVENT-SSID-TEMP-DISABLED "
	// WpaEventReenabled as defined in wpactrl/wpa_ctrl.h:68
	WpaEventReenabled = "CTRL-EVENT-SSID-REENABLED "
	// WpaEventScanStarted as defined in wpactrl/wpa_ctrl.h:70
	WpaEventScanStarted = "CTRL-EVENT-SCAN-STARTED "
	// WpaEventScanResults as defined in wpactrl/wpa_ctrl.h:72
	WpaEventScanResults = "CTRL-EVENT-SCAN-RESULTS "
	// WpaEventScanFailed as defined in wpactrl/wpa_ctrl.h:74
	WpaEventScanFailed = "CTRL-EVENT-SCAN-FAILED "
	// WpaEventStateChange as defined in wpactrl/wpa_ctrl.h:76
	WpaEventStateChange = "CTRL-EVENT-STATE-CHANGE "
	// WpaEventBssAdded as defined in wpactrl/wpa_ctrl.h:78
	WpaEventBssAdded = "CTRL-EVENT-BSS-ADDED "
	// WpaEventBssRemoved as defined in wpactrl/wpa_ctrl.h:80
	WpaEventBssRemoved = "CTRL-EVENT-BSS-REMOVED "
	// WpaEventNetworkNotFound as defined in wpactrl/wpa_ctrl.h:82
	WpaEventNetworkNotFound = "CTRL-EVENT-NETWORK-NOT-FOUND "
	// WpaEventSignalChange as defined in wpactrl/wpa_ctrl.h:84
	WpaEventSignalChange = "CTRL-EVENT-SIGNAL-CHANGE "
	// WpaEventBeaconLoss as defined in wpactrl/wpa_ctrl.h:86
	WpaEventBeaconLoss = "CTRL-EVENT-BEACON-LOSS "
	// WpaEventRegdomChange as defined in wpactrl/wpa_ctrl.h:88
	WpaEventRegdomChange = "CTRL-EVENT-REGDOM-CHANGE "
	// WpaEventChannelSwitch as defined in wpactrl/wpa_ctrl.h:90
	WpaEventChannelSwitch = "CTRL-EVENT-CHANNEL-SWITCH "
	// WpaEventSubnetStatusUpdate as defined in wpactrl/wpa_ctrl.h:103
	WpaEventSubnetStatusUpdate = "CTRL-EVENT-SUBNET-STATUS-UPDATE "
	// IbssRsnCompleted as defined in wpactrl/wpa_ctrl.h:106
	IbssRsnCompleted = "IBSS-RSN-COMPLETED "
	// WpaEventFreqConflict as defined in wpactrl/wpa_ctrl.h:113
	WpaEventFreqConflict = "CTRL-EVENT-FREQ-CONFLICT "
	// WpaEventAvoidFreq as defined in wpactrl/wpa_ctrl.h:115
	WpaEventAvoidFreq = "CTRL-EVENT-AVOID-FREQ "
	// WpsEventOverlap as defined in wpactrl/wpa_ctrl.h:117
	WpsEventOverlap = "WPS-OVERLAP-DETECTED "
	// WpsEventApAvailablePbc as defined in wpactrl/wpa_ctrl.h:119
	WpsEventApAvailablePbc = "WPS-AP-AVAILABLE-PBC "
	// WpsEventApAvailableAuth as defined in wpactrl/wpa_ctrl.h:121
	WpsEventApAvailableAuth = "WPS-AP-AVAILABLE-AUTH "
	// WpsEventApAvailablePin as defined in wpactrl/wpa_ctrl.h:124
	WpsEventApAvailablePin = "WPS-AP-AVAILABLE-PIN "
	// WpsEventApAvailable as defined in wpactrl/wpa_ctrl.h:126
	WpsEventApAvailable = "WPS-AP-AVAILABLE "
	// WpsEventCredReceived as defined in wpactrl/wpa_ctrl.h:128
	WpsEventCredReceived = "WPS-CRED-RECEIVED "
	// WpsEventM2d as defined in wpactrl/wpa_ctrl.h:130
	WpsEventM2d = "WPS-M2D "
	// WpsEventFail as defined in wpactrl/wpa_ctrl.h:132
	WpsEventFail = "WPS-FAIL "
	// WpsEventSuccess as defined in wpactrl/wpa_ctrl.h:134
	WpsEventSuccess = "WPS-SUCCESS "
	// WpsEventTimeout as defined in wpactrl/wpa_ctrl.h:136
	WpsEventTimeout = "WPS-TIMEOUT "
	// WpsEventActive as defined in wpactrl/wpa_ctrl.h:138
	WpsEventActive = "WPS-PBC-ACTIVE "
	// WpsEventDisable as defined in wpactrl/wpa_ctrl.h:140
	WpsEventDisable = "WPS-PBC-DISABLE "
	// WpsEventEnrolleeSeen as defined in wpactrl/wpa_ctrl.h:142
	WpsEventEnrolleeSeen = "WPS-ENROLLEE-SEEN "
	// WpsEventOpenNetwork as defined in wpactrl/wpa_ctrl.h:144
	WpsEventOpenNetwork = "WPS-OPEN-NETWORK "
	// WpsEventErApAdd as defined in wpactrl/wpa_ctrl.h:147
	WpsEventErApAdd = "WPS-ER-AP-ADD "
	// WpsEventErApRemove as defined in wpactrl/wpa_ctrl.h:148
	WpsEventErApRemove = "WPS-ER-AP-REMOVE "
	// WpsEventErEnrolleeAdd as defined in wpactrl/wpa_ctrl.h:149
	WpsEventErEnrolleeAdd = "WPS-ER-ENROLLEE-ADD "
	// WpsEventErEnrolleeRemove as defined in wpactrl/wpa_ctrl.h:150
	WpsEventErEnrolleeRemove = "WPS-ER-ENROLLEE-REMOVE "
	// WpsEventErApSettings as defined in wpactrl/wpa_ctrl.h:151
	WpsEventErApSettings = "WPS-ER-AP-SETTINGS "
	// WpsEventErSetSelReg as defined in wpactrl/wpa_ctrl.h:152
	WpsEventErSetSelReg = "WPS-ER-AP-SET-SEL-REG "
	// DppEventAuthSuccess as defined in wpactrl/wpa_ctrl.h:155
	DppEventAuthSuccess = "DPP-AUTH-SUCCESS "
	// DppEventNotCompatible as defined in wpactrl/wpa_ctrl.h:156
	DppEventNotCompatible = "DPP-NOT-COMPATIBLE "
	// DppEventResponsePending as defined in wpactrl/wpa_ctrl.h:157
	DppEventResponsePending = "DPP-RESPONSE-PENDING "
	// DppEventScanPeerQrCode as defined in wpactrl/wpa_ctrl.h:158
	DppEventScanPeerQrCode = "DPP-SCAN-PEER-QR-CODE "
	// DppEventConfReceived as defined in wpactrl/wpa_ctrl.h:159
	DppEventConfReceived = "DPP-CONF-RECEIVED "
	// DppEventConfSent as defined in wpactrl/wpa_ctrl.h:160
	DppEventConfSent = "DPP-CONF-SENT "
	// DppEventConfFailed as defined in wpactrl/wpa_ctrl.h:161
	DppEventConfFailed = "DPP-CONF-FAILED "
	// DppEventConfobjSsid as defined in wpactrl/wpa_ctrl.h:162
	DppEventConfobjSsid = "DPP-CONFOBJ-SSID "
	// DppEventConfobjPass as defined in wpactrl/wpa_ctrl.h:163
	DppEventConfobjPass = "DPP-CONFOBJ-PASS "
	// DppEventConfobjPsk as defined in wpactrl/wpa_ctrl.h:164
	DppEventConfobjPsk = "DPP-CONFOBJ-PSK "
	// DppEventConnector as defined in wpactrl/wpa_ctrl.h:165
	DppEventConnector = "DPP-CONNECTOR "
	// DppEventCSignKey as defined in wpactrl/wpa_ctrl.h:166
	DppEventCSignKey = "DPP-C-SIGN-KEY "
	// DppEventNetAccessKey as defined in wpactrl/wpa_ctrl.h:167
	DppEventNetAccessKey = "DPP-NET-ACCESS-KEY "
	// DppEventMissingConnector as defined in wpactrl/wpa_ctrl.h:168
	DppEventMissingConnector = "DPP-MISSING-CONNECTOR "
	// DppEventNetworkID as defined in wpactrl/wpa_ctrl.h:169
	DppEventNetworkID = "DPP-NETWORK-ID "
	// DppEventRx as defined in wpactrl/wpa_ctrl.h:170
	DppEventRx = "DPP-RX "
	// DppEventTx as defined in wpactrl/wpa_ctrl.h:171
	DppEventTx = "DPP-TX "
	// DppEventTxStatus as defined in wpactrl/wpa_ctrl.h:172
	DppEventTxStatus = "DPP-TX-STATUS "
	// DppEventFail as defined in wpactrl/wpa_ctrl.h:173
	DppEventFail = "DPP-FAIL "
	// MeshGroupStarted as defined in wpactrl/wpa_ctrl.h:176
	MeshGroupStarted = "MESH-GROUP-STARTED "
	// MeshGroupRemoved as defined in wpactrl/wpa_ctrl.h:177
	MeshGroupRemoved = "MESH-GROUP-REMOVED "
	// MeshPeerConnected as defined in wpactrl/wpa_ctrl.h:178
	MeshPeerConnected = "MESH-PEER-CONNECTED "
	// MeshPeerDisconnected as defined in wpactrl/wpa_ctrl.h:179
	MeshPeerDisconnected = "MESH-PEER-DISCONNECTED "
	// MeshSaeAuthFailure as defined in wpactrl/wpa_ctrl.h:181
	MeshSaeAuthFailure = "MESH-SAE-AUTH-FAILURE "
	// MeshSaeAuthBlocked as defined in wpactrl/wpa_ctrl.h:182
	MeshSaeAuthBlocked = "MESH-SAE-AUTH-BLOCKED "
	// P2pEventDeviceFound as defined in wpactrl/wpa_ctrl.h:190
	P2pEventDeviceFound = "P2P-DEVICE-FOUND "
	// P2pEventDeviceLost as defined in wpactrl/wpa_ctrl.h:193
	P2pEventDeviceLost = "P2P-DEVICE-LOST "
	// P2pEventGoNegRequest as defined in wpactrl/wpa_ctrl.h:197
	P2pEventGoNegRequest = "P2P-GO-NEG-REQUEST "
	// P2pEventGoNegSuccess as defined in wpactrl/wpa_ctrl.h:198
	P2pEventGoNegSuccess = "P2P-GO-NEG-SUCCESS "
	// P2pEventGoNegFailure as defined in wpactrl/wpa_ctrl.h:199
	P2pEventGoNegFailure = "P2P-GO-NEG-FAILURE "
	// P2pEventGroupFormationSuccess as defined in wpactrl/wpa_ctrl.h:200
	P2pEventGroupFormationSuccess = "P2P-GROUP-FORMATION-SUCCESS "
	// P2pEventGroupFormationFailure as defined in wpactrl/wpa_ctrl.h:201
	P2pEventGroupFormationFailure = "P2P-GROUP-FORMATION-FAILURE "
	// P2pEventGroupStarted as defined in wpactrl/wpa_ctrl.h:202
	P2pEventGroupStarted = "P2P-GROUP-STARTED "
	// P2pEventGroupRemoved as defined in wpactrl/wpa_ctrl.h:203
	P2pEventGroupRemoved = "P2P-GROUP-REMOVED "
	// P2pEventCrossConnectEnable as defined in wpactrl/wpa_ctrl.h:204
	P2pEventCrossConnectEnable = "P2P-CROSS-CONNECT-ENABLE "
	// P2pEventCrossConnectDisable as defined in wpactrl/wpa_ctrl.h:205
	P2pEventCrossConnectDisable = "P2P-CROSS-CONNECT-DISABLE "
	// P2pEventProvDiscShowPin as defined in wpactrl/wpa_ctrl.h:207
	P2pEventProvDiscShowPin = "P2P-PROV-DISC-SHOW-PIN "
	// P2pEventProvDiscEnterPin as defined in wpactrl/wpa_ctrl.h:209
	P2pEventProvDiscEnterPin = "P2P-PROV-DISC-ENTER-PIN "
	// P2pEventProvDiscPbcReq as defined in wpactrl/wpa_ctrl.h:211
	P2pEventProvDiscPbcReq = "P2P-PROV-DISC-PBC-REQ "
	// P2pEventProvDiscPbcResp as defined in wpactrl/wpa_ctrl.h:213
	P2pEventProvDiscPbcResp = "P2P-PROV-DISC-PBC-RESP "
	// P2pEventProvDiscFailure as defined in wpactrl/wpa_ctrl.h:215
	P2pEventProvDiscFailure = "P2P-PROV-DISC-FAILURE"
	// P2pEventServDiscReq as defined in wpactrl/wpa_ctrl.h:217
	P2pEventServDiscReq = "P2P-SERV-DISC-REQ "
	// P2pEventServDiscResp as defined in wpactrl/wpa_ctrl.h:219
	P2pEventServDiscResp = "P2P-SERV-DISC-RESP "
	// P2pEventServAspResp as defined in wpactrl/wpa_ctrl.h:220
	P2pEventServAspResp = "P2P-SERV-ASP-RESP "
	// P2pEventInvitationReceived as defined in wpactrl/wpa_ctrl.h:221
	P2pEventInvitationReceived = "P2P-INVITATION-RECEIVED "
	// P2pEventInvitationResult as defined in wpactrl/wpa_ctrl.h:222
	P2pEventInvitationResult = "P2P-INVITATION-RESULT "
	// P2pEventInvitationAccepted as defined in wpactrl/wpa_ctrl.h:223
	P2pEventInvitationAccepted = "P2P-INVITATION-ACCEPTED "
	// P2pEventFindStopped as defined in wpactrl/wpa_ctrl.h:224
	P2pEventFindStopped = "P2P-FIND-STOPPED "
	// P2pEventPersistentPskFail as defined in wpactrl/wpa_ctrl.h:225
	P2pEventPersistentPskFail = "P2P-PERSISTENT-PSK-FAIL id="
	// P2pEventPresenceResponse as defined in wpactrl/wpa_ctrl.h:226
	P2pEventPresenceResponse = "P2P-PRESENCE-RESPONSE "
	// P2pEventNfcBothGo as defined in wpactrl/wpa_ctrl.h:227
	P2pEventNfcBothGo = "P2P-NFC-BOTH-GO "
	// P2pEventNfcPeerClient as defined in wpactrl/wpa_ctrl.h:228
	P2pEventNfcPeerClient = "P2P-NFC-PEER-CLIENT "
	// P2pEventNfcWhileClient as defined in wpactrl/wpa_ctrl.h:229
	P2pEventNfcWhileClient = "P2P-NFC-WHILE-CLIENT "
	// P2pEventFallbackToGoNeg as defined in wpactrl/wpa_ctrl.h:230
	P2pEventFallbackToGoNeg = "P2P-FALLBACK-TO-GO-NEG "
	// P2pEventFallbackToGoNegEnabled as defined in wpactrl/wpa_ctrl.h:231
	P2pEventFallbackToGoNegEnabled = "P2P-FALLBACK-TO-GO-NEG-ENABLED "
	// EssDisassocImminent as defined in wpactrl/wpa_ctrl.h:234
	EssDisassocImminent = "ESS-DISASSOC-IMMINENT "
	// P2pEventRemoveAndReformGroup as defined in wpactrl/wpa_ctrl.h:235
	P2pEventRemoveAndReformGroup = "P2P-REMOVE-AND-REFORM-GROUP "
	// P2pEventP2psProvisionStart as defined in wpactrl/wpa_ctrl.h:237
	P2pEventP2psProvisionStart = "P2PS-PROV-START "
	// P2pEventP2psProvisionDone as defined in wpactrl/wpa_ctrl.h:238
	P2pEventP2psProvisionDone = "P2PS-PROV-DONE "
	// InterworkingAp as defined in wpactrl/wpa_ctrl.h:240
	InterworkingAp = "INTERWORKING-AP "
	// InterworkingBlacklisted as defined in wpactrl/wpa_ctrl.h:241
	InterworkingBlacklisted = "INTERWORKING-BLACKLISTED "
	// InterworkingNoMatch as defined in wpactrl/wpa_ctrl.h:242
	InterworkingNoMatch = "INTERWORKING-NO-MATCH "
	// InterworkingAlreadyConnected as defined in wpactrl/wpa_ctrl.h:243
	InterworkingAlreadyConnected = "INTERWORKING-ALREADY-CONNECTED "
	// InterworkingSelected as defined in wpactrl/wpa_ctrl.h:244
	InterworkingSelected = "INTERWORKING-SELECTED "
	// CredAdded as defined in wpactrl/wpa_ctrl.h:247
	CredAdded = "CRED-ADDED "
	// CredModified as defined in wpactrl/wpa_ctrl.h:249
	CredModified = "CRED-MODIFIED "
	// CredRemoved as defined in wpactrl/wpa_ctrl.h:251
	CredRemoved = "CRED-REMOVED "
	// GasResponseInfo as defined in wpactrl/wpa_ctrl.h:253
	GasResponseInfo = "GAS-RESPONSE-INFO "
	// GasQueryStart as defined in wpactrl/wpa_ctrl.h:255
	GasQueryStart = "GAS-QUERY-START "
	// GasQueryDone as defined in wpactrl/wpa_ctrl.h:257
	GasQueryDone = "GAS-QUERY-DONE "
	// RxAnqp as defined in wpactrl/wpa_ctrl.h:262
	RxAnqp = "RX-ANQP "
	// RxHs20Anqp as defined in wpactrl/wpa_ctrl.h:263
	RxHs20Anqp = "RX-HS20-ANQP "
	// RxHs20AnqpIcon as defined in wpactrl/wpa_ctrl.h:264
	RxHs20AnqpIcon = "RX-HS20-ANQP-ICON "
	// RxHs20Icon as defined in wpactrl/wpa_ctrl.h:265
	RxHs20Icon = "RX-HS20-ICON "
	// RxMboAnqp as defined in wpactrl/wpa_ctrl.h:266
	RxMboAnqp = "RX-MBO-ANQP "
	// Hs20SubscriptionRemediation as defined in wpactrl/wpa_ctrl.h:268
	Hs20SubscriptionRemediation = "HS20-SUBSCRIPTION-REMEDIATION "
	// Hs20DeauthImminentNotice as defined in wpactrl/wpa_ctrl.h:269
	Hs20DeauthImminentNotice = "HS20-DEAUTH-IMMINENT-NOTICE "
	// RrmEventNeighborRepRxed as defined in wpactrl/wpa_ctrl.h:274
	RrmEventNeighborRepRxed = "RRM-NEIGHBOR-REP-RECEIVED "
	// RrmEventNeighborRepFailed as defined in wpactrl/wpa_ctrl.h:275
	RrmEventNeighborRepFailed = "RRM-NEIGHBOR-REP-REQUEST-FAILED "
	// WpsEventPinNeeded as defined in wpactrl/wpa_ctrl.h:278
	WpsEventPinNeeded = "WPS-PIN-NEEDED "
	// WpsEventNewApSettings as defined in wpactrl/wpa_ctrl.h:279
	WpsEventNewApSettings = "WPS-NEW-AP-SETTINGS "
	// WpsEventRegSuccess as defined in wpactrl/wpa_ctrl.h:280
	WpsEventRegSuccess = "WPS-REG-SUCCESS "
	// WpsEventApSetupLocked as defined in wpactrl/wpa_ctrl.h:281
	WpsEventApSetupLocked = "WPS-AP-SETUP-LOCKED "
	// WpsEventApSetupUnlocked as defined in wpactrl/wpa_ctrl.h:282
	WpsEventApSetupUnlocked = "WPS-AP-SETUP-UNLOCKED "
	// WpsEventApPinEnabled as defined in wpactrl/wpa_ctrl.h:283
	WpsEventApPinEnabled = "WPS-AP-PIN-ENABLED "
	// WpsEventApPinDisabled as defined in wpactrl/wpa_ctrl.h:284
	WpsEventApPinDisabled = "WPS-AP-PIN-DISABLED "
	// ApStaConnected as defined in wpactrl/wpa_ctrl.h:285
	ApStaConnected = "AP-STA-CONNECTED "
	// ApStaDisconnected as defined in wpactrl/wpa_ctrl.h:286
	ApStaDisconnected = "AP-STA-DISCONNECTED "
	// ApStaPossiblePskMismatch as defined in wpactrl/wpa_ctrl.h:287
	ApStaPossiblePskMismatch = "AP-STA-POSSIBLE-PSK-MISMATCH "
	// ApStaPollOk as defined in wpactrl/wpa_ctrl.h:288
	ApStaPollOk = "AP-STA-POLL-OK "
	// ApRejectedMaxSta as defined in wpactrl/wpa_ctrl.h:290
	ApRejectedMaxSta = "AP-REJECTED-MAX-STA "
	// ApRejectedBlockedSta as defined in wpactrl/wpa_ctrl.h:291
	ApRejectedBlockedSta = "AP-REJECTED-BLOCKED-STA "
	// ApEventEnabled as defined in wpactrl/wpa_ctrl.h:293
	ApEventEnabled = "AP-ENABLED "
	// ApEventDisabled as defined in wpactrl/wpa_ctrl.h:294
	ApEventDisabled = "AP-DISABLED "
	// InterfaceEnabled as defined in wpactrl/wpa_ctrl.h:296
	InterfaceEnabled = "INTERFACE-ENABLED "
	// InterfaceDisabled as defined in wpactrl/wpa_ctrl.h:297
	InterfaceDisabled = "INTERFACE-DISABLED "
	// AcsEventStarted as defined in wpactrl/wpa_ctrl.h:299
	AcsEventStarted = "ACS-STARTED "
	// AcsEventCompleted as defined in wpactrl/wpa_ctrl.h:300
	AcsEventCompleted = "ACS-COMPLETED "
	// AcsEventFailed as defined in wpactrl/wpa_ctrl.h:301
	AcsEventFailed = "ACS-FAILED "
	// DfsEventRadarDetected as defined in wpactrl/wpa_ctrl.h:303
	DfsEventRadarDetected = "DFS-RADAR-DETECTED "
	// DfsEventNewChannel as defined in wpactrl/wpa_ctrl.h:304
	DfsEventNewChannel = "DFS-NEW-CHANNEL "
	// DfsEventCacStart as defined in wpactrl/wpa_ctrl.h:305
	DfsEventCacStart = "DFS-CAC-START "
	// DfsEventCacCompleted as defined in wpactrl/wpa_ctrl.h:306
	DfsEventCacCompleted = "DFS-CAC-COMPLETED "
	// DfsEventNopFinished as defined in wpactrl/wpa_ctrl.h:307
	DfsEventNopFinished = "DFS-NOP-FINISHED "
	// DfsEventPreCacExpired as defined in wpactrl/wpa_ctrl.h:308
	DfsEventPreCacExpired = "DFS-PRE-CAC-EXPIRED "
	// ApCsaFinished as defined in wpactrl/wpa_ctrl.h:310
	ApCsaFinished = "AP-CSA-FINISHED "
	// P2pEventListenOffloadStop as defined in wpactrl/wpa_ctrl.h:312
	P2pEventListenOffloadStop = "P2P-LISTEN-OFFLOAD-STOPPED "
	// P2pListenOffloadStopReason as defined in wpactrl/wpa_ctrl.h:313
	P2pListenOffloadStopReason = "P2P-LISTEN-OFFLOAD-STOP-REASON "
	// BssTmResp as defined in wpactrl/wpa_ctrl.h:316
	BssTmResp = "BSS-TM-RESP "
	// MboCellPreference as defined in wpactrl/wpa_ctrl.h:319
	MboCellPreference = "MBO-CELL-PREFERENCE "
	// MboTransitionReason as defined in wpactrl/wpa_ctrl.h:322
	MboTransitionReason = "MBO-TRANSITION-REASON "
	// PmksaCacheAdded as defined in wpactrl/wpa_ctrl.h:330
	PmksaCacheAdded = "PMKSA-CACHE-ADDED "
	// PmksaCacheRemoved as defined in wpactrl/wpa_ctrl.h:332
	PmksaCacheRemoved = "PMKSA-CACHE-REMOVED "
	// FilsHlpRx as defined in wpactrl/wpa_ctrl.h:336
	FilsHlpRx = "FILS-HLP-RX "
	// WpaBssMaskAll as defined in wpactrl/wpa_ctrl.h:340
	WpaBssMaskAll = 0xFFFDFFFF
	// WpaCtrlIfacePort as defined in wpactrl/wpa_ctrl.h:541
	WpaCtrlIfacePort = 9877
	// WpaCtrlIfacePortLimit as defined in wpactrl/wpa_ctrl.h:542
	WpaCtrlIfacePortLimit = 50
	// WpaGlobalCtrlIfacePort as defined in wpactrl/wpa_ctrl.h:543
	WpaGlobalCtrlIfacePort = 9878
	// WpaGlobalCtrlIfacePortLimit as defined in wpactrl/wpa_ctrl.h:544
	WpaGlobalCtrlIfacePortLimit = 20
)

// WpaVendorElemFrame as declared in wpactrl/wpa_ctrl.h:369
type WpaVendorElemFrame int32

// WpaVendorElemFrame enumeration from wpactrl/wpa_ctrl.h:369
const (
	VendorElemProbeReqP2p WpaVendorElemFrame = iota
	VendorElemProbeRespP2p
	VendorElemProbeRespP2pGo
	VendorElemBeaconP2pGo
	VendorElemP2pPdReq
	VendorElemP2pPdResp
	VendorElemP2pGoNegReq
	VendorElemP2pGoNegResp
	VendorElemP2pGoNegConf
	VendorElemP2pInvReq
	VendorElemP2pInvResp
	VendorElemP2pAssocReq
	VendorElemP2pAssocResp
	VendorElemAssocReq
	VendorElemProbeReq
	NumVendorElemFrames
)
