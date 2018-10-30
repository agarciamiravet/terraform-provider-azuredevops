package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider returns a terraform.ResourceProvider
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"azuredevops_project":      resourceAzureDevOpsProject(),
			"azuredevops_workitemtask": resourceAzureDevOpsWorkItemTask(),
		},
	}
}
