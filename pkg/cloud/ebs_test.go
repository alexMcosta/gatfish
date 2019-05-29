package cloud

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type mockedVolume struct {
	ec2iface.EC2API
}

// WHY AM I TRYING TO MOCK THIS EVIL THING?
func (m *mockedVolume) DescribeVolumes(*ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	type Volume struct {
		Attachments      []*VolumeAttachment `locationName:"attachmentSet" locationNameList:"item" type:"list"`
		AvailabilityZone *string             `locationName:"availabilityZone" type:"string"`
		CreateTime       *time.Time          `locationName:"createTime" type:"timestamp"`
		Encrypted        *bool               `locationName:"encrypted" type:"boolean"`
		Iops             *int64              `locationName:"iops" type:"integer"`
		KmsKeyId         *string             `locationName:"kmsKeyId" type:"string"`
		Size             *int64              `locationName:"size" type:"integer"`
		SnapshotId       *string             `locationName:"snapshotId" type:"string"`
		State            *string             `locationName:"status" type:"string" enum:"VolumeState"`
		Tags             []*Tag              `locationName:"tagSet" locationNameList:"item" type:"list"`
		VolumeId         *string             `locationName:"volumeId" type:"string"`
		VolumeType       *string             `locationName:"volumeType" type:"string" enum:"VolumeType"`
	}

	v := Volume{
		AvailabilityZone: "us-west-1b",
		CreateTime:       "2019-05-29 04:05:45.082 +0000 UTC",
		Encrypted:        false,
		Iops:             300,
		Size:             100,
		SnapshotId:       "",
		State:            "available",
		Tags: []*Tag{
			Key:   "Name",
			Value: "Alex",
		}, {
			Key:   "Origin",
			Value: "Portland",
		},
		VolumeId:   "vol-090a7b45e2059a989",
		VolumeType: "gp2",
	}

	return
}

func TestEBSCompliance(t *testing.T) {
	//Mock the service
	mockSvc := &mockEC2Client{}

	tags := make(map[string]string)
	tags["CAT"] = "REQUIRED"
	got := EBSCompliance(mockSvc, tags)
	want := "vol-1234567"

	if got != want {
		t.Errorf("wanted calls %s got %s", want, got)
	}
}
