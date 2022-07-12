terraform {
  required_providers {
    aws = {
      version = ">=1.2"
      source  = "hashicorp/aws"
    }
  }
}

terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "NonCoder"
	  token = "GWzAAi0SxtOzJQ.atlasv1.ceiiZVagysMFoPqWNzCPH2Fhj8c6kxIeZztNwc0euOrlLUUshROySqzk8vj80u6uLzg"
	
    workspaces {
      name = "AlphaClientServer"
    }
  }
}
