package bfs

import "container/list"

// Definition for a binary tree node.
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
    var result []float64
    if root == nil {
        return result
    }
    
    // Initialize queue for BFS
    queue := list.New()
    queue.PushBack(root)
    
    for queue.Len() > 0 {
        levelSum := 0
        levelCount := queue.Len()
        
        // Process all nodes at the current level
        for i := 0; i < levelCount; i++ {
            node := queue.Remove(queue.Front()).(*TreeNode)
            levelSum += node.Val
            
            // Add the children of the current node to the queue for the next level
            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
        }
        
        // Calculate the average of the current level
        average := float64(levelSum) / float64(levelCount)
        result = append(result, average)
    }
    
    return result
}
