package cloud

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type MockEC2API struct {
	ec2iface.EC2API       // embedding of the interface is needed to skip implementation of all methods
	DescribeVolumesMethod func(*ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error)
}

func (m *MockEC2API) DescribeVolumes(in *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	if m.DescribeVolumesMethod != nil {
		return m.DescribeVolumesMethod(in)
	}
	volumeID := "vol-123456"
	tagKey := "Name"
	tagValue := "DIRE"
	tag := ec2.Tag{Key: &tagKey, Value: &tagValue}

	return &ec2.DescribeVolumesOutput{Volumes: []*ec2.Volume{{VolumeId: &volumeID, Tags: []*ec2.Tag{&tag}}}}, nil
}

func TestEBSCompliance(t *testing.T) {
	//Mock the service
	mockSvc := &MockEC2API{}

	tags := make(map[string]string)
	tags["Name"] = "Cat"
	got := EBSCompliance(mockSvc, tags)
	want := "vol-1234567"

	if got != want {
		t.Errorf("wanted calls %s got %s", want, got)
	}
}