provider "aws" {
  region = "us-east-1"  # Altere para a região desejada
}

resource "aws_ecr_repository" "fiap_rocks" {
  name = "fiap-rocks"
}