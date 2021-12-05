package main

import ("fmt"
        "net/http"
        "html/template")

type User struct{
  Name string
  Age uint16
  Money uint16
  Avg_grades, Sadness float64
  Hobbies []string
}

func (u User) getAllInfo() string {
  return fmt.Sprintf("User NAME IS: %s. He is %d and he has" +
     "money: %d", u.Name, u.Age, u.Money)
}

func (z User) getAllInfo1() string{

  return fmt.Sprintf("%s %d %d", z.Name, z.Age, z.Money)
}

func (u *User) setNewName(newName string){
  u.Name = newName
}

func home_page(page http.ResponseWriter, r *http.Request) {
  bob := User{"bob", 25, 40, 4.2, 0.3, []string{"Football", "Scate"}}
  //fmt.Fprintf(page, "<b>Main Text</b>")
//  bob.setNewName("alex")
//  fmt.Fprintf(page, bob.getAllInfo())
tmpl, _ := template.ParseFiles("templates/home_page.html")
tmpl.Execute(page, bob)
}

func zalaev_page(w http.ResponseWriter, r *http.Request){
  gog := User{"gog", 223, 22, 5.5, 5.3, []string{"ss", "sss"}}
  //fmt.Fprintf(w, gog.getAllInfo1())
  tmpl, _ := template.ParseFiles("templates/zalaev.html")
  tmpl.Execute(w, gog)
}
func contacts_page(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Contacts page")
}

func handleRequest(){
  http.HandleFunc("/", home_page)
  http.HandleFunc("/contacts/", contacts_page)
  http.HandleFunc("/contacts/zalaev",zalaev_page)
  http.ListenAndServe(":8080", nil)
}

func main(){
  //var bob User =
//  bob := User{name: "bob", age:, money: -59, avg_grades: 1.1, sadness: 0.9}
  handleRequest()
}
for fmt.Scan(&a); i<10; fmt.Scan(&a){
    workArray[i]=a
    i++
}
    for j=0;j<6;j++{
    for fmt.Scan(&b);j<6; fmt.Scan(&a){
 arr[j] = a
    j++
      }
}
for j=0;i<6;i++{
    fmt.Printf("%d ",arr[j])
}
    /*
        //for j=0;j<6;j++{
      //  fmt.Println(arr[j])
 //   }
*/
    for j=0;j<6;j++{
        if j==0{
        s= workArray[arr[j]]
        t= workArray[arr[j+1]]
        workArray[arr[i]] = t
        workArray[arr[i+1]] = s
    }

    if j==2{
        s= workArray[arr[j]]
        t= workArray[arr[j+1]]
        workArray[arr[j]] = t
        workArray[arr[j+1]] = s
    }
    if j==4{
        s= workArray[arr[j]]
        t= workArray[arr[j+1]]
        workArray[arr[j]] = t
        workArray[arr[j+1]] = s
    }
}
