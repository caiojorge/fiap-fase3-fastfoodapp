provider "aws" {
  region = var.aws_region

  # Se precisar assumir RoleLab:
  # assume_role {
  #   role_arn = "arn:aws:iam::<IDCONTA>:role/RoleLab"
  # }
}

