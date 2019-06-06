package cloud

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// EBSCompliance searches EBS volumes to see if they have the requested tag keys
// If they do not it will store the volume ID and the level of severity and output the results
func EBSCompliance(svc ec2iface.EC2API, tags map[string]string) map[string][]string {

	culpritIDs := make(map[string][]string)

	// Go get the volumes
	volumes, err := svc.DescribeVolumes(&ec2.DescribeVolumesInput{})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
				os.Exit(1)
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	// Check every volume tag for the Tag name in question
	for tagN, tagV := range tags {
		for _, vol := range volumes.Volumes {
			if vol.Tags == nil {
				culpritIDs[tagV] = append(culpritIDs[tagV], *vol.VolumeId)
			}
			for _, valTag := range vol.Tags {
				switch *valTag.Key {
				case tagN:
					continue
				default:
					culpritIDs[tagV] = append(culpritIDs[tagV], *vol.VolumeId)
				}
			}
		}
	}
	// fmt.Printf("%v", volumes)
	return culpritIDs
}
