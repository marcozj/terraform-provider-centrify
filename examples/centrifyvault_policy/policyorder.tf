data "centrifyvault_policy" "Invited_Users" {
    name = "Invited Users"
}

data "centrifyvault_policy" "User_Login_Policy" {
    name = "LAB User Login Policy"
}

data "centrifyvault_policy" "Default_Policy" {
    name = "Default Policy"
}

data "centrifyvault_policy" "Deny_Login_Policy" {
    name = "LAB Deny Login Policy"
}

resource "centrifyvault_policyorder" "policy_order" {
    policy_order = [
        data.centrifyvault_policy.Invited_Users.id,
        data.centrifyvault_policy.User_Login_Policy.id,
        centrifyvault_policy.test_policy.id,
        data.centrifyvault_policy.Default_Policy.id,
        data.centrifyvault_policy.Deny_Login_Policy.id,
    ]
}
