resource "helm_release" "device-service" {
  name       = "device-service"
  namespace  = "default"
  chart      = "../charts/smart-home/device-service"
  timeout    = 1800  # Increase the timeout to 30 minutes
  force_update = true
  version = "1.0.0"
}