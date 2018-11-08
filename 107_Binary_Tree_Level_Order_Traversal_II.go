package problem0107

func levelOrderBottom(root *main.TreeNode) [][]int {

	res := [][]int{}

	var dfs func(*main.TreeNode, int)
	dfs = func(root *main.TreeNode, level int) {
		if root == nil {
			return
		}

		// 新增一層
		if level >= len(res) {
			res = append([][]int{[]int{}}, res...)
		}
		n := len(res)
		res[n-level-1] = append(res[n-level-1], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
