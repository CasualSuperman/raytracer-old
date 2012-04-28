package trace

import (
	"math"
	"os"
	"raytracer/color"
	"raytracer/debug"
	"raytracer/log"
	"raytracer/view"
	"runtime"
)

type pixelSegment struct {
	Min, Max struct {
		X, Y int
	}
}

var goroutines = runtime.NumCPU()

// Generates all the pixels in a range. Indicates completion on the given
// channel when done.
func makePixelSegment(done chan bool, m *view.Model, i color.Image,
	s *pixelSegment) {

	// Loop over the rows.
	for y := s.Min.Y; y < s.Max.Y; y++ {
		// Loop over the columns.
		for x := s.Min.X; x < s.Max.X; x++ {
			if debug.IMAGE {
				log.Printf("Calculating pixel (%d, %d)\n", x, y)
			}
			// Create the pixel
			makePixel(m, x, y, i)
		}
	}
	// Notify completion
	done <- true
	// Exit
}

// Creates an image from the given model, using its lights, shapes, and
// projection.
func MakeImage(m *view.Model) {
	// Make an image with the specified dimensions
	image := color.New(m.Projection.Pixel.Width,
		m.Projection.Pixel.Height)

	if debug.IMAGE {
		log.Println(*m)
	}

	if debug.ANY {
		// If any debugging is turned on, we want to do this in a single thread.
		// That prevents the log messages from showing up inside each other.
		// Also, loop backwards, so our pixels match up with Dr Kreahling's.
		for y := image.Height() - 1; y >= 0; y-- {
			for x := 0; x < image.Width(); x++ {
				makePixel(m, x, y, image)
			}
		}
	} else {
		// Otherwise, split up the work between the cores!
		work := make([]pixelSegment, goroutines*goroutines)

		// Split it in to the number of cores we have squared
		// If we had just split it into the number of cores we have (say four)
		// Then we run the risk of all the objects being in one quarter of the
		// screen, and that section of work taking a long time while the other
		// cores idle.
		for i := 0; i < goroutines; i++ {
			for j := 0; j < goroutines; j++ {
				// Store the segment we're accessing.
				segment := &work[i*goroutines+j]

				// Calculate the column range.
				segment.Min.X = image.Width() / goroutines * (i + 0)
				segment.Max.X = image.Width() / goroutines * (i + 1)

				// Calculate the row range.
				segment.Min.Y = image.Height() / goroutines * (j + 0)
				segment.Max.Y = image.Height() / goroutines * (j + 1)
			}
		}

		// Make a channel so the goroutines can communicate when they are finished.
		done := make(chan bool)

		// Kick off as many goroutines as we have cores.
		for i := 0; i < goroutines; i++ {
			go makePixelSegment(done, m, image, &work[i])
		}

		// Start a new goroutine every time we get a result back, keep the CPU busy
		for i := goroutines; i < goroutines*goroutines; i++ {
			<-done
			go makePixelSegment(done, m, image, &work[i])
		}

		// Wait for the last few to finish
		for i := 0; i < goroutines; i++ {
			<-done
		}
	}

	// Print out the image
	image.PPM(os.Stdout)
}

// Given a model and pixel column and row, and an image, calculates the
// position of the pixel in the world, raytracing using that location, and
// storing it in the given image.
func makePixel(m *view.Model, x, y int, i color.Image) {
	// Get the pixel's world location.
	base := mapPixToWorld(m, x, y)
	// Create storage space for the pixel
	p := color.Color{0, 0, 0}
	// We haven't travelled anywhere yet.
	dist := 0.0

	// Trace using the pixel
	rayTrace(m, base, &p, &dist, nil)

	// Cap the pixel's intensity
	p.R = math.Min(1, p.R)
	p.G = math.Min(1, p.G)
	p.B = math.Min(1, p.B)

	if debug.COLOR {
		log.Println("Color after intensity cap:", p)
	}

	pixelInImage := i.GetPixel(x, y)

	*pixelInImage.R = uint8(p.R * 255)
	*pixelInImage.G = uint8(p.G * 255)
	*pixelInImage.B = uint8(p.B * 255)
}
