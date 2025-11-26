package question06

import "strings"

func convert1(s string, numRows int) string {
	// 特殊情况处理
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// 每一行用 strings.Builder 保存，效率更高
	rows := make([]strings.Builder, numRows)
	row := 0           // 当前行
	goingDown := false // 当前方向：true 表示向下，false 表示向上

	for _, ch := range s {
		rows[row].WriteRune(ch)

		// 在第 0 行（最上面）或者第 numRows-1 行（最下面）时切换方向
		if row == 0 || row == numRows-1 {
			goingDown = !goingDown
		}

		// 根据当前方向移动到下一行
		if goingDown {
			row++
		} else {
			row--
		}
	}

	// 把所有行拼接起来
	var result strings.Builder
	for _, sb := range rows {
		result.WriteString(sb.String())
	}

	return result.String()
}
