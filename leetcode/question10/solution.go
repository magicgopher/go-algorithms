package question10

// 递归 + 记忆化
func isMatch1(s string, p string) bool {
	m, n := len(s), len(p)
	memo := make([][]int, m+1) // -1未计算 0false 1true
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if j == n {
			return i == m
		}
		if memo[i][j] != -1 {
			return memo[i][j] == 1
		}

		res := false
		if j+1 < n && p[j+1] == '*' { // 遇到 x*
			// 跳过这个 x* （零次）
			res = dfs(i, j+2)
			// 或者用一次 x* 匹配（前提是能匹配上当前字符）
			if i < m && (p[j] == '.' || p[j] == s[i]) {
				res = res || dfs(i+1, j) // 注意这里还是 j，不是 j+2
			}
		} else {
			// 普通字符或 .
			if i < m && (p[j] == '.' || p[j] == s[i]) {
				res = dfs(i+1, j+1)
			}
		}

		memo[i][j] = 0
		if res {
			memo[i][j] = 1
		}
		return res
	}

	return dfs(0, 0)
}

// dp
func isMatch2(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	// 处理空字符串匹配 p 的情况（只能靠 x* 匹配 0 次）
	for j := 2; j <= n; j += 2 {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// 情况1：* 匹配 0 次 → 看 dp[i][j-2]
				dp[i][j] = dp[i][j-2]

				// 情况2：* 匹配 1次或多次 → 前提是 p[j-2] 能匹配 s[i-1]，且 dp[i-1][j] 为 true
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j] // 关键！这里是 dp[i-1][j]，不是 dp[i-1][j-2]
				}
			} else {
				// 普通字符或 .
				if p[j-1] == '.' || p[j-1] == s[i-1] {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}
	return dp[m][n]
}

// 动态规划 + 滚动数组
func isMatch3(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i <= m; i++ {
		prev := dp[0]    // 保存 dp[i-1][0]
		dp[0] = (i == 0) // 只有空字符串能匹配空模式

		for j := 1; j <= n; j++ {
			temp := dp[j] // 保存上一行的 dp[i-1][j]

			if p[j-1] == '*' {
				dp[j] = dp[j-2] // 零次
				if i > 0 && (p[j-2] == '.' || p[j-2] == s[i-1]) {
					dp[j] = dp[j] || temp // temp 是 dp[i-1][j]，表示“多次”
				}
			} else {
				if i > 0 && (p[j-1] == '.' || p[j-1] == s[i-1]) {
					dp[j] = prev // prev 是 dp[i-1][j-1]
				} else {
					dp[j] = false
				}
			}
			prev = temp // 更新 prev 为 dp[i-1][j]，供下一个 j 使用
		}
	}
	return dp[n]
}
