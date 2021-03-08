package alicloud

import (
	"fmt"
	"testing"

	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccAlicloudEcsCommand_basic(t *testing.T) {
	var v map[string]interface{}
	resourceId := "alicloud_ecs_command.default"
	ra := resourceAttrInit(resourceId, AlicloudEcsCommandMap)
	rc := resourceCheckInitWithDescribeMethod(resourceId, &v, func() interface{} {
		return &EcsService{testAccProvider.Meta().(*connectivity.AliyunClient)}
	}, "DescribeEcsCommand")
	rac := resourceAttrCheckInit(rc, ra)
	testAccCheck := rac.resourceAttrMapUpdateSet()
	rand := acctest.RandIntRange(10000, 99999)
	name := fmt.Sprintf("tf-testAcc%sAlicloudEcsCommand%d", defaultRegionToTest, rand)
	testAccConfig := resourceTestAccConfigFunc(resourceId, name, AlicloudEcsCommandBasicDependence)
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},

		IDRefreshName: resourceId,
		Providers:     testAccProviders,
		CheckDestroy:  rac.checkResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccConfig(map[string]interface{}{
					"command_content": "bHMK",
					"description":     "For Terraform Test",
					"name":            name,
					"type":            "RunShellScript",
					"working_dir":     "/root",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"command_content": "bHMK",
						"description":     "For Terraform Test",
						"name":            name,
						"type":            "RunShellScript",
						"working_dir":     "/root",
					}),
				),
			},
			{
				ResourceName:      resourceId,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"name": name + "Update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"name": name + "Update",
					}),
				),
			},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"working_dir": "/home/",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"working_dir": "/home/",
			//		}),
			//	),
			//},
			{
				Config: testAccConfig(map[string]interface{}{
					"timeout": "70",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"timeout": "70",
					}),
				),
			},
			{
				Config: testAccConfig(map[string]interface{}{
					"description": "For Terraform Test Update",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						"description": "For Terraform Test Update",
					}),
				),
			},
			//{
			//	Config: testAccConfig(map[string]interface{}{
			//		"command_content": "cHdk",
			//	}),
			//	Check: resource.ComposeTestCheckFunc(
			//		testAccCheck(map[string]string{
			//			"command_content": "cHdk",
			//		}),
			//	),
			//},
			{
				Config: testAccConfig(map[string]interface{}{
					//"command_content": "bHMK",
					"description": "For Terraform Test",
					"name":        name,
					//"working_dir":     "/root",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheck(map[string]string{
						//"command_content": "bHMK",
						"description": "For Terraform Test",
						"name":        name,
						//"working_dir":     "/root",
					}),
				),
			},
		},
	})
}

var AlicloudEcsCommandMap = map[string]string{
	"enable_parameter": "false",
}

func AlicloudEcsCommandBasicDependence(name string) string {
	return ""
}
