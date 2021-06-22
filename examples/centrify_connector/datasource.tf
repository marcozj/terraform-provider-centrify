data "centrify_connector" "connector1" {
    name = "connector_host1" // Connector name registered in Centrify
    status = "Active"
}

output "connector1_machine_name" {
  value = data.centrify_connector.connector1.machine_name
}
output "connector1_ssh_service" {
  value = data.centrify_connector.connector1.ssh_service
}
output "connector1_rdp_service" {
  value = data.centrify_connector.connector1.rdp_service
}
output "connector1_ad_proxy" {
  value = data.centrify_connector.connector1.ad_proxy
}
output "connector1_app_gateway" {
  value = data.centrify_connector.connector1.app_gateway
}
output "connector1_http_api_service" {
  value = data.centrify_connector.connector1.http_api_service
}
output "connector1_ldap_proxy" {
  value = data.centrify_connector.connector1.ldap_proxy
}
output "connector1_radius_service" {
  value = data.centrify_connector.connector1.radius_service
}
output "connector1_radius_external_service" {
  value = data.centrify_connector.connector1.radius_external_service
}
output "connector1_version" {
  value = data.centrify_connector.connector1.version
}
output "connector1_vpc_identifier" {
  value = data.centrify_connector.connector1.vpc_identifier
}
output "connector1_forest" {
  value = data.centrify_connector.connector1.forest
}
output "connector1_dns_host_name" {
  value = data.centrify_connector.connector1.dns_host_name
}
