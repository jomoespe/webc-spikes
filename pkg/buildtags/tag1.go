// +build tag1 !tag2

package buildtags

func Salutation() (name string) {
	name = "from tag1"
	return
}
