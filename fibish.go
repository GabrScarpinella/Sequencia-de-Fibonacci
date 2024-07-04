package main

import "fmt"

func main(){
	var a, b, c, d int;
	a=1;
	b=1;
	d=0;
	fmt.Printf("%d, %d, ", a, b);
	FIB:
		c=a;
		a=b;
		b+=c;
		d++;
		fmt.Printf("%d, ", b);
	if d<10{goto FIB;}
}