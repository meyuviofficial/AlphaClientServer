# Defining the Terraform Providers
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.0.0"
    }
  }
}

# Establishing the Backend connectivity to TF Cloud
terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "NonCoder"

    workspaces {
      name = "AlphaClientServer"
    }
  }
}

provider "azurerm" {
  features {}
}