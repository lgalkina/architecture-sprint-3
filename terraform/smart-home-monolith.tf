resource "helm_release" "smart-home-monolith" {
  name       = "smart-home-monolith"
  namespace  = "default"
  chart      = "../charts/smart-home-monolith"
  timeout    = 600  # Increase the timeout to 10 minutes
}