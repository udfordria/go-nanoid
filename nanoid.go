package nanoid

import "crypto/rand"

const NanoidDefaultSize = 21

var NanoidAlphabet = []rune("_-0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func SizedNanoid(size int) string {
	bytes := make([]byte, size)
	rand.Read(bytes) // !FIXME : error skipped
	id := make([]rune, size)
	for i := 0; i < size; i++ {
		id[i] = NanoidAlphabet[bytes[i]&63]
	}
	return string(id)
}

func Nanoid() string {
	return SizedNanoid(NanoidDefaultSize)
}
