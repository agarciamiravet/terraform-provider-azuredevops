package main

import (
	"strconv"

	azuredevopsapi "github.com/agarciamiravet/go-azuredevopsapi"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAzureDevOpsWorkItemTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceWorkItemCreate,
		Read:   resourceWorkItemRead,
		Update: resourceWorkItemUpdate,
		Delete: resourceWorkItemDelete,

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
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceWorkItemCreate(d *schema.ResourceData, m interface{}) error {

	organization := d.Get("organization").(string)
	pat := d.Get("pat").(string)
	projectname := d.Get("projectname").(string)
	title := d.Get("title").(string)

	data := azuredevopsapi.CreateWorkItem(pat, organization, projectname, title)

	d.SetId(strconv.Itoa(data.ID))

	return resourceProjectRead(d, m)
}

func resourceWorkItemRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceWorkItemUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceProjectRead(d, m)
}

func resourceWorkItemDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
