package util

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vpc "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vpc/v20170312"
)

func GetLocation() string {
	credential := common.NewCredential(
		"AKIDf8KGvkNkslOkslOlunQnumOkpgHchX8T",
		"75nEOHtEI5csrZ1D8sNefRyGI4asugCS",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "vpc.tencentcloudapi.com"
	client, _ := vpc.NewClient(credential, "ap-beijing", cpf)

	request := vpc.NewDescribeIpGeolocationInfosRequest()

	request.AddressIps = common.StringPtrs([]string{"222.190.107.198"})
	request.Fields = &vpc.IpField{
		Country:  common.BoolPtr(true),
		Province: common.BoolPtr(true),
		City:     common.BoolPtr(true),
		Region:   common.BoolPtr(true),
		Isp:      common.BoolPtr(true),
	}

	response, err := client.DescribeIpGeolocationInfos(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)

	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
	return response.ToJsonString()
}
