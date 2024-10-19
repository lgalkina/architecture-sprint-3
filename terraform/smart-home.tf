resource "helm_release" "smart-home" {
  name       = "smart-home"
  namespace  = "default"
  chart      = "../charts/smart-home"
}