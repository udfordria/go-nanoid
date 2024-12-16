package __tests__

import (
	"log"
	"testing"

	"github.com/udfordria/nanoid"
)

func TestNanoid(t *testing.T) {
	id := nanoid.Nanoid()

	log.Println(id)
}
