package panel

import (
	"fmt"
	"path"
	"time"

	serverv1 "github.com/perfect-panel/ppanel-node/api/server/v1"
)

type NodeInfo struct {
	Id                     int
	Type                   string
	PushInterval           int
	PullInterval           int
	TrafficReportThreshold int
	ACMEEmail              string
	ACMECADirURL           string
	Protocol               *Protocol
}

type ServerPushStatusRequest struct {
	Cpu       float64 `json:"cpu"`
	Mem       float64 `json:"mem"`
	Disk      float64 `json:"disk"`
	UpdatedAt int64   `json:"updated_at"`
}

type NodeStatus struct {
	CPU    float64
	Mem    float64
	Disk   float64
	Uptime uint64
}

func (c *NodeClient) ReportNodeStatus(nodeStatus *NodeStatus) (err error) {
	if c.UseProtobuf {
		return c.reportNodeStatusProtobuf(nodeStatus)
	}
	p := "/v1/server/status"
	status := ServerPushStatusRequest{
		Cpu:       nodeStatus.CPU,
		Mem:       nodeStatus.Mem,
		Disk:      nodeStatus.Disk,
		UpdatedAt: time.Now().UnixMilli(),
	}
	r, err := c.Client.R().SetBody(status).ForceContentType("application/json").Post(p)
	if err != nil {
		return fmt.Errorf("访问 %s 失败: %v", path.Join(c.APIHost+p), err.Error())
	}
	return checkPanelResponse(r, path.Join(c.APIHost+p))
}

func (c *NodeClient) reportNodeStatusProtobuf(nodeStatus *NodeStatus) error {
	const p = "/v1/server/status"
	request := c.Client.R()
	if err := setProtobufRequestBody(request, &serverv1.PushServerStatusRequest{
		Cpu:       nodeStatus.CPU,
		Mem:       nodeStatus.Mem,
		Disk:      nodeStatus.Disk,
		UpdatedAt: time.Now().UnixMilli(),
	}); err != nil {
		return err
	}
	r, err := request.Post(p)
	if err != nil {
		return fmt.Errorf("访问 %s 失败: %v", path.Join(c.APIHost+p), err)
	}
	return checkPanelResponse(r, path.Join(c.APIHost+p))
}
