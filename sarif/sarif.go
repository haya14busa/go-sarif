// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    sarif, err := UnmarshalSarif(bytes)
//    bytes, err = sarif.Marshal()

package sarif

import "encoding/json"

func UnmarshalSarif(data []byte) (Sarif, error) {
	var r Sarif
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Sarif) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Static Analysis Results Format (SARIF) Version 2.1.0-rtm.4 JSON Schema: a standard format
// for the output of static analysis tools.
type Sarif struct {
	Schema                   *string              `json:"$schema,omitempty"`                 // The URI of the JSON schema corresponding to the version.
	InlineExternalProperties []ExternalProperties `json:"inlineExternalProperties,omitempty"`// References to external property files that share data between runs.
	Properties               *PropertyBag         `json:"properties,omitempty"`              // Key/value pairs that provide additional information about the log file.
	Runs                     []Run                `json:"runs"`                              // The set of runs contained in this log file.
	Version                  Version              `json:"version"`                           // The SARIF format version of this log file.
}

// The top-level element of an external property file.
type ExternalProperties struct {
	Addresses              []Address            `json:"addresses,omitempty"`             // Addresses that will be merged with a separate run.
	Artifacts              []Artifact           `json:"artifacts,omitempty"`             // An array of artifact objects that will be merged with a separate run.
	Conversion             *Conversion          `json:"conversion,omitempty"`            // A conversion object that will be merged with a separate run.
	Driver                 *ToolComponent       `json:"driver,omitempty"`                // The analysis tool object that will be merged with a separate run.
	Extensions             []ToolComponent      `json:"extensions,omitempty"`            // Tool extensions that will be merged with a separate run.
	ExternalizedProperties *PropertyBag         `json:"externalizedProperties,omitempty"`// Key/value pairs that provide additional information that will be merged with a separate; run.
	Graphs                 []Graph              `json:"graphs,omitempty"`                // An array of graph objects that will be merged with a separate run.
	GUID                   *string              `json:"guid,omitempty"`                  // A stable, unique identifer for this external properties object, in the form of a GUID.
	Invocations            []Invocation         `json:"invocations,omitempty"`           // Describes the invocation of the analysis tool that will be merged with a separate run.
	LogicalLocations       []LogicalLocation    `json:"logicalLocations,omitempty"`      // An array of logical locations such as namespaces, types or functions that will be merged; with a separate run.
	Policies               []ToolComponent      `json:"policies,omitempty"`              // Tool policies that will be merged with a separate run.
	Properties             *PropertyBag         `json:"properties,omitempty"`            // Key/value pairs that provide additional information about the external properties.
	Results                []Result             `json:"results,omitempty"`               // An array of result objects that will be merged with a separate run.
	RunGUID                *string              `json:"runGuid,omitempty"`               // A stable, unique identifer for the run associated with this external properties object,; in the form of a GUID.
	Schema                 *string              `json:"schema,omitempty"`                // The URI of the JSON schema corresponding to the version of the external property file; format.
	Taxonomies             []ToolComponent      `json:"taxonomies,omitempty"`            // Tool taxonomies that will be merged with a separate run.
	ThreadFlowLocations    []ThreadFlowLocation `json:"threadFlowLocations,omitempty"`   // An array of threadFlowLocation objects that will be merged with a separate run.
	Translations           []ToolComponent      `json:"translations,omitempty"`          // Tool translations that will be merged with a separate run.
	Version                *Version             `json:"version,omitempty"`               // The SARIF format version of this external properties object.
	WebRequests            []WebRequest         `json:"webRequests,omitempty"`           // Requests that will be merged with a separate run.
	WebResponses           []WebResponse        `json:"webResponses,omitempty"`          // Responses that will be merged with a separate run.
}

// A physical or virtual address, or a range of addresses, in an 'addressable region'
// (memory or a binary file).
//
// The address of the location.
type Address struct {
	AbsoluteAddress    *int64       `json:"absoluteAddress,omitempty"`   // The address expressed as a byte offset from the start of the addressable region.
	FullyQualifiedName *string      `json:"fullyQualifiedName,omitempty"`// A human-readable fully qualified name that is associated with the address.
	Index              *int64       `json:"index,omitempty"`             // The index within run.addresses of the cached object for this address.
	Kind               *string      `json:"kind,omitempty"`              // An open-ended string that identifies the address kind. 'data', 'function',; 'header','instruction', 'module', 'page', 'section', 'segment', 'stack', 'stackFrame',; 'table' are well-known values.
	Length             *int64       `json:"length,omitempty"`            // The number of bytes in this range of addresses.
	Name               *string      `json:"name,omitempty"`              // A name that is associated with the address, e.g., '.text'.
	OffsetFromParent   *int64       `json:"offsetFromParent,omitempty"`  // The byte offset of this address from the absolute or relative address of the parent; object.
	ParentIndex        *int64       `json:"parentIndex,omitempty"`       // The index within run.addresses of the parent object.
	Properties         *PropertyBag `json:"properties,omitempty"`        // Key/value pairs that provide additional information about the address.
	RelativeAddress    *int64       `json:"relativeAddress,omitempty"`   // The address expressed as a byte offset from the absolute address of the top-most parent; object.
}

// Key/value pairs that provide additional information about the address.
//
// Key/value pairs that provide additional information about the object.
//
// Key/value pairs that provide additional information about the artifact content.
//
// Key/value pairs that provide additional information about the message.
//
// Key/value pairs that provide additional information about the artifact location.
//
// Key/value pairs that provide additional information about the artifact.
//
// Contains configuration information specific to a report.
//
// Key/value pairs that provide additional information about the reporting configuration.
//
// Key/value pairs that provide additional information about the reporting descriptor
// reference.
//
// Key/value pairs that provide additional information about the toolComponentReference.
//
// Key/value pairs that provide additional information about the configuration override.
//
// Key/value pairs that provide additional information about the invocation.
//
// Key/value pairs that provide additional information about the exception.
//
// Key/value pairs that provide additional information about the region.
//
// Key/value pairs that provide additional information about the logical location.
//
// Key/value pairs that provide additional information about the physical location.
//
// Key/value pairs that provide additional information about the location.
//
// Key/value pairs that provide additional information about the location relationship.
//
// Key/value pairs that provide additional information about the stack frame.
//
// Key/value pairs that provide additional information about the stack.
//
// Key/value pairs that provide additional information about the notification.
//
// Key/value pairs that provide additional information about the conversion.
//
// Key/value pairs that provide additional information about the report.
//
// Key/value pairs that provide additional information about the tool component.
//
// Key/value pairs that provide additional information about the translation metadata.
//
// Key/value pairs that provide additional information about the tool.
//
// Key/value pairs that provide additional information that will be merged with a separate
// run.
//
// Key/value pairs that provide additional information about the edge.
//
// Key/value pairs that provide additional information about the node.
//
// Key/value pairs that provide additional information about the graph.
//
// Key/value pairs that provide additional information about the external properties.
//
// Key/value pairs that provide additional information about the attachment.
//
// Key/value pairs that provide additional information about the rectangle.
//
// Key/value pairs that provide additional information about the code flow.
//
// Key/value pairs that provide additional information about the threadflow location.
//
// Key/value pairs that provide additional information about the request.
//
// Key/value pairs that provide additional information about the response.
//
// Key/value pairs that provide additional information about the thread flow.
//
// Key/value pairs that provide additional information about the change.
//
// Key/value pairs that provide additional information about the replacement.
//
// Key/value pairs that provide additional information about the fix.
//
// Key/value pairs that provide additional information about the edge traversal.
//
// Key/value pairs that provide additional information about the graph traversal.
//
// Key/value pairs that provide additional information about the result.
//
// Key/value pairs that provide additional information about the suppression.
//
// Key/value pairs that provide additional information about the log file.
//
// Key/value pairs that provide additional information about the run automation details.
//
// Key/value pairs that provide additional information about the external property file.
//
// Key/value pairs that provide additional information about the external property files.
//
// Key/value pairs that provide additional information about the run.
//
// Key/value pairs that provide additional information about the special locations.
//
// Key/value pairs that provide additional information about the version control details.
type PropertyBag struct {
	Tags []string `json:"tags,omitempty"`// A set of distinct strings that provide additional information.
}

// A single artifact. In some cases, this artifact might be nested within another artifact.
type Artifact struct {
	Contents            *ArtifactContent  `json:"contents,omitempty"`           // The contents of the artifact.
	Description         *Message          `json:"description,omitempty"`        // A short description of the artifact.
	Encoding            *string           `json:"encoding,omitempty"`           // Specifies the encoding for an artifact object that refers to a text file.
	Hashes              map[string]string `json:"hashes,omitempty"`             // A dictionary, each of whose keys is the name of a hash function and each of whose values; is the hashed value of the artifact produced by the specified hash function.
	LastModifiedTimeUTC *string           `json:"lastModifiedTimeUtc,omitempty"`// The Coordinated Universal Time (UTC) date and time at which the artifact was most; recently modified. See "Date/time properties" in the SARIF spec for the required format.
	Length              *int64            `json:"length,omitempty"`             // The length of the artifact in bytes.
	Location            *ArtifactLocation `json:"location,omitempty"`           // The location of the artifact.
	MIMEType            *string           `json:"mimeType,omitempty"`           // The MIME type (RFC 2045) of the artifact.
	Offset              *int64            `json:"offset,omitempty"`             // The offset in bytes of the artifact within its containing artifact.
	ParentIndex         *int64            `json:"parentIndex,omitempty"`        // Identifies the index of the immediate parent of the artifact, if this artifact is nested.
	Properties          *PropertyBag      `json:"properties,omitempty"`         // Key/value pairs that provide additional information about the artifact.
	Roles               []Role            `json:"roles,omitempty"`              // The role or roles played by the artifact in the analysis.
	SourceLanguage      *string           `json:"sourceLanguage,omitempty"`     // Specifies the source language for any artifact object that refers to a text file that; contains source code.
}

// The contents of the artifact.
//
// Represents the contents of an artifact.
//
// The portion of the artifact contents within the specified region.
//
// The body of the request.
//
// The body of the response.
//
// The content to insert at the location specified by the 'deletedRegion' property.
type ArtifactContent struct {
	Binary     *string                   `json:"binary,omitempty"`    // MIME Base64-encoded content from a binary artifact, or from a text artifact in its; original encoding.
	Properties *PropertyBag              `json:"properties,omitempty"`// Key/value pairs that provide additional information about the artifact content.
	Rendered   *MultiformatMessageString `json:"rendered,omitempty"`  // An alternate rendered representation of the artifact (e.g., a decompiled representation; of a binary region).
	Text       *string                   `json:"text,omitempty"`      // UTF-8-encoded content from a text artifact.
}

// An alternate rendered representation of the artifact (e.g., a decompiled representation
// of a binary region).
//
// A message string or message format string rendered in multiple formats.
//
// A comprehensive description of the tool component.
//
// A description of the report. Should, as far as possible, provide details sufficient to
// enable resolution of any problem indicated by the result.
//
// Provides the primary documentation for the report, useful when there is no online
// documentation.
//
// A concise description of the report. Should be a single sentence that is understandable
// when visible space is limited to a single line of text.
//
// A brief description of the tool component.
//
// A comprehensive description of the translation metadata.
//
// A brief description of the translation metadata.
type MultiformatMessageString struct {
	Markdown   *string      `json:"markdown,omitempty"`  // A Markdown message string or format string.
	Properties *PropertyBag `json:"properties,omitempty"`// Key/value pairs that provide additional information about the message.
	Text       string       `json:"text"`                // A plain text message string or format string.
}

// A short description of the artifact.
//
// A short description of the artifact location.
//
// A message relevant to the region.
//
// A message relevant to the location.
//
// A description of the location relationship.
//
// A message relevant to this call stack.
//
// A message that describes the condition that was encountered.
//
// A description of the reporting descriptor relationship.
//
// A description of the graph.
//
// A short description of the edge.
//
// A short description of the node.
//
// A message describing the role played by the attachment.
//
// A message relevant to the rectangle.
//
// A message relevant to the code flow.
//
// A message relevant to the thread flow.
//
// A message that describes the proposed fix, enabling viewers to present the proposed
// change to an end user.
//
// A description of this graph traversal.
//
// A message to display to the user as the edge is traversed.
//
// A message that describes the result. The first sentence of the message only will be
// displayed when visible space is limited.
//
// A description of the identity and role played within the engineering system by this
// object's containing run object.
//
// Encapsulates a message intended to be read by the end user.
type Message struct {
	Arguments  []string     `json:"arguments,omitempty"` // An array of strings to substitute into the message string.
	ID         *string      `json:"id,omitempty"`        // The identifier for this message.
	Markdown   *string      `json:"markdown,omitempty"`  // A Markdown message string.
	Properties *PropertyBag `json:"properties,omitempty"`// Key/value pairs that provide additional information about the message.
	Text       *string      `json:"text,omitempty"`      // A plain text message string.
}

// The location of the artifact.
//
// Specifies the location of an artifact.
//
// An absolute URI specifying the location of the analysis tool's executable.
//
// A file containing the standard error stream from the process that was invoked.
//
// A file containing the standard input stream to the process that was invoked.
//
// A file containing the standard output stream from the process that was invoked.
//
// A file containing the interleaved standard output and standard error stream from the
// process that was invoked.
//
// The working directory for the analysis tool run.
//
// Identifies the artifact that the analysis tool was instructed to scan. This need not be
// the same as the artifact where the result actually occurred.
//
// The location of the attachment.
//
// The location of the artifact to change.
//
// The location of the external property file.
//
// Provides a suggestion to SARIF consumers to display file paths relative to the specified
// location.
//
// The location in the local file system to which the root of the repository was mapped at
// the time of the analysis.
type ArtifactLocation struct {
	Description *Message     `json:"description,omitempty"`// A short description of the artifact location.
	Index       *int64       `json:"index,omitempty"`      // The index within the run artifacts array of the artifact object associated with the; artifact location.
	Properties  *PropertyBag `json:"properties,omitempty"` // Key/value pairs that provide additional information about the artifact location.
	URI         *string      `json:"uri,omitempty"`        // A string containing a valid relative or absolute URI.
	URIBaseID   *string      `json:"uriBaseId,omitempty"`  // A string which indirectly specifies the absolute URI with respect to which a relative URI; in the "uri" property is interpreted.
}

// A conversion object that will be merged with a separate run.
//
// Describes how a converter transformed the output of a static analysis tool from the
// analysis tool's native output format into the SARIF format.
//
// A conversion object that describes how a converter transformed an analysis tool's native
// reporting format into the SARIF format.
type Conversion struct {
	AnalysisToolLogFiles []ArtifactLocation `json:"analysisToolLogFiles,omitempty"`// The locations of the analysis tool's per-run log files.
	Invocation           *Invocation        `json:"invocation,omitempty"`          // An invocation object that describes the invocation of the converter.
	Properties           *PropertyBag       `json:"properties,omitempty"`          // Key/value pairs that provide additional information about the conversion.
	Tool                 Tool               `json:"tool"`                          // A tool object that describes the converter.
}

// An invocation object that describes the invocation of the converter.
//
// The runtime environment of the analysis tool run.
type Invocation struct {
	Account                            *string                 `json:"account,omitempty"`                           // The account that ran the analysis tool.
	Arguments                          []string                `json:"arguments,omitempty"`                         // An array of strings, containing in order the command line arguments passed to the tool; from the operating system.
	CommandLine                        *string                 `json:"commandLine,omitempty"`                       // The command line used to invoke the tool.
	EndTimeUTC                         *string                 `json:"endTimeUtc,omitempty"`                        // The Coordinated Universal Time (UTC) date and time at which the run ended. See "Date/time; properties" in the SARIF spec for the required format.
	EnvironmentVariables               map[string]string       `json:"environmentVariables,omitempty"`              // The environment variables associated with the analysis tool process, expressed as; key/value pairs.
	ExecutableLocation                 *ArtifactLocation       `json:"executableLocation,omitempty"`                // An absolute URI specifying the location of the analysis tool's executable.
	ExecutionSuccessful                bool                    `json:"executionSuccessful"`                         // Specifies whether the tool's execution completed successfully.
	ExitCode                           *int64                  `json:"exitCode,omitempty"`                          // The process exit code.
	ExitCodeDescription                *string                 `json:"exitCodeDescription,omitempty"`               // The reason for the process exit.
	ExitSignalName                     *string                 `json:"exitSignalName,omitempty"`                    // The name of the signal that caused the process to exit.
	ExitSignalNumber                   *int64                  `json:"exitSignalNumber,omitempty"`                  // The numeric value of the signal that caused the process to exit.
	Machine                            *string                 `json:"machine,omitempty"`                           // The machine that hosted the analysis tool run.
	NotificationConfigurationOverrides []ConfigurationOverride `json:"notificationConfigurationOverrides,omitempty"`// An array of configurationOverride objects that describe notifications related runtime; overrides.
	ProcessID                          *int64                  `json:"processId,omitempty"`                         // The process id for the analysis tool run.
	ProcessStartFailureMessage         *string                 `json:"processStartFailureMessage,omitempty"`        // The reason given by the operating system that the process failed to start.
	Properties                         *PropertyBag            `json:"properties,omitempty"`                        // Key/value pairs that provide additional information about the invocation.
	ResponseFiles                      []ArtifactLocation      `json:"responseFiles,omitempty"`                     // The locations of any response files specified on the tool's command line.
	RuleConfigurationOverrides         []ConfigurationOverride `json:"ruleConfigurationOverrides,omitempty"`        // An array of configurationOverride objects that describe rules related runtime overrides.
	StartTimeUTC                       *string                 `json:"startTimeUtc,omitempty"`                      // The Coordinated Universal Time (UTC) date and time at which the run started. See; "Date/time properties" in the SARIF spec for the required format.
	Stderr                             *ArtifactLocation       `json:"stderr,omitempty"`                            // A file containing the standard error stream from the process that was invoked.
	Stdin                              *ArtifactLocation       `json:"stdin,omitempty"`                             // A file containing the standard input stream to the process that was invoked.
	Stdout                             *ArtifactLocation       `json:"stdout,omitempty"`                            // A file containing the standard output stream from the process that was invoked.
	StdoutStderr                       *ArtifactLocation       `json:"stdoutStderr,omitempty"`                      // A file containing the interleaved standard output and standard error stream from the; process that was invoked.
	ToolConfigurationNotifications     []Notification          `json:"toolConfigurationNotifications,omitempty"`    // A list of conditions detected by the tool that are relevant to the tool's configuration.
	ToolExecutionNotifications         []Notification          `json:"toolExecutionNotifications,omitempty"`        // A list of runtime conditions detected by the tool during the analysis.
	WorkingDirectory                   *ArtifactLocation       `json:"workingDirectory,omitempty"`                  // The working directory for the analysis tool run.
}

// Information about how a specific rule or notification was reconfigured at runtime.
type ConfigurationOverride struct {
	Configuration ReportingConfiguration       `json:"configuration"`       // Specifies how the rule or notification was configured during the scan.
	Descriptor    ReportingDescriptorReference `json:"descriptor"`          // A reference used to locate the descriptor whose configuration was overridden.
	Properties    *PropertyBag                 `json:"properties,omitempty"`// Key/value pairs that provide additional information about the configuration override.
}

// Specifies how the rule or notification was configured during the scan.
//
// Information about a rule or notification that can be configured at runtime.
//
// Default reporting configuration information.
type ReportingConfiguration struct {
	Enabled    *bool        `json:"enabled,omitempty"`   // Specifies whether the report may be produced during the scan.
	Level      *Level       `json:"level,omitempty"`     // Specifies the failure level for the report.
	Parameters *PropertyBag `json:"parameters,omitempty"`// Contains configuration information specific to a report.
	Properties *PropertyBag `json:"properties,omitempty"`// Key/value pairs that provide additional information about the reporting configuration.
	Rank       *float64     `json:"rank,omitempty"`      // Specifies the relative priority of the report. Used for analysis output only.
}

// A reference used to locate the descriptor whose configuration was overridden.
//
// A reference used to locate the rule descriptor associated with this notification.
//
// A reference used to locate the descriptor relevant to this notification.
//
// A reference to the related reporting descriptor.
//
// A reference used to locate the rule descriptor relevant to this result.
//
// Information about how to locate a relevant reporting descriptor.
type ReportingDescriptorReference struct {
	GUID          *string                 `json:"guid,omitempty"`         // A guid that uniquely identifies the descriptor.
	ID            *string                 `json:"id,omitempty"`           // The id of the descriptor.
	Index         *int64                  `json:"index,omitempty"`        // The index into an array of descriptors in toolComponent.ruleDescriptors,; toolComponent.notificationDescriptors, or toolComponent.taxonomyDescriptors, depending on; context.
	Properties    *PropertyBag            `json:"properties,omitempty"`   // Key/value pairs that provide additional information about the reporting descriptor; reference.
	ToolComponent *ToolComponentReference `json:"toolComponent,omitempty"`// A reference used to locate the toolComponent associated with the descriptor.
}

// A reference used to locate the toolComponent associated with the descriptor.
//
// Identifies a particular toolComponent object, either the driver or an extension.
//
// The component which is strongly associated with this component. For a translation, this
// refers to the component which has been translated. For an extension, this is the driver
// that provides the extension's plugin model.
type ToolComponentReference struct {
	GUID       *string      `json:"guid,omitempty"`      // The 'guid' property of the referenced toolComponent.
	Index      *int64       `json:"index,omitempty"`     // An index into the referenced toolComponent in tool.extensions.
	Name       *string      `json:"name,omitempty"`      // The 'name' property of the referenced toolComponent.
	Properties *PropertyBag `json:"properties,omitempty"`// Key/value pairs that provide additional information about the toolComponentReference.
}

// Describes a condition relevant to the tool itself, as opposed to being relevant to a
// target being analyzed by the tool.
type Notification struct {
	AssociatedRule *ReportingDescriptorReference `json:"associatedRule,omitempty"`// A reference used to locate the rule descriptor associated with this notification.
	Descriptor     *ReportingDescriptorReference `json:"descriptor,omitempty"`    // A reference used to locate the descriptor relevant to this notification.
	Exception      *Exception                    `json:"exception,omitempty"`     // The runtime exception, if any, relevant to this notification.
	Level          *Level                        `json:"level,omitempty"`         // A value specifying the severity level of the notification.
	Locations      []Location                    `json:"locations,omitempty"`     // The locations relevant to this notification.
	Message        Message                       `json:"message"`                 // A message that describes the condition that was encountered.
	Properties     *PropertyBag                  `json:"properties,omitempty"`    // Key/value pairs that provide additional information about the notification.
	ThreadID       *int64                        `json:"threadId,omitempty"`      // The thread identifier of the code that generated the notification.
	TimeUTC        *string                       `json:"timeUtc,omitempty"`       // The Coordinated Universal Time (UTC) date and time at which the analysis tool generated; the notification.
}

// The runtime exception, if any, relevant to this notification.
//
// Describes a runtime exception encountered during the execution of an analysis tool.
type Exception struct {
	InnerExceptions []Exception  `json:"innerExceptions,omitempty"`// An array of exception objects each of which is considered a cause of this exception.
	Kind            *string      `json:"kind,omitempty"`           // A string that identifies the kind of exception, for example, the fully qualified type; name of an object that was thrown, or the symbolic name of a signal.
	Message         *string      `json:"message,omitempty"`        // A message that describes the exception.
	Properties      *PropertyBag `json:"properties,omitempty"`     // Key/value pairs that provide additional information about the exception.
	Stack           *Stack       `json:"stack,omitempty"`          // The sequence of function calls leading to the exception.
}

// The sequence of function calls leading to the exception.
//
// A call stack that is relevant to a result.
//
// The call stack leading to this location.
type Stack struct {
	Frames     []StackFrame `json:"frames"`              // An array of stack frames that represents a sequence of calls, rendered in reverse; chronological order, that comprise the call stack.
	Message    *Message     `json:"message,omitempty"`   // A message relevant to this call stack.
	Properties *PropertyBag `json:"properties,omitempty"`// Key/value pairs that provide additional information about the stack.
}

// A function call within a stack trace.
type StackFrame struct {
	Location   *Location    `json:"location,omitempty"`  // The location to which this stack frame refers.
	Module     *string      `json:"module,omitempty"`    // The name of the module that contains the code of this stack frame.
	Parameters []string     `json:"parameters,omitempty"`// The parameters of the call that is executing.
	Properties *PropertyBag `json:"properties,omitempty"`// Key/value pairs that provide additional information about the stack frame.
	ThreadID   *int64       `json:"threadId,omitempty"`  // The thread identifier of the stack frame.
}

// The location to which this stack frame refers.
//
// A location within a programming artifact.
//
// A code location associated with the node.
//
// The code location.
//
// Identifies the location associated with the suppression.
type Location struct {
	Annotations      []Region               `json:"annotations,omitempty"`     // A set of regions relevant to the location.
	ID               *int64                 `json:"id,omitempty"`              // Value that distinguishes this location from all other locations within a single result; object.
	LogicalLocations []LogicalLocation      `json:"logicalLocations,omitempty"`// The logical locations associated with the result.
	Message          *Message               `json:"message,omitempty"`         // A message relevant to the location.
	PhysicalLocation *PhysicalLocation      `json:"physicalLocation,omitempty"`// Identifies the artifact and region.
	Properties       *PropertyBag           `json:"properties,omitempty"`      // Key/value pairs that provide additional information about the location.
	Relationships    []LocationRelationship `json:"relationships,omitempty"`   // An array of objects that describe relationships between this location and others.
}

// A region within an artifact where a result was detected.
//
// Specifies a portion of the artifact that encloses the region. Allows a viewer to display
// additional context around the region.
//
// Specifies a portion of the artifact.
//
// The region of the artifact to delete.
type Region struct {
	ByteLength     *int64           `json:"byteLength,omitempty"`    // The length of the region in bytes.
	ByteOffset     *int64           `json:"byteOffset,omitempty"`    // The zero-based offset from the beginning of the artifact of the first byte in the region.
	CharLength     *int64           `json:"charLength,omitempty"`    // The length of the region in characters.
	CharOffset     *int64           `json:"charOffset,omitempty"`    // The zero-based offset from the beginning of the artifact of the first character in the; region.
	EndColumn      *int64           `json:"endColumn,omitempty"`     // The column number of the character following the end of the region.
	EndLine        *int64           `json:"endLine,omitempty"`       // The line number of the last character in the region.
	Message        *Message         `json:"message,omitempty"`       // A message relevant to the region.
	Properties     *PropertyBag     `json:"properties,omitempty"`    // Key/value pairs that provide additional information about the region.
	Snippet        *ArtifactContent `json:"snippet,omitempty"`       // The portion of the artifact contents within the specified region.
	SourceLanguage *string          `json:"sourceLanguage,omitempty"`// Specifies the source language, if any, of the portion of the artifact specified by the; region object.
	StartColumn    *int64           `json:"startColumn,omitempty"`   // The column number of the first character in the region.
	StartLine      *int64           `json:"startLine,omitempty"`     // The line number of the first character in the region.
}

// A logical location of a construct that produced a result.
type LogicalLocation struct {
	DecoratedName      *string      `json:"decoratedName,omitempty"`     // The machine-readable name for the logical location, such as a mangled function name; provided by a C++ compiler that encodes calling convention, return type and other details; along with the function name.
	FullyQualifiedName *string      `json:"fullyQualifiedName,omitempty"`// The human-readable fully qualified name of the logical location.
	Index              *int64       `json:"index,omitempty"`             // The index within the logical locations array.
	Kind               *string      `json:"kind,omitempty"`              // The type of construct this logical location component refers to. Should be one of; 'function', 'member', 'module', 'namespace', 'parameter', 'resource', 'returnType',; 'type', 'variable', 'object', 'array', 'property', 'value', 'element', 'text',; 'attribute', 'comment', 'declaration', 'dtd' or 'processingInstruction', if any of those; accurately describe the construct.
	Name               *string      `json:"name,omitempty"`              // Identifies the construct in which the result occurred. For example, this property might; contain the name of a class or a method.
	ParentIndex        *int64       `json:"parentIndex,omitempty"`       // Identifies the index of the immediate parent of the construct in which the result was; detected. For example, this property might point to a logical location that represents; the namespace that holds a type.
	Properties         *PropertyBag `json:"properties,omitempty"`        // Key/value pairs that provide additional information about the logical location.
}

// Identifies the artifact and region.
//
// A physical location relevant to a result. Specifies a reference to a programming artifact
// together with a range of bytes or characters within that artifact.
type PhysicalLocation struct {
	Address          *Address          `json:"address,omitempty"`         // The address of the location.
	ArtifactLocation *ArtifactLocation `json:"artifactLocation,omitempty"`// The location of the artifact.
	ContextRegion    *Region           `json:"contextRegion,omitempty"`   // Specifies a portion of the artifact that encloses the region. Allows a viewer to display; additional context around the region.
	Properties       *PropertyBag      `json:"properties,omitempty"`      // Key/value pairs that provide additional information about the physical location.
	Region           *Region           `json:"region,omitempty"`          // Specifies a portion of the artifact.
}

// Information about the relation of one location to another.
type LocationRelationship struct {
	Description *Message     `json:"description,omitempty"`// A description of the location relationship.
	Kinds       []string     `json:"kinds,omitempty"`      // A set of distinct strings that categorize the relationship. Well-known kinds include; 'includes', 'isIncludedBy' and 'relevant'.
	Properties  *PropertyBag `json:"properties,omitempty"` // Key/value pairs that provide additional information about the location relationship.
	Target      int64        `json:"target"`               // A reference to the related location.
}

// A tool object that describes the converter.
//
// The analysis tool that was run.
//
// Information about the tool or tool pipeline that generated the results in this run. A run
// can only contain results produced by a single tool or tool pipeline. A run can aggregate
// results from multiple log files, as long as context around the tool run (tool
// command-line arguments and the like) is identical for all aggregated files.
type Tool struct {
	Driver     ToolComponent   `json:"driver"`              // The analysis tool that was run.
	Extensions []ToolComponent `json:"extensions,omitempty"`// Tool extensions that contributed to or reconfigured the analysis tool that was run.
	Properties *PropertyBag    `json:"properties,omitempty"`// Key/value pairs that provide additional information about the tool.
}

// The analysis tool that was run.
//
// A component, such as a plug-in or the driver, of the analysis tool that was run.
//
// The analysis tool object that will be merged with a separate run.
type ToolComponent struct {
	AssociatedComponent                         *ToolComponentReference             `json:"associatedComponent,omitempty"`                        // The component which is strongly associated with this component. For a translation, this; refers to the component which has been translated. For an extension, this is the driver; that provides the extension's plugin model.
	Contents                                    []Content                           `json:"contents,omitempty"`                                   // The kinds of data contained in this object.
	DottedQuadFileVersion                       *string                             `json:"dottedQuadFileVersion,omitempty"`                      // The binary version of the tool component's primary executable file expressed as four; non-negative integers separated by a period (for operating systems that express file; versions in this way).
	DownloadURI                                 *string                             `json:"downloadUri,omitempty"`                                // The absolute URI from which the tool component can be downloaded.
	FullDescription                             *MultiformatMessageString           `json:"fullDescription,omitempty"`                            // A comprehensive description of the tool component.
	FullName                                    *string                             `json:"fullName,omitempty"`                                   // The name of the tool component along with its version and any other useful identifying; information, such as its locale.
	GlobalMessageStrings                        map[string]MultiformatMessageString `json:"globalMessageStrings,omitempty"`                       // A dictionary, each of whose keys is a resource identifier and each of whose values is a; multiformatMessageString object, which holds message strings in plain text and; (optionally) Markdown format. The strings can include placeholders, which can be used to; construct a message in combination with an arbitrary number of additional string; arguments.
	GUID                                        *string                             `json:"guid,omitempty"`                                       // A unique identifer for the tool component in the form of a GUID.
	InformationURI                              *string                             `json:"informationUri,omitempty"`                             // The absolute URI at which information about this version of the tool component can be; found.
	IsComprehensive                             *bool                               `json:"isComprehensive,omitempty"`                            // Specifies whether this object contains a complete definition of the localizable and/or; non-localizable data for this component, as opposed to including only data that is; relevant to the results persisted to this log file.
	Language                                    *string                             `json:"language,omitempty"`                                   // The language of the messages emitted into the log file during this run (expressed as an; ISO 639-1 two-letter lowercase language code) and an optional region (expressed as an ISO; 3166-1 two-letter uppercase subculture code associated with a country or region). The; casing is recommended but not required (in order for this data to conform to RFC5646).
	LocalizedDataSemanticVersion                *string                             `json:"localizedDataSemanticVersion,omitempty"`               // The semantic version of the localized strings defined in this component; maintained by; components that provide translations.
	Locations                                   []ArtifactLocation                  `json:"locations,omitempty"`                                  // An array of the artifactLocation objects associated with the tool component.
	MinimumRequiredLocalizedDataSemanticVersion *string                             `json:"minimumRequiredLocalizedDataSemanticVersion,omitempty"`// The minimum value of localizedDataSemanticVersion required in translations consumed by; this component; used by components that consume translations.
	Name                                        string                              `json:"name"`                                                 // The name of the tool component.
	Notifications                               []ReportingDescriptor               `json:"notifications,omitempty"`                              // An array of reportingDescriptor objects relevant to the notifications related to the; configuration and runtime execution of the tool component.
	Organization                                *string                             `json:"organization,omitempty"`                               // The organization or company that produced the tool component.
	Product                                     *string                             `json:"product,omitempty"`                                    // A product suite to which the tool component belongs.
	ProductSuite                                *string                             `json:"productSuite,omitempty"`                               // A localizable string containing the name of the suite of products to which the tool; component belongs.
	Properties                                  *PropertyBag                        `json:"properties,omitempty"`                                 // Key/value pairs that provide additional information about the tool component.
	ReleaseDateUTC                              *string                             `json:"releaseDateUtc,omitempty"`                             // A string specifying the UTC date (and optionally, the time) of the component's release.
	Rules                                       []ReportingDescriptor               `json:"rules,omitempty"`                                      // An array of reportingDescriptor objects relevant to the analysis performed by the tool; component.
	SemanticVersion                             *string                             `json:"semanticVersion,omitempty"`                            // The tool component version in the format specified by Semantic Versioning 2.0.
	ShortDescription                            *MultiformatMessageString           `json:"shortDescription,omitempty"`                           // A brief description of the tool component.
	SupportedTaxonomies                         []ToolComponentReference            `json:"supportedTaxonomies,omitempty"`                        // An array of toolComponentReference objects to declare the taxonomies supported by the; tool component.
	Taxa                                        []ReportingDescriptor               `json:"taxa,omitempty"`                                       // An array of reportingDescriptor objects relevant to the definitions of both standalone; and tool-defined taxonomies.
	TranslationMetadata                         *TranslationMetadata                `json:"translationMetadata,omitempty"`                        // Translation metadata, required for a translation, not populated by other component types.
	Version                                     *string                             `json:"version,omitempty"`                                    // The tool component version, in whatever format the component natively provides.
}

// Metadata that describes a specific report produced by the tool, as part of the analysis
// it provides or its runtime reporting.
type ReportingDescriptor struct {
	DefaultConfiguration *ReportingConfiguration             `json:"defaultConfiguration,omitempty"`// Default reporting configuration information.
	DeprecatedGuids      []string                            `json:"deprecatedGuids,omitempty"`     // An array of unique identifies in the form of a GUID by which this report was known in; some previous version of the analysis tool.
	DeprecatedIDS        []string                            `json:"deprecatedIds,omitempty"`       // An array of stable, opaque identifiers by which this report was known in some previous; version of the analysis tool.
	DeprecatedNames      []string                            `json:"deprecatedNames,omitempty"`     // An array of readable identifiers by which this report was known in some previous version; of the analysis tool.
	FullDescription      *MultiformatMessageString           `json:"fullDescription,omitempty"`     // A description of the report. Should, as far as possible, provide details sufficient to; enable resolution of any problem indicated by the result.
	GUID                 *string                             `json:"guid,omitempty"`                // A unique identifer for the reporting descriptor in the form of a GUID.
	Help                 *MultiformatMessageString           `json:"help,omitempty"`                // Provides the primary documentation for the report, useful when there is no online; documentation.
	HelpURI              *string                             `json:"helpUri,omitempty"`             // A URI where the primary documentation for the report can be found.
	ID                   string                              `json:"id"`                            // A stable, opaque identifier for the report.
	MessageStrings       map[string]MultiformatMessageString `json:"messageStrings,omitempty"`      // A set of name/value pairs with arbitrary names. Each value is a multiformatMessageString; object, which holds message strings in plain text and (optionally) Markdown format. The; strings can include placeholders, which can be used to construct a message in combination; with an arbitrary number of additional string arguments.
	Name                 *string                             `json:"name,omitempty"`                // A report identifier that is understandable to an end user.
	Properties           *PropertyBag                        `json:"properties,omitempty"`          // Key/value pairs that provide additional information about the report.
	Relationships        []ReportingDescriptorRelationship   `json:"relationships,omitempty"`       // An array of objects that describe relationships between this reporting descriptor and; others.
	ShortDescription     *MultiformatMessageString           `json:"shortDescription,omitempty"`    // A concise description of the report. Should be a single sentence that is understandable; when visible space is limited to a single line of text.
}

// Information about the relation of one reporting descriptor to another.
type ReportingDescriptorRelationship struct {
	Description *Message                     `json:"description,omitempty"`// A description of the reporting descriptor relationship.
	Kinds       []string                     `json:"kinds,omitempty"`      // A set of distinct strings that categorize the relationship. Well-known kinds include; 'canPrecede', 'canFollow', 'willPrecede', 'willFollow', 'superset', 'subset', 'equal',; 'disjoint', 'relevant', and 'incomparable'.
	Properties  *PropertyBag                 `json:"properties,omitempty"` // Key/value pairs that provide additional information about the reporting descriptor; reference.
	Target      ReportingDescriptorReference `json:"target"`               // A reference to the related reporting descriptor.
}

// Translation metadata, required for a translation, not populated by other component
// types.
//
// Provides additional metadata related to translation.
type TranslationMetadata struct {
	DownloadURI      *string                   `json:"downloadUri,omitempty"`     // The absolute URI from which the translation metadata can be downloaded.
	FullDescription  *MultiformatMessageString `json:"fullDescription,omitempty"` // A comprehensive description of the translation metadata.
	FullName         *string                   `json:"fullName,omitempty"`        // The full name associated with the translation metadata.
	InformationURI   *string                   `json:"informationUri,omitempty"`  // The absolute URI from which information related to the translation metadata can be; downloaded.
	Name             string                    `json:"name"`                      // The name associated with the translation metadata.
	Properties       *PropertyBag              `json:"properties,omitempty"`      // Key/value pairs that provide additional information about the translation metadata.
	ShortDescription *MultiformatMessageString `json:"shortDescription,omitempty"`// A brief description of the translation metadata.
}

// A network of nodes and directed edges that describes some aspect of the structure of the
// code (for example, a call graph).
type Graph struct {
	Description *Message     `json:"description,omitempty"`// A description of the graph.
	Edges       []Edge       `json:"edges,omitempty"`      // An array of edge objects representing the edges of the graph.
	Nodes       []Node       `json:"nodes,omitempty"`      // An array of node objects representing the nodes of the graph.
	Properties  *PropertyBag `json:"properties,omitempty"` // Key/value pairs that provide additional information about the graph.
}

// Represents a directed edge in a graph.
type Edge struct {
	ID           string       `json:"id"`                  // A string that uniquely identifies the edge within its graph.
	Label        *Message     `json:"label,omitempty"`     // A short description of the edge.
	Properties   *PropertyBag `json:"properties,omitempty"`// Key/value pairs that provide additional information about the edge.
	SourceNodeID string       `json:"sourceNodeId"`        // Identifies the source node (the node at which the edge starts).
	TargetNodeID string       `json:"targetNodeId"`        // Identifies the target node (the node at which the edge ends).
}

// Represents a node in a graph.
type Node struct {
	Children   []Node       `json:"children,omitempty"`  // Array of child nodes.
	ID         string       `json:"id"`                  // A string that uniquely identifies the node within its graph.
	Label      *Message     `json:"label,omitempty"`     // A short description of the node.
	Location   *Location    `json:"location,omitempty"`  // A code location associated with the node.
	Properties *PropertyBag `json:"properties,omitempty"`// Key/value pairs that provide additional information about the node.
}

// A result produced by an analysis tool.
type Result struct {
	AnalysisTarget      *ArtifactLocation              `json:"analysisTarget,omitempty"`     // Identifies the artifact that the analysis tool was instructed to scan. This need not be; the same as the artifact where the result actually occurred.
	Attachments         []AttachmentElement            `json:"attachments,omitempty"`        // A set of artifacts relevant to the result.
	BaselineState       *BaselineState                 `json:"baselineState,omitempty"`      // The state of a result relative to a baseline of a previous run.
	CodeFlows           []CodeFlow                     `json:"codeFlows,omitempty"`          // An array of 'codeFlow' objects relevant to the result.
	CorrelationGUID     *string                        `json:"correlationGuid,omitempty"`    // A stable, unique identifier for the equivalence class of logically identical results to; which this result belongs, in the form of a GUID.
	Fingerprints        map[string]string              `json:"fingerprints,omitempty"`       // A set of strings each of which individually defines a stable, unique identity for the; result.
	Fixes               []Fix                          `json:"fixes,omitempty"`              // An array of 'fix' objects, each of which represents a proposed fix to the problem; indicated by the result.
	Graphs              []Graph                        `json:"graphs,omitempty"`             // An array of zero or more unique graph objects associated with the result.
	GraphTraversals     []GraphTraversal               `json:"graphTraversals,omitempty"`    // An array of one or more unique 'graphTraversal' objects.
	GUID                *string                        `json:"guid,omitempty"`               // A stable, unique identifer for the result in the form of a GUID.
	HostedViewerURI     *string                        `json:"hostedViewerUri,omitempty"`    // An absolute URI at which the result can be viewed.
	Kind                *ResultKind                    `json:"kind,omitempty"`               // A value that categorizes results by evaluation state.
	Level               *Level                         `json:"level,omitempty"`              // A value specifying the severity level of the result.
	Locations           []Location                     `json:"locations,omitempty"`          // The set of locations where the result was detected. Specify only one location unless the; problem indicated by the result can only be corrected by making a change at every; specified location.
	Message             Message                        `json:"message"`                      // A message that describes the result. The first sentence of the message only will be; displayed when visible space is limited.
	OccurrenceCount     *int64                         `json:"occurrenceCount,omitempty"`    // A positive integer specifying the number of times this logically unique result was; observed in this run.
	PartialFingerprints map[string]string              `json:"partialFingerprints,omitempty"`// A set of strings that contribute to the stable, unique identity of the result.
	Properties          *PropertyBag                   `json:"properties,omitempty"`         // Key/value pairs that provide additional information about the result.
	Provenance          *ResultProvenance              `json:"provenance,omitempty"`         // Information about how and when the result was detected.
	Rank                *float64                       `json:"rank,omitempty"`               // A number representing the priority or importance of the result.
	RelatedLocations    []Location                     `json:"relatedLocations,omitempty"`   // A set of locations relevant to this result.
	Rule                *ReportingDescriptorReference  `json:"rule,omitempty"`               // A reference used to locate the rule descriptor relevant to this result.
	RuleID              *string                        `json:"ruleId,omitempty"`             // The stable, unique identifier of the rule, if any, to which this notification is; relevant. This member can be used to retrieve rule metadata from the rules dictionary, if; it exists.
	RuleIndex           *int64                         `json:"ruleIndex,omitempty"`          // The index within the tool component rules array of the rule object associated with this; result.
	Stacks              []Stack                        `json:"stacks,omitempty"`             // An array of 'stack' objects relevant to the result.
	Suppressions        []Suppression                  `json:"suppressions,omitempty"`       // A set of suppressions relevant to this result.
	Taxa                []ReportingDescriptorReference `json:"taxa,omitempty"`               // An array of references to taxonomy reporting descriptors that are applicable to the; result.
	WebRequest          *WebRequest                    `json:"webRequest,omitempty"`         // A web request associated with this result.
	WebResponse         *WebResponse                   `json:"webResponse,omitempty"`        // A web response associated with this result.
	WorkItemUris        []string                       `json:"workItemUris,omitempty"`       // The URIs of the work items associated with this result.
}

// An artifact relevant to a result.
type AttachmentElement struct {
	ArtifactLocation ArtifactLocation `json:"artifactLocation"`     // The location of the attachment.
	Description      *Message         `json:"description,omitempty"`// A message describing the role played by the attachment.
	Properties       *PropertyBag     `json:"properties,omitempty"` // Key/value pairs that provide additional information about the attachment.
	Rectangles       []Rectangle      `json:"rectangles,omitempty"` // An array of rectangles specifying areas of interest within the image.
	Regions          []Region         `json:"regions,omitempty"`    // An array of regions of interest within the attachment.
}

// An area within an image.
type Rectangle struct {
	Bottom     *float64     `json:"bottom,omitempty"`    // The Y coordinate of the bottom edge of the rectangle, measured in the image's natural; units.
	Left       *float64     `json:"left,omitempty"`      // The X coordinate of the left edge of the rectangle, measured in the image's natural units.
	Message    *Message     `json:"message,omitempty"`   // A message relevant to the rectangle.
	Properties *PropertyBag `json:"properties,omitempty"`// Key/value pairs that provide additional information about the rectangle.
	Right      *float64     `json:"right,omitempty"`     // The X coordinate of the right edge of the rectangle, measured in the image's natural; units.
	Top        *float64     `json:"top,omitempty"`       // The Y coordinate of the top edge of the rectangle, measured in the image's natural units.
}

// A set of threadFlows which together describe a pattern of code execution relevant to
// detecting a result.
type CodeFlow struct {
	Message     *Message     `json:"message,omitempty"`   // A message relevant to the code flow.
	Properties  *PropertyBag `json:"properties,omitempty"`// Key/value pairs that provide additional information about the code flow.
	ThreadFlows []ThreadFlow `json:"threadFlows"`         // An array of one or more unique threadFlow objects, each of which describes the progress; of a program through a thread of execution.
}

// Describes a sequence of code locations that specify a path through a single thread of
// execution such as an operating system or fiber.
type ThreadFlow struct {
	ID             *string                             `json:"id,omitempty"`            // An string that uniquely identifies the threadFlow within the codeFlow in which it occurs.
	ImmutableState map[string]MultiformatMessageString `json:"immutableState,omitempty"`// Values of relevant expressions at the start of the thread flow that remain constant.
	InitialState   map[string]MultiformatMessageString `json:"initialState,omitempty"`  // Values of relevant expressions at the start of the thread flow that may change during; thread flow execution.
	Locations      []ThreadFlowLocation                `json:"locations"`               // A temporally ordered array of 'threadFlowLocation' objects, each of which describes a; location visited by the tool while producing the result.
	Message        *Message                            `json:"message,omitempty"`       // A message relevant to the thread flow.
	Properties     *PropertyBag                        `json:"properties,omitempty"`    // Key/value pairs that provide additional information about the thread flow.
}

// A location visited by an analysis tool while simulating or monitoring the execution of a
// program.
type ThreadFlowLocation struct {
	ExecutionOrder   *int64                              `json:"executionOrder,omitempty"`  // An integer representing the temporal order in which execution reached this location.
	ExecutionTimeUTC *string                             `json:"executionTimeUtc,omitempty"`// The Coordinated Universal Time (UTC) date and time at which this location was executed.
	Importance       *Importance                         `json:"importance,omitempty"`      // Specifies the importance of this location in understanding the code flow in which it; occurs. The order from most to least important is "essential", "important",; "unimportant". Default: "important".
	Index            *int64                              `json:"index,omitempty"`           // The index within the run threadFlowLocations array.
	Kinds            []string                            `json:"kinds,omitempty"`           // A set of distinct strings that categorize the thread flow location. Well-known kinds; include 'acquire', 'release', 'enter', 'exit', 'call', 'return', 'branch', 'implicit',; 'false', 'true', 'caution', 'danger', 'unknown', 'unreachable', 'taint', 'function',; 'handler', 'lock', 'memory', 'resource', 'scope' and 'value'.
	Location         *Location                           `json:"location,omitempty"`        // The code location.
	Module           *string                             `json:"module,omitempty"`          // The name of the module that contains the code that is executing.
	NestingLevel     *int64                              `json:"nestingLevel,omitempty"`    // An integer representing a containment hierarchy within the thread flow.
	Properties       *PropertyBag                        `json:"properties,omitempty"`      // Key/value pairs that provide additional information about the threadflow location.
	Stack            *Stack                              `json:"stack,omitempty"`           // The call stack leading to this location.
	State            map[string]MultiformatMessageString `json:"state,omitempty"`           // A dictionary, each of whose keys specifies a variable or expression, the associated value; of which represents the variable or expression value. For an annotation of kind; 'continuation', for example, this dictionary might hold the current assumed values of a; set of global variables.
	Taxa             []ReportingDescriptorReference      `json:"taxa,omitempty"`            // An array of references to rule or taxonomy reporting descriptors that are applicable to; the thread flow location.
	WebRequest       *WebRequest                         `json:"webRequest,omitempty"`      // A web request associated with this thread flow location.
	WebResponse      *WebResponse                        `json:"webResponse,omitempty"`     // A web response associated with this thread flow location.
}

// A web request associated with this thread flow location.
//
// Describes an HTTP request.
//
// A web request associated with this result.
type WebRequest struct {
	Body       *ArtifactContent  `json:"body,omitempty"`      // The body of the request.
	Headers    map[string]string `json:"headers,omitempty"`   // The request headers.
	Index      *int64            `json:"index,omitempty"`     // The index within the run.webRequests array of the request object associated with this; result.
	Method     *string           `json:"method,omitempty"`    // The HTTP method. Well-known values are 'GET', 'PUT', 'POST', 'DELETE', 'PATCH', 'HEAD',; 'OPTIONS', 'TRACE', 'CONNECT'.
	Parameters map[string]string `json:"parameters,omitempty"`// The request parameters.
	Properties *PropertyBag      `json:"properties,omitempty"`// Key/value pairs that provide additional information about the request.
	Protocol   *string           `json:"protocol,omitempty"`  // The request protocol. Example: 'http'.
	Target     *string           `json:"target,omitempty"`    // The target of the request.
	Version    *string           `json:"version,omitempty"`   // The request version. Example: '1.1'.
}

// A web response associated with this thread flow location.
//
// Describes the response to an HTTP request.
//
// A web response associated with this result.
type WebResponse struct {
	Body               *ArtifactContent  `json:"body,omitempty"`              // The body of the response.
	Headers            map[string]string `json:"headers,omitempty"`           // The response headers.
	Index              *int64            `json:"index,omitempty"`             // The index within the run.webResponses array of the response object associated with this; result.
	NoResponseReceived *bool             `json:"noResponseReceived,omitempty"`// Specifies whether a response was received from the server.
	Properties         *PropertyBag      `json:"properties,omitempty"`        // Key/value pairs that provide additional information about the response.
	Protocol           *string           `json:"protocol,omitempty"`          // The response protocol. Example: 'http'.
	ReasonPhrase       *string           `json:"reasonPhrase,omitempty"`      // The response reason. Example: 'Not found'.
	StatusCode         *int64            `json:"statusCode,omitempty"`        // The response status code. Example: 451.
	Version            *string           `json:"version,omitempty"`           // The response version. Example: '1.1'.
}

// A proposed fix for the problem represented by a result object. A fix specifies a set of
// artifacts to modify. For each artifact, it specifies a set of bytes to remove, and
// provides a set of new bytes to replace them.
type Fix struct {
	ArtifactChanges []ArtifactChange `json:"artifactChanges"`      // One or more artifact changes that comprise a fix for a result.
	Description     *Message         `json:"description,omitempty"`// A message that describes the proposed fix, enabling viewers to present the proposed; change to an end user.
	Properties      *PropertyBag     `json:"properties,omitempty"` // Key/value pairs that provide additional information about the fix.
}

// A change to a single artifact.
type ArtifactChange struct {
	ArtifactLocation ArtifactLocation `json:"artifactLocation"`    // The location of the artifact to change.
	Properties       *PropertyBag     `json:"properties,omitempty"`// Key/value pairs that provide additional information about the change.
	Replacements     []Replacement    `json:"replacements"`        // An array of replacement objects, each of which represents the replacement of a single; region in a single artifact specified by 'artifactLocation'.
}

// The replacement of a single region of an artifact.
type Replacement struct {
	DeletedRegion   Region           `json:"deletedRegion"`            // The region of the artifact to delete.
	InsertedContent *ArtifactContent `json:"insertedContent,omitempty"`// The content to insert at the location specified by the 'deletedRegion' property.
	Properties      *PropertyBag     `json:"properties,omitempty"`     // Key/value pairs that provide additional information about the replacement.
}

// Represents a path through a graph.
type GraphTraversal struct {
	Description      *Message                            `json:"description,omitempty"`     // A description of this graph traversal.
	EdgeTraversals   []EdgeTraversal                     `json:"edgeTraversals,omitempty"`  // The sequences of edges traversed by this graph traversal.
	ImmutableState   map[string]MultiformatMessageString `json:"immutableState,omitempty"`  // Values of relevant expressions at the start of the graph traversal that remain constant; for the graph traversal.
	InitialState     map[string]MultiformatMessageString `json:"initialState,omitempty"`    // Values of relevant expressions at the start of the graph traversal that may change during; graph traversal.
	Properties       *PropertyBag                        `json:"properties,omitempty"`      // Key/value pairs that provide additional information about the graph traversal.
	ResultGraphIndex *int64                              `json:"resultGraphIndex,omitempty"`// The index within the result.graphs to be associated with the result.
	RunGraphIndex    *int64                              `json:"runGraphIndex,omitempty"`   // The index within the run.graphs to be associated with the result.
}

// Represents the traversal of a single edge during a graph traversal.
type EdgeTraversal struct {
	EdgeID            string                              `json:"edgeId"`                     // Identifies the edge being traversed.
	FinalState        map[string]MultiformatMessageString `json:"finalState,omitempty"`       // The values of relevant expressions after the edge has been traversed.
	Message           *Message                            `json:"message,omitempty"`          // A message to display to the user as the edge is traversed.
	Properties        *PropertyBag                        `json:"properties,omitempty"`       // Key/value pairs that provide additional information about the edge traversal.
	StepOverEdgeCount *int64                              `json:"stepOverEdgeCount,omitempty"`// The number of edge traversals necessary to return from a nested graph.
}

// Information about how and when the result was detected.
//
// Contains information about how and when a result was detected.
type ResultProvenance struct {
	ConversionSources     []PhysicalLocation `json:"conversionSources,omitempty"`    // An array of physicalLocation objects which specify the portions of an analysis tool's; output that a converter transformed into the result.
	FirstDetectionRunGUID *string            `json:"firstDetectionRunGuid,omitempty"`// A GUID-valued string equal to the automationDetails.guid property of the run in which the; result was first detected.
	FirstDetectionTimeUTC *string            `json:"firstDetectionTimeUtc,omitempty"`// The Coordinated Universal Time (UTC) date and time at which the result was first; detected. See "Date/time properties" in the SARIF spec for the required format.
	InvocationIndex       *int64             `json:"invocationIndex,omitempty"`      // The index within the run.invocations array of the invocation object which describes the; tool invocation that detected the result.
	LastDetectionRunGUID  *string            `json:"lastDetectionRunGuid,omitempty"` // A GUID-valued string equal to the automationDetails.guid property of the run in which the; result was most recently detected.
	LastDetectionTimeUTC  *string            `json:"lastDetectionTimeUtc,omitempty"` // The Coordinated Universal Time (UTC) date and time at which the result was most recently; detected. See "Date/time properties" in the SARIF spec for the required format.
	Properties            *PropertyBag       `json:"properties,omitempty"`           // Key/value pairs that provide additional information about the result.
}

// A suppression that is relevant to a result.
type Suppression struct {
	GUID          *string         `json:"guid,omitempty"`         // A stable, unique identifer for the suprression in the form of a GUID.
	Justification *string         `json:"justification,omitempty"`// A string representing the justification for the suppression.
	Kind          SuppressionKind `json:"kind"`                   // A string that indicates where the suppression is persisted.
	Location      *Location       `json:"location,omitempty"`     // Identifies the location associated with the suppression.
	Properties    *PropertyBag    `json:"properties,omitempty"`   // Key/value pairs that provide additional information about the suppression.
	State         *State          `json:"state,omitempty"`        // A string that indicates the state of the suppression.
}

// Describes a single run of an analysis tool, and contains the reported output of that run.
type Run struct {
	Addresses                      []Address                       `json:"addresses,omitempty"`                     // Addresses associated with this run instance, if any.
	Artifacts                      []Artifact                      `json:"artifacts,omitempty"`                     // An array of artifact objects relevant to the run.
	AutomationDetails              *RunAutomationDetails           `json:"automationDetails,omitempty"`             // Automation details that describe this run.
	BaselineGUID                   *string                         `json:"baselineGuid,omitempty"`                  // The 'guid' property of a previous SARIF 'run' that comprises the baseline that was used; to compute result 'baselineState' properties for the run.
	ColumnKind                     *ColumnKind                     `json:"columnKind,omitempty"`                    // Specifies the unit in which the tool measures columns.
	Conversion                     *Conversion                     `json:"conversion,omitempty"`                    // A conversion object that describes how a converter transformed an analysis tool's native; reporting format into the SARIF format.
	DefaultEncoding                *string                         `json:"defaultEncoding,omitempty"`               // Specifies the default encoding for any artifact object that refers to a text file.
	DefaultSourceLanguage          *string                         `json:"defaultSourceLanguage,omitempty"`         // Specifies the default source language for any artifact object that refers to a text file; that contains source code.
	ExternalPropertyFileReferences *ExternalPropertyFileReferences `json:"externalPropertyFileReferences,omitempty"`// References to external property files that should be inlined with the content of a root; log file.
	Graphs                         []Graph                         `json:"graphs,omitempty"`                        // An array of zero or more unique graph objects associated with the run.
	Invocations                    []Invocation                    `json:"invocations,omitempty"`                   // Describes the invocation of the analysis tool.
	Language                       *string                         `json:"language,omitempty"`                      // The language of the messages emitted into the log file during this run (expressed as an; ISO 639-1 two-letter lowercase culture code) and an optional region (expressed as an ISO; 3166-1 two-letter uppercase subculture code associated with a country or region). The; casing is recommended but not required (in order for this data to conform to RFC5646).
	LogicalLocations               []LogicalLocation               `json:"logicalLocations,omitempty"`              // An array of logical locations such as namespaces, types or functions.
	NewlineSequences               []string                        `json:"newlineSequences,omitempty"`              // An ordered list of character sequences that were treated as line breaks when computing; region information for the run.
	OriginalURIBaseIDS             map[string]ArtifactLocation     `json:"originalUriBaseIds,omitempty"`            // The artifact location specified by each uriBaseId symbol on the machine where the tool; originally ran.
	Policies                       []ToolComponent                 `json:"policies,omitempty"`                      // Contains configurations that may potentially override both; reportingDescriptor.defaultConfiguration (the tool's default severities) and; invocation.configurationOverrides (severities established at run-time from the command; line).
	Properties                     *PropertyBag                    `json:"properties,omitempty"`                    // Key/value pairs that provide additional information about the run.
	RedactionTokens                []string                        `json:"redactionTokens,omitempty"`               // An array of strings used to replace sensitive information in a redaction-aware property.
	Results                        []Result                        `json:"results,omitempty"`                       // The set of results contained in an SARIF log. The results array can be omitted when a run; is solely exporting rules metadata. It must be present (but may be empty) if a log file; represents an actual scan.
	RunAggregates                  []RunAutomationDetails          `json:"runAggregates,omitempty"`                 // Automation details that describe the aggregate of runs to which this run belongs.
	SpecialLocations               *SpecialLocations               `json:"specialLocations,omitempty"`              // A specialLocations object that defines locations of special significance to SARIF; consumers.
	Taxonomies                     []ToolComponent                 `json:"taxonomies,omitempty"`                    // An array of toolComponent objects relevant to a taxonomy in which results are categorized.
	ThreadFlowLocations            []ThreadFlowLocation            `json:"threadFlowLocations,omitempty"`           // An array of threadFlowLocation objects cached at run level.
	Tool                           Tool                            `json:"tool"`                                    // Information about the tool or tool pipeline that generated the results in this run. A run; can only contain results produced by a single tool or tool pipeline. A run can aggregate; results from multiple log files, as long as context around the tool run (tool; command-line arguments and the like) is identical for all aggregated files.
	Translations                   []ToolComponent                 `json:"translations,omitempty"`                  // The set of available translations of the localized data provided by the tool.
	VersionControlProvenance       []VersionControlDetails         `json:"versionControlProvenance,omitempty"`      // Specifies the revision in version control of the artifacts that were scanned.
	WebRequests                    []WebRequest                    `json:"webRequests,omitempty"`                   // An array of request objects cached at run level.
	WebResponses                   []WebResponse                   `json:"webResponses,omitempty"`                  // An array of response objects cached at run level.
}

// Automation details that describe this run.
//
// Information that describes a run's identity and role within an engineering system process.
type RunAutomationDetails struct {
	CorrelationGUID *string      `json:"correlationGuid,omitempty"`// A stable, unique identifier for the equivalence class of runs to which this object's; containing run object belongs in the form of a GUID.
	Description     *Message     `json:"description,omitempty"`    // A description of the identity and role played within the engineering system by this; object's containing run object.
	GUID            *string      `json:"guid,omitempty"`           // A stable, unique identifer for this object's containing run object in the form of a GUID.
	ID              *string      `json:"id,omitempty"`             // A hierarchical string that uniquely identifies this object's containing run object.
	Properties      *PropertyBag `json:"properties,omitempty"`     // Key/value pairs that provide additional information about the run automation details.
}

// References to external property files that should be inlined with the content of a root
// log file.
type ExternalPropertyFileReferences struct {
	Addresses              []ExternalPropertyFileReference `json:"addresses,omitempty"`             // An array of external property files containing run.addresses arrays to be merged with the; root log file.
	Artifacts              []ExternalPropertyFileReference `json:"artifacts,omitempty"`             // An array of external property files containing run.artifacts arrays to be merged with the; root log file.
	Conversion             *ExternalPropertyFileReference  `json:"conversion,omitempty"`            // An external property file containing a run.conversion object to be merged with the root; log file.
	Driver                 *ExternalPropertyFileReference  `json:"driver,omitempty"`                // An external property file containing a run.driver object to be merged with the root log; file.
	Extensions             []ExternalPropertyFileReference `json:"extensions,omitempty"`            // An array of external property files containing run.extensions arrays to be merged with; the root log file.
	ExternalizedProperties *ExternalPropertyFileReference  `json:"externalizedProperties,omitempty"`// An external property file containing a run.properties object to be merged with the root; log file.
	Graphs                 []ExternalPropertyFileReference `json:"graphs,omitempty"`                // An array of external property files containing a run.graphs object to be merged with the; root log file.
	Invocations            []ExternalPropertyFileReference `json:"invocations,omitempty"`           // An array of external property files containing run.invocations arrays to be merged with; the root log file.
	LogicalLocations       []ExternalPropertyFileReference `json:"logicalLocations,omitempty"`      // An array of external property files containing run.logicalLocations arrays to be merged; with the root log file.
	Policies               []ExternalPropertyFileReference `json:"policies,omitempty"`              // An array of external property files containing run.policies arrays to be merged with the; root log file.
	Properties             *PropertyBag                    `json:"properties,omitempty"`            // Key/value pairs that provide additional information about the external property files.
	Results                []ExternalPropertyFileReference `json:"results,omitempty"`               // An array of external property files containing run.results arrays to be merged with the; root log file.
	Taxonomies             []ExternalPropertyFileReference `json:"taxonomies,omitempty"`            // An array of external property files containing run.taxonomies arrays to be merged with; the root log file.
	ThreadFlowLocations    []ExternalPropertyFileReference `json:"threadFlowLocations,omitempty"`   // An array of external property files containing run.threadFlowLocations arrays to be; merged with the root log file.
	Translations           []ExternalPropertyFileReference `json:"translations,omitempty"`          // An array of external property files containing run.translations arrays to be merged with; the root log file.
	WebRequests            []ExternalPropertyFileReference `json:"webRequests,omitempty"`           // An array of external property files containing run.requests arrays to be merged with the; root log file.
	WebResponses           []ExternalPropertyFileReference `json:"webResponses,omitempty"`          // An array of external property files containing run.responses arrays to be merged with the; root log file.
}

// An external property file containing a run.conversion object to be merged with the root
// log file.
//
// An external property file containing a run.driver object to be merged with the root log
// file.
//
// An external property file containing a run.properties object to be merged with the root
// log file.
//
// Contains information that enables a SARIF consumer to locate the external property file
// that contains the value of an externalized property associated with the run.
type ExternalPropertyFileReference struct {
	GUID       *string           `json:"guid,omitempty"`      // A stable, unique identifer for the external property file in the form of a GUID.
	ItemCount  *int64            `json:"itemCount,omitempty"` // A non-negative integer specifying the number of items contained in the external property; file.
	Location   *ArtifactLocation `json:"location,omitempty"`  // The location of the external property file.
	Properties *PropertyBag      `json:"properties,omitempty"`// Key/value pairs that provide additional information about the external property file.
}

// A specialLocations object that defines locations of special significance to SARIF
// consumers.
//
// Defines locations of special significance to SARIF consumers.
type SpecialLocations struct {
	DisplayBase *ArtifactLocation `json:"displayBase,omitempty"`// Provides a suggestion to SARIF consumers to display file paths relative to the specified; location.
	Properties  *PropertyBag      `json:"properties,omitempty"` // Key/value pairs that provide additional information about the special locations.
}

// Specifies the information necessary to retrieve a desired revision from a version control
// system.
type VersionControlDetails struct {
	AsOfTimeUTC   *string           `json:"asOfTimeUtc,omitempty"`// A Coordinated Universal Time (UTC) date and time that can be used to synchronize an; enlistment to the state of the repository at that time.
	Branch        *string           `json:"branch,omitempty"`     // The name of a branch containing the revision.
	MappedTo      *ArtifactLocation `json:"mappedTo,omitempty"`   // The location in the local file system to which the root of the repository was mapped at; the time of the analysis.
	Properties    *PropertyBag      `json:"properties,omitempty"` // Key/value pairs that provide additional information about the version control details.
	RepositoryURI string            `json:"repositoryUri"`        // The absolute URI of the repository.
	RevisionID    *string           `json:"revisionId,omitempty"` // A string that uniquely and permanently identifies the revision within the repository.
	RevisionTag   *string           `json:"revisionTag,omitempty"`// A tag that has been applied to the revision.
}

type Role string
const (
	Added Role = "added"
	AnalysisTarget Role = "analysisTarget"
	Attachment Role = "attachment"
	DebugOutputFile Role = "debugOutputFile"
	Deleted Role = "deleted"
	Directory Role = "directory"
	Driver Role = "driver"
	Extension Role = "extension"
	MemoryContents Role = "memoryContents"
	Modified Role = "modified"
	Policy Role = "policy"
	ReferencedOnCommandLine Role = "referencedOnCommandLine"
	Renamed Role = "renamed"
	ResponseFile Role = "responseFile"
	ResultFile Role = "resultFile"
	StandardStream Role = "standardStream"
	Taxonomy Role = "taxonomy"
	ToolSpecifiedConfiguration Role = "toolSpecifiedConfiguration"
	TracedFile Role = "tracedFile"
	Translation Role = "translation"
	Uncontrolled Role = "uncontrolled"
	Unmodified Role = "unmodified"
	UserSpecifiedConfiguration Role = "userSpecifiedConfiguration"
)

// Specifies the failure level for the report.
//
// A value specifying the severity level of the notification.
//
// A value specifying the severity level of the result.
type Level string
const (
	Error Level = "error"
	None Level = "none"
	Note Level = "note"
	Warning Level = "warning"
)

type Content string
const (
	LocalizedData Content = "localizedData"
	NonLocalizedData Content = "nonLocalizedData"
)

// The state of a result relative to a baseline of a previous run.
type BaselineState string
const (
	Absent BaselineState = "absent"
	New BaselineState = "new"
	Unchanged BaselineState = "unchanged"
	Updated BaselineState = "updated"
)

// Specifies the importance of this location in understanding the code flow in which it
// occurs. The order from most to least important is "essential", "important",
// "unimportant". Default: "important".
type Importance string
const (
	Essential Importance = "essential"
	Important Importance = "important"
	Unimportant Importance = "unimportant"
)

// A value that categorizes results by evaluation state.
type ResultKind string
const (
	Fail ResultKind = "fail"
	Informational ResultKind = "informational"
	NotApplicable ResultKind = "notApplicable"
	Open ResultKind = "open"
	Pass ResultKind = "pass"
	Review ResultKind = "review"
)

// A string that indicates where the suppression is persisted.
type SuppressionKind string
const (
	External SuppressionKind = "external"
	InSource SuppressionKind = "inSource"
)

// A string that indicates the state of the suppression.
type State string
const (
	Accepted State = "accepted"
	Rejected State = "rejected"
	UnderReview State = "underReview"
)

// The SARIF format version of this external properties object.
//
// The SARIF format version of this log file.
type Version string
const (
	The210 Version = "2.1.0"
)

// Specifies the unit in which the tool measures columns.
type ColumnKind string
const (
	UnicodeCodePoints ColumnKind = "unicodeCodePoints"
	Utf16CodeUnits ColumnKind = "utf16CodeUnits"
)
