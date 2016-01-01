package scrubber

func Scrub(in string) (scrubbed string) {
	var count int = -1
	for _, val := range in {
		if val != ' ' && val != '	' {
			scrubbed += string(val)
			count = 0
		} else if count == 0 {
			scrubbed += " "
			count++
		}
	}
	return
}

// Label - (Opcode - Comment) || (Opcode) || (Comment)
// Opcode - Comment
// Comment
//func FineScrub(in string) (fineScrubbed string) {

//}
