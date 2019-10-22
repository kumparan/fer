package init

import (
	"fmt"
	"os"

	"github.com/kumparan/fer/installer"
)

// InitChangelog :nodoc:
func InitChangelog() {
	fmt.Println("Generate changelog configuration")
	err := os.Mkdir(".chglog", os.ModePerm)
	if err != nil {
		installer.ProgressBar(1)
		fmt.Println(err)
		fmt.Println("fail generate chglog")
	}
}
