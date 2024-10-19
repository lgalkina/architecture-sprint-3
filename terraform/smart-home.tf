resource "helm_release" "smart-home" {
  name       = "smart-home"
  namespace  = "default"
  chart      = "../charts/smart-home"
  timeout    = 600  # Increase the timeout to 10 minutes
}