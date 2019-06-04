package cloud

import (
	"reflect"
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
	volumeID := "vol-1234567"
	tagKey := "Name"
	tagValue := "Baron"
	tag := ec2.Tag{Key: &tagKey, Value: &tagValue}

	return &ec2.DescribeVolumesOutput{Volumes: []*ec2.Volume{{VolumeId: &volumeID, Tags: []*ec2.Tag{&tag}}}}, nil
}

func TestEBSCompliance(t *testing.T) {
	//Mock the service
	mockSvc := &MockEC2API{}

	want := make(map[string][]string)

	//Mock the output of resources resources/tags.go
	tags := make(map[string]string)
	tags["Notname"] = "DIRE"

	got := EBSCompliance(mockSvc, tags)
	want["DIRE"] = append(want["DIRE"], "vol-1234567")

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted calls %v got %v", want, got)
	}
}
