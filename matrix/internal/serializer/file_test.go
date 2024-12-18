package serializer

import (
	"os"
	"testing"

	"matrix/internal/fns"

	"github.com/stretchr/testify/require"
)

func TestProtoToBinaryFile(t *testing.T) {
	filename := "../tmp/test_output_file.bin"

	laptop := fns.NewLaptop()

	err := ProtoToBinaryFile(laptop, filename)
	require.NoError(t, err, "ProtoToBinaryFile should not return an error")

	_, err = os.Stat(filename)
	require.NoError(t, err, "File should be created")

	defer os.Remove(filename)
}
