data "centrify_authenticationprofile" "new_device" {
    name = "Default New Device Login Profile"
}
output "id" {
    value = data.centrify_authenticationprofile.new_device.id
}
output "uuid" {
    value = data.centrify_authenticationprofile.new_device.uuid
}
output "name" {
    value = data.centrify_authenticationprofile.new_device.name
}
output "pass_through_duration" {
    value = data.centrify_authenticationprofile.new_device.pass_through_duration
}
output "number_of_questions" {
    value = data.centrify_authenticationprofile.new_device.additional_data[0].number_of_questions
}
output "challenges" {
    value = data.centrify_authenticationprofile.new_device.challenges
}