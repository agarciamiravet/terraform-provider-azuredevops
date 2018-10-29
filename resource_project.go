package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// resourceIssue is used to define a JIRA issue
func resourceAzureDevOpsProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
			"organization": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"pat": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"projectname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"summary": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceProjectCreate(d *schema.ResourceData, m interface{}) error {
	return resourceProjectRead(d, m)
}

func resourceProjectRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceProjectUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceProjectRead(d, m)
}

func resourceProjectDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
