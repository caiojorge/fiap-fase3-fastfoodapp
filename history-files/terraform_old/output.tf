output "ecr_repository_url" {
  description = "URL do repositório ECR"
  value       = aws_ecr_repository.fiap_rocks.repository_url
}