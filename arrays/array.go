package array

// fill an empty array using a loop
func fill_array(in [3]string) [3]string {
	result := [3]string{"", "", ""}
	for i := 0; i < len(in); i++ {
		result[i] = in[i]
	}
	return result
}
