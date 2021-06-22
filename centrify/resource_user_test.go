package centrify

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/hashicorp/terraform/helper/acctest"
	vault "github.com/marcozj/golang-sdk/platform"
	"github.com/marcozj/golang-sdk/restapi"
)

func TestAccResourceUserCreation(t *testing.T) {
	rName := acctest.RandomWithPrefix("test-user@")
	resourceName := "centrify_user.testuser"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBasicValExists(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "username", rName),
					resource.TestCheckResourceAttr(resourceName, "email", "test@example.com"),
					resource.TestCheckResourceAttr(resourceName, "displayname", "Test User"),
					resource.TestCheckResourceAttr(resourceName, "password", "TestUser@123"),
					resource.TestCheckResourceAttr(resourceName, "description", "Test User"),
				),
			},
		},
	})
}

func testAccBasicValExists(rName string) string {

	return fmt.Sprintf(`resource "centrifyvault_user" "testuser" {
		username = %[1]q
		email = "test@example.com"
		displayname = "Test User"
		description = "Test User"
		password = "TestUser@123"
	}`, rName)
}

func testAccCheckUserDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*restapi.RestClient)
	object := vault.NewUser(client)
	for _, res := range s.RootModule().Resources {
		if res.Type != "centrify_user" {
			continue
		}
		object.ID = res.Primary.ID
		err := object.Read()
		if err == nil {
			return fmt.Errorf("User Still Exists")
		}

		if err != nil {
			notFoundErr := "not found"
			expectedErr := regexp.MustCompile(notFoundErr)
			if !expectedErr.Match([]byte(err.Error())) {
				return fmt.Errorf("expected %s, got %s", notFoundErr, err)
			}
		}
	}
	return nil

}
