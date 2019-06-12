#terraform-provider-azuredevops


This repository contains a provider to work with Azure devOps.Next weeks I will update with new resources like create repositories, work items ...  

# Warning

Before using the token you need encode it in base64.

# Resources

- azuredevops_project: 
   Create a new project in Azure Devops Organization

- azuredevops_workitemtask: 
   Create a new wotkitem of type task in a project in Azure Devops Organization
   
# Usage

resource "azuredevops_project" "projecttest" {

     organization = "[organization_azuredevops_name]"
     pat = "[Personal_access_token_in_base64]"
     projectname = "[name_project_to_create]"
     description = "[description_project]"
     summary = "[short_description]"
}

resource "azuredevops_workitemtask" "workitemtasktest" {

     organization = "[organization_azuredevops_name]"
     pat = "[Personal_access_token_in_base64]"
     projectname = "[name_project_to_create]"
     title = "[title_workitem]"     
}

