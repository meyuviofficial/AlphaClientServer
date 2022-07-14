locals {
  count = var.NodeCount
}
resource "azurerm_virtual_network" "AlphaClientVNET" {
  name                = var.VirtualNetwork
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.AlphaClientServer.location
  resource_group_name = azurerm_resource_group.AlphaClientServer.name
}

resource "azurerm_subnet" "AlphaClientSubnet" {
  name                 = var.Subnet
  resource_group_name  = azurerm_resource_group.AlphaClientServer.name
  virtual_network_name = azurerm_virtual_network.AlphaClientVNET.name
  address_prefixes     = ["10.0.2.0/24"]
}

resource "azurerm_network_interface" "VirtualMachineNIC" {
  count               = local.count
  name                = "${var.VirtualMachineName}-nic-${count.index}"
  location            = azurerm_resource_group.AlphaClientServer.location
  resource_group_name = azurerm_resource_group.AlphaClientServer.name

  ip_configuration {
    name                          = "VirtualMachine-Config"
    subnet_id                     = azurerm_subnet.AlphaClientSubnet.id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_virtual_machine" "AlphaClient" {
  count                 = local.count
  name                  = "${var.VirtualMachineName}-${count.index}"
  location              = azurerm_resource_group.AlphaClientServer.location
  resource_group_name   = azurerm_resource_group.AlphaClientServer.name
  network_interface_ids = [element(azurerm_network_interface.VirtualMachineNIC.*.id, count.index)]
  vm_size               = "Standard_DS1_v2"

  delete_os_disk_on_termination = true

  delete_data_disks_on_termination = true

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }
  storage_os_disk {
    name              = "OsDisk-${count.index}"
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }
  os_profile {
    computer_name  = var.VirtualMachineName
    admin_username = var.Admin
    admin_password = var.Password
  }
  os_profile_linux_config {
    disable_password_authentication = false
  }
}
