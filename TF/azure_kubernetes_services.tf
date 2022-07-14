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
  admin_enabled       = true
}