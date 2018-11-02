package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider returns a terraform.ResourceProvider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"organization": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Azure DevOps organization",
			},

			"pat": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password access token",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"azuredevops_project":      resourceAzureDevOpsProject(),
			"azuredevops_workitemtask": resourceAzureDevOpsWorkItemTask(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		Organization: d.Get("organization").(string),
		Pat:          d.Get("pat").(string),
	}
	return config.Client()
}
