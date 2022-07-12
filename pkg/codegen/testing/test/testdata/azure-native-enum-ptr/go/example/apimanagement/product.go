// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Product details.
// API Version: 2020-12-01.
//
// ## Example Usage
// ### ApiManagementCreateProduct
//
// ```go
// package main
//
// import (
// 	apimanagement "github.com/pulumi/pulumi-azure-native/sdk/go/azure/apimanagement"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apimanagement.NewProduct(ctx, "product", &apimanagement.ProductArgs{
// 			DisplayName:       pulumi.String("Test Template ProductName 4"),
// 			ProductId:         pulumi.String("testproduct"),
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ServiceName:       pulumi.String("apimService1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
//
// ## Import
//
// An existing resource can be imported using its type token, name, and identifier, e.g.
//
// ```sh
// $ pulumi import azure-native:apimanagement:Product testproduct /subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/testproduct
// ```
type Product struct {
	pulumi.CustomResourceState

	// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of false.
	ApprovalRequired pulumi.BoolPtrOutput `pulumi:"approvalRequired"`
	// Product description. May include HTML formatting tags.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Product name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
	SubscriptionRequired pulumi.BoolPtrOutput `pulumi:"subscriptionRequired"`
	// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of false.
	SubscriptionsLimit pulumi.IntPtrOutput `pulumi:"subscriptionsLimit"`
	// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
	Terms pulumi.StringPtrOutput `pulumi:"terms"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProduct registers a new resource with the given unique name, arguments, and options.
func NewProduct(ctx *pulumi.Context,
	name string, args *ProductArgs, opts ...pulumi.ResourceOption) (*Product, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20160707:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20161010:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:Product"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:Product"),
		},
	})
	opts = append(opts, aliases)
	var resource Product
	err := ctx.RegisterResource("example:apimanagement:Product", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProduct gets an existing Product resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductState, opts ...pulumi.ResourceOption) (*Product, error) {
	var resource Product
	err := ctx.ReadResource("example:apimanagement:Product", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Product resources.
type productState struct {
}

type ProductState struct {
}

func (ProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*productState)(nil)).Elem()
}

type productArgs struct {
	// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of false.
	ApprovalRequired *bool `pulumi:"approvalRequired"`
	// Product description. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// Product name.
	DisplayName string `pulumi:"displayName"`
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId *string `pulumi:"productId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
	State *ProductStateEnum `pulumi:"state"`
	// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
	SubscriptionRequired *bool `pulumi:"subscriptionRequired"`
	// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of false.
	SubscriptionsLimit *int `pulumi:"subscriptionsLimit"`
	// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
	Terms *string `pulumi:"terms"`
}

// The set of arguments for constructing a Product resource.
type ProductArgs struct {
	// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of false.
	ApprovalRequired pulumi.BoolPtrInput
	// Product description. May include HTML formatting tags.
	Description pulumi.StringPtrInput
	// Product name.
	DisplayName pulumi.StringInput
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
	State ProductStateEnumPtrInput
	// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
	SubscriptionRequired pulumi.BoolPtrInput
	// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of false.
	SubscriptionsLimit pulumi.IntPtrInput
	// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
	Terms pulumi.StringPtrInput
}

func (ProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productArgs)(nil)).Elem()
}

type ProductInput interface {
	pulumi.Input

	ToProductOutput() ProductOutput
	ToProductOutputWithContext(ctx context.Context) ProductOutput
}

func (*Product) ElementType() reflect.Type {
	return reflect.TypeOf((**Product)(nil)).Elem()
}

func (i *Product) ToProductOutput() ProductOutput {
	return i.ToProductOutputWithContext(context.Background())
}

func (i *Product) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProductOutput)
}

type ProductOutput struct{ *pulumi.OutputState }

func (ProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Product)(nil)).Elem()
}

func (o ProductOutput) ToProductOutput() ProductOutput {
	return o
}

func (o ProductOutput) ToProductOutputWithContext(ctx context.Context) ProductOutput {
	return o
}

// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of false.
func (o ProductOutput) ApprovalRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.BoolPtrOutput { return v.ApprovalRequired }).(pulumi.BoolPtrOutput)
}

// Product description. May include HTML formatting tags.
func (o ProductOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Product name.
func (o ProductOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Product) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource name.
func (o ProductOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Product) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
func (o ProductOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
func (o ProductOutput) SubscriptionRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.BoolPtrOutput { return v.SubscriptionRequired }).(pulumi.BoolPtrOutput)
}

// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of false.
func (o ProductOutput) SubscriptionsLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.IntPtrOutput { return v.SubscriptionsLimit }).(pulumi.IntPtrOutput)
}

// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
func (o ProductOutput) Terms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Product) pulumi.StringPtrOutput { return v.Terms }).(pulumi.StringPtrOutput)
}

// Resource type for API Management resource.
func (o ProductOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Product) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProductOutput{})
}
