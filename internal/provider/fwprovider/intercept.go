// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fwprovider

import (
	"context"
	"fmt"

	"github.com/hashicorp/aws-sdk-go-base/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/isometry/terraform-provider-faws/internal/conns"
	"github.com/isometry/terraform-provider-faws/internal/errs"
	"github.com/isometry/terraform-provider-faws/internal/framework/flex"
	"github.com/isometry/terraform-provider-faws/internal/slices"
	tftags "github.com/isometry/terraform-provider-faws/internal/tags"
	"github.com/isometry/terraform-provider-faws/internal/tfresource"
	"github.com/isometry/terraform-provider-faws/internal/types"
	"github.com/isometry/terraform-provider-faws/internal/types/option"
	"github.com/isometry/terraform-provider-faws/names"
)

type interceptorFunc[Request, Response any] func(context.Context, Request, *Response, *conns.AWSClient, when, diag.Diagnostics) (context.Context, diag.Diagnostics)

// A data source interceptor is functionality invoked during the data source's CRUD request lifecycle.
// If a Before interceptor returns Diagnostics indicating an error occurred then
// no further interceptors in the chain are run and neither is the schema's method.
// In other cases all interceptors in the chain are run.
type dataSourceInterceptor interface {
	// read is invoke for a Read call.
	read(context.Context, datasource.ReadRequest, *datasource.ReadResponse, *conns.AWSClient, when, diag.Diagnostics) (context.Context, diag.Diagnostics)
}

type dataSourceInterceptors []dataSourceInterceptor

type dataSourceInterceptorReadFunc interceptorFunc[datasource.ReadRequest, datasource.ReadResponse]

// read returns a slice of interceptors that run on data source Read.
func (s dataSourceInterceptors) read() []dataSourceInterceptorReadFunc {
	return slices.ApplyToAll(s, func(e dataSourceInterceptor) dataSourceInterceptorReadFunc {
		return e.read
	})
}

type ephemeralResourceInterceptor interface {
	// TODO implement me
}

type ephemeralResourceInterceptors []ephemeralResourceInterceptor

type resourceCRUDRequest interface {
	resource.CreateRequest | resource.ReadRequest | resource.UpdateRequest | resource.DeleteRequest
}
type resourceCRUDResponse interface {
	resource.CreateResponse | resource.ReadResponse | resource.UpdateResponse | resource.DeleteResponse
}

// A resource interceptor is functionality invoked during the resource's CRUD request lifecycle.
// If a Before interceptor returns Diagnostics indicating an error occurred then
// no further interceptors in the chain are run and neither is the schema's method.
// In other cases all interceptors in the chain are run.
type resourceInterceptor interface {
	// create is invoke for a Create call.
	create(context.Context, resource.CreateRequest, *resource.CreateResponse, *conns.AWSClient, when, diag.Diagnostics) (context.Context, diag.Diagnostics)
	// read is invoke for a Read call.
	read(context.Context, resource.ReadRequest, *resource.ReadResponse, *conns.AWSClient, when, diag.Diagnostics) (context.Context, diag.Diagnostics)
	// update is invoke for an Update call.
	update(context.Context, resource.UpdateRequest, *resource.UpdateResponse, *conns.AWSClient, when, diag.Diagnostics) (context.Context, diag.Diagnostics)
	// delete is invoke for a Delete call.
	delete(context.Context, resource.DeleteRequest, *resource.DeleteResponse, *conns.AWSClient, when, diag.Diagnostics) (context.Context, diag.Diagnostics)
}

type resourceInterceptors []resourceInterceptor

type resourceInterceptorFunc[Request resourceCRUDRequest, Response resourceCRUDResponse] interceptorFunc[Request, Response]

// create returns a slice of interceptors that run on resource Create.
func (s resourceInterceptors) create() []resourceInterceptorFunc[resource.CreateRequest, resource.CreateResponse] {
	return slices.ApplyToAll(s, func(e resourceInterceptor) resourceInterceptorFunc[resource.CreateRequest, resource.CreateResponse] {
		return e.create
	})
}

// read returns a slice of interceptors that run on resource Read.
func (s resourceInterceptors) read() []resourceInterceptorFunc[resource.ReadRequest, resource.ReadResponse] {
	return slices.ApplyToAll(s, func(e resourceInterceptor) resourceInterceptorFunc[resource.ReadRequest, resource.ReadResponse] {
		return e.read
	})
}

// update returns a slice of interceptors that run on resource Update.
func (s resourceInterceptors) update() []resourceInterceptorFunc[resource.UpdateRequest, resource.UpdateResponse] {
	return slices.ApplyToAll(s, func(e resourceInterceptor) resourceInterceptorFunc[resource.UpdateRequest, resource.UpdateResponse] {
		return e.update
	})
}

// delete returns a slice of interceptors that run on resource Delete.
func (s resourceInterceptors) delete() []resourceInterceptorFunc[resource.DeleteRequest, resource.DeleteResponse] {
	return slices.ApplyToAll(s, func(e resourceInterceptor) resourceInterceptorFunc[resource.DeleteRequest, resource.DeleteResponse] {
		return e.delete
	})
}

// when represents the point in the CRUD request lifecycle that an interceptor is run.
// Multiple values can be ORed together.
type when uint16

const (
	Before  when = 1 << iota // Interceptor is invoked before call to method in schema
	After                    // Interceptor is invoked after successful call to method in schema
	OnError                  // Interceptor is invoked after unsuccessful call to method in schema
	Finally                  // Interceptor is invoked after After or OnError
)

// TODO Share the intercepted handler logic between data sources and resources..

// interceptedDataSourceHandler returns a handler that invokes the specified data source Read handler, running any interceptors.
func interceptedDataSourceReadHandler(interceptors []dataSourceInterceptorReadFunc, f func(context.Context, datasource.ReadRequest, *datasource.ReadResponse) diag.Diagnostics, meta *conns.AWSClient) func(context.Context, datasource.ReadRequest, *datasource.ReadResponse) diag.Diagnostics {
	return func(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) diag.Diagnostics {
		var diags diag.Diagnostics
		// Before interceptors are run first to last.
		forward := interceptors

		when := Before
		for _, v := range forward {
			ctx, diags = v(ctx, request, response, meta, when, diags)

			// Short circuit if any Before interceptor errors.
			if diags.HasError() {
				return diags
			}
		}

		// All other interceptors are run last to first.
		reverse := slices.Reverse(forward)
		diags = f(ctx, request, response)

		if diags.HasError() {
			when = OnError
		} else {
			when = After
		}
		for _, v := range reverse {
			ctx, diags = v(ctx, request, response, meta, when, diags)
		}

		when = Finally
		for _, v := range reverse {
			ctx, diags = v(ctx, request, response, meta, when, diags)
		}

		return diags
	}
}

// interceptedResourceHandler returns a handler that invokes the specified resource CRUD handler, running any interceptors.
func interceptedResourceHandler[Request resourceCRUDRequest, Response resourceCRUDResponse](interceptors []resourceInterceptorFunc[Request, Response], f func(context.Context, Request, *Response) diag.Diagnostics, meta *conns.AWSClient) func(context.Context, Request, *Response) diag.Diagnostics {
	return func(ctx context.Context, request Request, response *Response) diag.Diagnostics {
		var diags diag.Diagnostics
		// Before interceptors are run first to last.
		forward := interceptors

		when := Before
		for _, v := range forward {
			ctx, diags = v(ctx, request, response, meta, when, diags)

			// Short circuit if any Before interceptor errors.
			if diags.HasError() {
				return diags
			}
		}

		// All other interceptors are run last to first.
		reverse := slices.Reverse(forward)
		diags = f(ctx, request, response)

		if diags.HasError() {
			when = OnError
		} else {
			when = After
		}
		for _, v := range reverse {
			ctx, diags = v(ctx, request, response, meta, when, diags)
		}

		when = Finally
		for _, v := range reverse {
			ctx, diags = v(ctx, request, response, meta, when, diags)
		}

		return diags
	}
}

// tagsDataSourceInterceptor implements transparent tagging for data sources.
type tagsDataSourceInterceptor struct {
	tags *types.ServicePackageResourceTags
}

func (r tagsDataSourceInterceptor) read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse, meta *conns.AWSClient, when when, diags diag.Diagnostics) (context.Context, diag.Diagnostics) {
	if r.tags == nil {
		return ctx, diags
	}

	inContext, ok := conns.FromContext(ctx)
	if !ok {
		return ctx, diags
	}

	sp := meta.ServicePackage(ctx, inContext.ServicePackageName)
	if sp == nil {
		return ctx, diags
	}

	serviceName, err := names.HumanFriendly(sp.ServicePackageName())
	if err != nil {
		serviceName = "<service>"
	}

	resourceName := inContext.ResourceName
	if resourceName == "" {
		resourceName = "<thing>"
	}

	tagsInContext, ok := tftags.FromContext(ctx)
	if !ok {
		return ctx, diags
	}

	switch when {
	case Before:
		var configTags tftags.Map
		diags.Append(request.Config.GetAttribute(ctx, path.Root(names.AttrTags), &configTags)...)
		if diags.HasError() {
			return ctx, diags
		}

		tags := tftags.New(ctx, configTags)

		tagsInContext.TagsIn = option.Some(tags)

	case After:
		// If the R handler didn't set tags, try and read them from the service API.
		if tagsInContext.TagsOut.IsNone() {
			if identifierAttribute := r.tags.IdentifierAttribute; identifierAttribute != "" {
				var identifier string
				diags.Append(response.State.GetAttribute(ctx, path.Root(identifierAttribute), &identifier)...)
				if diags.HasError() {
					return ctx, diags
				}

				// If the service package has a generic resource list tags methods, call it.
				var err error
				if v, ok := sp.(interface {
					ListTags(context.Context, any, string) error
				}); ok {
					err = v.ListTags(ctx, meta, identifier) // Sets tags in Context
				} else if v, ok := sp.(interface {
					ListTags(context.Context, any, string, string) error
				}); ok && r.tags.ResourceType != "" {
					err = v.ListTags(ctx, meta, identifier, r.tags.ResourceType) // Sets tags in Context
				} else {
					tflog.Warn(ctx, "No ListTags method found", map[string]interface{}{
						"ServicePackage": sp.ServicePackageName(),
						"ResourceType":   r.tags.ResourceType,
					})
				}

				// ISO partitions may not support tagging, giving error.
				if errs.IsUnsupportedOperationInPartitionError(meta.Partition(ctx), err) {
					return ctx, diags
				}

				if sp.ServicePackageName() == names.DynamoDB && err != nil {
					// When a DynamoDB Table is `ARCHIVED`, ListTags returns `ResourceNotFoundException`.
					if tfresource.NotFound(err) || tfawserr.ErrMessageContains(err, "UnknownOperationException", "Tagging is not currently supported in DynamoDB Local.") {
						err = nil
					}
				}

				if err != nil {
					diags.AddError(fmt.Sprintf("listing tags for %s %s (%s)", serviceName, resourceName, identifier), err.Error())
					return ctx, diags
				}
			}
		}

		tags := tagsInContext.TagsOut.UnwrapOrDefault()

		// Remove any provider configured ignore_tags and system tags from those returned from the service API.
		stateTags := flex.FlattenFrameworkStringValueMapLegacy(ctx, tags.IgnoreSystem(sp.ServicePackageName()).IgnoreConfig(tagsInContext.IgnoreConfig).Map())
		diags.Append(response.State.SetAttribute(ctx, path.Root(names.AttrTags), tftags.NewMapFromMapValue(stateTags))...)

		if diags.HasError() {
			return ctx, diags
		}
	}

	return ctx, diags
}

// tagsResourceInterceptor implements transparent tagging for resources.
type tagsResourceInterceptor struct {
	tags *types.ServicePackageResourceTags
}

func (r tagsResourceInterceptor) create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse, meta *conns.AWSClient, when when, diags diag.Diagnostics) (context.Context, diag.Diagnostics) {
	if r.tags == nil {
		return ctx, diags
	}

	inContext, ok := conns.FromContext(ctx)
	if !ok {
		return ctx, diags
	}

	tagsInContext, ok := tftags.FromContext(ctx)
	if !ok {
		return ctx, diags
	}

	switch when {
	case Before:
		var planTags tftags.Map
		diags.Append(request.Plan.GetAttribute(ctx, path.Root(names.AttrTags), &planTags)...)

		if diags.HasError() {
			return ctx, diags
		}

		// Merge the resource's configured tags with any provider configured default_tags.
		tags := tagsInContext.DefaultConfig.MergeTags(tftags.New(ctx, planTags))
		// Remove system tags.
		tags = tags.IgnoreSystem(inContext.ServicePackageName)

		tagsInContext.TagsIn = option.Some(tags)
	case After:
		// Set values for unknowns.
		// Remove any provider configured ignore_tags and system tags from those passed to the service API.
		// Computed tags_all include any provider configured default_tags.
		stateTagsAll := flex.FlattenFrameworkStringValueMapLegacy(ctx, tagsInContext.TagsIn.MustUnwrap().IgnoreSystem(inContext.ServicePackageName).IgnoreConfig(tagsInContext.IgnoreConfig).Map())
		diags.Append(response.State.SetAttribute(ctx, path.Root(names.AttrTagsAll), tftags.NewMapFromMapValue(stateTagsAll))...)

		if diags.HasError() {
			return ctx, diags
		}
	}

	return ctx, diags
}

func (r tagsResourceInterceptor) read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse, meta *conns.AWSClient, when when, diags diag.Diagnostics) (context.Context, diag.Diagnostics) {
	if r.tags == nil {
		return ctx, diags
	}

	inContext, ok := conns.FromContext(ctx)
	if !ok {
		return ctx, diags
	}

	sp := meta.ServicePackage(ctx, inContext.ServicePackageName)
	if sp == nil {
		return ctx, diags
	}

	serviceName, err := names.HumanFriendly(sp.ServicePackageName())
	if err != nil {
		serviceName = "<service>"
	}

	resourceName := inContext.ResourceName
	if resourceName == "" {
		resourceName = "<thing>"
	}

	tagsInContext, ok := tftags.FromContext(ctx)
	if !ok {
		return ctx, diags
	}

	switch when {
	case After:
		// Will occur on a refresh when the resource does not exist in AWS and needs to be recreated, e.g. "_disappears" tests.
		if response.State.Raw.IsNull() {
			return ctx, diags
		}

		// If the R handler didn't set tags, try and read them from the service API.
		if tagsInContext.TagsOut.IsNone() {
			if identifierAttribute := r.tags.IdentifierAttribute; identifierAttribute != "" {
				var identifier string

				diags.Append(response.State.GetAttribute(ctx, path.Root(identifierAttribute), &identifier)...)

				if diags.HasError() {
					return ctx, diags
				}

				// Some old resources may not have the required attribute set after Read:
				// https://github.com/isometry/terraform-provider-faws/issues/31180
				if identifier != "" {
					// If the service package has a generic resource list tags methods, call it.
					var err error

					if v, ok := sp.(tftags.ServiceTagLister); ok {
						err = v.ListTags(ctx, meta, identifier) // Sets tags in Context
					} else if v, ok := sp.(tftags.ResourceTypeTagLister); ok {
						if r.tags.ResourceType == "" {
							tflog.Error(ctx, "ListTags method requires ResourceType but none set", map[string]interface{}{
								"ServicePackage": sp.ServicePackageName(),
							})
						} else {
							err = v.ListTags(ctx, meta, identifier, r.tags.ResourceType) // Sets tags in Context
						}
					} else {
						tflog.Warn(ctx, "No ListTags method found", map[string]interface{}{
							"ServicePackage": sp.ServicePackageName(),
							"ResourceType":   r.tags.ResourceType,
						})
					}

					// ISO partitions may not support tagging, giving error.
					if errs.IsUnsupportedOperationInPartitionError(meta.Partition(ctx), err) {
						return ctx, diags
					}

					if err != nil {
						diags.AddError(fmt.Sprintf("listing tags for %s %s (%s)", serviceName, resourceName, identifier), err.Error())

						return ctx, diags
					}
				}
			}
		}

		apiTags := tagsInContext.TagsOut.UnwrapOrDefault()

		// AWS APIs often return empty lists of tags when none have been configured.
		var stateTags tftags.Map
		response.State.GetAttribute(ctx, path.Root(names.AttrTags), &stateTags)
		// Remove any provider configured ignore_tags and system tags from those returned from the service API.
		// The resource's configured tags do not include any provider configured default_tags.
		if v := apiTags.IgnoreSystem(sp.ServicePackageName()).IgnoreConfig(tagsInContext.IgnoreConfig).ResolveDuplicatesFramework(ctx, tagsInContext.DefaultConfig, tagsInContext.IgnoreConfig, response, &diags).Map(); len(v) > 0 {
			stateTags = tftags.NewMapFromMapValue(flex.FlattenFrameworkStringValueMapLegacy(ctx, v))
		}
		diags.Append(response.State.SetAttribute(ctx, path.Root(names.AttrTags), &stateTags)...)

		if diags.HasError() {
			return ctx, diags
		}

		// Computed tags_all do.
		stateTagsAll := flex.FlattenFrameworkStringValueMapLegacy(ctx, apiTags.IgnoreSystem(sp.ServicePackageName()).IgnoreConfig(tagsInContext.IgnoreConfig).Map())
		diags.Append(response.State.SetAttribute(ctx, path.Root(names.AttrTagsAll), tftags.NewMapFromMapValue(stateTagsAll))...)

		if diags.HasError() {
			return ctx, diags
		}
	}

	return ctx, diags
}

func (r tagsResourceInterceptor) update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse, meta *conns.AWSClient, when when, diags diag.Diagnostics) (context.Context, diag.Diagnostics) {
	if r.tags == nil {
		return ctx, diags
	}

	inContext, ok := conns.FromContext(ctx)
	if !ok {
		return ctx, diags
	}

	sp := meta.ServicePackage(ctx, inContext.ServicePackageName)
	if sp == nil {
		return ctx, diags
	}

	serviceName, err := names.HumanFriendly(sp.ServicePackageName())
	if err != nil {
		serviceName = "<service>"
	}

	resourceName := inContext.ResourceName
	if resourceName == "" {
		resourceName = "<thing>"
	}

	tagsInContext, ok := tftags.FromContext(ctx)
	if !ok {
		return ctx, diags
	}

	switch when {
	case Before:
		var planTags tftags.Map
		diags.Append(request.Plan.GetAttribute(ctx, path.Root(names.AttrTags), &planTags)...)

		if diags.HasError() {
			return ctx, diags
		}

		// Merge the resource's configured tags with any provider configured default_tags.
		tags := tagsInContext.DefaultConfig.MergeTags(tftags.New(ctx, planTags))
		// Remove system tags.
		tags = tags.IgnoreSystem(sp.ServicePackageName())

		tagsInContext.TagsIn = option.Some(tags)

		var oldTagsAll, newTagsAll tftags.Map

		diags.Append(request.State.GetAttribute(ctx, path.Root(names.AttrTagsAll), &oldTagsAll)...)

		if diags.HasError() {
			return ctx, diags
		}

		diags.Append(request.Plan.GetAttribute(ctx, path.Root(names.AttrTagsAll), &newTagsAll)...)

		if diags.HasError() {
			return ctx, diags
		}

		if !newTagsAll.Equal(oldTagsAll) {
			if identifierAttribute := r.tags.IdentifierAttribute; identifierAttribute != "" {
				var identifier string

				diags.Append(request.Plan.GetAttribute(ctx, path.Root(identifierAttribute), &identifier)...)

				if diags.HasError() {
					return ctx, diags
				}

				// Some old resources may not have the required attribute set after Read:
				// https://github.com/isometry/terraform-provider-faws/issues/31180
				if identifier != "" {
					// If the service package has a generic resource update tags methods, call it.
					var err error

					if v, ok := sp.(tftags.ServiceTagUpdater); ok {
						err = v.UpdateTags(ctx, meta, identifier, oldTagsAll, newTagsAll)
					} else if v, ok := sp.(tftags.ResourceTypeTagUpdater); ok && r.tags.ResourceType != "" {
						err = v.UpdateTags(ctx, meta, identifier, r.tags.ResourceType, oldTagsAll, newTagsAll)
					} else {
						tflog.Warn(ctx, "No UpdateTags method found", map[string]interface{}{
							"ServicePackage": sp.ServicePackageName(),
							"ResourceType":   r.tags.ResourceType,
						})
					}

					// ISO partitions may not support tagging, giving error.
					if errs.IsUnsupportedOperationInPartitionError(meta.Partition(ctx), err) {
						return ctx, diags
					}

					if err != nil {
						diags.AddError(fmt.Sprintf("updating tags for %s %s (%s)", serviceName, resourceName, identifier), err.Error())

						return ctx, diags
					}
				}
			}
			// TODO If the only change was to tags it would be nice to not call the resource's U handler.
		}
	}

	return ctx, diags
}

func (r tagsResourceInterceptor) delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse, meta *conns.AWSClient, when when, diags diag.Diagnostics) (context.Context, diag.Diagnostics) {
	return ctx, diags
}
