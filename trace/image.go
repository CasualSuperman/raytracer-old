package trace

import (
	"os"
	"raytracer/color"
	"raytracer/debug"
	"raytracer/log"
	"raytracer/view"
	"runtime"
)

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
				makeAntiAliasedPixel(m, x, y, image)
			}
		}
	} else {
		numCpus := runtime.NumCPU()

		// Generate some pixel sections based on the number of cores we have.
		work := generateWorkSegments(image.Height(), image.Width(), numCpus)

		// Make a channel so the goroutines can communicate when they are finished.
		done := make(chan bool)

		// Kick off as many goroutines as we have cores.
		for i := 0; i < numCpus; i++ {
			go makePixelSegment(done, m, image, &work[i])
		}

		// Start a new goroutine every time we get a result back, keep the CPU busy
		for i := numCpus; i < numCpus*numCpus; i++ {
			<-done
			go makePixelSegment(done, m, image, &work[i])
		}

		// Wait for the last few to finish
		for i := 0; i < numCpus; i++ {
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
	rayTrace(m, base, &p, dist, nil)

	// Cap the pixel's intensity
	p.Cap(1)

	if debug.COLOR {
		log.Println("Color after intensity cap:", p)
	}

	pixelInImage := i.GetPixel(x, y)

	*pixelInImage.R = uint8(p.R * 255)
	*pixelInImage.G = uint8(p.G * 255)
	*pixelInImage.B = uint8(p.B * 255)
}
