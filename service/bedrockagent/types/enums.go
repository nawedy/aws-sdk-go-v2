// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActionGroupSignature string

// Enum values for ActionGroupSignature
const (
	ActionGroupSignatureAmazonUserinput       ActionGroupSignature = "AMAZON.UserInput"
	ActionGroupSignatureAmazonCodeinterpreter ActionGroupSignature = "AMAZON.CodeInterpreter"
)

// Values returns all known values for ActionGroupSignature. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ActionGroupSignature) Values() []ActionGroupSignature {
	return []ActionGroupSignature{
		"AMAZON.UserInput",
		"AMAZON.CodeInterpreter",
	}
}

type ActionGroupState string

// Enum values for ActionGroupState
const (
	ActionGroupStateEnabled  ActionGroupState = "ENABLED"
	ActionGroupStateDisabled ActionGroupState = "DISABLED"
)

// Values returns all known values for ActionGroupState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ActionGroupState) Values() []ActionGroupState {
	return []ActionGroupState{
		"ENABLED",
		"DISABLED",
	}
}

type AgentAliasStatus string

// Enum values for AgentAliasStatus
const (
	AgentAliasStatusCreating AgentAliasStatus = "CREATING"
	AgentAliasStatusPrepared AgentAliasStatus = "PREPARED"
	AgentAliasStatusFailed   AgentAliasStatus = "FAILED"
	AgentAliasStatusUpdating AgentAliasStatus = "UPDATING"
	AgentAliasStatusDeleting AgentAliasStatus = "DELETING"
)

// Values returns all known values for AgentAliasStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AgentAliasStatus) Values() []AgentAliasStatus {
	return []AgentAliasStatus{
		"CREATING",
		"PREPARED",
		"FAILED",
		"UPDATING",
		"DELETING",
	}
}

type AgentStatus string

// Enum values for AgentStatus
const (
	AgentStatusCreating    AgentStatus = "CREATING"
	AgentStatusPreparing   AgentStatus = "PREPARING"
	AgentStatusPrepared    AgentStatus = "PREPARED"
	AgentStatusNotPrepared AgentStatus = "NOT_PREPARED"
	AgentStatusDeleting    AgentStatus = "DELETING"
	AgentStatusFailed      AgentStatus = "FAILED"
	AgentStatusVersioning  AgentStatus = "VERSIONING"
	AgentStatusUpdating    AgentStatus = "UPDATING"
)

// Values returns all known values for AgentStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (AgentStatus) Values() []AgentStatus {
	return []AgentStatus{
		"CREATING",
		"PREPARING",
		"PREPARED",
		"NOT_PREPARED",
		"DELETING",
		"FAILED",
		"VERSIONING",
		"UPDATING",
	}
}

type ChunkingStrategy string

// Enum values for ChunkingStrategy
const (
	ChunkingStrategyFixedSize    ChunkingStrategy = "FIXED_SIZE"
	ChunkingStrategyNone         ChunkingStrategy = "NONE"
	ChunkingStrategyHierarchical ChunkingStrategy = "HIERARCHICAL"
	ChunkingStrategySemantic     ChunkingStrategy = "SEMANTIC"
)

// Values returns all known values for ChunkingStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ChunkingStrategy) Values() []ChunkingStrategy {
	return []ChunkingStrategy{
		"FIXED_SIZE",
		"NONE",
		"HIERARCHICAL",
		"SEMANTIC",
	}
}

type ConfluenceAuthType string

// Enum values for ConfluenceAuthType
const (
	ConfluenceAuthTypeBasic                   ConfluenceAuthType = "BASIC"
	ConfluenceAuthTypeOauth2ClientCredentials ConfluenceAuthType = "OAUTH2_CLIENT_CREDENTIALS"
)

// Values returns all known values for ConfluenceAuthType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConfluenceAuthType) Values() []ConfluenceAuthType {
	return []ConfluenceAuthType{
		"BASIC",
		"OAUTH2_CLIENT_CREDENTIALS",
	}
}

type ConfluenceHostType string

// Enum values for ConfluenceHostType
const (
	ConfluenceHostTypeSaas ConfluenceHostType = "SAAS"
)

// Values returns all known values for ConfluenceHostType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConfluenceHostType) Values() []ConfluenceHostType {
	return []ConfluenceHostType{
		"SAAS",
	}
}

type ConversationRole string

// Enum values for ConversationRole
const (
	ConversationRoleUser      ConversationRole = "user"
	ConversationRoleAssistant ConversationRole = "assistant"
)

// Values returns all known values for ConversationRole. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConversationRole) Values() []ConversationRole {
	return []ConversationRole{
		"user",
		"assistant",
	}
}

type CrawlFilterConfigurationType string

// Enum values for CrawlFilterConfigurationType
const (
	CrawlFilterConfigurationTypePattern CrawlFilterConfigurationType = "PATTERN"
)

// Values returns all known values for CrawlFilterConfigurationType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CrawlFilterConfigurationType) Values() []CrawlFilterConfigurationType {
	return []CrawlFilterConfigurationType{
		"PATTERN",
	}
}

type CreationMode string

// Enum values for CreationMode
const (
	CreationModeDefault    CreationMode = "DEFAULT"
	CreationModeOverridden CreationMode = "OVERRIDDEN"
)

// Values returns all known values for CreationMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CreationMode) Values() []CreationMode {
	return []CreationMode{
		"DEFAULT",
		"OVERRIDDEN",
	}
}

type CustomControlMethod string

// Enum values for CustomControlMethod
const (
	CustomControlMethodReturnControl CustomControlMethod = "RETURN_CONTROL"
)

// Values returns all known values for CustomControlMethod. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (CustomControlMethod) Values() []CustomControlMethod {
	return []CustomControlMethod{
		"RETURN_CONTROL",
	}
}

type DataDeletionPolicy string

// Enum values for DataDeletionPolicy
const (
	DataDeletionPolicyRetain DataDeletionPolicy = "RETAIN"
	DataDeletionPolicyDelete DataDeletionPolicy = "DELETE"
)

// Values returns all known values for DataDeletionPolicy. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataDeletionPolicy) Values() []DataDeletionPolicy {
	return []DataDeletionPolicy{
		"RETAIN",
		"DELETE",
	}
}

type DataSourceStatus string

// Enum values for DataSourceStatus
const (
	DataSourceStatusAvailable          DataSourceStatus = "AVAILABLE"
	DataSourceStatusDeleting           DataSourceStatus = "DELETING"
	DataSourceStatusDeleteUnsuccessful DataSourceStatus = "DELETE_UNSUCCESSFUL"
)

// Values returns all known values for DataSourceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceStatus) Values() []DataSourceStatus {
	return []DataSourceStatus{
		"AVAILABLE",
		"DELETING",
		"DELETE_UNSUCCESSFUL",
	}
}

type DataSourceType string

// Enum values for DataSourceType
const (
	DataSourceTypeS3         DataSourceType = "S3"
	DataSourceTypeWeb        DataSourceType = "WEB"
	DataSourceTypeConfluence DataSourceType = "CONFLUENCE"
	DataSourceTypeSalesforce DataSourceType = "SALESFORCE"
	DataSourceTypeSharepoint DataSourceType = "SHAREPOINT"
)

// Values returns all known values for DataSourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DataSourceType) Values() []DataSourceType {
	return []DataSourceType{
		"S3",
		"WEB",
		"CONFLUENCE",
		"SALESFORCE",
		"SHAREPOINT",
	}
}

type EmbeddingDataType string

// Enum values for EmbeddingDataType
const (
	EmbeddingDataTypeFloat32 EmbeddingDataType = "FLOAT32"
	EmbeddingDataTypeBinary  EmbeddingDataType = "BINARY"
)

// Values returns all known values for EmbeddingDataType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EmbeddingDataType) Values() []EmbeddingDataType {
	return []EmbeddingDataType{
		"FLOAT32",
		"BINARY",
	}
}

type FlowConnectionType string

// Enum values for FlowConnectionType
const (
	FlowConnectionTypeData        FlowConnectionType = "Data"
	FlowConnectionTypeConditional FlowConnectionType = "Conditional"
)

// Values returns all known values for FlowConnectionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FlowConnectionType) Values() []FlowConnectionType {
	return []FlowConnectionType{
		"Data",
		"Conditional",
	}
}

type FlowNodeIODataType string

// Enum values for FlowNodeIODataType
const (
	FlowNodeIODataTypeString  FlowNodeIODataType = "String"
	FlowNodeIODataTypeNumber  FlowNodeIODataType = "Number"
	FlowNodeIODataTypeBoolean FlowNodeIODataType = "Boolean"
	FlowNodeIODataTypeObject  FlowNodeIODataType = "Object"
	FlowNodeIODataTypeArray   FlowNodeIODataType = "Array"
)

// Values returns all known values for FlowNodeIODataType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FlowNodeIODataType) Values() []FlowNodeIODataType {
	return []FlowNodeIODataType{
		"String",
		"Number",
		"Boolean",
		"Object",
		"Array",
	}
}

type FlowNodeType string

// Enum values for FlowNodeType
const (
	FlowNodeTypeInput          FlowNodeType = "Input"
	FlowNodeTypeOutput         FlowNodeType = "Output"
	FlowNodeTypeKnowledgeBase  FlowNodeType = "KnowledgeBase"
	FlowNodeTypeCondition      FlowNodeType = "Condition"
	FlowNodeTypeLex            FlowNodeType = "Lex"
	FlowNodeTypePrompt         FlowNodeType = "Prompt"
	FlowNodeTypeLambdaFunction FlowNodeType = "LambdaFunction"
	FlowNodeTypeStorage        FlowNodeType = "Storage"
	FlowNodeTypeAgent          FlowNodeType = "Agent"
	FlowNodeTypeRetrieval      FlowNodeType = "Retrieval"
	FlowNodeTypeIterator       FlowNodeType = "Iterator"
	FlowNodeTypeCollector      FlowNodeType = "Collector"
)

// Values returns all known values for FlowNodeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FlowNodeType) Values() []FlowNodeType {
	return []FlowNodeType{
		"Input",
		"Output",
		"KnowledgeBase",
		"Condition",
		"Lex",
		"Prompt",
		"LambdaFunction",
		"Storage",
		"Agent",
		"Retrieval",
		"Iterator",
		"Collector",
	}
}

type FlowStatus string

// Enum values for FlowStatus
const (
	FlowStatusFailed      FlowStatus = "Failed"
	FlowStatusPrepared    FlowStatus = "Prepared"
	FlowStatusPreparing   FlowStatus = "Preparing"
	FlowStatusNotPrepared FlowStatus = "NotPrepared"
)

// Values returns all known values for FlowStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FlowStatus) Values() []FlowStatus {
	return []FlowStatus{
		"Failed",
		"Prepared",
		"Preparing",
		"NotPrepared",
	}
}

type FlowValidationSeverity string

// Enum values for FlowValidationSeverity
const (
	FlowValidationSeverityWarning FlowValidationSeverity = "Warning"
	FlowValidationSeverityError   FlowValidationSeverity = "Error"
)

// Values returns all known values for FlowValidationSeverity. Note that this can
// be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FlowValidationSeverity) Values() []FlowValidationSeverity {
	return []FlowValidationSeverity{
		"Warning",
		"Error",
	}
}

type FlowValidationType string

// Enum values for FlowValidationType
const (
	FlowValidationTypeCyclicConnection                FlowValidationType = "CyclicConnection"
	FlowValidationTypeDuplicateConnections            FlowValidationType = "DuplicateConnections"
	FlowValidationTypeDuplicateConditionExpression    FlowValidationType = "DuplicateConditionExpression"
	FlowValidationTypeUnreachableNode                 FlowValidationType = "UnreachableNode"
	FlowValidationTypeUnknownConnectionSource         FlowValidationType = "UnknownConnectionSource"
	FlowValidationTypeUnknownConnectionSourceOutput   FlowValidationType = "UnknownConnectionSourceOutput"
	FlowValidationTypeUnknownConnectionTarget         FlowValidationType = "UnknownConnectionTarget"
	FlowValidationTypeUnknownConnectionTargetInput    FlowValidationType = "UnknownConnectionTargetInput"
	FlowValidationTypeUnknownConnectionCondition      FlowValidationType = "UnknownConnectionCondition"
	FlowValidationTypeMalformedConditionExpression    FlowValidationType = "MalformedConditionExpression"
	FlowValidationTypeMalformedNodeInputExpression    FlowValidationType = "MalformedNodeInputExpression"
	FlowValidationTypeMismatchedNodeInputType         FlowValidationType = "MismatchedNodeInputType"
	FlowValidationTypeMismatchedNodeOutputType        FlowValidationType = "MismatchedNodeOutputType"
	FlowValidationTypeIncompatibleConnectionDataType  FlowValidationType = "IncompatibleConnectionDataType"
	FlowValidationTypeMissingConnectionConfiguration  FlowValidationType = "MissingConnectionConfiguration"
	FlowValidationTypeMissingDefaultCondition         FlowValidationType = "MissingDefaultCondition"
	FlowValidationTypeMissingEndingNodes              FlowValidationType = "MissingEndingNodes"
	FlowValidationTypeMissingNodeConfiguration        FlowValidationType = "MissingNodeConfiguration"
	FlowValidationTypeMissingNodeInput                FlowValidationType = "MissingNodeInput"
	FlowValidationTypeMissingNodeOutput               FlowValidationType = "MissingNodeOutput"
	FlowValidationTypeMissingStartingNodes            FlowValidationType = "MissingStartingNodes"
	FlowValidationTypeMultipleNodeInputConnections    FlowValidationType = "MultipleNodeInputConnections"
	FlowValidationTypeUnfulfilledNodeInput            FlowValidationType = "UnfulfilledNodeInput"
	FlowValidationTypeUnsatisfiedConnectionConditions FlowValidationType = "UnsatisfiedConnectionConditions"
	FlowValidationTypeUnspecified                     FlowValidationType = "Unspecified"
)

// Values returns all known values for FlowValidationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (FlowValidationType) Values() []FlowValidationType {
	return []FlowValidationType{
		"CyclicConnection",
		"DuplicateConnections",
		"DuplicateConditionExpression",
		"UnreachableNode",
		"UnknownConnectionSource",
		"UnknownConnectionSourceOutput",
		"UnknownConnectionTarget",
		"UnknownConnectionTargetInput",
		"UnknownConnectionCondition",
		"MalformedConditionExpression",
		"MalformedNodeInputExpression",
		"MismatchedNodeInputType",
		"MismatchedNodeOutputType",
		"IncompatibleConnectionDataType",
		"MissingConnectionConfiguration",
		"MissingDefaultCondition",
		"MissingEndingNodes",
		"MissingNodeConfiguration",
		"MissingNodeInput",
		"MissingNodeOutput",
		"MissingStartingNodes",
		"MultipleNodeInputConnections",
		"UnfulfilledNodeInput",
		"UnsatisfiedConnectionConditions",
		"Unspecified",
	}
}

type IngestionJobFilterAttribute string

// Enum values for IngestionJobFilterAttribute
const (
	IngestionJobFilterAttributeStatus IngestionJobFilterAttribute = "STATUS"
)

// Values returns all known values for IngestionJobFilterAttribute. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IngestionJobFilterAttribute) Values() []IngestionJobFilterAttribute {
	return []IngestionJobFilterAttribute{
		"STATUS",
	}
}

type IngestionJobFilterOperator string

// Enum values for IngestionJobFilterOperator
const (
	IngestionJobFilterOperatorEq IngestionJobFilterOperator = "EQ"
)

// Values returns all known values for IngestionJobFilterOperator. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IngestionJobFilterOperator) Values() []IngestionJobFilterOperator {
	return []IngestionJobFilterOperator{
		"EQ",
	}
}

type IngestionJobSortByAttribute string

// Enum values for IngestionJobSortByAttribute
const (
	IngestionJobSortByAttributeStatus    IngestionJobSortByAttribute = "STATUS"
	IngestionJobSortByAttributeStartedAt IngestionJobSortByAttribute = "STARTED_AT"
)

// Values returns all known values for IngestionJobSortByAttribute. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IngestionJobSortByAttribute) Values() []IngestionJobSortByAttribute {
	return []IngestionJobSortByAttribute{
		"STATUS",
		"STARTED_AT",
	}
}

type IngestionJobStatus string

// Enum values for IngestionJobStatus
const (
	IngestionJobStatusStarting   IngestionJobStatus = "STARTING"
	IngestionJobStatusInProgress IngestionJobStatus = "IN_PROGRESS"
	IngestionJobStatusComplete   IngestionJobStatus = "COMPLETE"
	IngestionJobStatusFailed     IngestionJobStatus = "FAILED"
	IngestionJobStatusStopping   IngestionJobStatus = "STOPPING"
	IngestionJobStatusStopped    IngestionJobStatus = "STOPPED"
)

// Values returns all known values for IngestionJobStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (IngestionJobStatus) Values() []IngestionJobStatus {
	return []IngestionJobStatus{
		"STARTING",
		"IN_PROGRESS",
		"COMPLETE",
		"FAILED",
		"STOPPING",
		"STOPPED",
	}
}

type KnowledgeBaseState string

// Enum values for KnowledgeBaseState
const (
	KnowledgeBaseStateEnabled  KnowledgeBaseState = "ENABLED"
	KnowledgeBaseStateDisabled KnowledgeBaseState = "DISABLED"
)

// Values returns all known values for KnowledgeBaseState. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseState) Values() []KnowledgeBaseState {
	return []KnowledgeBaseState{
		"ENABLED",
		"DISABLED",
	}
}

type KnowledgeBaseStatus string

// Enum values for KnowledgeBaseStatus
const (
	KnowledgeBaseStatusCreating           KnowledgeBaseStatus = "CREATING"
	KnowledgeBaseStatusActive             KnowledgeBaseStatus = "ACTIVE"
	KnowledgeBaseStatusDeleting           KnowledgeBaseStatus = "DELETING"
	KnowledgeBaseStatusUpdating           KnowledgeBaseStatus = "UPDATING"
	KnowledgeBaseStatusFailed             KnowledgeBaseStatus = "FAILED"
	KnowledgeBaseStatusDeleteUnsuccessful KnowledgeBaseStatus = "DELETE_UNSUCCESSFUL"
)

// Values returns all known values for KnowledgeBaseStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseStatus) Values() []KnowledgeBaseStatus {
	return []KnowledgeBaseStatus{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"UPDATING",
		"FAILED",
		"DELETE_UNSUCCESSFUL",
	}
}

type KnowledgeBaseStorageType string

// Enum values for KnowledgeBaseStorageType
const (
	KnowledgeBaseStorageTypeOpensearchServerless KnowledgeBaseStorageType = "OPENSEARCH_SERVERLESS"
	KnowledgeBaseStorageTypePinecone             KnowledgeBaseStorageType = "PINECONE"
	KnowledgeBaseStorageTypeRedisEnterpriseCloud KnowledgeBaseStorageType = "REDIS_ENTERPRISE_CLOUD"
	KnowledgeBaseStorageTypeRds                  KnowledgeBaseStorageType = "RDS"
	KnowledgeBaseStorageTypeMongoDbAtlas         KnowledgeBaseStorageType = "MONGO_DB_ATLAS"
)

// Values returns all known values for KnowledgeBaseStorageType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseStorageType) Values() []KnowledgeBaseStorageType {
	return []KnowledgeBaseStorageType{
		"OPENSEARCH_SERVERLESS",
		"PINECONE",
		"REDIS_ENTERPRISE_CLOUD",
		"RDS",
		"MONGO_DB_ATLAS",
	}
}

type KnowledgeBaseType string

// Enum values for KnowledgeBaseType
const (
	KnowledgeBaseTypeVector KnowledgeBaseType = "VECTOR"
)

// Values returns all known values for KnowledgeBaseType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (KnowledgeBaseType) Values() []KnowledgeBaseType {
	return []KnowledgeBaseType{
		"VECTOR",
	}
}

type MemoryType string

// Enum values for MemoryType
const (
	MemoryTypeSessionSummary MemoryType = "SESSION_SUMMARY"
)

// Values returns all known values for MemoryType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (MemoryType) Values() []MemoryType {
	return []MemoryType{
		"SESSION_SUMMARY",
	}
}

type OrchestrationType string

// Enum values for OrchestrationType
const (
	OrchestrationTypeDefault             OrchestrationType = "DEFAULT"
	OrchestrationTypeCustomOrchestration OrchestrationType = "CUSTOM_ORCHESTRATION"
)

// Values returns all known values for OrchestrationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (OrchestrationType) Values() []OrchestrationType {
	return []OrchestrationType{
		"DEFAULT",
		"CUSTOM_ORCHESTRATION",
	}
}

type ParsingStrategy string

// Enum values for ParsingStrategy
const (
	ParsingStrategyBedrockFoundationModel ParsingStrategy = "BEDROCK_FOUNDATION_MODEL"
)

// Values returns all known values for ParsingStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ParsingStrategy) Values() []ParsingStrategy {
	return []ParsingStrategy{
		"BEDROCK_FOUNDATION_MODEL",
	}
}

type PromptState string

// Enum values for PromptState
const (
	PromptStateEnabled  PromptState = "ENABLED"
	PromptStateDisabled PromptState = "DISABLED"
)

// Values returns all known values for PromptState. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PromptState) Values() []PromptState {
	return []PromptState{
		"ENABLED",
		"DISABLED",
	}
}

type PromptTemplateType string

// Enum values for PromptTemplateType
const (
	PromptTemplateTypeText PromptTemplateType = "TEXT"
	PromptTemplateTypeChat PromptTemplateType = "CHAT"
)

// Values returns all known values for PromptTemplateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PromptTemplateType) Values() []PromptTemplateType {
	return []PromptTemplateType{
		"TEXT",
		"CHAT",
	}
}

type PromptType string

// Enum values for PromptType
const (
	PromptTypePreProcessing                   PromptType = "PRE_PROCESSING"
	PromptTypeOrchestration                   PromptType = "ORCHESTRATION"
	PromptTypePostProcessing                  PromptType = "POST_PROCESSING"
	PromptTypeKnowledgeBaseResponseGeneration PromptType = "KNOWLEDGE_BASE_RESPONSE_GENERATION"
)

// Values returns all known values for PromptType. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (PromptType) Values() []PromptType {
	return []PromptType{
		"PRE_PROCESSING",
		"ORCHESTRATION",
		"POST_PROCESSING",
		"KNOWLEDGE_BASE_RESPONSE_GENERATION",
	}
}

type RequireConfirmation string

// Enum values for RequireConfirmation
const (
	RequireConfirmationEnabled  RequireConfirmation = "ENABLED"
	RequireConfirmationDisabled RequireConfirmation = "DISABLED"
)

// Values returns all known values for RequireConfirmation. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (RequireConfirmation) Values() []RequireConfirmation {
	return []RequireConfirmation{
		"ENABLED",
		"DISABLED",
	}
}

type SalesforceAuthType string

// Enum values for SalesforceAuthType
const (
	SalesforceAuthTypeOauth2ClientCredentials SalesforceAuthType = "OAUTH2_CLIENT_CREDENTIALS"
)

// Values returns all known values for SalesforceAuthType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SalesforceAuthType) Values() []SalesforceAuthType {
	return []SalesforceAuthType{
		"OAUTH2_CLIENT_CREDENTIALS",
	}
}

type SharePointAuthType string

// Enum values for SharePointAuthType
const (
	SharePointAuthTypeOauth2ClientCredentials SharePointAuthType = "OAUTH2_CLIENT_CREDENTIALS"
)

// Values returns all known values for SharePointAuthType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SharePointAuthType) Values() []SharePointAuthType {
	return []SharePointAuthType{
		"OAUTH2_CLIENT_CREDENTIALS",
	}
}

type SharePointHostType string

// Enum values for SharePointHostType
const (
	SharePointHostTypeOnline SharePointHostType = "ONLINE"
)

// Values returns all known values for SharePointHostType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SharePointHostType) Values() []SharePointHostType {
	return []SharePointHostType{
		"ONLINE",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "ASCENDING"
	SortOrderDescending SortOrder = "DESCENDING"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"ASCENDING",
		"DESCENDING",
	}
}

type StepType string

// Enum values for StepType
const (
	StepTypePostChunking StepType = "POST_CHUNKING"
)

// Values returns all known values for StepType. Note that this can be expanded in
// the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (StepType) Values() []StepType {
	return []StepType{
		"POST_CHUNKING",
	}
}

type Type string

// Enum values for Type
const (
	TypeString  Type = "string"
	TypeNumber  Type = "number"
	TypeInteger Type = "integer"
	TypeBoolean Type = "boolean"
	TypeArray   Type = "array"
)

// Values returns all known values for Type. Note that this can be expanded in the
// future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (Type) Values() []Type {
	return []Type{
		"string",
		"number",
		"integer",
		"boolean",
		"array",
	}
}

type WebScopeType string

// Enum values for WebScopeType
const (
	WebScopeTypeHostOnly   WebScopeType = "HOST_ONLY"
	WebScopeTypeSubdomains WebScopeType = "SUBDOMAINS"
)

// Values returns all known values for WebScopeType. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WebScopeType) Values() []WebScopeType {
	return []WebScopeType{
		"HOST_ONLY",
		"SUBDOMAINS",
	}
}
