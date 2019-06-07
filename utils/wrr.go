package utils

//加权轮询算法 weighted round robin
// 客服权重默认都为0
// 返回当前匹配到的客服信息



type Node struct {
	Uid string
	Current int
	Weight int
}


func Wrr(customers []*Node)  (best *Node){
	if len(customers) == 0{
		return
	}
	weight_sum := 0
	for _,node := range customers{
		if node == nil{
			return
		}
		weight_sum+=node.Weight
		node.Current +=node.Weight
		if best == nil || node.Current >best.Current{
			best = node
		}
	}
	if best == nil{
		return
	}
	best.Current -= weight_sum

	return
}

