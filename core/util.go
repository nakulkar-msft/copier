package copier

import (
    "errors"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
)
const (
    MiB = 1024 * 1024
    defaultBlockBlobBlockSize = 8 * MiB
    blockSizeThreshold = 256 * MiB
    MaxNumberOfBlocksPerBlob = blockblob.MaxBlocks
    maxBlobSize = blockblob.MaxBlocks * blockblob.MaxStageBlockBytes
)

var ErrFileTooLarge = errors.New("file too large")

func getBlockSize(sourceSize int64) (int64, error) {
    if (sourceSize > maxBlobSize) {
        return int64(0), ErrFileTooLarge
    }

	for blockSize := defaultBlockBlobBlockSize; blockSize <= blockSizeThreshold; blockSize *= 2 {
        if sourceSize <= int64(MaxNumberOfBlocksPerBlob * blockSize) {
            return int64(blockSize), nil
        }
    }

    return ((sourceSize - 1) / MaxNumberOfBlocksPerBlob) + 1, nil
}