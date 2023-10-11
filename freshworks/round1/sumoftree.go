package main 

//func main(){}


type Node struct {
	data int
	left *Node
	right *Node
}


var result int 

func dep(root *Node ) int{
	if root ==nil{
		return 0
	}
	if root.left ==nil && root.right== nil{
		return root.data
	}
	sumofLeftTree:= dep(root.left)
	sumofRightTree:=dep(root.right)

	rootSum:= sumofLeftTree+sumofRightTree+root.data




	//tmax:= max(sumofLeftTree,sumofRightTree)
	//tmax=max(tmax,rootSum)

	result = max(result,rootSum)
	return rootSum

}

func findlargestSum(root *Node) int{
	if root ==nil{
		return 0
	}
	if root.left ==nil && root.right== nil{
		return root.data
	}

	dep(root)

	return result


}


func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}


/*Given a binary tree, task is to find subtree with maximum sum in tree. Return the max subtree sum.

 	      	  1
            /    \
          -2      3
          / \    /  \
         6   5  -9   2
*/