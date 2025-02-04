// Code generated by smithy-go-codegen DO NOT EDIT.

package notifications

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateChannel struct {
}

func (*validateOpAssociateChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateEventRule struct {
}

func (*validateOpCreateEventRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateEventRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateEventRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateEventRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateNotificationConfiguration struct {
}

func (*validateOpCreateNotificationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateNotificationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateNotificationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateNotificationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEventRule struct {
}

func (*validateOpDeleteEventRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEventRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEventRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEventRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteNotificationConfiguration struct {
}

func (*validateOpDeleteNotificationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteNotificationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteNotificationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteNotificationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeregisterNotificationHub struct {
}

func (*validateOpDeregisterNotificationHub) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeregisterNotificationHub) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeregisterNotificationHubInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeregisterNotificationHubInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateChannel struct {
}

func (*validateOpDisassociateChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetEventRule struct {
}

func (*validateOpGetEventRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetEventRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetEventRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetEventRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetNotificationConfiguration struct {
}

func (*validateOpGetNotificationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetNotificationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetNotificationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetNotificationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetNotificationEvent struct {
}

func (*validateOpGetNotificationEvent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetNotificationEvent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetNotificationEventInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetNotificationEventInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListChannels struct {
}

func (*validateOpListChannels) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListChannels) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListChannelsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListChannelsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListEventRules struct {
}

func (*validateOpListEventRules) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListEventRules) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListEventRulesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListEventRulesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRegisterNotificationHub struct {
}

func (*validateOpRegisterNotificationHub) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRegisterNotificationHub) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RegisterNotificationHubInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRegisterNotificationHubInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateEventRule struct {
}

func (*validateOpUpdateEventRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateEventRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateEventRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateEventRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateNotificationConfiguration struct {
}

func (*validateOpUpdateNotificationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateNotificationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateNotificationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateNotificationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateChannel{}, middleware.After)
}

func addOpCreateEventRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateEventRule{}, middleware.After)
}

func addOpCreateNotificationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateNotificationConfiguration{}, middleware.After)
}

func addOpDeleteEventRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteEventRule{}, middleware.After)
}

func addOpDeleteNotificationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteNotificationConfiguration{}, middleware.After)
}

func addOpDeregisterNotificationHubValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeregisterNotificationHub{}, middleware.After)
}

func addOpDisassociateChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateChannel{}, middleware.After)
}

func addOpGetEventRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetEventRule{}, middleware.After)
}

func addOpGetNotificationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetNotificationConfiguration{}, middleware.After)
}

func addOpGetNotificationEventValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetNotificationEvent{}, middleware.After)
}

func addOpListChannelsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListChannels{}, middleware.After)
}

func addOpListEventRulesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListEventRules{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpRegisterNotificationHubValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRegisterNotificationHub{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateEventRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateEventRule{}, middleware.After)
}

func addOpUpdateNotificationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateNotificationConfiguration{}, middleware.After)
}

func validateOpAssociateChannelInput(v *AssociateChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateChannelInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.NotificationConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NotificationConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateEventRuleInput(v *CreateEventRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateEventRuleInput"}
	if v.NotificationConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NotificationConfigurationArn"))
	}
	if v.Source == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Source"))
	}
	if v.EventType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventType"))
	}
	if v.Regions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Regions"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateNotificationConfigurationInput(v *CreateNotificationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateNotificationConfigurationInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Description == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Description"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEventRuleInput(v *DeleteEventRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEventRuleInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteNotificationConfigurationInput(v *DeleteNotificationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteNotificationConfigurationInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeregisterNotificationHubInput(v *DeregisterNotificationHubInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeregisterNotificationHubInput"}
	if v.NotificationHubRegion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NotificationHubRegion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateChannelInput(v *DisassociateChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateChannelInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.NotificationConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NotificationConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetEventRuleInput(v *GetEventRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetEventRuleInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetNotificationConfigurationInput(v *GetNotificationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetNotificationConfigurationInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetNotificationEventInput(v *GetNotificationEventInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetNotificationEventInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListChannelsInput(v *ListChannelsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListChannelsInput"}
	if v.NotificationConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NotificationConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListEventRulesInput(v *ListEventRulesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListEventRulesInput"}
	if v.NotificationConfigurationArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NotificationConfigurationArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRegisterNotificationHubInput(v *RegisterNotificationHubInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RegisterNotificationHubInput"}
	if v.NotificationHubRegion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NotificationHubRegion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateEventRuleInput(v *UpdateEventRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateEventRuleInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateNotificationConfigurationInput(v *UpdateNotificationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateNotificationConfigurationInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
