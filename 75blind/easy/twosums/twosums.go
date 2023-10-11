
package main 

//two sums with O(n2)
func twoSum(nums []int, target int) []int {
    for i:=0 ; i <len(nums) ; i++{
        for j:=i+1;j<len(nums); j++{
            if nums[i]+nums[j] == target{
                return []int{i,j}
            }
        }
    }
    return nil
}

//two sum hashmap
func twoSumMap(nums []int, target int)[]int{
	numMp:= make(map[int]int)
	for i,num:= range nums{
		val,ok:=numMp[target-num]
		if ok{
			return []int{i,val}
		}
		numMp[num]=i
	}
	return nil
}
