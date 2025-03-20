resource "aws_ecr_repository" "fiap_rocks" {
  name                 = var.ecr_repository_name
  image_tag_mutability = "MUTABLE"
}

resource "aws_ecr_lifecycle_policy" "fiap_rocks_lifecycle" {
  repository = aws_ecr_repository.fiap_rocks.name
  policy     = <<EOF
{
  "rules": [
    {
      "rulePriority": 1,
      "description": "Remove old images",
      "selection": {
        "tagStatus": "any",
        "countType": "imageCountMoreThan",
        "countNumber": 10
      },
      "action": {
        "type": "expire"
      }
    }
  ]
}
EOF
}