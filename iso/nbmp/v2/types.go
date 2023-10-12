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

const (
	// TODO: update this with the new spec
	SchemaURI = "urn:mpeg:mpegi:nbmp:2023"
)

const (
	WorkflowDescriptionDocumentMIMEType = "application/mpeg-nbmp-wdd+json"
	WorkflowDescriptionDocumentExt      = "wdd"

	TaskDescriptionDocumentMIMEType = "application/mpeg-nbmp-tdd+json"
	TaskDescriptionDocumentExt      = "tdd"

	FunctionDescriptionDocumentMIMEType = "application/mpeg-nbmp-fdd+json"
	FunctionDescriptionDocumentExt      = "fdd"
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

// This descriptor provides a scheme identifier to identify the base scheme used for its descriptors.
type Scheme struct {
	// identifies the scheme for this document
	// It shall be a valid URI according to IETF RFC 3986.
	URI base.URI `json:"uri"`
}

// This descriptor provides general details about the underlying resource.
type General struct {
	// unique string in the scope of repository/workflow of the resource
	ID string `json:"id"`

	// name for identifying the resource
	Name string `json:"name"`

	// a human-readable description for the resource
	Description string `json:"description"`

	// rank of function/function group among functions with the same functionality
	// A higher number means a higher rank.
	// +optional
	Rank *uint64 `json:"rank,omitempty"`

	// URN must match ^urn:mpeg:mpegi:nbmp:(2([0-9]{3})):([a-zA-Z0-9_]+)$
	// TODO: Annex C (informative) describes urn:mpeg:mpegi:nbmp:... as brand name to identify MPEG compatible functions.
	//       Can this be interpreted that non-compatible custom functions should use a different namespace thus not
	//       matching the pattern?
	// +optional
	NBMPBrand *base.URI `json:"nbmp-brand,omitempty"`

	// date and time of publication of this resource
	// +optional
	PublishedTime *time.Time `json:"published-time,omitempty"`

	// priority information for the resource
	// shall not be present for function description documents
	// +optional
	Priority *uint64 `json:"priority,omitempty"`

	// +optional
	Location *string `json:"location,omitempty"`

	// +optional
	TaskGroup []TaskGroupItem `json:"task-group,omitempty"`

	// input-ports and output-ports objects specify the endpoints where the data is communicated from NBMP sources to
	// tasks; and between tasks; and from tasks to NBMP sinks.
	//
	// shall not be present for workflow description documents
	// TODO: the JSON schema requires this field; the NBMP specification forbids this field for workflow description
	//       documents ("shall not"); the NBMP specification requires this parameter for general descriptors. All are
	//       normative.
	InputPorts []Port `json:"input-ports,omitempty"`

	// input-ports and output-ports objects specify the endpoints where the data is communicated from NBMP sources to
	// tasks; and between tasks; and from tasks to NBMP sinks.
	//
	// shall not be present for workflow description documents
	// TODO: the JSON schema requires this field; the NBMP specification forbids this field for workflow description
	//       documents ("shall not"); the NBMP specification requires this parameter for general descriptors. All are
	//       normative.
	OutputPorts []Port `json:"output-ports,omitempty"`

	// value ‘true’ indicates containing descriptor describes a function group or task workflow
	// If the value is ‘true’, a connection-map object shall exist in this description.
	// default is false
	// this is a pointer as this field should not be present in workflow description documents
	// +optional
	IsGroup *bool `json:"is-group,omitempty"`

	// default is false
	// +optional
	Nonessential bool `json:"nonessential,omitempty"`

	// +optional
	State *State `json:"state"`
}

type State string

var (
	InstantiatedState State = "instantiated"
	IdleState         State = "idle"
	RunningState      State = "running"
	InErrorState      State = "in-error"
	DestroyedState    State = "destroyed"
)

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

var (
	DistanceGroupType GroupType = "distance"
	SyncGroupType     GroupType = "sync"
	VirtualGroupType  GroupType = "virtual"
)

type GroupMode string

var (
	SynchronousGroupMode  GroupMode = "synchronous"
	AsynchronousGroupMode GroupMode = "asynchronous"
)

type Port struct {
	// unique string among all port-names of this resource defining the logic name for input or output
	PortName string `json:"port-name"`

	// The bind object specifies how to associate a port name to a stream, either input or output. For NBMP functions,
	// they provide static information about the port names and their binding data formats and protocols. For NBMP tasks,
	// they provide information about the needs for connections between ports and input and output streams by NBMP
	// workflow manager.
	Bind PortBinding `json:"bind"`
}

type PortBinding struct {
	// TODO: this is required in the 2nd edition JSON schema but not in the 2020 NBMP specification. Forcing this field
	//       probably does not make sens for function descriptions?
	// +optional
	StreamID *string `json:"stream-id,omitempty"`

	Name string `json:"name"`

	// +optional
	Keywords []string `json:"keywords,omitempty"`
}

// This descriptor provides the parameters of the underlying resource’s inputs. The input descriptor consists of two
// arrays of objects: one for the media inputs and one for metadata inputs.
type Input struct {
	// +optional
	MediaParameters []MediaParameter `json:"media-parameters,omitempty"`

	// +optional
	MetadataParameters []MetadataParameter `json:"metadata-parameters,omitempty"`
}

// This descriptor provides the parameters of the underlying resource’s outputs. The output descriptor consists of two
// arrays of objects: one for the media outputs and one for metadata outputs.
type Output struct {
	// +optional
	MediaParameters []MediaParameter `json:"media-parameters,omitempty"`

	// +optional
	MetadataParameters []MetadataParameter `json:"metadata-parameters,omitempty"`
}

type MediaParameter struct {
	// unique identifier, with the scope of function or task or workflow, to identify the media or metadata stream
	//
	// For functions, it is defined in the function descriptor. For tasks, it is assigned by the workflow manager. For
	// workflows, it is assigned by the NBMP source.
	//
	// TODO: the JSON schema requires this field; the NBMP specification forbids this field for function description
	//       documents ("shall not"); the NBMP specification requires this parameter for Input/Output descriptors. All are
	//       normative.
	//
	// shall not be present for function description documents
	StreamID string `json:"stream-id"`

	// string name assigned to this input
	//
	// For functions, it is defined in the function descriptor. For tasks, it is assigned by the workflow manager. For
	// workflows, it is assigned by the NBMP source.
	Name string `json:"name"`

	// list of keywords describing this input properties
	//
	// The keyword should be human-readable.
	Keywords []string `json:"keywords"`

	// MIME media type of media or metadata in IANA registry
	MimeType string `json:"mime-type"`

	// format of video
	//
	// The parameter list is defined using generic parameter representation of subclause 9.22.1.
	// +optional
	VideoFormat []Parameter `json:"video-format,omitempty"`

	// format of audio
	//
	// The parameter list is defined using generic parameter representation of subclause 9.22.1.
	// +optional
	AudioFormat []Parameter `json:"audio-format,omitempty"`

	// format of image
	//
	// The parameter list is defined using generic parameter representation of subclause 9.22.1.
	// +optional
	ImageFormat []Parameter `json:"image-format,omitempty"`

	// 'codecs' and 'profiles' parameters, as defined in IETF RFC 6381
	// +optional
	CodecType *string `json:"codec-type,omitempty"`

	// protocol for delivery of or access to media including protocol parameters such as port number(s)
	//
	// Ingest protocol for timed metadata including protocol parameters such as the port number(s). Example: HTTP. When
	// the workflow manager receives this information, it takes the responsibility of returning back with the protocol
	// endpoint information of the appropriate media processing entity to the media source so the media source can ingest
	// metadata using that protocol.
	//
	// NOTE This is only applicable for media and timed metadata.
	Protocol string `json:"protocol"`

	// default is "push"
	// +optional
	Mode *MediaAccessMode `json:"mode,omitempty"`

	// maximum accepted throughput by this resource
	// +optional
	Throughput *uint64 `json:"throughput,omitempty"`

	// minimum input buffer size
	// +optional
	Buffersize *uint64 `json:"buffersize,omitempty"`

	// +optional
	AvailabilityDuration *uint64 `json:"availability-duration,omitempty"`

	// must be >= 1
	// +optional
	Timeout *uint64 `json:"timeout,omitempty"`

	// URL (according to IETF RFC 3986) of the server where the media will be sent to/from or the location to/from where
	// the media can be fetched to/from
	//
	// NOTE When this parameter is missing for a workflow, the workflow manager can assign origination/destination
	// information of the appropriate media processing entity to the media source so the media source/sink can ingest
	// media using that protocol.
	CachingServerURL base.URI `json:"caching-server-url"`

	// must not be set for Outputs
	// +optional
	CompletionTimeout *uint64 `json:"completion-timeout,omitempty"`
}

type MetadataParameter struct {
	// unique identifier, with the scope of function or task or workflow, to identify the media or metadata stream
	//
	// For functions, it is defined in the function descriptor. For tasks, it is assigned by the workflow manager. For
	// workflows, it is assigned by the NBMP source.
	//
	// TODO: the JSON schema requires this field; the NBMP specification forbids this field for function description
	//       documents ("shall not"); the NBMP specification requires this parameter for Input/Output descriptors. All are
	//       normative.
	//
	// shall not be present for function description documents
	StreamID string `json:"stream-id"`

	// string name assigned to this input
	//
	// For functions, it is defined in the function descriptor. For tasks, it is assigned by the workflow manager. For
	// workflows, it is assigned by the NBMP source.
	Name string `json:"name"`

	// list of keywords describing this input properties
	//
	// The keyword should be human-readable.
	Keywords []string `json:"keywords"`

	// MIME media type of media or metadata in IANA registry
	MimeType string `json:"mime-type"`

	// 'codecs' and 'profiles' parameters, as defined in IETF RFC 6381
	// +optional
	CodecType *string `json:"codec-type,omitempty"`

	// protocol for delivery of or access to media including protocol parameters such as port number(s)
	//
	// Ingest protocol for timed metadata including protocol parameters such as the port number(s). Example: HTTP. When
	// the workflow manager receives this information, it takes the responsibility of returning back with the protocol
	// endpoint information of the appropriate media processing entity to the media source so the media source can ingest
	// metadata using that protocol.
	//
	// NOTE This is only applicable for media and timed metadata.
	Protocol string `json:"protocol"`

	// default is "push"
	// +optional
	Mode *MediaAccessMode `json:"mode,omitempty"`

	// maximum size of metadata in each fetch or push accepted by this input
	// +optional
	MaxSize *uint64 `json:"max-size,omitempty"`

	// minimum interval between two fetch or push accepted by this input
	// +optional
	MinInterval *uint64 `json:"min-interval,omitempty"`

	// +optional
	AvailabilityDuration *uint64 `json:"availability-duration,omitempty"`

	// must be >= 1
	// +optional
	Timeout *uint64 `json:"timeout,omitempty"`

	// URL (according to IETF RFC 3986) of the server where the media will be sent from or the location from where the
	// media can be fetched from
	//
	// NOTE When this parameter is missing for a workflow, the workflow manager can assign origination information of the
	// appropriate media processing entity to the media source so the media source can ingest media using that protocol.
	// +optional
	CachingServerURL *base.URI `json:"caching-server-url,omitempty"`

	// URL (according to IETF RFC 3986) or scheme identifier of metadata
	//
	// The schema-uri refers to a metadata-dictionary object that consists of a set of parameter-value pairs. The
	// parameters’ names, data types and value ranges are defined by the metadata scheme owner.
	//
	// +optional
	SchemeURI *base.URI `json:"scheme-uri,omitempty"`

	// must not be set for Outputs
	// +optional
	CompletionTimeout *uint64 `json:"completion-timeout,omitempty"`
}

type MediaAccessMode string

var (
	PushMediaAccessMode MediaAccessMode = "push"
	PullMediaAccessMode MediaAccessMode = "pull"
)

// This descriptor provides high-level details about the requested media processing of a workflow by listing the set of
// tasks to be performed on the input media data.
type Processing struct {
	// list of keywords that can be used to execute a search in the function repository
	// TODO: the JSON schema requires this field; the NBMP specification forbids this field for task description
	//       documents ("shall not"); the NBMP specification requires this parameter for processing descriptors. All are
	//       normative.
	Keywords []string `json:"keywords"`

	Image []ProcessingImage `json:"image"`

	// resource’s start time
	// shall not be present for function description documents
	// +optional
	StartTime *time.Time `json:"start-time,omitempty"`

	// The array of connection-map object provides a description of the media workflow DAG, i.e. the connection
	// information between different tasks in the graph. Each element in this array represents an edge in the DAG.
	//
	// shall not be present for task description documents
	//
	// +optional
	ConnectionMap []ConnectionMapping `json:"connection-map,omitempty"`

	// shall not be present for task description documents
	// +optional
	FunctionRestrictions []FunctionRestriction `json:"function-restrictions,omitempty"`
}

type ProcessingImage struct {
	// flag indicating whether the image is static or dynamic
	//
	// A value of ‘true’ indicates the image is built dynamically.
	// default is false
	// +optional
	IsDynamic bool `json:"is-dynamic"`

	// pointer to the resource implementation, according to IETF RFC 3986
	// TODO: the JSON schema requires this field; the NBMP specification forbids this field for workflow description
	//       documents ("shall not"); the NBMP specification requires this parameter for processing descriptors. All are
	//       normative.
	URL base.URI `json:"url"`

	// +optional
	StaticImageInfo *StaticImageInfo `json:"static-image-info,omitempty"`

	// +optional
	DynamicImageInfo *DynamicImageInfo `json:"dynamic-image-info,omitempty"`
}

type StaticImageInfo struct {
	// operation system
	OS string `json:"os"`

	// version number of operation system
	Version string `json:"version"`

	// hardware architecture
	Architecture string `json:"architecture"`

	// environment
	Environment string `json:"environment"`

	// URL (according to IETF RFC 3986) defining the patching scheme for this image
	// +optional
	PatchURL *string `json:"patch-url,omitempty"`

	// +optional
	PatchScript map[string]interface{} `json:"patch-script,omitempty"`
}

type DynamicImageInfo struct {
	// URL (according to IETF RFC 3986) defining information object scheme or information needed for dynamic build
	Scheme base.URI `json:"scheme"`

	Information map[string]interface{} `json:"information"`
}

type ConnectionMapping struct {
	ConnectionID string `json:"connection-id"`

	From ConnectionMappingPort `json:"from"`

	To ConnectionMappingPort `json:"to"`

	// +optional
	Flowcontrol *FlowcontrolRequirement `json:"flowcontrol,omitempty"`

	// specifies the deployment of the 2 connected tasks
	//
	// When the value is True, the 2 tasks shall be deployed into the same MPE, Otherwise, the deployment is determined by
	// the workflow manager based on available resources.
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
	// specifies function’s id
	ID string `json:"id"`

	// specifies identifier for one instance of a function
	//
	// An instance of a function shall have unique restrictions in a function group. This identifier shall be unique for
	// each instance in the same function group.
	//
	// NOTE If a function is used more than once in one function group with identical restrictions, these restrictions can
	// be defined by one instance of that function.
	Instance string `json:"instance"`

	// specifies function’s logic port name
	PortName string `json:"port-name"`

	// restrictions to the input descriptor parameters
	// This object shall not be present in “from” objects.
	// +optional
	InputRestrictions *Input `json:"input-restrictions,omitempty"`

	// restrictions to the output descriptor parameters
	// This object shall not be present in “to” objects.
	// +optional
	OutputRestrictions *Output `json:"output-restrictions,omitempty"`
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
	Blacklist []Blacklist `json:"blacklist,omitempty"`
}

type Blacklist string

var (
	RequirementBlacklist     Blacklist = "requirement"
	ClientAssistantBlacklist Blacklist = "client-assistant"
	FailOverBlacklist        Blacklist = "fail-over"
	MonitoringBlacklist      Blacklist = "monitoring"
	ReportingBlacklist       Blacklist = "reporting"
	NotificationBlacklist    Blacklist = "notification"
	SecurityBlacklist        Blacklist = "security"
)

// This descriptor provides requirements parameters that can be configured for the underlying resource.
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
	// typical expected delay for the resource (in millisecond)
	// For workflows this specifies the end-to-end delay requirements for the workflow.
	// For tasks this specifies the delay requirements for the task.
	// +optional
	TypicalDelay *uint64 `json:"typical-delay,omitempty"`

	// minimum delay (i.e. amount time from input to output sample) adequate for this resource  (in millisecond)
	// +optional
	MinDelay *uint64 `json:"min-delay,omitempty"`

	// maximum delay required for this resource (in millisecond)
	// +optional
	MaxDelay *uint64 `json:"max-delay,omitempty"`

	// minimum bandwidth required for this resource (in bits/second)
	// +optional
	MinThroughput *uint64 `json:"min-throughput,omitempty"`

	// maximum bandwidth adequate for this resource (in bits/second)
	// +optional
	MaxThroughput *uint64 `json:"max-throughput,omitempty"`

	// averaging window used to calculate the throughput (in millisecond)
	// The default is one second.
	// +optional
	AveragingWindow *uint64 `json:"averaging-window,omitempty"`
}

type HardwareRequirement struct {
	// number of vcpus to be reserved for the execution of a task
	// +optional
	VCPU *uint64 `json:"vcpu,omitempty"`

	// number of vgpus to be reserved for the execution of a task
	// +optional
	VGPU *uint64 `json:"vgpu,omitempty"`

	// memory to be reserved for the execution of a task (in megabytes)
	// +optional
	RAM *uint64 `json:"ram,omitempty"`

	// size of local disk to be used by a workflow or a task (in gigabytes)
	// +optional
	Disk *uint64 `json:"disk,omitempty"`

	// identifier of the geographical location of the data center in which the task is to be executed
	//
	// The location is represented by a two-letter (alpha2) country code (ISO 3166-1) optionally followed by ‘-‘ and the
	// postal code conforming to the country’s postal code standard.
	// +optional
	Placement *HardwareRequirementPlacement `json:"placement,omitempty"`
}

type HardwareRequirementPlacement string // pattern (^[A-Z]{2}$)|(^[A-Z]{2}-.*)

type SecurityRequirement struct {
	// indicates if TLS or DTLS shall be used for the transport of media data
	// default is false
	// +optional
	TLS bool `json:"tls"`

	// indicates if IPSec tunnel model shall be used for the transport of media data
	// default is false
	// +optional
	IPsec bool `json:"ipsec"`

	// indicates if MPEG common encryption (ISO/IEC 23001-7) shall be used for the exchange of media data
	// default is false
	// +optional
	CENC bool `json:"cenc"`
}

type WorkflowTaskRequirement struct {
	// whether functions can be fused or split by the NBMP workflow manager
	//
	// When fused or enhanced, some system tasks may be added or dropped automatically and dynamically.
	// default is false
	// +optional
	FunctionFusible bool `json:"function-fusible"`

	// whether the inputs and outputs of a task can be modified or enhanced with system-provided built-in functions such
	// as media transcoding, media transport buffering for synchronization, or transporting media/metadata data between
	// tasks/MPEs over networks
	// default is false
	// +optional
	FunctionEnhancable bool `json:"function-enhancable"`

	// defines workflow execution modes
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

var (
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

var (
	PnormTaskSplitEfficiencyNorm  TaskSplitEfficiencyNorm = "pnorm"
	CustomTaskSplitEfficiencyNorm TaskSplitEfficiencyNorm = "custom"
)

type ResourceEstimatorsRequirement struct {
	DefaultValues []DefaultValue `json:"default-values"`

	// +optional
	ComputationalEstimator *string `json:"computational-estimator,omitempty"`

	// +optional
	MemoryEstimator *string `json:"memory-estimator,omitempty"`

	// +optional
	BandwidthEstimator *string `json:"bandwidth-estimator,omitempty"`
}

type DefaultValue struct {
	// name of the input, output or configuration parameter
	Name string `json:"name"`

	// default value of the input, output or configuration parameter
	Value string `json:"value"`
}

// This descriptor provides configuration information for the underlying resource.
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
	// must be nil for array datatype
	// +optional
	Values []ParameterValue `json:"values,omitempty"`

	// must be nil for non-array datatypes
	// must be set for array datatype
	// TODO: this is an array in the JSON schema, but must (likely) be an JSON schema object.
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

var (
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

// This descriptor provides information for a delayed startup of the underlying resource.
type StartupDelay struct {
	// amount of delay before task startup in seconds
	StartupDelayValue uint64 `json:"startup-delay-value"`
}

// This descriptor provides client assistance information for the underlying resource.
type ClientAssistant struct {
	// indicates whether the resource requires/supports client monitoring/assistance
	//
	// If client-assistance-flag is set to true in WDD: the workflow manager shall insert a measurement function to
	// collect client assistance information from the client. The workflow manager shall connect the workflow tasks with
	// functions that support client assistance. If the workflow manager cannot support client assistance, the workflow
	// construction shall fail.
	//
	// If client-assistance-flag is set to false in WDD: The insertion of measurement function in workflow is optional.
	//
	// If client-assistance-flag is set to true in function description: the function cannot be instantiated without
	// client assistance information.
	//
	// If client-assistance-flag is set to false in function description: client assistance information is optional.
	// default is false
	ClientAssistanceFlag bool `json:"client-assistance-flag"`

	// object representing the list of measurements to be collected
	//
	// Each element of the list represents a measurement to be collected.
	//
	// Following is a sample list of measurements to be supported:
	// * viewPortCollection
	// * deviceCapabilityCollection
	// * userPreferencesCollection
	//
	// The object may include the frequency of collection for each parameter.
	//
	// The elements of this objects shall be described using parameter schema
	//
	// shall not be present for function description documents
	//
	// +optional
	MeasurementCollectionList map[string]interface{} `json:"measurement-collection-list,omitempty"`

	// list of objects where each object represents different type information from NBMP source
	//
	// The elements of this objects shall be described using parameter schema.
	//
	// shall not be present for function description documents
	//
	// +optional
	SourceAssistanceInformation map[string]interface{} `json:"source-assistance-information,omitempty"`
}

type Failover struct {
	// indicates action upon failover of underlying resource
	// default is "exit"
	FailoverMode FailoverMode `json:"failover-mode"`

	// indicates the amount of delay before starting fail-over (in seconds)
	//
	// When failover-mode value is ‘restart-immediately’ or ‘exit’: this value is considered to be 0.
	//
	// For other failover-mode values: this value defines the amount of time before failover is taken.
	FailoverDelay uint64 `json:"failover-delay"`

	// URL (according to IETF RFC 3986) to an external/internal instruction file for backup deployment that needs to be
	// executed upon failover
	// +optional
	BackupDeploymentURL *base.URI `json:"backup-deployment-url,omitempty"`

	// URL (according to IETF RFC 3986) of storage where the state information is persisted
	//
	// This information is optional from the media source. The workflow manager can allocate some storage and use it for
	// state information persistence.
	// +optional
	PersistenceURL *base.URI `json:"persistence-url,omitempty"`

	// defines how often the state information is written to the persistence-url (in seconds)
	// +optional
	PersistenceInterval *uint64 `json:"persistence-interval,omitempty"`
}

type FailoverMode string

var (
	// restart the resource
	RestartImmediatelyFailoverMode FailoverMode = "restart-immediately"

	// restart the resource after a certain delay
	RestartWithDelayFailoverMode FailoverMode = "restart-with-delay"

	// restart the resource based on available state persistence information
	ContinueWithLastGoodStateFailoverMode FailoverMode = "continue-with-last-good-state"

	// execute backup deployment script given by backup-deployment-url below
	ExecuteBackupDeploymentFailoverMode FailoverMode = "execute-backup-deployment"

	// exit the resource
	ExitFailoverMode FailoverMode = "exit"
)

// This descriptor provides events for the underlying resource. For a function, this descriptor describes the events
// that can be monitored, reported or used in notification in the task or workflow implementing this function.
type Event struct {
	// event’s name
	// +optional
	Name *string `json:"name,omitempty"`

	// humanly readable definition of this event
	// +optional
	Definition *string `json:"definition,omitempty"`

	// unique identifier for event, according to IETF RFC 3986
	// +optional
	URL *base.URI `json:"url,omitempty"`
}

// This descriptor provides variables for the underlying resource. For a function, this descriptor describes the
// variables that can be monitored and/or reported in the function.
type Variable struct {
	// variable’s name
	Name string `json:"name"`

	// humanly readable definition of the variable.
	Definition string `json:"definition"`

	// unit the variable is measured in
	Unit string `json:"unit"`

	// type of variable
	VarType VariableType `json:"var-type"`

	// value of variable
	// +optional
	Value *string `json:"value,omitempty"`

	// minimum value of variable’s range
	// +optional
	Min *int64 `json:"min,omitempty"`

	// maximum value of variable’s range
	// +optional
	Max *int64 `json:"max,omitempty"`

	// unique identifier for this variable, according to IETF RFC 3986
	// +optional
	URL *base.URI `json:"url,omitempty"`

	// +optional
	Children []Variable `json:"children,omitempty"`
}

type VariableType string

var (
	StringVariableType  VariableType = "string"
	IntegerVariableType VariableType = "integer"
	FloatVariableType   VariableType = "float"
	BooleanVariableType VariableType = "boolean"
	NumberVariableType  VariableType = "number"
)

// This descriptor provides monitoring information for the underlying resource.
type Monitoring struct {
	// +optional
	Event []Event `json:"event,omitempty"`

	// shall not be present for workflow description documents
	// +optional
	Variable []Variable `json:"variable,omitempty"`

	// +optional
	SystemEvents []map[string]interface{} `json:"system-events,omitempty"`

	// +optional
	SystemVariables []map[string]interface{} `json:"system-variables,omitempty"`
}

// This descriptor provides reporting information for the underlying resource.
type Reporting struct {
	// +optional
	Event []Event `json:"event,omitempty"`

	// +optional
	Variable []Variable `json:"variable,omitempty"`

	// +optional
	SystemEvents []map[string]interface{} `json:"system-events,omitempty"`

	// +optional
	SystemVariables []map[string]interface{} `json:"system-variables,omitempty"`

	// describes the type of report, defined by the NBMP source or workflow manager
	ReportType string `json:"report-type"`

	// indicates how often the reports needs to be generated and reported
	ReportingInterval uint64 `json:"reporting-interval"`

	// start time for reporting
	ReportStartTime time.Time `json:"report-start-time"`

	// URL (according to IETF RFC 3986) of an external repository where the reports need to be reported/deposited
	URL base.URI `json:"url"`

	// type of delivery methods that are supported for reporting
	// default is "HTTP POST"
	DeliveryMethod DeliveryMethod `json:"delivery-method"`
}

type DeliveryMethod string

var (
	HTTP_POSTDeliveryMethod DeliveryMethod = "HTTP POST"
)

// This descriptor provides notification information for the underlying resource.
type Notification struct {
	// +optional
	Event []Event `json:"event,omitempty"`

	// +optional
	Variable []Variable `json:"variable,omitempty"`

	// +optional
	SystemEvents []map[string]interface{} `json:"system-events,omitempty"`

	// +optional
	SystemVariables []map[string]interface{} `json:"system-variables,omitempty"`

	// time for notification
	NotificationTime time.Time `json:"notification-time"`

	// The level of severity defined by the NBMP source/workflow manager
	SeverityLevel string `json:"severity-level"`

	// Type of notification this resource can produce/send.
	NotificationType []NotificationType `json:"notification-type"`

	// URLs (according to IETF RFC 3986) where the resources intend to receive notifications
	URLs []base.URI `json:"urls"`

	// Interval at which notifications needs to be delivered (in milliseconds).
	//
	// Notification interval of zero indicates that the notification should be sent as soon as the corresponding event is
	// observed
	//
	// Value greater than 0: Any value greater than 0 indicates the interval after which the notification is delivered.
	// default is 0 +optional
	NotificationInterval uint64 `json:"notification-interval"`
}

type NotificationType string

var (
	// Indicates capability to send congestion notification information
	CongestionNotificationType NotificationType = "congestion"

	// Indicates capability to send application-specific notification information
	ApplicationNotificationType NotificationType = "application"

	// Indicates capability to send system-specific notification information
	SystemNotificationType NotificationType = "system"
)

// This descriptor provides assertion information for validating the underlying resource.
type Assertion struct {
	// minimum priority above which all assertions with higher priority shall be processed
	MinPriority uint64 `json:"min-priority"`

	// common action for all lower priority assertions i.e. assertions whose priority is less than minpriority
	MinPriorityAction AssertionAction `json:"min-priority-action"`

	// indicates whether the resource supports providing verification information for validating function assertions
	// default is false
	// +optional
	SupportVerification bool `json:"support-verification"`

	// acknowledgement for verification
	// +optional
	VerificationAcknowledgement *string `json:"verification-acknowledgement,omitempty"`

	// public certificate for signature verification
	// +optional
	Certificate *string `json:"certificate,omitempty"`

	Assertion []AssertionItem `json:"assertion"`
}

type AssertionItem struct {
	// name of the parameter to be checked
	Name string `json:"name"`

	ValuePredicate AssertionValuePredicate `json:"value-predicate"`
}

type AssertionValuePredicate struct {
	// condition against which the parameter will be checked with the given value.
	EvaluationCondition AssertionEvaluationCondition `json:"evaluation-condition"`

	// value against which the parameter value will be checked
	CheckValue map[string]interface{} `json:"check-value"`

	// aggregation function to be used while evaluating assertion across all tasks in the workflow
	Aggregation AssertionValuePredicateAggregation `json:"aggregation"`

	// offset limit that the parameter can deviate from the given value for the evaluation condition to evaluate to a
	// success
	// +optional
	Offset *string `json:"offset,omitempty"`

	// priority of assertion
	Priority uint64 `json:"priority"`

	// action to perform if the evaluation has failed
	//
	// The above actions may only be set for workflow and shall not be set for any task. A task may ignore any of these
	// actions if it is requested.
	Action AssertionAction `json:"action"`

	// parameters for an action represented using ‘action’ is performed
	//
	// The action ‘wait’ has the parameter ‘wait-time’ which indicates the time to wait in milliseconds.
	// +optional
	ActionParameters []string `json:"action-parameters,omitempty"`
}

type AssertionEvaluationCondition string

var (
	// provides description to create assertions that check the quality of media processing
	QualityAssertionEvaluationCondition AssertionEvaluationCondition = "quality"

	// provides description to create assertions that check the computational requirements of media processing
	ComputationalAssertionEvaluationCondition AssertionEvaluationCondition = "computational"

	// provides description to create assertions that check whether the workflow input is of certain kind
	InputAssertionEvaluationCondition AssertionEvaluationCondition = "input"

	// provides description to create assertions that check whether the workflow output is of certain kind
	OutputAssertionEvaluationCondition AssertionEvaluationCondition = "output"
)

type AssertionValuePredicateAggregation string

var (
	// aggregate based on sum over parameters of individual tasks
	SumAssertionValuePredicateAggregation AssertionValuePredicateAggregation = "sum"

	// aggregate based on minimum
	MinAssertionValuePredicateAggregation AssertionValuePredicateAggregation = "min"

	// aggregate based on maximum
	MaxAssertionValuePredicateAggregation AssertionValuePredicateAggregation = "max"

	// aggregate based on average
	AvgAssertionValuePredicateAggregation AssertionValuePredicateAggregation = "avg"
)

type AssertionAction string

var (
	// rebuild the workflow
	RebuildAssertionAction AssertionAction = "rebuild"

	// restart the workflow with the same tasks to satisfy the assertion
	RestartAssertionAction AssertionAction = "restart"

	// wait for a certain time to continue execution of the workflow
	WaitAssertionAction AssertionAction = "wait"
)

// This descriptor provides information for the request sent by a task.  It can be used for identifying the repeated and
// identical requests, the priority of request compared to other received requests, and the requesting task.
type Request struct {
	// unique unsigned integer to identify this NBMP request
	//
	// NBMP requests with identical ids are equivalent and only one may be processed.
	RequestID uint64 `json:"request-id"`

	// unsigned integer indicating the priority of request compared to requests with different id’s value
	//
	// Lower number means higher priority.
	// +optional
	Priority *uint64 `json:"priority,omitempty"`

	// value of requesting task’s general descriptor’s id
	TaskID string `json:"task-id"`
}

// This descriptor indicated whether a description or a descriptor is fulfilled during the processing of a request.
type Acknowledge struct {
	// indicates the status of the item
	Status AcknowledgeStatus `json:"status"`

	// object names of the subitems which are not supported
	// +optional
	Unsupported []string `json:"unsupported,omitempty"`

	// object names of subitems that are not accommodated
	// +optional
	Failed []string `json:"failed,omitempty"`

	// object names of subitems that are not accommodated to the requested extends
	// +optional
	Partial []string `json:"partial,omitempty"`
}

type AcknowledgeStatus string

var (
	// The request was fulfilled for this item.
	FulfilledAcknowledgeStatus AcknowledgeStatus = "fulfilled"

	// The request was not fulfilled for this item.
	FailedAcknowledgeStatus AcknowledgeStatus = "failed"

	// This request is not supported for this item.
	NotSupportedAcknowledgeStatus AcknowledgeStatus = "not-supported"

	// This request was partially fulfilled for this item. In this case, the subitem’s value(s) are the actual values that
	// are fulfilled.
	PartiallyFulfilledAcknowledgeStatus AcknowledgeStatus = "partially-fulfilled"
)

// This descriptor provides the list of function repositories to be used in creating a workflow.
type Repository struct {
	// provides the mode for the repository preference
	//
	// If more than one repository is listed, then the order of listing indicates the preference order of use from high to
	// low, i.e. the first repository is the most preferred one.
	//
	// default is "available"
	// +optional
	Mode *RepositoryMode `json:"mode,omitempty"`

	Location []RepositoryLocation `json:"location"`
}

type RepositoryMode string

var (
	// only the listed repositories in this descriptor shall be used for deployment of the workflow.
	StrictRepositoryMode RepositoryMode = "strict"

	// the listed repositories in this descriptor shall be used first for deployment of the workflow. If a function is not
	// found in these repositories, a different repository may be used.
	PreferredRepositoryMode RepositoryMode = "preferred"

	// the listed repositories in this descriptor may be used first for deployment of the workflow. Other repositories may
	// also be used.
	AvailableRepositoryMode RepositoryMode = "available"
)

type RepositoryLocation struct {
	// location of the repository. It shall be a valid URL according to IETF RFC 3986.
	URL base.URI `json:"url"`

	// name of the repository
	Name string `json:"name"`

	// provides a human readable description for the repository
	Description string `json:"description"`
}

// These should have zero or multiple authentication descriptors, depending on the concrete use cases, inputs, and
// outputs.
type Security struct {
	// identifier can be used by input, output, processing descriptors
	Name string `json:"name"`

	// scope of the authentication, authorization and encryption on different resources.
	// default is "data"
	// +optional
	Scope *SecurityScope `json:"scope,omitempty"`

	// suggested authentication, authorization, and encryption methods or protocols by names
	//
	// Multiple methods or protocols can be used with specific parameters.
	//
	// Sample methods are access token, JSON web token (JWT), single-sign-on (SSO) like OAuth1/2, SAML1/2, client certificate, server certificate
	AuthenticationMethod string `json:"authentication-method"`

	// authority URL (according to IETF RFC 3986) for authentication and authorization, if provided
	// +optional
	AuthorityURL *base.URI `json:"authority-url,omitempty"`

	// trusted certificate, X.509 certificate, if the certificate method is specified
	// +optional
	Certificate *string `json:"certificate,omitempty"`

	// access token like HMAC, key wrapped key, or security key to a KMS (key management system), if token-based method is specified.
	// +optional
	AuthToken *string `json:"auth-token,omitempty"`

	// client grants if token-based method is specified
	// +optional
	ClientGrants *string `json:"client-grants,omitempty"`

	// period of media source for which the authentication token is applicable in format defined by IETF RFC 3339:2002,
	// Section 5.6
	// +optional
	AuthTokenExpires *time.Time `json:"auth-token-expires,omitempty"`

	// token to renew the auth-token after it is expired
	// +optional
	AuthTokenRenew *string `json:"auth-token-renew,omitempty"`

	// flag whether or not an auth-token gets rotated and renewed
	// default is false
	// +optional
	AuthTokenRotation bool `json:"auth-token-rotation"`
}

type SecurityScope string

var (
	// parameters for media and metadata
	DataSecurityScope SecurityScope = "data"

	// parameters for NBMP functions
	FunctionSecurityScope SecurityScope = "function"

	// parameters for NBMP tasks
	TaskSecurityScope SecurityScope = "task"
)

// This descriptor provides information for stateful and stateless step operation of the underlying resource. A resource
// with step descriptor shall include one metadata for each input for receiving the sequence number and/or
// start/duration of each input instance. Each input instance has a duration equal to “segment-duration”.
type Step struct {
	// defining the resource running mode
	// default is "stream"
	// +optional
	StepMode *StepMode `json:"step-mode,omitempty"`

	// default is false
	// +optional
	VariableDuration bool `json:"variable-duration"`

	// duration for which the output(s) of resource are independent to any inputs outside of the corresponding duration
	// (in microseconds).
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

	// number of segment-duration the resource is operating in a stateless fashion
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

var (
	// continuous play
	StreamStepMode StepMode = "stream"

	// maintain the state of tasks at end each step
	StatefulStepMode StepMode = "stateful"

	// run in stateless mode without the need for maintaining state
	StatelessStepMode StepMode = "stateless"
)

type SegmentMetadataSupportedFormat string

var (
	NBMPLocationBytestream2022SegmentMetadataSupportedFormat SegmentMetadataSupportedFormat = "nbmp-location-bytestream-2022"
	NBMPSequenceBytestream2022SegmentMetadataSupportedFormat SegmentMetadataSupportedFormat = "nbmp-sequence-bytestream-2022"
	NBMPLocationJSON2022SegmentMetadataSupportedFormat       SegmentMetadataSupportedFormat = "nbmp-location-json-2022"
	NBMPSequenceJSON2022SegmentMetadataSupportedFormat       SegmentMetadataSupportedFormat = "nbmp-sequence-json-2022"
)

type HigherDimensionsDescription string

var (
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

var (
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

var (
	MPEScalingType        ScalingType = "MPE"
	SplitMergeScalingType ScalingType = "split-merge"
)

type ScalingStatus string

var (
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

var (
	DurationScheduleType ScheduleType = "duration"
	SegmentScheduleType  ScheduleType = "segment"
)

type ScheduleStatus string

var (
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
