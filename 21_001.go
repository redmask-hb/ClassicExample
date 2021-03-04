/*
输入某二叉树的前序遍历和中序遍历的结果，请重建出该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建二叉树并返回。
 */


func reConstructBinaryTree( pre []int ,  vin []int ) *TreeNode {
  index:=0
	cp:=&index
	return buildTree(pre, vin, cp)
}

func buildTree(pre []int, vin []int, index *int) *TreeNode {
    if len(vin)==0{
        return nil
    }
	value := pre[*index]
	node := &TreeNode{Val: value}    
    if *index<len(pre)-1 {
		*index++
	}
	if len(vin)>1 {
		tag := 0
		for k, v := range vin {
			if v == value {
				tag = k
			}
		}		
		leftVin := vin[:tag]
		node.Left = buildTree(pre, leftVin, index)
		if tag < len(vin) {
				rightVin := vin[tag+1:]
				node.Right = buildTree(pre, rightVin, index)
		}
	}
	return node
}

