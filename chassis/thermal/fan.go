package thermal

import (
	"github.com/MauveSoftware/ilo4_exporter/common"
)

type Fan struct {
	Name           string        `json:"FanName"`
	CurrentReading float64       `json:"CurrentReading"`
	Status         common.Status `json:"Status"`
}
