package medium

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	column := len(s) / (numRows + numRows - 2) * (numRows - 1)
	c := len(s) % (numRows + numRows - 2)
	if c <= numRows {
		column++
	} else {
		column += 1 + c - numRows
	}

	record := make([][]byte, numRows)
	for i := 0; i < numRows; i++ { // 初始化
		arr := make([]byte, column)
		record[i] = arr
	}

	i, j := 0, 0
	for _, v := range []byte(s) { // 赋值
		record[i][j] = v
		if j%(numRows-1) == 0 && i < numRows-1 {
			i++
			continue
		}
		if j%(numRows-1) == 0 {
			j++
			i--
			continue
		}
		if i >= 1 {
			j++
			i--
			continue
		}
	}

	// 遍历
	res := []byte("")
	var flag byte
	for i := range record {
		for j := range record[i] {
			if record[i][j] != flag {
				res = append(res, record[i][j])
			}
		}
	}

	return string(res)
}
