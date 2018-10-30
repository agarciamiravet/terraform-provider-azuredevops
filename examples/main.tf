resource "azuredevops_project" "projectTest" {

     oraganization = "MyOrganization"
     pat = "PersonalAccessToken"
     projectname = "Projectname"
     description = "Description"
     summary = "Summary"
}

resource "azuredevops_workitemtask" "terraformworkitemtest" {

     oraganization = "MyOrganization"
     pat = "PersonalAccessToken"
     projectname = "Projectname"
     title = "My First Work Item"     
}
