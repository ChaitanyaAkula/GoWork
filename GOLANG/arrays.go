package main
import(
	"fmt"
)
func main(){
	var matrix1[100][100] int 
	var matrix2[100][100] int
	var sum[100][100] int
	var row,col int
	fmt.Print("enter number of rows and columns :")	
	fmt.Scanln(&row,&col)
	fmt.Println("")
	fmt.Printf("MATRIX1 /n")
	fmt.Println("enter the elements for Matrix1 is =")
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			
			fmt.Scanln(&matrix1[i][j])
		
		}
	}
	fmt.Println("")
	fmt.Printf("MATRIX2 /n")
			fmt.Println("enter the elements for Matrix2 is =  ")

	
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			fmt.Scanln(&matrix2[i][j])
			
			}

	}
	if row==col{
		for i:=0;i<row;i++{
			for j:=0;j<col;j++{
			
				sum[i][j]= matrix1[i][j]+matrix2[i][j]
			}
			
		}
	}else {
		fmt.Println("2 matrices rows and columns are not equal")
	}
	
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			//fmt.Println("")
			fmt.Printf("%d",sum[i][j])
			
			if(j==col-1){
				fmt.Println("")
			}

		}
	}
//	sum1:=add(matrix1,matrix2,row,col)
	//fmt.Println(sum1)
}
/*func add(matrix1 [][]int,matrix2 [][]int,row int,col int)[][]int{
	var sum[20][20] int

	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			sum[i][j] = matrix1[i][j]+matrix2[i][j]
		}
	}
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			fmt.Println("")
			fmt.Printf("%d",sum[i][j])
		}
	}
	return sum
}*/