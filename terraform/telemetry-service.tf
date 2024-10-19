resource "helm_release" "telemetry-service" {
  name       = "telemetry-service"
  namespace  = "default"
  chart      = "../charts/smart-home/telemetry-service"
  timeout    = 1800  # Increase the timeout to 30 minutes
  force_update = true
  version = "1.0.0"
}