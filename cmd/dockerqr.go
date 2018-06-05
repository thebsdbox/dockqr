package cmd

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	qre "github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
	qri "github.com/tuotoo/qrcode"
)

var dockerfile, qrfile string

func init() {
	dockerqr.AddCommand(qrbuild)
	dockerqr.AddCommand(qrimport)
	qrbuild.Flags().StringVarP(&dockerfile, "dockerfile", "d", os.Getenv("QR_DOCKERFILE"), "Path to a dockerfile")
	qrbuild.Flags().StringVarP(&qrfile, "qrfile", "q", os.Getenv("QR_QRFILE"), "Path to a QR image")

}

var dockerqr = &cobra.Command{
	Use:   "dockerqr",
	Short: "Builds and exports dockerfiles to/from QR files",
}

// Execute -
func Execute() {
	if err := dockerqr.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// qrbuild - Creates a QR code from a dockerfile //TODO
var qrbuild = &cobra.Command{
	Use:   "qrbuild",
	Short: "Creates a QR image from a Dockerfile ",
	Run: func(cmd *cobra.Command, args []string) {
		if dockerfile == "" {
			log.Debugf("No Dockerfile specified, looking in current directory")
		}
		if qrfile == "" {
			log.Debugf("No output QR file specified, using default")
			qrfile = "dockerfile.png"
		}

		err := qre.WriteFile("test", qre.Medium, 256, qrfile)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	},
}

// qrimport - Imports an image and processes it
var qrimport = &cobra.Command{
	Use:   "qrimport",
	Short: "Creates a QR image from a Dockerfile ",
	Run: func(cmd *cobra.Command, args []string) {
		if dockerfile == "" {
			log.Debugf("No Dockerfile specified, looking in current directory")
		}
		if qrfile == "" {
			log.Debugf("No output QR file specified, using default")
			qrfile = "dockerfile.png"
		}

		fi, err := os.Open("dockerfile.png")
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		defer fi.Close()
		qrmatrix, err := qri.Decode(fi)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		fmt.Printf("%s\n", qrmatrix.Content)
	},
}
