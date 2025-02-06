// /internal/modules/ads/inserter.go
package ads

import (
	"fmt"
	"os/exec"
)

type AdInserter struct{}

func (ai *AdInserter) InsertAd(videoPath, adText, adImage, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vf", "drawtext=text="+adText, outputPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to insert ad: %w", err)
	}
	return nil
}
