package other

type KnightTourProblem struct {
	x int
	y int
	visited []bool
	isFinished bool
}

type ChessboardPoint struct {
	x int
	y int
}

func NewKnightTourProblem(x int, y int) *KnightTourProblem {
	knightTourProblem := &KnightTourProblem{x: x,y: y}
	knightTourProblem.visited = make([]bool,x * y)
	return knightTourProblem
}

//使用递归回溯的方式遍历整个棋盘走法
func (knightTourProblem KnightTourProblem) TraverseChessboard(chessboard [][]int,row,column,step int) {
	chessboard[row][column] = step
	knightTourProblem.visited[knightTourProblem.y * row + column] = true
	points := knightTourProblem.getNext(ChessboardPoint{x: row,y: column})
	knightTourProblem.nextSort(points,0,len(points)-1)
	for len(points) > 0 {
		point := points[0]
		points = points[1:]
		if !knightTourProblem.visited[knightTourProblem.y * point.x + point.y]{
			knightTourProblem.TraverseChessboard(chessboard,point.x,point.y,step+1)
		}
	}
	if step < knightTourProblem.x * knightTourProblem.y && !knightTourProblem.isFinished{
		chessboard[row][column] = 0
		knightTourProblem.visited[knightTourProblem.y*row+column] = false
	}else{
		knightTourProblem.isFinished = true
	}
}

//获取当前点可以跳跃的下一步点集合
func (knightTourProblem KnightTourProblem) getNext(current ChessboardPoint) []ChessboardPoint {
	points := make([]ChessboardPoint,0)
	var point ChessboardPoint
	point.x = current.x-2
	point.y = current.y-1
	if point.x >=0 && point.y>=0{
		points = append(points,ChessboardPoint{x: point.x,y: point.y})
	}
	point.x = current.x-1
	point.y = current.y-2
	if point.x >=0 && point.y>=0{
		points = append(points,ChessboardPoint{x: point.x,y: point.y})
	}
	point.x = current.x+1
	point.y = current.y-2
	if point.x < knightTourProblem.x && point.y>=0{
		points = append(points,ChessboardPoint{x: point.x,y: point.y})
	}
	point.x = current.x+2
	point.y = current.y-1
	if point.x < knightTourProblem.x && point.y>=0{
		points = append(points,ChessboardPoint{x: point.x,y: point.y})
	}
	point.x = current.x-2
	point.y = current.y+1
	if point.x >=0 && point.y < knightTourProblem.y{
		points = append(points,ChessboardPoint{x: point.x,y: point.y})
	}
	point.x = current.x-1
	point.y = current.y+2
	if point.x >=0 && point.y < knightTourProblem.y{
		points = append(points,ChessboardPoint{x: point.x,y: point.y})
	}
	point.x = current.x+1
	point.y = current.y+2
	if point.x < knightTourProblem.x && point.y < knightTourProblem.y{
		points = append(points,ChessboardPoint{x: point.x,y: point.y})
	}
	point.x = current.x+2
	point.y = current.y+1
	if point.x < knightTourProblem.x && point.y < knightTourProblem.y{
		points = append(points,ChessboardPoint{x: point.x,y: point.y})
	}
	return points
}

// 对下一步可以跳跃的点进行排序
func (knightTourProblem KnightTourProblem) nextSort(points []ChessboardPoint,left,right int) {
	l := left
	r := right
	middle := left + (right-left) / 2
	base := len(knightTourProblem.getNext(points[middle]))
	for l < r {
		for len(knightTourProblem.getNext(points[l])) < base {
			l++
		}
		for len(knightTourProblem.getNext(points[r])) > base {
			r--
		}
		if l >= r{
			break
		}
		points[l],points[r] = points[r],points[l]
		if len(knightTourProblem.getNext(points[l])) == base{
			l++
		}
		if len(knightTourProblem.getNext(points[r])) == base{
			r--
		}
	}
	if l==r{
		l++
		r--
	}
	if r > left{
		knightTourProblem.nextSort(points,left,r)
	}
	if l < right{
		knightTourProblem.nextSort(points,right,l)
	}
}