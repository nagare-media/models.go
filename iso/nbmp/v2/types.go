/*
Copyright 2022 The nagare media authors

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

package v2

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/nagare-media/models.go/base"
)

type Function struct {
	// +optional
	Scheme *Scheme `json:"scheme,omitempty"`

	General General `json:"general"`

	Input Input `json:"input"`

	Output Output `json:"output"`

	// +optional
	Processing *Processing `json:"processing,omitempty"`

	// +optional
	Requirement *Requirement `json:"requirement,omitempty"`

	// +optional
	Configuration *Configuration `json:"configuration,omitempty"`

	// +optional
	Step *Step `json:"step,omitempty"`

	// +optional
	ClientAssistant *ClientAssistant `json:"client-assistant,omitempty"`

	// +optional
	Assertion *Assertion `json:"assertion,omitempty"`

	// +optional
	Variables []Variable `json:"variables,omitempty"`

	// +optional
	Events []Event `json:"events,omitempty"`

	// +optional
	Security *Security `json:"security,omitempty"`
}

type MediaProcessingEntityCapabilities struct {
	// +optional
	Scheme *Scheme `json:"scheme,omitempty"`

	General General `json:"general"`

	// +optional
	Capabilities *Capabilities `json:"capabilities,omitempty"`

	// +optional
	Variables []Variable `json:"variables,omitempty"`

	// +optional
	Events []Event `json:"events,omitempty"`

	// +optional
	Monitoring *Monitoring `json:"monitoring,omitempty"`

	// +optional
	Reporting *Reporting `json:"reporting,omitempty"`

	// +optional
	Notification *Notification `json:"notification,omitempty"`
}

type Workflow struct {
	// +optional
	Scheme *Scheme `json:"scheme,omitempty"`

	General General `json:"general"`

	// +optional
	Repository *Repository `json:"repository,omitempty"`

	Input Input `json:"input"`

	Output Output `json:"output"`

	Processing Processing `json:"processing"`

	Requirement Requirement `json:"requirement"`

	// +optional
	Step *Step `json:"step,omitempty"`

	// +optional
	ClientAssistant *ClientAssistant `json:"client-assistant,omitempty"`

	// +optional
	Failover *Failover `json:"failover,omitempty"`

	// +optional
	Monitoring *Monitoring `json:"monitoring,omitempty"`

	// +optional
	Assertion *Assertion `json:"assertion,omitempty"`

	// +optional
	Reporting *Reporting `json:"reporting,omitempty"`

	// +optional
	Notification *Notification `json:"notification,omitempty"`

	// +optional
	Acknowledge *Acknowledge `json:"acknowledge,omitempty"`

	// +optional
	Security *Security `json:"security,omitempty"`

	// +optional
	Scale *Scale `json:"scale,omitempty"`

	// +optional
	Schedule *Schedule `json:"schedule,omitempty"`
}

type Task struct {
	// +optional
	Scheme *Scheme `json:"scheme,omitempty"`

	General General `json:"general"`

	Input Input `json:"input"`

	Output Output `json:"output"`

	Processing Processing `json:"processing"`

	Requirement Requirement `json:"requirement"`

	// +optional
	Configuration *Configuration `json:"configuration,omitempty"`

	// +optional
	Step *Step `json:"step,omitempty"`

	// +optional
	StartupDelay *StartupDelay `json:"startup-delay,omitempty"`

	// +optional
	ClientAssistant *ClientAssistant `json:"client-assistant,omitempty"`

	// +optional
	Failover *Failover `json:"failover,omitempty"`

	// +optional
	Monitoring *Monitoring `json:"monitoring,omitempty"`

	// +optional
	Assertion *Assertion `json:"assertion,omitempty"`

	// +optional
	Reporting *Reporting `json:"reporting,omitempty"`

	// +optional
	Notification *Notification `json:"notification,omitempty"`

	// +optional
	Acknowledge *Acknowledge `json:"acknowledge,omitempty"`

	// +optional
	Security *Security `json:"security,omitempty"`

	// +optional
	Scale *Scale `json:"scale,omitempty"`

	// +optional
	Schedule *Schedule `json:"schedule,omitempty"`
}

type Scheme struct {
	URI base.URI `json:"uri"`
}

type General struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	// +optional
	Rank *uint64 `json:"rank,omitempty"`

	// URN must match ^urn:mpeg:mpegi:nbmp:(2([0-9]{3})):([a-zA-Z0-9_]+)$
	// +optional
	NBMPBrand *base.URI `json:"nbmp-brand,omitempty"`

	// +optional
	PublishedTime *time.Time `json:"published-time,omitempty"`

	// +optional
	Priority *float64 `json:"priority,omitempty"`

	// +optional
	Location *string `json:"location,omitempty"`

	// +optional
	TaskGroup []TaskGroupItem `json:"task-group,omitempty"`

	// +optional
	InputPorts []Port `json:"input-ports"`

	// +optional
	OutputPorts []Port `json:"output-ports"`

	// default is false
	// +optional
	IsGroup bool `json:"is-group"`

	// default is false
	// +optional
	Nonessential bool `json:"nonessential"`

	// +optional
	State *string `json:"state"`
}

type TaskGroupItem struct {
	GroupID string `json:"group-id"`

	TaskID []string `json:"task-id"`

	// default is "distance"
	// +optional
	GroupType *GroupType `json:"group-type,omitempty"`

	// default is "synchronous"
	// +optional
	GroupMode *GroupMode `json:"group-mode,omitempty"`

	// default is false
	// +optional
	NetZero bool `json:"net-zero"`
}

type GroupType string

const (
	DistanceGroupType GroupType = "distance"
	SyncGroupType     GroupType = "sync"
	VirtualGroupType  GroupType = "virtual"
)

type GroupMode string

const (
	SynchronousGroupMode  GroupMode = "synchronous"
	AsynchronousGroupMode GroupMode = "asynchronous"
)

type Port struct {
	PortName string `json:"port-name"`

	Bind PortBinding `json:"bind"`
}

type PortBinding struct {
	StreamID string `json:"stream-id"`

	Name string `json:"name"`

	// +optional
	Keywords []string `json:"keywords,omitempty"`
}

type Input struct {
	// +optional
	MediaParameters []MediaParameter `json:"media-parameters,omitempty"`

	// +optional
	MetadataParameters []MetadataParameter `json:"metadata-parameters,omitempty"`
}

type Output struct {
	// +optional
	MediaParameters []MediaParameter `json:"media-parameters,omitempty"`

	// +optional
	MetadataParameters []MetadataParameter `json:"metadata-parameters,omitempty"`
}

type MediaParameter struct {
	StreamID string `json:"stream-id"`

	Name string `json:"name"`

	Keywords []string `json:"keywords"`

	MimeType string `json:"mime-type"`

	// +optional
	VideoFormat []Parameter `json:"video-format,omitempty"`

	// +optional
	AudioFormat []Parameter `json:"audio-format,omitempty"`

	// +optional
	ImageFormat []Parameter `json:"image-format,omitempty"`

	// +optional
	CodecType *string `json:"codec-type,omitempty"`

	Protocol string `json:"protocol"`

	// default is "push"
	// +optional
	Mode *MediaAccessMode `json:"mode,omitempty"`

	// +optional
	Throughput *uint64 `json:"throughput,omitempty"`

	// +optional
	Buffersize *uint64 `json:"buffersize,omitempty"`

	// +optional
	AvailabilityDuration *uint64 `json:"availability-duration,omitempty"`

	// must be >= 1
	// +optional
	Timeout *uint64 `json:"timeout,omitempty"`

	CachingServerURL base.URI `json:"caching-server-url"`

	// must not be set for Outputs
	// +optional
	CompletionTimeout *uint64 `json:"completion-timeout,omitempty"`
}

type MetadataParameter struct {
	StreamID string `json:"stream-id"`

	Name string `json:"name"`

	Keywords []string `json:"keywords"`

	MimeType string `json:"mime-type"`

	// +optional
	CodecType *string `json:"codec-type,omitempty"`

	Protocol string `json:"protocol"`

	// default is "push"
	// +optional
	Mode *MediaAccessMode `json:"mode,omitempty"`

	// +optional
	MaxSize *uint64 `json:"max-size,omitempty"`

	// +optional
	MinInterval *uint64 `json:"min-interval,omitempty"`

	// +optional
	AvailabilityDuration *uint64 `json:"availability-duration,omitempty"`

	// +optional
	Timeout *uint64 `json:"timeout,omitempty"` // >= 1

	// +optional
	CachingServerURL *base.URI `json:"caching-server-url,omitempty"`

	// +optional
	SchemeURI *base.URI `json:"scheme-uri,omitempty"`

	// must not be set for Outputs
	// +optional
	CompletionTimeout *uint64 `json:"completion-timeout,omitempty"`
}

type MediaAccessMode string

const (
	PushMediaAccessMode MediaAccessMode = "push"
	PullMediaAccessMode MediaAccessMode = "pull"
)

type Processing struct {
	Keywords []string `json:"keywords"`

	Image []ProcessingImage `json:"image"`

	// +optional
	StartTime *time.Time `json:"start-time,omitempty"`

	// +optional
	ConnectionMap []ConnectionMapping `json:"connection-map,omitempty"`

	// +optional
	FunctionRestrictions []FunctionRestriction `json:"function-restrictions,omitempty"`
}

type ProcessingImage struct {
	// default is false
	// +optional
	IsDynamic bool `json:"is-dynamic"`

	URL base.URI `json:"url"`

	// +optional
	StaticImageInfo *StaticImageInfo `json:"static-image-info,omitempty"`

	// +optional
	DynamicImageInfo *DynamicImageInfo `json:"dynamic-image-info,omitempty"`
}

type StaticImageInfo struct {
	OS string `json:"os"`

	Version string `json:"version"`

	Architecture string `json:"architecture"`

	Environment string `json:"environment"`

	// +optional
	PatchURL *string `json:"patch-url,omitempty"`

	// +optional
	PatchScript map[string]interface{} `json:"patch-script,omitempty"`
}

type DynamicImageInfo struct {
	Scheme base.URI `json:"scheme"`

	Information map[string]interface{} `json:"information"`
}

type ConnectionMapping struct {
	ConnectionID string `json:"connection-id"`

	From ConnectionMappingFrom `json:"from"`

	To ConnectionMappingTo `json:"to"`

	// +optional
	Flowcontrol *FlowcontrolRequirement `json:"flowcontrol,omitempty"`

	// default is false
	// +optional
	CoLocated bool `json:"co-located"`

	// default is true (TODO)
	// +optional
	Breakable *bool `json:"breakable,omitempty"`

	// +optional
	OtherParameters []Parameter `json:"other-parameters,omitempty"`
}

type ConnectionMappingPort struct {
	ID string `json:"id"`

	Instance string `json:"instance"`

	PortName string `json:"port-name"`
}

type ConnectionMappingFrom struct {
	ConnectionMappingPort

	// +optional
	OutputRestrictions *Output `json:"output-restrictions,omitempty"`
}

type ConnectionMappingTo struct {
	ConnectionMappingPort

	// +optional
	InputRestrictions *Input `json:"input-restrictions,omitempty"`
}

type FunctionRestriction struct {
	Instance string `json:"instance"`

	// +optional
	General *General `json:"general,omitempty"`

	// +optional
	Processing *Processing `json:"processing,omitempty"`

	// +optional
	Requirements *Requirement `json:"requirements,omitempty"`

	// +optional
	Configuration *Configuration `json:"configuration,omitempty"`

	// +optional
	ClientAssistant *ClientAssistant `json:"client-assistant,omitempty"`

	// +optional
	Failover *Failover `json:"failover,omitempty"`

	// +optional
	Monitoring *Monitoring `json:"monitoring,omitempty"`

	// +optional
	Reporting *Reporting `json:"reporting,omitempty"`

	// +optional
	Notification *Notification `json:"notification,omitempty"`

	// +optional
	Step *Step `json:"step,omitempty"`

	// +optional
	Security *Security `json:"security,omitempty"`

	// +optional
	Blacklist []FunctionBlacklist `json:"blacklist,omitempty"`
}

type FunctionBlacklist string

const (
	RequirementFunctionBlacklist     FunctionBlacklist = "requirement"
	ClientAssistantFunctionBlacklist FunctionBlacklist = "client-assistant"
	FailOverFunctionBlacklist        FunctionBlacklist = "fail-over"
	MonitoringFunctionBlacklist      FunctionBlacklist = "monitoring"
	ReportingFunctionBlacklist       FunctionBlacklist = "reporting"
	NotificationFunctionBlacklist    FunctionBlacklist = "notification"
	SecurityFunctionBlacklist        FunctionBlacklist = "security"
)

type Requirement struct {
	// +optional
	Flowcontrol *FlowcontrolRequirement `json:"flowcontrol,omitempty"`

	// +optional
	Hardware *HardwareRequirement `json:"hardware,omitempty"`

	// +optional
	Security *SecurityRequirement `json:"security,omitempty"`

	// +optional
	WorkflowTask *WorkflowTaskRequirement `json:"workflow-task,omitempty"`

	// +optional
	ResourceEstimators *ResourceEstimatorsRequirement `json:"resource-estimators,omitempty"`
}

type FlowcontrolRequirement struct {
	// +optional
	TypicalDelay *uint64 `json:"typical-delay,omitempty"`

	// +optional
	MinDelay *uint64 `json:"min-delay,omitempty"`

	// +optional
	MaxDelay *uint64 `json:"max-delay,omitempty"`

	// +optional
	MinThroughput *uint64 `json:"min-throughput,omitempty"`

	// +optional
	MaxThroughput *uint64 `json:"max-throughput,omitempty"`

	// +optional
	AveragingWindow *uint64 `json:"averaging-window,omitempty"`
}

type HardwareRequirement struct {
	// +optional
	VCPU *uint64 `json:"vcpu,omitempty"`

	// +optional
	VGPU *uint64 `json:"vgpu,omitempty"`

	// +optional
	RAM *uint64 `json:"ram,omitempty"`

	// +optional
	Disk *uint64 `json:"disk,omitempty"`

	// +optional
	Placement *HardwareRequirementPlacement `json:"placement,omitempty"`
}

type HardwareRequirementPlacement string // pattern (^[A-Z]{2}$)|(^[A-Z]{2}-.*)

type SecurityRequirement struct {
	// default is false
	// +optional
	TLS bool `json:"tls"`

	// default is false
	// +optional
	IPsec bool `json:"ipsec"`

	// default is false
	// +optional
	CENC bool `json:"cenc"`
}

type WorkflowTaskRequirement struct {
	// default is false
	// +optional
	FunctionFusible bool `json:"function-fusible"`

	// default is false
	// +optional
	FunctionEnhancable bool `json:"function-enhancable"`

	// default is "streaming"
	// +optional
	ExecutionMode *ExecutionMode `json:"execution-mode,omitempty"`

	// +optional
	Proximity []TaskProximityRequirement `json:"proximity,omitempty"`

	// +optional
	ProximityEquation *TaskProximityEquation `json:"proximity-equation,omitempty"`

	// +optional
	SplitEfficiency *TaskSplitEfficiency `json:"split-efficiency,omitempty"`
}

type ExecutionMode string

const (
	StreamingExecutionMode ExecutionMode = "streaming"
	StepExecutionMode      ExecutionMode = "step"
	HybridExecutionMode    ExecutionMode = "hybrid"
)

type TaskProximityRequirement struct {
	OtherTaskID string `json:"other-task-id"`

	// must be >= -1
	Distance int64 `json:"distance"`
}

type TaskProximityEquation struct {
	DistanceParameters []Variable `json:"distance-parameters"`

	DistanceEquation string `json:"distance-equation"`
}

type TaskSplitEfficiency struct {
	// default is "pnorm"
	// +optional
	SplitNorm *TaskSplitEfficiencyNorm `json:"split-norm,omitempty"`

	// default is "2"
	// +optional
	SplitEquation *string `json:"split-equation,omitempty"`

	// +optional
	SplitResult *uint64 `json:"split-result,omitempty"`
}

type TaskSplitEfficiencyNorm string

const (
	PnormTaskSplitEfficiencyNorm  TaskSplitEfficiencyNorm = "pnorm"
	CustomTaskSplitEfficiencyNorm TaskSplitEfficiencyNorm = "custom"
)

type ResourceEstimatorsRequirement struct {
	DefaultValues []Value `json:"default-values"`

	// +optional
	ComputationalEstimator *string `json:"computational-estimator,omitempty"`

	// +optional
	MemoryEstimator *string `json:"memory-estimator,omitempty"`

	// +optional
	BandwidthEstimator *string `json:"bandwidth-estimator,omitempty"`
}

type Value struct {
	Name string `json:"name"`

	Value string `json:"value"`
}

type Configuration struct {
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Name string `json:"name"`

	ID int64 `json:"id"`

	// TODO: spelling
	// +optional
	Description *string `json:"discription,omitempty"`

	Datatype Datatype `json:"datatype"`

	// List of ids that shall exist for this parameter to become valid
	// +optional
	Conditions []int64 `json:"conditions,omitempty"`

	// List of configuration ids that shall not exist for this parameter to become valid
	// +optional
	Exclusions []int64 `json:"exclusions,omitempty"`

	// must be set for non-array datatypes
	// must be nil for array datatypes
	// +optional
	Values []ParameterValue `json:"values,omitempty"`

	// must be nil for non-array datatypes
	// must be set for array datatypes
	// +optional
	Schema map[string]interface{} `json:"schema,omitempty"`
}

func (p *Parameter) UnmarshalJSON(data []byte) error {
	type PartialParameter struct {
		Name        string                 `json:"name"`
		ID          int64                  `json:"id"`
		Description *string                `json:"discription,omitempty"` // TODO: spelling
		Datatype    Datatype               `json:"datatype"`
		Conditions  []int64                `json:"conditions,omitempty"`
		Exclusions  []int64                `json:"exclusions,omitempty"`
		Values      []json.RawMessage      `json:"values,omitempty"`
		Schema      map[string]interface{} `json:"schema,omitempty"`
	}

	pp := PartialParameter{}
	err := json.Unmarshal(data, &pp)
	if err != nil {
		return err
	}

	p.Name = pp.Name
	p.ID = pp.ID
	p.Description = pp.Description
	p.Datatype = pp.Datatype
	p.Conditions = pp.Conditions
	p.Exclusions = pp.Exclusions
	p.Schema = pp.Schema
	p.Values = make([]ParameterValue, len(pp.Values))

	for i, rawVal := range pp.Values {
		var val ParameterValue
		switch p.Datatype {
		case BooleanDatatype:
			val = &BooleanParameterValue{}
		case IntegerDatatype:
			val = &IntegerParameterValue{}
		case NumberDatatype:
			val = &NumberParameterValue{}
		case StringDatatype:
			val = &StringParameterValue{}
		case ArrayDatatype:
			return errors.New("Parameter.UnmarshalJSON: specifying values is invalid for array datatype")
		default:
			return errors.New("Parameter.UnmarshalJSON: values datatype unspecified or unknown")
		}

		err = json.Unmarshal(rawVal, val)
		if err != nil {
			return err
		}
		p.Values[i] = val
	}

	return nil
}

var (
	_ json.Unmarshaler = &Parameter{}
)

type Datatype string

const (
	BooleanDatatype Datatype = "boolean"
	IntegerDatatype Datatype = "integer"
	NumberDatatype  Datatype = "number"
	StringDatatype  Datatype = "string"
	ArrayDatatype   Datatype = "array"
)

type ParameterValue interface{}

var (
	_ ParameterValue = &BooleanParameterValue{}
	_ ParameterValue = &IntegerParameterValue{}
	_ ParameterValue = &NumberParameterValue{}
	_ ParameterValue = &StringParameterValue{}
)

type BooleanParameterValue struct {
	Name string `json:"name"`

	ID int64 `json:"id"`

	Restrictions bool `json:"restrictions"`
}

type IntegerParameterValue struct {
	Name string `json:"name"`

	ID int64 `json:"id"`

	Restrictions IntegerParameterValueRestrictions `json:"restrictions"`
}

type IntegerParameterValueRestrictions struct {
	// +optional
	MinValue *int64 `json:"min-value,omitempty"`

	// +optional
	MaxValue *int64 `json:"max-value,omitempty"`

	// +optional
	Increment *int64 `json:"increment,omitempty"`
}

type NumberParameterValue struct {
	Name string `json:"name"`

	ID int64 `json:"id"`

	Restrictions NumberParameterValueRestrictions `json:"restrictions"`
}

type NumberParameterValueRestrictions struct {
	// +optional
	MinValue *float64 `json:"min-value,omitempty"`

	// +optional
	MaxValue *float64 `json:"max-value,omitempty"`

	// +optional
	Increment *float64 `json:"increment,omitempty"`
}

type StringParameterValue struct {
	Name string `json:"name"`

	ID int64 `json:"id"`

	Restrictions []string `json:"restrictions"`
}

type StartupDelay struct {
	StartupDelayValue uint64 `json:"startup-delay-value"`
}

type ClientAssistant struct {
	// default is false
	ClientAssistanceFlag bool `json:"client-assistance-flag"`

	// +optional
	MeasurementCollectionList map[string]interface{} `json:"measurement-collection-list,omitempty"`

	// +optional
	SourceAssistanceInformation map[string]interface{} `json:"source-assistance-information,omitempty"`
}

type Failover struct {
	FailoverMode FailoverMode `json:"failover-mode"`

	FailoverDelay uint64 `json:"failover-delay"`

	// +optional
	BackupDeploymentURL *base.URI `json:"backup-deployment-url,omitempty"`

	// +optional
	PersistenceURL *base.URI `json:"persistence-url,omitempty"`

	// +optional
	PersistenceInterval *uint64 `json:"persistence-interval,omitempty"`
}

type FailoverMode string

const (
	RestartImmediatelyFailoverMode        FailoverMode = "restart-immediately"
	RestartWithDelayFailoverMode          FailoverMode = "restart-with-delay"
	ContinueWithLastGoodStateFailoverMode FailoverMode = "continue-with-last-good-state"
	ExecuteBackupDeploymentFailoverMode   FailoverMode = "execute-backup-deployment"
	ExitFailoverMode                      FailoverMode = "exit"
)

type Event struct {
	// +optional
	Name *string `json:"name,omitempty"`

	// +optional
	Definition *string `json:"definition,omitempty"`

	// +optional
	URL *base.URI `json:"url,omitempty"`
}

type Variable struct {
	Name string `json:"name"`

	Definition string `json:"definition"`

	Unit string `json:"unit"`

	VarType VariableType `json:"var-type"`

	// +optional
	Value *string `json:"value,omitempty"`

	// +optional
	Min *int64 `json:"min,omitempty"`

	// +optional
	Max *int64 `json:"max,omitempty"`

	// +optional
	URL *base.URI `json:"url,omitempty"`

	// +optional
	Children []Variable `json:"children,omitempty"`
}

type VariableType string

const (
	StringVariableType  VariableType = "string"
	IntegerVariableType VariableType = "integer"
	FloatVariableType   VariableType = "float"
	BooleanVariableType VariableType = "boolean"
	NumberVariableType  VariableType = "number"
)

type Monitoring struct {
	// +optional
	Event []Event `json:"event,omitempty"`

	// +optional
	Variable []Variable `json:"variable,omitempty"`

	// +optional
	SystemEvents []map[string]interface{} `json:"system-events,omitempty"`

	// +optional
	SystemVariables []map[string]interface{} `json:"system-variables,omitempty"`
}

type Reporting struct {
	// +optional
	Event []Event `json:"event,omitempty"`

	// +optional
	Variable []Variable `json:"variable,omitempty"`

	// +optional
	SystemEvents []map[string]interface{} `json:"system-events,omitempty"`

	// +optional
	SystemVariables []map[string]interface{} `json:"system-variables,omitempty"`

	ReportType string `json:"report-type"`

	ReportingInterval uint64 `json:"reporting-interval"`

	ReportStartTime time.Time `json:"report-start-time"`

	URL base.URI `json:"url"`

	// default is "HTTP POST"
	DeliveryMethod DeliveryMethod `json:"delivery-method"`
}

type DeliveryMethod string

const (
	HTTP_POSTDeliveryMethod DeliveryMethod = "HTTP POST"
)

type Notification struct {
	// +optional
	Event []Event `json:"event,omitempty"`

	// +optional
	Variable []Variable `json:"variable,omitempty"`

	// +optional
	SystemEvents []map[string]interface{} `json:"system-events,omitempty"`

	// +optional
	SystemVariables []map[string]interface{} `json:"system-variables,omitempty"`

	NotificationTime time.Time `json:"notification-time"`

	SeverityLevel string `json:"severity-level"`

	NotificationType []NotificationType `json:"notification-type"`

	URLs []base.URI `json:"urls"`

	// default is 0
	// +optional
	NotificationInterval uint64 `json:"notification-interval"`
}

type NotificationType string

const (
	CongestionNotificationType  NotificationType = "congestion"
	ApplicationNotificationType NotificationType = "application"
	SystemNotificationType      NotificationType = "system"
)

type Assertion struct {
	MinPriority       uint64 `json:"min-priority"`
	MinPriorityAction string `json:"min-priority-action"`

	// default is false
	// +optional
	SupportVerification bool `json:"support-verification"`

	// +optional
	VerificationAcknowledgement *string `json:"verification-acknowledgement,omitempty"`

	// +optional
	Certificate *string `json:"certificate,omitempty"`

	Assertion []AssertionItem `json:"assertion"`
}

type AssertionItem struct {
	Name string `json:"name"`

	ValuePredicate AssertionValuePredicate `json:"value-predicate"`
}

type AssertionValuePredicate struct {
	EvaluationCondition AssertionValuePredicateEvaluationCondition `json:"evaluation-condition"`

	CheckValue map[string]interface{} `json:"check-value"`

	Aggregation AssertionValuePredicateAggregation `json:"aggregation"`

	// +optional
	Offset *string `json:"offset,omitempty"`

	Priority uint64 `json:"priority"`

	Action AssertionValuePredicateAction `json:"action"`

	// +optional
	ActionParameters []string `json:"action-parameters,omitempty"`
}

type AssertionValuePredicateEvaluationCondition string

const (
	QualityAssertionValuePredicateEvaluationCondition       AssertionValuePredicateEvaluationCondition = "quality"
	ComputationalAssertionValuePredicateEvaluationCondition AssertionValuePredicateEvaluationCondition = "computational"
	InputAssertionValuePredicateEvaluationCondition         AssertionValuePredicateEvaluationCondition = "input"
	OutputAssertionValuePredicateEvaluationCondition        AssertionValuePredicateEvaluationCondition = "output"
)

type AssertionValuePredicateAggregation string

const (
	SumAssertionValuePredicateAggregation AssertionValuePredicateAggregation = "sum"
	MinAssertionValuePredicateAggregation AssertionValuePredicateAggregation = "min"
	MaxAssertionValuePredicateAggregation AssertionValuePredicateAggregation = "max"
	AvgAssertionValuePredicateAggregation AssertionValuePredicateAggregation = "avg"
)

type AssertionValuePredicateAction string

const (
	RebuildAssertionValuePredicateAction AssertionValuePredicateAction = "rebuild"
	RestartAssertionValuePredicateAction AssertionValuePredicateAction = "restart"
	WaitAssertionValuePredicateAction    AssertionValuePredicateAction = "wait"
)

type Request struct {
	RequestID uint64 `json:"request-id"`

	// +optional
	Priority *uint64 `json:"priority,omitempty"`

	TaskID string `json:"task-id"`
}

type Acknowledge struct {
	Status AcknowledgeStatus `json:"status"`

	// +optional
	Unsupported []string `json:"unsupported,omitempty"`

	// +optional
	Failed []string `json:"failed,omitempty"`

	// +optional
	Partial []string `json:"partial,omitempty"`
}

type AcknowledgeStatus string

const (
	FulfilledAcknowledgeStatus          AcknowledgeStatus = "fulfilled"
	FailedAcknowledgeStatus             AcknowledgeStatus = "failed"
	NotSupportedAcknowledgeStatus       AcknowledgeStatus = "not-supported"
	PartiallyFulfilledAcknowledgeStatus AcknowledgeStatus = "partially-fulfilled"
)

type Repository struct {
	// default is "available"
	// +optional
	Mode *RepositoryMode `json:"mode,omitempty"`

	Location []RepositoryLocation `json:"location"`
}

type RepositoryMode string

const (
	StrictRepositoryMode    RepositoryMode = "strict"
	PreferredRepositoryMode RepositoryMode = "preferred"
	AvailableRepositoryMode RepositoryMode = "available"
)

type RepositoryLocation struct {
	URL base.URI `json:"url"`

	Name string `json:"name"`

	Description string `json:"description"`
}

type Security struct {
	Name string `json:"name"`

	// default is "data"
	// +optional
	Scope *SecurityScope `json:"scope,omitempty"`

	AuthenticationMethod string `json:"authentication-method"`

	// +optional
	AuthorityURL *base.URI `json:"authority-url,omitempty"`

	// +optional
	Certificate *string `json:"certificate,omitempty"`

	// +optional
	AuthToken *string `json:"auth-token,omitempty"`

	// +optional
	ClientGrants *string `json:"client-grants,omitempty"`

	// +optional
	AuthTokenExpires *time.Time `json:"auth-token-expires,omitempty"`

	// +optional
	AuthTokenRenew *string `json:"auth-token-renew,omitempty"`

	// default is false
	// +optional
	AuthTokenRotation bool `json:"auth-token-rotation"`
}

type SecurityScope string

const (
	DataSecurityScope     SecurityScope = "data"
	FunctionSecurityScope SecurityScope = "function"
	TaskSecurityScope     SecurityScope = "task"
)

type Step struct {
	// default is "stream"
	// +optional
	StepMode *StepMode `json:"step-mode,omitempty"`

	// default is false
	// +optional
	VariableDuration bool `json:"variable-duration"`

	// +optional
	SegmentDuration *uint64 `json:"segment-duration,omitempty"`

	// default is false
	// +optional
	SegmentLocation bool `json:"segment-location"`

	// default is false
	// +optional
	SegmentSequence bool `json:"segment-sequence"`

	// +optional
	SegmentMetadataSupportedFormats []SegmentMetadataSupportedFormat `json:"segment-metadata-supported-formats,omitempty"`

	// +optional
	OperatingUnits *uint64 `json:"operating-units,omitempty"`

	// +optional
	TemporalOverlap *uint64 `json:"temporal-overlap,omitempty"`

	// +optional
	NumberOfDimensions *uint64 `json:"number-of-dimensions,omitempty"`

	// elements must be >= 1
	// +optional
	HigherDimensionSegmentDivisors []uint64 `json:"higher-dimension-segment-divisors,omitempty"`

	// +optional
	HigherDimensionsDescriptions []HigherDimensionsDescription `json:"higher-dimensions-descriptions,omitempty"`

	// +optional
	HigherDimensionsSegmentOrder []uint64 `json:"higher-dimensions-segment-order,omitempty"`

	// +optional
	HigherDimensionsOverlap []uint64 `json:"higher-dimension-overlap,omitempty"`

	// elements must be >= 1
	// +optional
	HigherDimensionsOperationUnits []uint64 `json:"higher-dimension-operation-units,omitempty"`
}

type StepMode string

const (
	StreamStepMode    StepMode = "stream"
	StatefulStepMode  StepMode = "stateful"
	StatelessStepMode StepMode = "stateless"
)

type SegmentMetadataSupportedFormat string

const (
	NBMPLocationBytestream2022SegmentMetadataSupportedFormat SegmentMetadataSupportedFormat = "nbmp-location-bytestream-2022"
	NBMPSequenceBytestream2022SegmentMetadataSupportedFormat SegmentMetadataSupportedFormat = "nbmp-sequence-bytestream-2022"
	NBMPLocationJSON2022SegmentMetadataSupportedFormat       SegmentMetadataSupportedFormat = "nbmp-location-json-2022"
	NBMPSequenceJSON2022SegmentMetadataSupportedFormat       SegmentMetadataSupportedFormat = "nbmp-sequence-json-2022"
)

type HigherDimensionsDescription string

const (
	WidthHigherDimensionsDescription  HigherDimensionsDescription = "width"
	HeightHigherDimensionsDescription HigherDimensionsDescription = "height"
	RGBHigherDimensionsDescription    HigherDimensionsDescription = "RGB"
	DepthHigherDimensionsDescription  HigherDimensionsDescription = "depth"
	YUVHigherDimensionsDescription    HigherDimensionsDescription = "YUV"
	V_PCCHigherDimensionsDescription  HigherDimensionsDescription = "V-PCC"
)

type Capabilities struct {

	// +optional
	ResourceAvailability []ResourceAvailabilityItem `json:"resource-availability,omitempty"`

	// +optional
	Placement *HardwareRequirementPlacement `json:"placement,omitempty"`

	// +optional
	Location *string `json:"location,omitempty"`

	// +optional
	Repository *Repository `json:"repository,omitempty"`

	// +optional
	Functions []Function `json:"functions,omitempty"`

	// +optional
	Connectivity []CapabilityConnectivity `json:"connectivity,omitempty"`

	// default is true (TODO)
	// +optional
	PersistencyCapabilities *bool `json:"persistency-capabilities,omitempty"`

	// default is false
	// +optional
	SecurePersistency bool `json:"secure-persistency"`

	// +optional
	PersistenceStorageURL []base.URI `json:"persistence-storage-url,omitempty"`
}

type ResourceAvailabilityItem struct {
	// +optional
	Key *ResourceAvailabilityItemKey `json:"key,omitempty"`

	// +optional
	AbsoluteValue *uint64 `json:"absolute-value,omitempty"`

	// must be >= 0 <= 100
	// +optional
	Availability *uint64 `json:"availability,omitempty"`
}

type ResourceAvailabilityItemKey string

const (
	VCPUResourceAvailabilityItemKey  ResourceAvailabilityItemKey = "vcpu"
	VGPUResourceAvailabilityItemKey  ResourceAvailabilityItemKey = "vgpu"
	RAMResourceAvailabilityItemKey   ResourceAvailabilityItemKey = "ram"
	DiskResourceAvailabilityItemKey  ResourceAvailabilityItemKey = "disk"
	PowerResourceAvailabilityItemKey ResourceAvailabilityItemKey = "power"
)

type CapabilityConnectivity struct {
	ID string `json:"id"`

	// +optional
	URL *base.URI `json:"url,omitempty"`

	// +optional
	Forward *CapabilityConnectivityProperties `json:"forward,omitempty"`

	// +optional
	Return *CapabilityConnectivityProperties `json:"return,omitempty"`
}

type CapabilityConnectivityProperties struct {
	// +optional
	MinDelay *uint64 `json:"min-delay,omitempty"`

	// +optional
	MaxThroughput *uint64 `json:"max-throughput,omitempty"`

	// +optional
	AveragingWindow *uint64 `json:"averaging-window,omitempty"`
}

type Scale struct {
	ID string `json:"id"`

	// +optional
	Description *string `json:"description,omitempty"`

	// +optional
	ScalingType *ScalingType `json:"scaling-type,omitempty"`

	// must be >= 1
	// default is 1
	// +optional
	ScalingFactor *uint64 `json:"scaling-factor,omitempty"`

	// default is "failed"
	Status ScalingStatus `json:"status"`

	// +optional
	TargetID *string `json:"target-id,omitempty"`
}

type ScalingType string

const (
	MPEScalingType        ScalingType = "MPE"
	SplitMergeScalingType ScalingType = "split-merge"
)

type ScalingStatus string

const (
	CapabilitiesScalingStatus ScalingStatus = "capabilities"
	ConsiderScalingStatus     ScalingStatus = "consider"
	RequestScalingStatus      ScalingStatus = "request"
	PassedScalingStatus       ScalingStatus = "passed"
	FailedScalingStatus       ScalingStatus = "failed"
)

type Schedule struct {
	ID string `json:"id"`

	// +optional
	Description *string `json:"description,omitempty"`

	// default is "duration"
	// +optional
	ScheduleType *ScheduleType `json:"schedule-type,omitempty"`

	// +optional
	ScheduleTable []ScheduleTableItem `json:"schedule-table,omitempty"`

	// elements must be >= 1
	// elements default is 1
	// +optional
	NumberOfSegments []int64 `json:"number-of-segments,omitempty"`

	// default is false
	// +optional
	Loop bool `json:"loop"`

	// default is "failed"
	Status ScheduleStatus `json:"status,omitempty"`
}

type ScheduleType string

const (
	DurationScheduleType ScheduleType = "duration"
	SegmentScheduleType  ScheduleType = "segment"
)

type ScheduleStatus string

const (
	CapabilitiesScheduleStatus ScheduleStatus = "capabilities"
	ConsiderScheduleStatus     ScheduleStatus = "consider"
	RequestScheduleStatus      ScheduleStatus = "request"
	PassedScheduleStatus       ScheduleStatus = "passed"
	FailedScheduleStatus       ScheduleStatus = "failed"
)

type ScheduleTableItem struct {
	TaskID string `json:"task-id"`

	// +optional
	StartTime *string `json:"start-time,omitempty"`

	// default is 1
	// +optional
	Duration *uint64 `json:"duration,omitempty"`

	// default is 1
	// +optional
	Timescale *uint64 `json:"timescale,omitempty"`
}

type SegmentLocation struct {
	Scale []uint64 `json:"scale"`

	Length []uint64 `json:"length"`

	Size uint64 `json:"size"`
}

type SegmentSequence struct {
	Sequence []uint64 `json:"sequence,omitempty"`

	Size uint64 `json:"size,omitempty"`
}
