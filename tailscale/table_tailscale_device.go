package tailscale

import (
	"context"

	"github.com/pulumi/pulumi-tailscale"
	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableTailscaleDevice(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "tailscale_device",
		Description: "A Tailscale Device represents the devices under a tailnet.",
		List: &plugin.ListConfig{
			Hydrate: listTailscaleDevices,
			// KeyColumns: []*plugin.KeyColumn{
			// 	{
			// 		Name:    "name",
			// 		Require: plugin.Optional,
			// 	},
			// },
		},
		// Get: &plugin.GetConfig{
		// 	Hydrate:    gettailscaleVendor,
		// 	KeyColumns: plugin.SingleColumn("id"),
		// },
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the device.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "id",
				Description: "An unique identifier of the device.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ID"),
			},
			// {
			// 	Name:        "description",
			// 	Description: "The description of the vendor.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "long_name",
			// 	Description: "The full name of the vendor.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "website_url",
			// 	Description: "The description of the vendor.",
			// 	Type:        proto.ColumnType_STRING,
			// 	Transform:   transform.FromField("WebsiteURL").NullIfZero(),
			// },
			// {
			// 	Name:        "alert_creation_default",
			// 	Description: "Specifies the default method for the alert creation.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "alert_creation_editable",
			// 	Description: "Indicates whether the default alert creation method can be editable, or not.",
			// 	Type:        proto.ColumnType_BOOL,
			// },
			// {
			// 	Name:        "generic_service_type",
			// 	Description: "Specifies the generic service type.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "html_url",
			// 	Description: "The API show URL at which the object is accessible.",
			// 	Type:        proto.ColumnType_STRING,
			// 	Transform:   transform.FromField("HTMLURL").NullIfZero(),
			// },
			// {
			// 	Name:        "integration_guide_url",
			// 	Description: "Specifies the URL of an integration guide for this vendor.",
			// 	Type:        proto.ColumnType_STRING,
			// 	Transform:   transform.FromField("IntegrationGuideURL").NullIfZero(),
			// },
			// {
			// 	Name:        "is_pdcef",
			// 	Description: "Indicates the tailscale Common Event Format(PD-CEF).",
			// 	Type:        proto.ColumnType_BOOL,
			// 	Transform:   transform.FromField("IsPDCEF").NullIfZero(),
			// },
			// {
			// 	Name:        "logo_url",
			// 	Description: "Specifies the URL of a logo identifying the vendor.",
			// 	Type:        proto.ColumnType_STRING,
			// 	Transform:   transform.FromField("LogoURL").NullIfZero(),
			// },
			// {
			// 	Name:        "self",
			// 	Description: "The API show URL at which the object is accessible.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "summary",
			// 	Description: "A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "thumbnail_url",
			// 	Description: "Specifies the URL of a small thumbnail image identifying the vendor.",
			// 	Type:        proto.ColumnType_STRING,
			// 	Transform:   transform.FromField("ThumbnailURL").NullIfZero(),
			// },
			// {
			// 	Name:        "type",
			// 	Description: "The type of object being created.",
			// 	Type:        proto.ColumnType_STRING,
			// },

			// Steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

//// LIST FUNCTION

func listTailscaleDevices(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Create client
	client, err := getSessionConfig(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("tailscale_device.listTailscaleDevices", "connection_error", err)
		return nil, err
	}

	req := tailscale.GetDevicesArgs{}

	// Additional Filters
	// if d.KeyColumnQuals["name"] != nil {
	// 	req.Query = d.KeyColumnQuals["name"].GetStringValue()
	// }

	// Retrieve the list of vendors
	// maxResult := uint(100)

	// Reduce the basic request limit down if the user has only requested a small number of rows
	// limit := d.QueryContext.Limit
	// if d.QueryContext.Limit != nil {
	// 	if uint(*limit) < maxResult {
	// 		maxResult = uint(*limit)
	// 	}
	// }
	// req.APIListObject.Limit = maxResult

	// devices, err := conn.AlertContact.GetAlertContacts(params)
	devices, err := tailscale.GetDevices(ctx, req)
	// if err != nil {
	// 	plugin.Logger(ctx).Error("listAlertContact", "api_error", err)
	// 	return nil, err
	// }

	// for _, item := range contacts.AlertContacts {
	// 	d.StreamListItem(ctx, item)

	// 	// Context can be cancelled due to manual cancellation or the limit has been hit
	// 	if d.QueryStatus.RowsRemaining(ctx) == 0 {
	// 		return nil, nil
	// 	}
	// }

	// return nil, nil
}

//// HYDRATE FUNCTIONS

// func gettailscaleVendor(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	// Create client
// 	client, err := getSessionConfig(ctx, d)
// 	if err != nil {
// 		plugin.Logger(ctx).Error("tailscale_vendor.gettailscaleVendor", "connection_error", err)
// 		return nil, err
// 	}
// 	id := d.KeyColumnQuals["id"].GetStringValue()

// 	// No inputs
// 	if id == "" {
// 		return nil, nil
// 	}

// 	getDetails := func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 		data, err := client.GetVendorWithContext(ctx, id)
// 		return data, err
// 	}
// 	getResponse, err := plugin.RetryHydrate(ctx, d, h, getDetails, &plugin.RetryConfig{ShouldRetryError: shouldRetryError})
// 	if err != nil {
// 		plugin.Logger(ctx).Error("tailscale_vendor.gettailscaleVendor", "query_error", err)

// 		if isNotFoundError(err) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	getResp := getResponse.(*tailscale.Vendor)

// 	return *getResp, nil
// }
