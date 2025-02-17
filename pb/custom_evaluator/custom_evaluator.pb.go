/*
Copyright The TestGrid Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: custom_evaluator.proto

package custom_evaluator

import (
	test_status "github.com/GoogleCloudPlatform/testgrid/pb/test_status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Comparison_Operator int32

const (
	// Unknown. May assume OP_EQ for legacy purposes, but should warn.
	Comparison_OP_UNKNOWN Comparison_Operator = 0
	// Equals operator.
	Comparison_OP_EQ Comparison_Operator = 1
	// Not equals operator.
	Comparison_OP_NE Comparison_Operator = 2
	// Comparison value less than TestResult's value
	Comparison_OP_LT Comparison_Operator = 3
	// Comparison value less than or equal TestResult's value
	Comparison_OP_LE Comparison_Operator = 4
	// Comparison value greater than TestResult's value
	Comparison_OP_GT Comparison_Operator = 5
	// Comparison value greater than or equal TestResult's value
	Comparison_OP_GE Comparison_Operator = 6
	// Regex match of Comparison.value string with the TestResult's evaluation
	// value string.
	Comparison_OP_REGEX Comparison_Operator = 7
	// Checks to see if the evaluation value string starts with the
	// Comparison.value string
	Comparison_OP_STARTS_WITH Comparison_Operator = 8
	// Checks to see if the evaluation value string is contained within the
	// Comparison.value string
	Comparison_OP_CONTAINS Comparison_Operator = 9
)

// Enum value maps for Comparison_Operator.
var (
	Comparison_Operator_name = map[int32]string{
		0: "OP_UNKNOWN",
		1: "OP_EQ",
		2: "OP_NE",
		3: "OP_LT",
		4: "OP_LE",
		5: "OP_GT",
		6: "OP_GE",
		7: "OP_REGEX",
		8: "OP_STARTS_WITH",
		9: "OP_CONTAINS",
	}
	Comparison_Operator_value = map[string]int32{
		"OP_UNKNOWN":     0,
		"OP_EQ":          1,
		"OP_NE":          2,
		"OP_LT":          3,
		"OP_LE":          4,
		"OP_GT":          5,
		"OP_GE":          6,
		"OP_REGEX":       7,
		"OP_STARTS_WITH": 8,
		"OP_CONTAINS":    9,
	}
)

func (x Comparison_Operator) Enum() *Comparison_Operator {
	p := new(Comparison_Operator)
	*p = x
	return p
}

func (x Comparison_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Comparison_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_custom_evaluator_proto_enumTypes[0].Descriptor()
}

func (Comparison_Operator) Type() protoreflect.EnumType {
	return &file_custom_evaluator_proto_enumTypes[0]
}

func (x Comparison_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Comparison_Operator.Descriptor instead.
func (Comparison_Operator) EnumDescriptor() ([]byte, []int) {
	return file_custom_evaluator_proto_rawDescGZIP(), []int{3, 0}
}

// A collection of Rule objects. Used to define many rules.
type RuleSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rules []*Rule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *RuleSet) Reset() {
	*x = RuleSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_evaluator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleSet) ProtoMessage() {}

func (x *RuleSet) ProtoReflect() protoreflect.Message {
	mi := &file_custom_evaluator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleSet.ProtoReflect.Descriptor instead.
func (*RuleSet) Descriptor() ([]byte, []int) {
	return file_custom_evaluator_proto_rawDescGZIP(), []int{0}
}

func (x *RuleSet) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// A single rule that describes how to evaluate a test_cases_pb2.TestResult
type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Multiple comparisons to run against a result. EVERY TestResultComparison
	// has to succeed for this Rule to succeed.
	TestResultComparisons []*TestResultComparison `protobuf:"bytes,1,rep,name=test_result_comparisons,json=testResultComparisons,proto3" json:"test_result_comparisons,omitempty"`
	// Required: The TestStatus to return if the comparison succeeds.
	ComputedStatus test_status.TestStatus `protobuf:"varint,3,opt,name=computed_status,json=computedStatus,proto3,enum=TestStatus" json:"computed_status,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_evaluator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_custom_evaluator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_custom_evaluator_proto_rawDescGZIP(), []int{1}
}

func (x *Rule) GetTestResultComparisons() []*TestResultComparison {
	if x != nil {
		return x.TestResultComparisons
	}
	return nil
}

func (x *Rule) GetComputedStatus() test_status.TestStatus {
	if x != nil {
		return x.ComputedStatus
	}
	return test_status.TestStatus(0)
}

// Describes how to get information the TestResult proto and how to compare the
// value against the comparison value.
type TestResultComparison struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required: This is the comparison that will be used as
	Comparison *Comparison `protobuf:"bytes,1,opt,name=comparison,proto3" json:"comparison,omitempty"`
	// Types that are assignable to TestResultInfo:
	//	*TestResultComparison_PropertyKey
	//	*TestResultComparison_TestResultField
	//	*TestResultComparison_TestResultErrorField
	TestResultInfo isTestResultComparison_TestResultInfo `protobuf_oneof:"test_result_info"`
}

func (x *TestResultComparison) Reset() {
	*x = TestResultComparison{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_evaluator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResultComparison) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResultComparison) ProtoMessage() {}

func (x *TestResultComparison) ProtoReflect() protoreflect.Message {
	mi := &file_custom_evaluator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResultComparison.ProtoReflect.Descriptor instead.
func (*TestResultComparison) Descriptor() ([]byte, []int) {
	return file_custom_evaluator_proto_rawDescGZIP(), []int{2}
}

func (x *TestResultComparison) GetComparison() *Comparison {
	if x != nil {
		return x.Comparison
	}
	return nil
}

func (m *TestResultComparison) GetTestResultInfo() isTestResultComparison_TestResultInfo {
	if m != nil {
		return m.TestResultInfo
	}
	return nil
}

func (x *TestResultComparison) GetPropertyKey() string {
	if x, ok := x.GetTestResultInfo().(*TestResultComparison_PropertyKey); ok {
		return x.PropertyKey
	}
	return ""
}

func (x *TestResultComparison) GetTestResultField() string {
	if x, ok := x.GetTestResultInfo().(*TestResultComparison_TestResultField); ok {
		return x.TestResultField
	}
	return ""
}

func (x *TestResultComparison) GetTestResultErrorField() string {
	if x, ok := x.GetTestResultInfo().(*TestResultComparison_TestResultErrorField); ok {
		return x.TestResultErrorField
	}
	return ""
}

type isTestResultComparison_TestResultInfo interface {
	isTestResultComparison_TestResultInfo()
}

type TestResultComparison_PropertyKey struct {
	// The name of the property to evaluate.
	// Properties are usually strings, so a string comparison is assumed and
	// required.
	PropertyKey string `protobuf:"bytes,2,opt,name=property_key,json=propertyKey,proto3,oneof"`
}

type TestResultComparison_TestResultField struct {
	// This will find the scalar field with the given name within the TestResult
	// proto. The value of that field will be used to evaluate.
	//
	// Accepted junit values for junit results are:
	//   name: name of the test case
	//   error_count: 1 if the test case has an error message
	//   failure_count: 1 if the test case has a failure message
	//
	// NOTE: Only supported for string and numerical values.
	TestResultField string `protobuf:"bytes,3,opt,name=test_result_field,json=testResultField,proto3,oneof"`
}

type TestResultComparison_TestResultErrorField struct {
	// This will find the field nested within the first error of the TestResult
	// proto. The value of that field will be used to evaluate.
	//
	// Accepted values for junit results are:
	//   exception_type: the failure and/or error message.
	//
	// NOTE: Only supported for string and numerical values
	TestResultErrorField string `protobuf:"bytes,4,opt,name=test_result_error_field,json=testResultErrorField,proto3,oneof"`
}

func (*TestResultComparison_PropertyKey) isTestResultComparison_TestResultInfo() {}

func (*TestResultComparison_TestResultField) isTestResultComparison_TestResultInfo() {}

func (*TestResultComparison_TestResultErrorField) isTestResultComparison_TestResultInfo() {}

// The method of comparison used for evaluation. Describes how to compare two
// values.
type Comparison struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required: Defines how to compare two attributes.
	// When the TestResult value is numerical, numerical_value will be used to
	// compare. When the TestResult value is a string, string_value will be used.
	Op Comparison_Operator `protobuf:"varint,1,opt,name=op,proto3,enum=Comparison_Operator" json:"op,omitempty"`
	// Types that are assignable to ComparisonValue:
	//	*Comparison_StringValue
	//	*Comparison_NumericalValue
	ComparisonValue isComparison_ComparisonValue `protobuf_oneof:"comparison_value"`
}

func (x *Comparison) Reset() {
	*x = Comparison{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_evaluator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comparison) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comparison) ProtoMessage() {}

func (x *Comparison) ProtoReflect() protoreflect.Message {
	mi := &file_custom_evaluator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comparison.ProtoReflect.Descriptor instead.
func (*Comparison) Descriptor() ([]byte, []int) {
	return file_custom_evaluator_proto_rawDescGZIP(), []int{3}
}

func (x *Comparison) GetOp() Comparison_Operator {
	if x != nil {
		return x.Op
	}
	return Comparison_OP_UNKNOWN
}

func (m *Comparison) GetComparisonValue() isComparison_ComparisonValue {
	if m != nil {
		return m.ComparisonValue
	}
	return nil
}

func (x *Comparison) GetStringValue() string {
	if x, ok := x.GetComparisonValue().(*Comparison_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Comparison) GetNumericalValue() float64 {
	if x, ok := x.GetComparisonValue().(*Comparison_NumericalValue); ok {
		return x.NumericalValue
	}
	return 0
}

type isComparison_ComparisonValue interface {
	isComparison_ComparisonValue()
}

type Comparison_StringValue struct {
	// For operations EQ, NE, REGEX, STARTS_WITH, CONTAINS
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Comparison_NumericalValue struct {
	// For operations EQ, NE, LT, LE, GT, GE
	NumericalValue float64 `protobuf:"fixed64,3,opt,name=numerical_value,json=numericalValue,proto3,oneof"`
}

func (*Comparison_StringValue) isComparison_ComparisonValue() {}

func (*Comparison_NumericalValue) isComparison_ComparisonValue() {}

var File_custom_evaluator_proto protoreflect.FileDescriptor

var file_custom_evaluator_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x07, 0x52, 0x75,
	0x6c, 0x65, 0x53, 0x65, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x22, 0x8b, 0x01, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x4d, 0x0a, 0x17, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x69, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69,
	0x73, 0x6f, 0x6e, 0x52, 0x15, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0xe3, 0x01, 0x0a, 0x14, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x37, 0x0a, 0x17, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x14, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x42, 0x12, 0x0a, 0x10, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x69, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x23, 0x0a, 0x0c, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x29, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0e, 0x6e, 0x75, 0x6d,
	0x65, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x08,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x50, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x50, 0x5f, 0x45,
	0x51, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x50, 0x5f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x09,
	0x0a, 0x05, 0x4f, 0x50, 0x5f, 0x4c, 0x54, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x50, 0x5f,
	0x4c, 0x45, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x50, 0x5f, 0x47, 0x54, 0x10, 0x05, 0x12,
	0x09, 0x0a, 0x05, 0x4f, 0x50, 0x5f, 0x47, 0x45, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x50,
	0x5f, 0x52, 0x45, 0x47, 0x45, 0x58, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x50, 0x5f, 0x53,
	0x54, 0x41, 0x52, 0x54, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b,
	0x4f, 0x50, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x09, 0x42, 0x12, 0x0a,
	0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x67, 0x72, 0x69, 0x64, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_custom_evaluator_proto_rawDescOnce sync.Once
	file_custom_evaluator_proto_rawDescData = file_custom_evaluator_proto_rawDesc
)

func file_custom_evaluator_proto_rawDescGZIP() []byte {
	file_custom_evaluator_proto_rawDescOnce.Do(func() {
		file_custom_evaluator_proto_rawDescData = protoimpl.X.CompressGZIP(file_custom_evaluator_proto_rawDescData)
	})
	return file_custom_evaluator_proto_rawDescData
}

var file_custom_evaluator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_custom_evaluator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_custom_evaluator_proto_goTypes = []interface{}{
	(Comparison_Operator)(0),     // 0: Comparison.Operator
	(*RuleSet)(nil),              // 1: RuleSet
	(*Rule)(nil),                 // 2: Rule
	(*TestResultComparison)(nil), // 3: TestResultComparison
	(*Comparison)(nil),           // 4: Comparison
	(test_status.TestStatus)(0),  // 5: TestStatus
}
var file_custom_evaluator_proto_depIdxs = []int32{
	2, // 0: RuleSet.rules:type_name -> Rule
	3, // 1: Rule.test_result_comparisons:type_name -> TestResultComparison
	5, // 2: Rule.computed_status:type_name -> TestStatus
	4, // 3: TestResultComparison.comparison:type_name -> Comparison
	0, // 4: Comparison.op:type_name -> Comparison.Operator
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_custom_evaluator_proto_init() }
func file_custom_evaluator_proto_init() {
	if File_custom_evaluator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_custom_evaluator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_custom_evaluator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_custom_evaluator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestResultComparison); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_custom_evaluator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comparison); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_custom_evaluator_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TestResultComparison_PropertyKey)(nil),
		(*TestResultComparison_TestResultField)(nil),
		(*TestResultComparison_TestResultErrorField)(nil),
	}
	file_custom_evaluator_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Comparison_StringValue)(nil),
		(*Comparison_NumericalValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_custom_evaluator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_custom_evaluator_proto_goTypes,
		DependencyIndexes: file_custom_evaluator_proto_depIdxs,
		EnumInfos:         file_custom_evaluator_proto_enumTypes,
		MessageInfos:      file_custom_evaluator_proto_msgTypes,
	}.Build()
	File_custom_evaluator_proto = out.File
	file_custom_evaluator_proto_rawDesc = nil
	file_custom_evaluator_proto_goTypes = nil
	file_custom_evaluator_proto_depIdxs = nil
}
