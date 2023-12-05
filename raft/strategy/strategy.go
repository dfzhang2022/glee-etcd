// package strategy

// import (
// 	"fmt"
// 	"sort"
// )

// // 定义接口
// type MyInterface interface {
// 	CalFlag(Ti_DM []int, T1_DM []int, T2_DM []int, LN int) bool
// 	CalBetterNodes(kMatrix, kMinus1Matrix [][]int) []int
// 	CalBestNode(CommitLatency []int, KMatrix [][]int, betterN []int, UseRate []float64, LN int) int
// }

// // 实现接口的结构体
// type MyStruct struct {
// }

// // 实现接口的方法
// func (ms MyStruct) CalFlag(Ti_DM []int, T1_DM []int, T2_DM []int, LN int) bool {
// 	n := len(Ti_DM)
// 	diff_1 := 0
// 	diff_2 := 0
// 	LNSum1 := 0
// 	LNSum2 := 0

// 	if n < 2 {
// 		//小于两个节点则不用考虑主动选主
// 		return false
// 	}

// 	for col := 0; col < n; col++ {
// 		if col != LN {
// 			diff_1 += Ti_DM[col] - T1_DM[col]
// 			LNSum1 += T1_DM[col]
// 		}
// 	}

// 	if LNSum1 == 0 {
// 		//0处理，说明无有效时延信息，无法主动选主
// 		return false
// 	}

// 	R_1 := float64(diff_1) / float64(LNSum1)
// 	fmt.Println(R_1)
// 	if R_1 < 0.3 {
// 		return false
// 	}
// 	for col := 0; col < n; col++ {
// 		if col != LN {
// 			diff_2 += T1_DM[col] - T2_DM[col]
// 			LNSum2 += T2_DM[col]
// 		}
// 	}

// 	if LNSum2 == 0 {
// 		//0处理，说明无有效时延信息，无法主动选主
// 		return false
// 	}

// 	R_2 := float64(diff_2) / float64(LNSum2)
// 	if R_2 < 0.3 {
// 		return false
// 	}
// 	return true
// }

// func (ms MyStruct) CalBetterNodes(kMatrix, kMinus1Matrix [][]int, LN int) []int {
// 	// 获取节点数量，且该数量一定是大于等于2的
// 	numNodes := len(kMatrix)

// 	// 初始化Delaylist，存放每个节点的平均时延
// 	Delaylist := make([]int, numNodes)

// 	// 计算k时刻的所有节点到其他节点的综合平均时延
// 	for i := 0; i < numNodes; i++ {
// 		totalDelay := 0
// 		for j := 0; j < numNodes; j++ {
// 			if i != j { // 排除节点到自身的时延
// 				totalDelay += kMatrix[i][j]
// 			}
// 		}
// 		// 平均时延不包括节点到自身
// 		Delaylist[i] = totalDelay / (numNodes - 1)
// 	}

// 	// 计算k-1时刻的所有节点到其他节点的综合平均时延
// 	for i := 0; i < numNodes; i++ {
// 		totalDelay := 0
// 		for j := 0; j < numNodes; j++ {
// 			if i != j { // 排除节点到自身的时延
// 				totalDelay += kMinus1Matrix[i][j]
// 			}
// 		}
// 		// 平均时延不包括节点到自身，且与k时刻的计算值进行加权平滑处理
// 		Delaylist[i] = 8*Delaylist[i]/10 + 2*(totalDelay/(numNodes-1))/10
// 	}

// 	// 初始化BetterN，存放加权平均时延比LN小的节点编号
// 	BetterN := []int{}

// 	// 比较加权平均时延，将比LN时延小的节点编号添加到BetterN中
// 	for i := 0; i < numNodes; i++ {
// 		if i != LN && Delaylist[i] < Delaylist[LN] {
// 			BetterN = append(BetterN, i)
// 		}
// 	}

// 	return BetterN
// }

// func (ms MyStruct) CalBestNode(CommitLatency int, KMatrix [][]int, betterN []int, UseRate []float64, LN int) int {
// 	//参数设置
// 	Alpha := 1.0
// 	n := len(KMatrix)
// 	m := len(betterN)
// 	bestN := LN
// 	bestProfit := 0.0

// 	//遍历betterN中所有节点，计算收益,若betterN为空，则bestN依旧是原LN
// 	for k := 0; k < m; k++ {
// 		//计算node成为newleader时的收益
// 		node := betterN[k]

// 		//求中位数mideumRTT
// 		var RTT []int
// 		RTT = append(RTT, KMatrix[node]...)
// 		sort.Ints(RTT)
// 		var mideumRTT int
// 		if n%2 == 1 {
// 			mideumRTT = RTT[n/2]
// 		} else {
// 			mideumRTT = (RTT[n/2] + RTT[n/2+1]) / 2
// 		}

// 		//┏ - ┏_new
// 		CommitLatencydiff := CommitLatency - mideumRTT

// 		// ∑ Hi*(RTT_leader - RTT_newleader)
// 		UseRatediff := 0.0
// 		for i := 0; i < n; i++ {
// 			RRTdiff := float64(KMatrix[LN][i] - KMatrix[node][i])
// 			UseRatediff += UseRate[i] * RRTdiff
// 		}

// 		//U = α*(┏ - ┏_new) + ∑ Hi*(RTT_leader - RTT_newleader)
// 		Profit := UseRatediff + Alpha*float64(CommitLatencydiff)

// 		if Profit > bestProfit {
// 			bestProfit = Profit
// 			bestN = node
// 		}
// 	}
// 	return bestN
// }

// func main() {
// 	// 创建 MyStruct 的实例
// 	myStructInstance := MyStruct{}

// 	// 将实例赋值给接口变量
// 	var myInterfaceVar MyInterface
// 	myInterfaceVar = myStructInstance

// 	// 使用接口方法
// 	IEflag := myInterfaceVar.CalFlag(Ti_DM, T1_DM, T2_DM, LN)
// 	betterN := myInterfaceVar.CalBetterNodes(kMatrix, kMinus1Matrix, LN)
// 	bestN := myInterfaceVar.CalBestNode(CommitLatency, KMatrix, betterN, UseRate, LN)
// }
