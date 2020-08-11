package main

func spiralOrder(matrix [][]int) []int {
	var result []int

	if len(matrix) == 0 {
		return result
	}

	// initial indicies
	rs := 0                  // row start
	re := len(matrix) - 1    // row end
	cs := 0                  // column start
	ce := len(matrix[0]) - 1 // column end

	for rs <= re && cs <= ce {
		// Move right
		for j := cs; j <= ce; j++ {
			result = append(result, matrix[rs][j])
		}

		// Jump down one row

		rs++
		fmt.Println(result)
		if rs > re {
			break
		}

		// Move down
		for i := rs; i <= re; i++ {
			result = append(result, matrix[i][ce])
		}

		// Move to the left one column, from the end
		ce--
		fmt.Println(result)
		if cs > ce {
			break
		}

		// Move left
		for j := ce; j >= cs; j-- {
			result = append(result, matrix[re][j])
		}
		fmt.Println(result)

		// Move up one row
		re--

		if rs > re {
			break
		}

		// Move up
		for i := re; i >= rs; i-- {
			result = append(result, matrix[i][cs])
		}
		fmt.Println(result)
		cs++
	}

	return result
}
