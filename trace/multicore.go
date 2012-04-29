package trace

import (
	"raytracer/color"
	"raytracer/view"
)

type pixelSegment struct {
	Min, Max struct {
		X, Y int
	}
}

// Generates all the pixels in a range. Indicates completion on the given
// channel when done.
func makePixelSegment(done chan bool, m *view.Model, i color.Image,
	s *pixelSegment) {

	// Loop over the rows.
	for y := s.Min.Y; y < s.Max.Y; y++ {
		// Loop over the columns.
		for x := s.Min.X; x < s.Max.X; x++ {
			// Create the pixel
			makeAntiAliasedPixel(m, x, y, i)
		}
	}
	// Notify completion
	done <- true
	// Exit
}

func generateWorkSegments(height, width, numCores int) []pixelSegment {
	// Otherwise, split up the work between the cores!
	work := make([]pixelSegment, numCores*numCores)

	// Split it in to the number of cores we have squared
	// If we had just split it into the number of cores we have (say four)
	// Then we run the risk of all the objects being in one quarter of the
	// screen, and that section of work taking a long time while the other
	// cores idle.
	for i := 0; i < numCores; i++ {
		for j := 0; j < numCores; j++ {
			// Store the segment we're accessing.
			segment := &work[i*numCores+j]

			// Calculate the column range.
			segment.Min.X = width / numCores * (i + 0)
			segment.Max.X = width / numCores * (i + 1)

			// Calculate the row range.
			segment.Min.Y = height / numCores * (j + 0)
			segment.Max.Y = height / numCores * (j + 1)
		}
	}

	return work
}
