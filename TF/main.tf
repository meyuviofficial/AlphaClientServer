terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.0.0"
    }
  }
}

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

resource "azurerm_resource_group" "AlphaClientServer" {
  name     = var.ResourceGroup
  location = var.Location
  tags = {
    Project = "AlphaClientServer"
  }
}

resource "azurerm_kubernetes_cluster" "AlphaClientServer_aks" {
  name                = var.AksName
  location            = azurerm_resource_group.AlphaClientServer.location
  resource_group_name = azurerm_resource_group.AlphaClientServer.name
  dns_prefix          = "${var.AksName}-DNS-k8s"

  default_node_pool {
    name            = "default"
    node_count      = 2
    vm_size         = "Standard_D2_v2"
    os_disk_size_gb = 30
  }

  service_principal {
    client_id     = var.ApplicationId
    client_secret = var.ClientSecret
  }
  tags = {
    Project = "AlphaClientServer"
  }
}

resource "azurerm_container_registry" "AlphaClientServer_acr" {
  name                = var.AcrName
  resource_group_name = azurerm_resource_group.AlphaClientServer.name
  location            = azurerm_resource_group.AlphaClientServer.location
  sku                 = "Basic"
  admin_enabled       = false
  georeplications {
    location                = "East US"
    zone_redundancy_enabled = true
    tags                    = {}
  }
}

resource "azurerm_role_assignment" "AlphaClientServer-RoleAssignment" {
  principal_id                     = azurerm_kubernetes_cluster.AlphaClientServer_aks.kubelet_identity[0].object_id
  role_definition_name             = "AcrPull"
  scope                            = azurerm_container_registry.AlphaClientServer_acr.id
  skip_service_principal_aad_check = true
}