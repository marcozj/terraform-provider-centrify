data "centrify_policy" "Invited_Users" {
    name = "Invited Users"
}

data "centrify_policy" "User_Login_Policy" {
    name = "LAB User Login Policy"
}

data "centrify_policy" "Default_Policy" {
    name = "Default Policy"
}

data "centrify_policy" "Deny_Login_Policy" {
    name = "LAB Deny Login Policy"
}

resource "centrify_policyorder" "policy_order" {
    policy_order = [
        data.centrify_policy.Invited_Users.id,
        data.centrify_policy.User_Login_Policy.id,
        centrify_policy.test_policy.id,
        data.centrify_policy.Default_Policy.id,
        data.centrify_policy.Deny_Login_Policy.id,
    ]
}
