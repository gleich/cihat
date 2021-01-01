package lights

import (
	"github.com/Matt-Gleich/cihat/pkg/api"
	"github.com/Matt-Gleich/logoru"
	"github.com/nathany/bobblehat/sense/screen"
	"github.com/nathany/bobblehat/sense/screen/color"
)

// Update the lights on the sense hat.
func UpdateLights(queryContent api.QueryOutline) {
	var (
		repos = queryContent.Data.Viewer.Repositories.Edges
		y     int
		x     int
		fb    = screen.NewFrameBuffer()
	)

	for _, repo := range repos {
		if y == 8 {
			break
		}
		var hasValidChecks bool
		checks := repo.Node.DefaultBranchRef.Target.CheckSuites.Nodes
		for _, check := range checks {
			if x == 8 {
				break
			}

			var c color.Color
			switch check.Status {
			case "COMPLETED":
				if check.Conclusion == "SUCCESS" {
					c = color.New(0, 255, 0)
				} else if check.Conclusion == "FAILURE" {
					c = color.New(255, 0, 0)
				}
			case "IN_PROGRESS":
				c = color.New(255, 255, 0)
			default:
				continue
			}
			hasValidChecks = true
			fb.SetPixel(x, y, c)
			x++
		}

		if hasValidChecks {
			x = 0
			y++
		}
	}

	err := screen.Draw(fb)
	if err != nil {
		logoru.Error("Failed to update the lights:", err)
	}

	logoru.Success("Updated lights!")
}
