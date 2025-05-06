package main
import (
  "fmt",
  "os",
  "strconv"
};
  
func readAFileContent() float64 {
  data, _ := os.ReadFile('file.txt');
  b := string(data);
  balance, _ := strconv.parseFloat(b, 64);
  return balance;
}
  
func writeAcontentFile(balance float64){
  b := fmt.Sprint(balance)
  os.WriteFile("file.txt", []byte(b), 0644)
}

func main() {
    var accountBalance = 1000;
    var deposit float64;
    var submit bool;
    fmt.Scan(&deposit);
    fmt.Scan(&submit)

    if submit {
      accountBalance += deposit;
      writeAcontentFile(accountBalance);
      balanceAmount := readAFileContent()
      fmt.Println(balanceAmount)
    }
}
