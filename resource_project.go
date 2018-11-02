package main

import (
	azuredevopsapi "github.com/agarciamiravet/go-azuredevopsapi"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAzureDevOpsProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceProjectCreate,
		Read:   resourceProjectRead,
		Update: resourceProjectUpdate,
		Delete: resourceProjectDelete,

		Schema: map[string]*schema.Schema{
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

	provider := m.(*Client)

	//organization := d.Get("organization").(string)
	//pat := d.Get("pat").(string)
	projectname := d.Get("projectname").(string)

	data := azuredevopsapi.CreateProject(provider.config.Pat, provider.config.Organization, projectname)

	d.SetId(data.ID)

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
