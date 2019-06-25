package main

import "fmt"

func main() {

	fmt.Println("hello world")

	var h string = "ozgur"
	fmt.Printf("h is type %T\n", h)

	fmt.Println(add(3, 5))
	fmt.Println(subt(5, 3))

	//birden fazla değer return edilebilir
	x, y := mult("hello", "world")
	fmt.Println(x, y)

	//for
	z := 1
	for z <= 5 {
		fmt.Println(z)
		z++
	}
	for z := 1; z <= 5; z++ {
		fmt.Println(z)
	}

	ifStatement()

	//Pointer example 1
	pointer()
	//Pointer example 2
	pointer2()
	//Pointer example 3
	pointer3()
}
func pointer3(){
/*new ile string tipinden nickP isimli bir Pointer tanımlanmıştır. Sadece bellekte bunun için yer ayrılmıştır. Bu nedenle atama öncesi güncel değeri yoktur. "Speedy Gonzales" atamasından sonra ise adres bölgesi veri ile doldurulmuştur*/
    nickP:=new(string)
    fmt.Println("Address ",nickP)
    fmt.Println("Current value",*nickP)
    *nickP="speedy gonzales"
    fmt.Printf("After assignment. '%s'\n",*nickP)
}

func pointer2() {
/*Olaylar main fonksiyonunda tanımlı age isimli değişken üzerinde gelişiyor. call parametre olarak değer bazlı değişken kullanmakta. Gelen age call fonksiyonunda değiştirilse bile main'deki age değerini etkilemiyor. callWithPointer içinse durum böyle değil. callWithPointer içerisine main fonksiyonundaki age değişkeninin adresi taşınıyor. Nitekim parametre bir pointer. Bu nedenle fonksiyon içerisinde yapılan değişim main içerisindeki age değişkeninde yapılmış oluyor. Konunun iyi anlaşılması için main'deki age ve fonksiyonlardaki parametre adreslerine dikkat etmenizi öneririm. Kodun çalışma zamanı çıktısı aşağıdaki gibidir.*/
    age:=41
    fmt.Println("main-age = ",age)
    fmt.Println("main-address of age",&age)
	call(age)
	fmt.Println("main-age = ",age)
	callWithPointer(&age)
    fmt.Println("main-age = ",age)

}
func callWithPointer(value *int){
    *value+=1
    fmt.Println("callWithPointer-value = ",*value)
    fmt.Println("callWithPointer-address of value",value)
}

func call(value int){
    value+=1
    fmt.Println("call-value = ",value)
    fmt.Println("call-address of value",&value)
}

func pointer() {
	/*pAge bir pointer oldu. Eşitliğin sağ tarafındaki & operatörü izleyen age değişkeninin bellek adresini yakalıyor. Bu nedenle pAge içeriğini ekrana bastığımızda bellek adresini görebiliyoruz. pAge isimli Pointer'ın işerat ettiği adres bölgesindeki değeri almak için * operatörü devreye giriyor. Sonrasında dikkat çekici bir durum söz konusu. *pAge değerinde bir değişiklik yapıp yaşa 41 sayısını atıyoruz. pAge bir pointer olduğundan üzerinden yapılan değişiklik aslında asıl işaret ettiği değişkenin değerinde gerçekleşmekte. Bu nedenle age değişkeni de artık 41 değerine sahip oluyor. point yine bir tamsayı değişkeni ve bellek adresini pPoint üzerinde taşıyoruz. * operatörü pointer'ın işaret ettiği değişken değerini verdiğinden matematik işlemlerine de dahip edebiliyoruz. *pPoint'in 2 ile çarpımında bu özellik ele alınıyor. Çok doğal olarak çarpma işlemi point değişkenini etkilemekte.*/
	age := 22
	pAge := &age
	fmt.Println("Address of age", pAge)
	fmt.Println("Age(from pointer)", *pAge)
	*pAge = 41
	fmt.Println("Age=", age)
	fmt.Println("Age(from pointer)", *pAge)

	point := 1000
	pPoint := &point
	*pPoint = *pPoint * 2
	fmt.Println("point=", point)
}

func ifStatement() {
	for y := 1; y <= 5; y++ {
		if y%2 == 0 {
			fmt.Println(y, "even number")
		} else {
			fmt.Println(y, "odd number")
		}
	}
}

func add(x int, y int) int {
	return x + y
}

func subt(a, b int) int {
	return a - b

}

func mult(a, b string) (string, string) {
	return a, b
}
