package test

import (
	"fmt"
	"github.com/mihang-laoban/PUB-GO-BASE/utils"
	"testing"
	"time"
)

func TestTimeConvert(t *testing.T) {
	now := time.Now().Unix()
	ti := utils.Unix2Time(now)
	fmt.Println(ti)
	ui := utils.Time2Unix(ti)
	fmt.Println(ui)
	fmt.Println(utils.Unix2Time(ui))
}

