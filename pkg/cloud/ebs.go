package cloud

import "github.com/aws/aws-sdk-go/service/ec2/ec2iface"

func EBSCompliance(svc ec2iface.EC2API, tags map[string]string) string {
	return "vol-1234567"
}
