package main

import (
	"fmt"
)

type barang struct {
	NamaBrg  string
	HargaBrg int
}

type belanja struct {
	JumlahBrg int
}

func main() {
	brg1 := barang{
		NamaBrg:  "Sabun Cuci Piring",
		HargaBrg: 50000,
	}
	brg2 := barang{
		NamaBrg:  "Sabun Mandi",
		HargaBrg: 25000,
	}
	brg3 := barang{
		NamaBrg:  "Shampoo",
		HargaBrg: 35000,
	}
	brg4 := barang{
		NamaBrg:  "Pasta Gigi",
		HargaBrg: 10000,
	}
	brg5 := barang{
		NamaBrg:  "Tissue",
		HargaBrg: 7500,
	}

	blj1 := belanja{
		JumlahBrg: 7,
	}
	blj2 := belanja{
		JumlahBrg: 9,
	}
	blj3 := belanja{
		JumlahBrg: 10,
	}
	blj4 := belanja{
		JumlahBrg: 11,
	}
	blj5 := belanja{
		JumlahBrg: 20,
	}

	var total int
	total = (brg1.HargaBrg * blj1.JumlahBrg) + (brg2.HargaBrg * blj2.JumlahBrg) +
		(brg3.HargaBrg * blj3.JumlahBrg) + (brg4.HargaBrg * blj4.JumlahBrg) +
		(brg5.HargaBrg * blj5.JumlahBrg)

	fmt.Println("---List Barang Belanja---")
	fmt.Println(brg1.NamaBrg, "=", blj1.JumlahBrg, "buah")
	fmt.Println(brg2.NamaBrg, "=", blj2.JumlahBrg, "buah")
	fmt.Println(brg3.NamaBrg, "=", blj3.JumlahBrg, "buah")
	fmt.Println(brg4.NamaBrg, "=", blj4.JumlahBrg, "buah")
	fmt.Println(brg5.NamaBrg, "=", blj5.JumlahBrg, "buah")
	fmt.Println("Total Biaya yang harus dibayar:", total)
}