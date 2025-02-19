// Code generated by "stringer -type=MemoryKind -trimprefix=MemoryKind"; DO NOT EDIT.

package common

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MemoryKindUnknown-0]
	_ = x[MemoryKindBoolValue-1]
	_ = x[MemoryKindAddressValue-2]
	_ = x[MemoryKindStringValue-3]
	_ = x[MemoryKindCharacterValue-4]
	_ = x[MemoryKindNumberValue-5]
	_ = x[MemoryKindArrayValueBase-6]
	_ = x[MemoryKindDictionaryValueBase-7]
	_ = x[MemoryKindCompositeValueBase-8]
	_ = x[MemoryKindSimpleCompositeValueBase-9]
	_ = x[MemoryKindOptionalValue-10]
	_ = x[MemoryKindNilValue-11]
	_ = x[MemoryKindVoidValue-12]
	_ = x[MemoryKindTypeValue-13]
	_ = x[MemoryKindPathValue-14]
	_ = x[MemoryKindCapabilityValue-15]
	_ = x[MemoryKindLinkValue-16]
	_ = x[MemoryKindStorageReferenceValue-17]
	_ = x[MemoryKindEphemeralReferenceValue-18]
	_ = x[MemoryKindInterpretedFunctionValue-19]
	_ = x[MemoryKindHostFunctionValue-20]
	_ = x[MemoryKindBoundFunctionValue-21]
	_ = x[MemoryKindBigInt-22]
	_ = x[MemoryKindSimpleCompositeValue-23]
	_ = x[MemoryKindPublishedValue-24]
	_ = x[MemoryKindAtreeArrayDataSlab-25]
	_ = x[MemoryKindAtreeArrayMetaDataSlab-26]
	_ = x[MemoryKindAtreeArrayElementOverhead-27]
	_ = x[MemoryKindAtreeMapDataSlab-28]
	_ = x[MemoryKindAtreeMapMetaDataSlab-29]
	_ = x[MemoryKindAtreeMapElementOverhead-30]
	_ = x[MemoryKindAtreeMapPreAllocatedElement-31]
	_ = x[MemoryKindAtreeEncodedSlab-32]
	_ = x[MemoryKindPrimitiveStaticType-33]
	_ = x[MemoryKindCompositeStaticType-34]
	_ = x[MemoryKindInterfaceStaticType-35]
	_ = x[MemoryKindVariableSizedStaticType-36]
	_ = x[MemoryKindConstantSizedStaticType-37]
	_ = x[MemoryKindDictionaryStaticType-38]
	_ = x[MemoryKindOptionalStaticType-39]
	_ = x[MemoryKindRestrictedStaticType-40]
	_ = x[MemoryKindReferenceStaticType-41]
	_ = x[MemoryKindCapabilityStaticType-42]
	_ = x[MemoryKindFunctionStaticType-43]
	_ = x[MemoryKindCadenceVoidValue-44]
	_ = x[MemoryKindCadenceOptionalValue-45]
	_ = x[MemoryKindCadenceBoolValue-46]
	_ = x[MemoryKindCadenceStringValue-47]
	_ = x[MemoryKindCadenceCharacterValue-48]
	_ = x[MemoryKindCadenceAddressValue-49]
	_ = x[MemoryKindCadenceIntValue-50]
	_ = x[MemoryKindCadenceNumberValue-51]
	_ = x[MemoryKindCadenceArrayValueBase-52]
	_ = x[MemoryKindCadenceArrayValueLength-53]
	_ = x[MemoryKindCadenceDictionaryValue-54]
	_ = x[MemoryKindCadenceKeyValuePair-55]
	_ = x[MemoryKindCadenceStructValueBase-56]
	_ = x[MemoryKindCadenceStructValueSize-57]
	_ = x[MemoryKindCadenceResourceValueBase-58]
	_ = x[MemoryKindCadenceResourceValueSize-59]
	_ = x[MemoryKindCadenceEventValueBase-60]
	_ = x[MemoryKindCadenceEventValueSize-61]
	_ = x[MemoryKindCadenceContractValueBase-62]
	_ = x[MemoryKindCadenceContractValueSize-63]
	_ = x[MemoryKindCadenceEnumValueBase-64]
	_ = x[MemoryKindCadenceEnumValueSize-65]
	_ = x[MemoryKindCadenceLinkValue-66]
	_ = x[MemoryKindCadencePathValue-67]
	_ = x[MemoryKindCadenceTypeValue-68]
	_ = x[MemoryKindCadenceCapabilityValue-69]
	_ = x[MemoryKindCadenceSimpleType-70]
	_ = x[MemoryKindCadenceOptionalType-71]
	_ = x[MemoryKindCadenceVariableSizedArrayType-72]
	_ = x[MemoryKindCadenceConstantSizedArrayType-73]
	_ = x[MemoryKindCadenceDictionaryType-74]
	_ = x[MemoryKindCadenceField-75]
	_ = x[MemoryKindCadenceParameter-76]
	_ = x[MemoryKindCadenceStructType-77]
	_ = x[MemoryKindCadenceResourceType-78]
	_ = x[MemoryKindCadenceEventType-79]
	_ = x[MemoryKindCadenceContractType-80]
	_ = x[MemoryKindCadenceStructInterfaceType-81]
	_ = x[MemoryKindCadenceResourceInterfaceType-82]
	_ = x[MemoryKindCadenceContractInterfaceType-83]
	_ = x[MemoryKindCadenceFunctionType-84]
	_ = x[MemoryKindCadenceReferenceType-85]
	_ = x[MemoryKindCadenceRestrictedType-86]
	_ = x[MemoryKindCadenceCapabilityType-87]
	_ = x[MemoryKindCadenceEnumType-88]
	_ = x[MemoryKindRawString-89]
	_ = x[MemoryKindAddressLocation-90]
	_ = x[MemoryKindBytes-91]
	_ = x[MemoryKindVariable-92]
	_ = x[MemoryKindCompositeTypeInfo-93]
	_ = x[MemoryKindCompositeField-94]
	_ = x[MemoryKindInvocation-95]
	_ = x[MemoryKindStorageMap-96]
	_ = x[MemoryKindStorageKey-97]
	_ = x[MemoryKindTypeToken-98]
	_ = x[MemoryKindErrorToken-99]
	_ = x[MemoryKindSpaceToken-100]
	_ = x[MemoryKindProgram-101]
	_ = x[MemoryKindIdentifier-102]
	_ = x[MemoryKindArgument-103]
	_ = x[MemoryKindBlock-104]
	_ = x[MemoryKindFunctionBlock-105]
	_ = x[MemoryKindParameter-106]
	_ = x[MemoryKindParameterList-107]
	_ = x[MemoryKindTransfer-108]
	_ = x[MemoryKindMembers-109]
	_ = x[MemoryKindTypeAnnotation-110]
	_ = x[MemoryKindDictionaryEntry-111]
	_ = x[MemoryKindFunctionDeclaration-112]
	_ = x[MemoryKindCompositeDeclaration-113]
	_ = x[MemoryKindInterfaceDeclaration-114]
	_ = x[MemoryKindEnumCaseDeclaration-115]
	_ = x[MemoryKindFieldDeclaration-116]
	_ = x[MemoryKindTransactionDeclaration-117]
	_ = x[MemoryKindImportDeclaration-118]
	_ = x[MemoryKindVariableDeclaration-119]
	_ = x[MemoryKindSpecialFunctionDeclaration-120]
	_ = x[MemoryKindPragmaDeclaration-121]
	_ = x[MemoryKindAssignmentStatement-122]
	_ = x[MemoryKindBreakStatement-123]
	_ = x[MemoryKindContinueStatement-124]
	_ = x[MemoryKindEmitStatement-125]
	_ = x[MemoryKindExpressionStatement-126]
	_ = x[MemoryKindForStatement-127]
	_ = x[MemoryKindIfStatement-128]
	_ = x[MemoryKindReturnStatement-129]
	_ = x[MemoryKindSwapStatement-130]
	_ = x[MemoryKindSwitchStatement-131]
	_ = x[MemoryKindWhileStatement-132]
	_ = x[MemoryKindBooleanExpression-133]
	_ = x[MemoryKindNilExpression-134]
	_ = x[MemoryKindStringExpression-135]
	_ = x[MemoryKindIntegerExpression-136]
	_ = x[MemoryKindFixedPointExpression-137]
	_ = x[MemoryKindArrayExpression-138]
	_ = x[MemoryKindDictionaryExpression-139]
	_ = x[MemoryKindIdentifierExpression-140]
	_ = x[MemoryKindInvocationExpression-141]
	_ = x[MemoryKindMemberExpression-142]
	_ = x[MemoryKindIndexExpression-143]
	_ = x[MemoryKindConditionalExpression-144]
	_ = x[MemoryKindUnaryExpression-145]
	_ = x[MemoryKindBinaryExpression-146]
	_ = x[MemoryKindFunctionExpression-147]
	_ = x[MemoryKindCastingExpression-148]
	_ = x[MemoryKindCreateExpression-149]
	_ = x[MemoryKindDestroyExpression-150]
	_ = x[MemoryKindReferenceExpression-151]
	_ = x[MemoryKindForceExpression-152]
	_ = x[MemoryKindPathExpression-153]
	_ = x[MemoryKindConstantSizedType-154]
	_ = x[MemoryKindDictionaryType-155]
	_ = x[MemoryKindFunctionType-156]
	_ = x[MemoryKindInstantiationType-157]
	_ = x[MemoryKindNominalType-158]
	_ = x[MemoryKindOptionalType-159]
	_ = x[MemoryKindReferenceType-160]
	_ = x[MemoryKindRestrictedType-161]
	_ = x[MemoryKindVariableSizedType-162]
	_ = x[MemoryKindPosition-163]
	_ = x[MemoryKindRange-164]
	_ = x[MemoryKindElaboration-165]
	_ = x[MemoryKindActivation-166]
	_ = x[MemoryKindActivationEntries-167]
	_ = x[MemoryKindVariableSizedSemaType-168]
	_ = x[MemoryKindConstantSizedSemaType-169]
	_ = x[MemoryKindDictionarySemaType-170]
	_ = x[MemoryKindOptionalSemaType-171]
	_ = x[MemoryKindRestrictedSemaType-172]
	_ = x[MemoryKindReferenceSemaType-173]
	_ = x[MemoryKindCapabilitySemaType-174]
	_ = x[MemoryKindOrderedMap-175]
	_ = x[MemoryKindOrderedMapEntryList-176]
	_ = x[MemoryKindOrderedMapEntry-177]
	_ = x[MemoryKindLast-178]
}

const _MemoryKind_name = "UnknownBoolValueAddressValueStringValueCharacterValueNumberValueArrayValueBaseDictionaryValueBaseCompositeValueBaseSimpleCompositeValueBaseOptionalValueNilValueVoidValueTypeValuePathValueCapabilityValueLinkValueStorageReferenceValueEphemeralReferenceValueInterpretedFunctionValueHostFunctionValueBoundFunctionValueBigIntSimpleCompositeValuePublishedValueAtreeArrayDataSlabAtreeArrayMetaDataSlabAtreeArrayElementOverheadAtreeMapDataSlabAtreeMapMetaDataSlabAtreeMapElementOverheadAtreeMapPreAllocatedElementAtreeEncodedSlabPrimitiveStaticTypeCompositeStaticTypeInterfaceStaticTypeVariableSizedStaticTypeConstantSizedStaticTypeDictionaryStaticTypeOptionalStaticTypeRestrictedStaticTypeReferenceStaticTypeCapabilityStaticTypeFunctionStaticTypeCadenceVoidValueCadenceOptionalValueCadenceBoolValueCadenceStringValueCadenceCharacterValueCadenceAddressValueCadenceIntValueCadenceNumberValueCadenceArrayValueBaseCadenceArrayValueLengthCadenceDictionaryValueCadenceKeyValuePairCadenceStructValueBaseCadenceStructValueSizeCadenceResourceValueBaseCadenceResourceValueSizeCadenceEventValueBaseCadenceEventValueSizeCadenceContractValueBaseCadenceContractValueSizeCadenceEnumValueBaseCadenceEnumValueSizeCadenceLinkValueCadencePathValueCadenceTypeValueCadenceCapabilityValueCadenceSimpleTypeCadenceOptionalTypeCadenceVariableSizedArrayTypeCadenceConstantSizedArrayTypeCadenceDictionaryTypeCadenceFieldCadenceParameterCadenceStructTypeCadenceResourceTypeCadenceEventTypeCadenceContractTypeCadenceStructInterfaceTypeCadenceResourceInterfaceTypeCadenceContractInterfaceTypeCadenceFunctionTypeCadenceReferenceTypeCadenceRestrictedTypeCadenceCapabilityTypeCadenceEnumTypeRawStringAddressLocationBytesVariableCompositeTypeInfoCompositeFieldInvocationStorageMapStorageKeyTypeTokenErrorTokenSpaceTokenProgramIdentifierArgumentBlockFunctionBlockParameterParameterListTransferMembersTypeAnnotationDictionaryEntryFunctionDeclarationCompositeDeclarationInterfaceDeclarationEnumCaseDeclarationFieldDeclarationTransactionDeclarationImportDeclarationVariableDeclarationSpecialFunctionDeclarationPragmaDeclarationAssignmentStatementBreakStatementContinueStatementEmitStatementExpressionStatementForStatementIfStatementReturnStatementSwapStatementSwitchStatementWhileStatementBooleanExpressionNilExpressionStringExpressionIntegerExpressionFixedPointExpressionArrayExpressionDictionaryExpressionIdentifierExpressionInvocationExpressionMemberExpressionIndexExpressionConditionalExpressionUnaryExpressionBinaryExpressionFunctionExpressionCastingExpressionCreateExpressionDestroyExpressionReferenceExpressionForceExpressionPathExpressionConstantSizedTypeDictionaryTypeFunctionTypeInstantiationTypeNominalTypeOptionalTypeReferenceTypeRestrictedTypeVariableSizedTypePositionRangeElaborationActivationActivationEntriesVariableSizedSemaTypeConstantSizedSemaTypeDictionarySemaTypeOptionalSemaTypeRestrictedSemaTypeReferenceSemaTypeCapabilitySemaTypeOrderedMapOrderedMapEntryListOrderedMapEntryLast"

var _MemoryKind_index = [...]uint16{0, 7, 16, 28, 39, 53, 64, 78, 97, 115, 139, 152, 160, 169, 178, 187, 202, 211, 232, 255, 279, 296, 314, 320, 340, 354, 372, 394, 419, 435, 455, 478, 505, 521, 540, 559, 578, 601, 624, 644, 662, 682, 701, 721, 739, 755, 775, 791, 809, 830, 849, 864, 882, 903, 926, 948, 967, 989, 1011, 1035, 1059, 1080, 1101, 1125, 1149, 1169, 1189, 1205, 1221, 1237, 1259, 1276, 1295, 1324, 1353, 1374, 1386, 1402, 1419, 1438, 1454, 1473, 1499, 1527, 1555, 1574, 1594, 1615, 1636, 1651, 1660, 1675, 1680, 1688, 1705, 1719, 1729, 1739, 1749, 1758, 1768, 1778, 1785, 1795, 1803, 1808, 1821, 1830, 1843, 1851, 1858, 1872, 1887, 1906, 1926, 1946, 1965, 1981, 2003, 2020, 2039, 2065, 2082, 2101, 2115, 2132, 2145, 2164, 2176, 2187, 2202, 2215, 2230, 2244, 2261, 2274, 2290, 2307, 2327, 2342, 2362, 2382, 2402, 2418, 2433, 2454, 2469, 2485, 2503, 2520, 2536, 2553, 2572, 2587, 2601, 2618, 2632, 2644, 2661, 2672, 2684, 2697, 2711, 2728, 2736, 2741, 2752, 2762, 2779, 2800, 2821, 2839, 2855, 2873, 2890, 2908, 2918, 2937, 2952, 2956}

func (i MemoryKind) String() string {
	if i >= MemoryKind(len(_MemoryKind_index)-1) {
		return "MemoryKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MemoryKind_name[_MemoryKind_index[i]:_MemoryKind_index[i+1]]
}
