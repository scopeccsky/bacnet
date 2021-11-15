// Code generated by "stringer -type=ObjectType"; DO NOT EDIT.

package bacnet

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AnalogInput-0]
	_ = x[AnalogOutput-1]
	_ = x[AnalogValue-2]
	_ = x[BinaryInput-3]
	_ = x[BinaryOutput-4]
	_ = x[BinaryValue-5]
	_ = x[Calendar-6]
	_ = x[Command-7]
	_ = x[BacnetDevice-8]
	_ = x[EventEnrollment-9]
	_ = x[File-10]
	_ = x[Group-11]
	_ = x[Loop-12]
	_ = x[MultiStateInput-13]
	_ = x[MultiStateOutput-14]
	_ = x[NotificationClass-15]
	_ = x[Program-16]
	_ = x[Schedule-17]
	_ = x[Averaging-18]
	_ = x[MultiStateValue-19]
	_ = x[Trendlog-20]
	_ = x[LifeSafetyPoint-21]
	_ = x[LifeSafetyZone-22]
	_ = x[Accumulator-23]
	_ = x[PulseConverter-24]
	_ = x[EventLog-25]
	_ = x[GlobalGroup-26]
	_ = x[TrendLogMultiple-27]
	_ = x[LoadControl-28]
	_ = x[StructuredView-29]
	_ = x[AccessDoor-30]
	_ = x[Timer-31]
	_ = x[AccessCredential-32]
	_ = x[AccessPoint-33]
	_ = x[AccessRights-34]
	_ = x[AccessUser-35]
	_ = x[AccessZone-36]
	_ = x[CredentialDataInput-37]
	_ = x[NetworkSecurity-38]
	_ = x[BitstringValue-39]
	_ = x[CharacterstringValue-40]
	_ = x[DatePatternValue-41]
	_ = x[DateValue-42]
	_ = x[DatetimePatternValue-43]
	_ = x[DatetimeValue-44]
	_ = x[IntegerValue-45]
	_ = x[LargeAnalogValue-46]
	_ = x[OctetstringValue-47]
	_ = x[PositiveIntegerValue-48]
	_ = x[TimePatternValue-49]
	_ = x[TimeValue-50]
	_ = x[NotificationForwarder-51]
	_ = x[AlertEnrollment-52]
	_ = x[Channel-53]
	_ = x[LightingOutput-54]
	_ = x[BinaryLightingOutput-55]
	_ = x[NetworkPort-56]
	_ = x[ProprietaryMin-128]
	_ = x[Proprietarymax-1023]
}

const (
	_ObjectType_name_0 = "AnalogInputAnalogOutputAnalogValueBinaryInputBinaryOutputBinaryValueCalendarCommandBacnetDeviceEventEnrollmentFileGroupLoopMultiStateInputMultiStateOutputNotificationClassProgramScheduleAveragingMultiStateValueTrendlogLifeSafetyPointLifeSafetyZoneAccumulatorPulseConverterEventLogGlobalGroupTrendLogMultipleLoadControlStructuredViewAccessDoorTimerAccessCredentialAccessPointAccessRightsAccessUserAccessZoneCredentialDataInputNetworkSecurityBitstringValueCharacterstringValueDatePatternValueDateValueDatetimePatternValueDatetimeValueIntegerValueLargeAnalogValueOctetstringValuePositiveIntegerValueTimePatternValueTimeValueNotificationForwarderAlertEnrollmentChannelLightingOutputBinaryLightingOutputNetworkPort"
	_ObjectType_name_1 = "ProprietaryMin"
	_ObjectType_name_2 = "Proprietarymax"
)

var (
	_ObjectType_index_0 = [...]uint16{0, 11, 23, 34, 45, 57, 68, 76, 83, 95, 110, 114, 119, 123, 138, 154, 171, 178, 186, 195, 210, 218, 233, 247, 258, 272, 280, 291, 307, 318, 332, 342, 347, 363, 374, 386, 396, 406, 425, 440, 454, 474, 490, 499, 519, 532, 544, 560, 576, 596, 612, 621, 642, 657, 664, 678, 698, 709}
)

func (i ObjectType) String() string {
	switch {
	case i <= 56:
		return _ObjectType_name_0[_ObjectType_index_0[i]:_ObjectType_index_0[i+1]]
	case i == 128:
		return _ObjectType_name_1
	case i == 1023:
		return _ObjectType_name_2
	default:
		return "ObjectType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}