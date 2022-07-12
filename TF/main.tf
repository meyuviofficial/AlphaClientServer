terraform {
  required_providers {
    aws = {
      version = ">=1.2"
      source  = "hashicorp/aws"
    }
  }
}

terraform {
  cloud {
    organization = "NonCoder"
    
    workspaces {
      name = "AlphaClientServer"
    }
  }
}
