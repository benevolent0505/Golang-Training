package main

import "fmt"

func main(){
	//int型の変数を2つ宣言
	a,b := 1,1
	
	//関数に変数を渡す
	//aは値をそのまま渡す。値渡し。
	//bはアドレス演算子を使ってポインタとして渡す。ポインタ渡し
	double(a,&b)
	
	//output
	fmt.Println("値渡し:",a)
	fmt.Println("参照渡し:",b)
}

//この関数のパラメータxは値のコピーを受け取り
//yはポインタ(正確にはポインタのコピー)を受け取る
func double(x int,y *int){
	//変数の値を変更
	x = x * 2
	//間接参照演算子を使用し、ポインタyが指し示す変数の値を変更
	*y = *y * 2
}
