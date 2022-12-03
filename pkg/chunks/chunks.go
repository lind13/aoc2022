package chunks

func Create[T comparable](arr []T, chunkSize int) [][]T {
	if len(arr) == 0 {
		return nil
	}

	chunks := make([][]T, 0, len(arr)/chunkSize+1)

	for chunkSize < len(arr) {
		arr, chunks = arr[chunkSize:], append(chunks, arr[0:chunkSize:chunkSize])
	}

	chunks = append(chunks, arr)

	return chunks
}
