package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetRegionOptStatusrequest represents the GetRegionOptStatusrequest schema from the OpenAPI specification
type GetRegionOptStatusrequest struct {
	Accountid string `json:"AccountId,omitempty"` // <p>Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation. If you don't specify this parameter, it defaults to the Amazon Web Services account of the identity used to call the operation. To use this parameter, the caller must be an identity in the <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account">organization's management account</a> or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html">all features enabled</a>, and the organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html">trusted access</a> enabled for the Account Management service, and optionally a <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html">delegated admin</a> account assigned.</p> <note> <p>The management account can't specify its own <code>AccountId</code>. It must call the operation in standalone context by not including the <code>AccountId</code> parameter.</p> </note> <p>To call this operation on an account that is not a member of an organization, don't specify this parameter. Instead, call the operation using an identity belonging to the account whose contacts you wish to retrieve or modify.</p>
	Regionname string `json:"RegionName"` // Specifies the Region-code for a given Region name (for example, <code>af-south-1</code>). This function will return the status of whatever Region you pass into this parameter.
}

// PutAlternateContactrequest represents the PutAlternateContactrequest schema from the OpenAPI specification
type PutAlternateContactrequest struct {
	Name string `json:"Name"` // Specifies a name for the alternate contact.
	Phonenumber string `json:"PhoneNumber"` // Specifies a phone number for the alternate contact.
	Title string `json:"Title"` // Specifies a title for the alternate contact.
	Accountid string `json:"AccountId,omitempty"` // <p>Specifies the 12 digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.</p> <p>If you do not specify this parameter, it defaults to the Amazon Web Services account of the identity used to call the operation.</p> <p>To use this parameter, the caller must be an identity in the <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account">organization's management account</a> or a delegated administrator account, and the specified account ID must be a member account in the same organization. The organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html">all features enabled</a>, and the organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html">trusted access</a> enabled for the Account Management service, and optionally a <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html">delegated admin</a> account assigned.</p> <note> <p>The management account can't specify its own <code>AccountId</code>; it must call the operation in standalone context by not including the <code>AccountId</code> parameter.</p> </note> <p>To call this operation on an account that is not a member of an organization, then don't specify this parameter, and call the operation using an identity belonging to the account whose contacts you wish to retrieve or modify.</p>
	Alternatecontacttype string `json:"AlternateContactType"` // Specifies which alternate contact you want to create or update.
	Emailaddress string `json:"EmailAddress"` // Specifies an email address for the alternate contact.
}

// Region represents the Region schema from the OpenAPI specification
type Region struct {
	Regionname interface{} `json:"RegionName,omitempty"`
	Regionoptstatus interface{} `json:"RegionOptStatus,omitempty"`
}

// DisableRegionrequest represents the DisableRegionrequest schema from the OpenAPI specification
type DisableRegionrequest struct {
	Accountid string `json:"AccountId,omitempty"` // <p>Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation. If you don't specify this parameter, it defaults to the Amazon Web Services account of the identity used to call the operation. To use this parameter, the caller must be an identity in the <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account">organization's management account</a> or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html">all features enabled</a>, and the organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html">trusted access</a> enabled for the Account Management service, and optionally a <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html">delegated admin</a> account assigned.</p> <note> <p>The management account can't specify its own <code>AccountId</code>. It must call the operation in standalone context by not including the <code>AccountId</code> parameter.</p> </note> <p>To call this operation on an account that is not a member of an organization, don't specify this parameter. Instead, call the operation using an identity belonging to the account whose contacts you wish to retrieve or modify.</p>
	Regionname string `json:"RegionName"` // Specifies the Region-code for a given Region name (for example, <code>af-south-1</code>). When you disable a Region, Amazon Web Services performs actions to deactivate that Region in your account, such as destroying IAM resources in the Region. This process takes a few minutes for most accounts, but this can take several hours. You cannot enable the Region until the disabling process is fully completed.
}

// GetAlternateContactResponseAlternateContact represents the GetAlternateContactResponseAlternateContact schema from the OpenAPI specification
type GetAlternateContactResponseAlternateContact struct {
	Name interface{} `json:"Name,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Title interface{} `json:"Title,omitempty"`
	Alternatecontacttype interface{} `json:"AlternateContactType,omitempty"`
	Emailaddress interface{} `json:"EmailAddress,omitempty"`
}

// AlternateContact represents the AlternateContact schema from the OpenAPI specification
type AlternateContact struct {
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Title interface{} `json:"Title,omitempty"`
	Alternatecontacttype interface{} `json:"AlternateContactType,omitempty"`
	Emailaddress interface{} `json:"EmailAddress,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// ListRegionsRequest represents the ListRegionsRequest schema from the OpenAPI specification
type ListRegionsRequest struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Regionoptstatuscontains interface{} `json:"RegionOptStatusContains,omitempty"`
}

// EnableRegionRequest represents the EnableRegionRequest schema from the OpenAPI specification
type EnableRegionRequest struct {
	Regionname interface{} `json:"RegionName"`
	Accountid interface{} `json:"AccountId,omitempty"`
}

// ListRegionsResponse represents the ListRegionsResponse schema from the OpenAPI specification
type ListRegionsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Regions interface{} `json:"Regions,omitempty"`
}

// GetRegionOptStatusRequest represents the GetRegionOptStatusRequest schema from the OpenAPI specification
type GetRegionOptStatusRequest struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Regionname interface{} `json:"RegionName"`
}

// PutContactInformationrequest represents the PutContactInformationrequest schema from the OpenAPI specification
type PutContactInformationrequest struct {
	Accountid string `json:"AccountId,omitempty"` // <p>Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation. If you don't specify this parameter, it defaults to the Amazon Web Services account of the identity used to call the operation. To use this parameter, the caller must be an identity in the <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account">organization's management account</a> or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html">all features enabled</a>, and the organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html">trusted access</a> enabled for the Account Management service, and optionally a <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html">delegated admin</a> account assigned.</p> <note> <p>The management account can't specify its own <code>AccountId</code>. It must call the operation in standalone context by not including the <code>AccountId</code> parameter.</p> </note> <p>To call this operation on an account that is not a member of an organization, don't specify this parameter. Instead, call the operation using an identity belonging to the account whose contacts you wish to retrieve or modify.</p>
	Contactinformation PutContactInformationrequestContactInformation `json:"ContactInformation"` // Contains the details of the primary contact information associated with an Amazon Web Services account.
}

// GetContactInformationrequest represents the GetContactInformationrequest schema from the OpenAPI specification
type GetContactInformationrequest struct {
	Accountid string `json:"AccountId,omitempty"` // <p>Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation. If you don't specify this parameter, it defaults to the Amazon Web Services account of the identity used to call the operation. To use this parameter, the caller must be an identity in the <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account">organization's management account</a> or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html">all features enabled</a>, and the organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html">trusted access</a> enabled for the Account Management service, and optionally a <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html">delegated admin</a> account assigned.</p> <note> <p>The management account can't specify its own <code>AccountId</code>. It must call the operation in standalone context by not including the <code>AccountId</code> parameter.</p> </note> <p>To call this operation on an account that is not a member of an organization, don't specify this parameter. Instead, call the operation using an identity belonging to the account whose contacts you wish to retrieve or modify.</p>
}

// ListRegionsrequest represents the ListRegionsrequest schema from the OpenAPI specification
type ListRegionsrequest struct {
	Accountid string `json:"AccountId,omitempty"` // <p>Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation. If you don't specify this parameter, it defaults to the Amazon Web Services account of the identity used to call the operation. To use this parameter, the caller must be an identity in the <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account">organization's management account</a> or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html">all features enabled</a>, and the organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html">trusted access</a> enabled for the Account Management service, and optionally a <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html">delegated admin</a> account assigned.</p> <note> <p>The management account can't specify its own <code>AccountId</code>. It must call the operation in standalone context by not including the <code>AccountId</code> parameter.</p> </note> <p>To call this operation on an account that is not a member of an organization, don't specify this parameter. Instead, call the operation using an identity belonging to the account whose contacts you wish to retrieve or modify.</p>
	Maxresults int `json:"MaxResults,omitempty"` // The total number of items to return in the command’s output. If the total number of items available is more than the value specified, a <code>NextToken</code> is provided in the command’s output. To resume pagination, provide the <code>NextToken</code> value in the <code>starting-token</code> argument of a subsequent command. Do not use the <code>NextToken</code> response element directly outside of the Amazon Web Services CLI. For usage examples, see <a href="http://docs.aws.amazon.com/cli/latest/userguide/pagination.html">Pagination</a> in the <i>Amazon Web Services Command Line Interface User Guide</i>.
	Nexttoken string `json:"NextToken,omitempty"` // A token used to specify where to start paginating. This is the <code>NextToken</code> from a previously truncated response. For usage examples, see <a href="http://docs.aws.amazon.com/cli/latest/userguide/pagination.html">Pagination</a> in the <i>Amazon Web Services Command Line Interface User Guide</i>.
	Regionoptstatuscontains []string `json:"RegionOptStatusContains,omitempty"` // A list of Region statuses (Enabling, Enabled, Disabling, Disabled, Enabled_by_default) to use to filter the list of Regions for a given account. For example, passing in a value of ENABLING will only return a list of Regions with a Region status of ENABLING.
}

// GetAlternateContactrequest represents the GetAlternateContactrequest schema from the OpenAPI specification
type GetAlternateContactrequest struct {
	Accountid string `json:"AccountId,omitempty"` // <p>Specifies the 12 digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.</p> <p>If you do not specify this parameter, it defaults to the Amazon Web Services account of the identity used to call the operation.</p> <p>To use this parameter, the caller must be an identity in the <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account">organization's management account</a> or a delegated administrator account, and the specified account ID must be a member account in the same organization. The organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html">all features enabled</a>, and the organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html">trusted access</a> enabled for the Account Management service, and optionally a <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html">delegated admin</a> account assigned.</p> <note> <p>The management account can't specify its own <code>AccountId</code>; it must call the operation in standalone context by not including the <code>AccountId</code> parameter.</p> </note> <p>To call this operation on an account that is not a member of an organization, then don't specify this parameter, and call the operation using an identity belonging to the account whose contacts you wish to retrieve or modify.</p>
	Alternatecontacttype string `json:"AlternateContactType"` // Specifies which alternate contact you want to retrieve.
}

// GetRegionOptStatusResponse represents the GetRegionOptStatusResponse schema from the OpenAPI specification
type GetRegionOptStatusResponse struct {
	Regionname interface{} `json:"RegionName,omitempty"`
	Regionoptstatus interface{} `json:"RegionOptStatus,omitempty"`
}

// PutContactInformationrequestContactInformation represents the PutContactInformationrequestContactInformation schema from the OpenAPI specification
type PutContactInformationrequestContactInformation struct {
	Stateorregion interface{} `json:"StateOrRegion,omitempty"`
	Websiteurl interface{} `json:"WebsiteUrl,omitempty"`
	City interface{} `json:"City,omitempty"`
	Companyname interface{} `json:"CompanyName,omitempty"`
	Countrycode interface{} `json:"CountryCode,omitempty"`
	Districtorcounty interface{} `json:"DistrictOrCounty,omitempty"`
	Addressline2 interface{} `json:"AddressLine2,omitempty"`
	Addressline1 interface{} `json:"AddressLine1,omitempty"`
	Addressline3 interface{} `json:"AddressLine3,omitempty"`
	Phonenumber interface{} `json:"PhoneNumber,omitempty"`
	Postalcode interface{} `json:"PostalCode,omitempty"`
	Fullname interface{} `json:"FullName,omitempty"`
}

// GetContactInformationResponseContactInformation represents the GetContactInformationResponseContactInformation schema from the OpenAPI specification
type GetContactInformationResponseContactInformation struct {
	Companyname interface{} `json:"CompanyName,omitempty"`
	Countrycode interface{} `json:"CountryCode"`
	Districtorcounty interface{} `json:"DistrictOrCounty,omitempty"`
	Addressline3 interface{} `json:"AddressLine3,omitempty"`
	City interface{} `json:"City"`
	Phonenumber interface{} `json:"PhoneNumber"`
	Postalcode interface{} `json:"PostalCode"`
	Addressline1 interface{} `json:"AddressLine1"`
	Fullname interface{} `json:"FullName"`
	Stateorregion interface{} `json:"StateOrRegion,omitempty"`
	Websiteurl interface{} `json:"WebsiteUrl,omitempty"`
	Addressline2 interface{} `json:"AddressLine2,omitempty"`
}

// DisableRegionRequest represents the DisableRegionRequest schema from the OpenAPI specification
type DisableRegionRequest struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Regionname interface{} `json:"RegionName"`
}

// GetContactInformationResponse represents the GetContactInformationResponse schema from the OpenAPI specification
type GetContactInformationResponse struct {
	Contactinformation GetContactInformationResponseContactInformation `json:"ContactInformation,omitempty"`
}

// DeleteAlternateContactrequest represents the DeleteAlternateContactrequest schema from the OpenAPI specification
type DeleteAlternateContactrequest struct {
	Accountid string `json:"AccountId,omitempty"` // <p>Specifies the 12 digit account ID number of the Amazon Web Services account that you want to access or modify with this operation.</p> <p>If you do not specify this parameter, it defaults to the Amazon Web Services account of the identity used to call the operation.</p> <p>To use this parameter, the caller must be an identity in the <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account">organization's management account</a> or a delegated administrator account, and the specified account ID must be a member account in the same organization. The organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html">all features enabled</a>, and the organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html">trusted access</a> enabled for the Account Management service, and optionally a <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html">delegated admin</a> account assigned.</p> <note> <p>The management account can't specify its own <code>AccountId</code>; it must call the operation in standalone context by not including the <code>AccountId</code> parameter.</p> </note> <p>To call this operation on an account that is not a member of an organization, then don't specify this parameter, and call the operation using an identity belonging to the account whose contacts you wish to retrieve or modify.</p>
	Alternatecontacttype string `json:"AlternateContactType"` // Specifies which of the alternate contacts to delete.
}

// PutContactInformationRequest represents the PutContactInformationRequest schema from the OpenAPI specification
type PutContactInformationRequest struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Contactinformation GetContactInformationResponseContactInformation `json:"ContactInformation"`
}

// GetAlternateContactRequest represents the GetAlternateContactRequest schema from the OpenAPI specification
type GetAlternateContactRequest struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Alternatecontacttype interface{} `json:"AlternateContactType"`
}

// PutAlternateContactRequest represents the PutAlternateContactRequest schema from the OpenAPI specification
type PutAlternateContactRequest struct {
	Phonenumber interface{} `json:"PhoneNumber"`
	Title interface{} `json:"Title"`
	Accountid interface{} `json:"AccountId,omitempty"`
	Alternatecontacttype interface{} `json:"AlternateContactType"`
	Emailaddress interface{} `json:"EmailAddress"`
	Name interface{} `json:"Name"`
}

// ContactInformation represents the ContactInformation schema from the OpenAPI specification
type ContactInformation struct {
	Addressline2 interface{} `json:"AddressLine2,omitempty"`
	Companyname interface{} `json:"CompanyName,omitempty"`
	Countrycode interface{} `json:"CountryCode"`
	Districtorcounty interface{} `json:"DistrictOrCounty,omitempty"`
	Addressline3 interface{} `json:"AddressLine3,omitempty"`
	City interface{} `json:"City"`
	Phonenumber interface{} `json:"PhoneNumber"`
	Postalcode interface{} `json:"PostalCode"`
	Addressline1 interface{} `json:"AddressLine1"`
	Fullname interface{} `json:"FullName"`
	Stateorregion interface{} `json:"StateOrRegion,omitempty"`
	Websiteurl interface{} `json:"WebsiteUrl,omitempty"`
}

// GetAlternateContactResponse represents the GetAlternateContactResponse schema from the OpenAPI specification
type GetAlternateContactResponse struct {
	Alternatecontact GetAlternateContactResponseAlternateContact `json:"AlternateContact,omitempty"`
}

// EnableRegionrequest represents the EnableRegionrequest schema from the OpenAPI specification
type EnableRegionrequest struct {
	Accountid string `json:"AccountId,omitempty"` // <p>Specifies the 12-digit account ID number of the Amazon Web Services account that you want to access or modify with this operation. If you don't specify this parameter, it defaults to the Amazon Web Services account of the identity used to call the operation. To use this parameter, the caller must be an identity in the <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account">organization's management account</a> or a delegated administrator account. The specified account ID must also be a member account in the same organization. The organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html">all features enabled</a>, and the organization must have <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html">trusted access</a> enabled for the Account Management service, and optionally a <a href="https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html">delegated admin</a> account assigned.</p> <note> <p>The management account can't specify its own <code>AccountId</code>. It must call the operation in standalone context by not including the <code>AccountId</code> parameter.</p> </note> <p>To call this operation on an account that is not a member of an organization, don't specify this parameter. Instead, call the operation using an identity belonging to the account whose contacts you wish to retrieve or modify.</p>
	Regionname string `json:"RegionName"` // Specifies the Region-code for a given Region name (for example, <code>af-south-1</code>). When you enable a Region, Amazon Web Services performs actions to prepare your account in that Region, such as distributing your IAM resources to the Region. This process takes a few minutes for most accounts, but it can take several hours. You cannot use the Region until this process is complete. Furthermore, you cannot disable the Region until the enabling process is fully completed.
}

// GetContactInformationRequest represents the GetContactInformationRequest schema from the OpenAPI specification
type GetContactInformationRequest struct {
	Accountid interface{} `json:"AccountId,omitempty"`
}

// DeleteAlternateContactRequest represents the DeleteAlternateContactRequest schema from the OpenAPI specification
type DeleteAlternateContactRequest struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Alternatecontacttype interface{} `json:"AlternateContactType"`
}
