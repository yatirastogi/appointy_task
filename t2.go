package main

import (
    "context"
    "html/template"
    "net/http"
    "path"
    "fmt"
    "log"
	//"reflect"
 
	"time"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
   // "strings"
)

type Fields struct {
	Id string
    Name  string
    Email string
    Password  string
}
type PostFields struct {
	Id string
	PostId string
    Caption  string
    ImageURL string
    TimePosted time.Time
}


func index(w http.ResponseWriter, r *http.Request){
	
    fp := path.Join("templates", "main.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func createpage(w http.ResponseWriter, r *http.Request) {
    
   
    fp := path.Join("templates", "create.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    
}

func loginpage(w http.ResponseWriter, r *http.Request) {
    
   
    fp := path.Join("templates", "login.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    
}
func postpicpage(w http.ResponseWriter, r *http.Request) {
    
   
    fp := path.Join("templates", "post.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
    

func homePage(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	clientOptions := options.Client().ApplyURI("mongodb+srv://rastogi_yati:iTuTOKUM3UKK6e98@cluster0.dokcw.mongodb.net/mydb?retryWrites=true&w=majority")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error while connecting")
	}
	 col := client.Database("instagram").Collection("newusers")
	 var result Fields 
	 
	 insta := client.Database("instagram")
	 newuser := insta.Collection("newusers")

	id := r.URL.Query().Get("id")
	if id ==""{//create user
		
	err = col.FindOne(context.TODO(), bson.M{"Id":r.FormValue("id")}).Decode(&result)
    if err != nil {//if username doesnot exixt
       //INSERT INTO DB
	   res, err := newuser.InsertOne(ctx, bson.D{
		{Key: "Name", Value: r.FormValue("username")},
		{Key: "Id", Value: r.FormValue("id")},
		{Key:"Email",Value:r.FormValue("email")},
		{Key:"Password",Value:r.FormValue("password")},
	})
	
			
		fmt.Println(res);
		fmt.Println(err);
		fmt.Fprintf(w, "Created Successfully")
    } else {
        // fmt.Println("FindOne() result:", result)
        // fmt.Println("FindOne() Name:", result.Name)
        // fmt.Println("FindOne() Dept:", result.Password)
		fmt.Fprintf(w, "Id Already Exists")
    }
}else{
	
	
	 err = col.FindOne(context.TODO(), bson.M{"Id":r.FormValue("id")}).Decode(&result)
	 if err == nil {
		//fmt.Println("FindOne() result:", result)
        
		fmt.Fprintf(w,"Username is "+result.Name)
        //fmt.Println("FindOne() Dept:", result.Password)
   
   } else {
	   fmt.Fprintf(w,"Id Not Found!")
   }


}

}


func getpostpage(w http.ResponseWriter, r *http.Request) {
    
   
    fp := path.Join("templates", "getpost.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    
}



func getallpostpage(w http.ResponseWriter, r *http.Request) {
    
   
    fp := path.Join("templates", "getallpost.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    
}
func postpic(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	clientOptions := options.Client().ApplyURI("mongodb+srv://rastogi_yati:iTuTOKUM3UKK6e98@cluster0.dokcw.mongodb.net/mydb?retryWrites=true&w=majority")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error while connecting")
	}
	// col := client.Database("instagram").Collection("posts")
	
	
	 insta := client.Database("instagram")
	 newuser := insta.Collection("posts")

	   res, err := newuser.InsertOne(ctx, bson.D{
		{Key: "Id", Value:r.FormValue("id")},
		{Key: "PostId", Value: r.FormValue("postid")},
		{Key:"Caption",Value:r.FormValue("caption")},
		{Key:"ImageURL",Value:r.FormValue("imgurl")},
		{Key:"TimePosted",Value:time.Now()},
	})
	
			
		fmt.Println(res);
		fmt.Println(err);
		fmt.Fprintf(w, "Created Successfully")




}



func homePage1(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the MY!")
    fmt.Println("Endpoint Hit: Yati")
}











func getpost(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	clientOptions := options.Client().ApplyURI("mongodb+srv://rastogi_yati:iTuTOKUM3UKK6e98@cluster0.dokcw.mongodb.net/mydb?retryWrites=true&w=majority")
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error while connecting")
	}
	 col := client.Database("instagram").Collection("posts")
	 var Posted PostFields 
	 
	// insta := client.Database("instagram")
	 //newuser := insta.Collection("posts")

	id := r.URL.Query().Get("id")
	if id !=""{//invalid id
		
	err = col.FindOne(context.TODO(), bson.M{"PostId":r.FormValue("id")}).Decode(&Posted)
    if err == nil {

 		
        fmt.Fprintf(w,"Imageurl %s \n", Posted.ImageURL)
        fmt.Fprintf(w,"Caption %s \n", Posted.Caption)
		fmt.Fprintf(w,"PostId %s \n", Posted.PostId)
		fmt.Fprintf(w,"Posted By %s \n ", Posted.Id)
		fmt.Println(err);
		
    } else{
		fmt.Fprintf(w, "No Posts Found")
	}
}
}




func getallpost(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	clientOptions := options.Client().ApplyURI("mongodb+srv://rastogi_yati:iTuTOKUM3UKK6e98@cluster0.dokcw.mongodb.net/mydb?retryWrites=true&w=majority")
//	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error while connecting")
	}
	// col := client.Database("instagram").Collection("posts")
	// var Posted PostFields 
	
	 insta := client.Database("instagram")
	posts := insta.Collection("posts")

	cursor, err := posts.Find(context.TODO(), bson.M{"Id":r.FormValue("id")})
if err != nil {
  log.Fatal(err)
}
var posteddata []PostFields
if err = cursor.All(context.TODO(), &posteddata); err != nil {
  log.Fatal(err)
}


for i, v := range posteddata {
	fmt.Fprintf(w,"%d. Post URL %+v\n",i, v.ImageURL)	
	fmt.Fprintf(w,"   Post Caption %+v\n", v.Caption)	
	fmt.Fprintf(w,"   Post Id %+v\n", v.PostId)	
	fmt.Fprintf(w,"   User Id %+v\n", v.Id)	
}



}
func handleRequests() {
    http.HandleFunc("/", index)
	http.HandleFunc("/createpage", createpage)
	http.HandleFunc("/new", homePage1)
	http.HandleFunc("/users", homePage)
	http.HandleFunc("/loginpage", loginpage)	
	//http.HandleFunc("/login", login)
	http.HandleFunc("/postpicpage", postpicpage)
	http.HandleFunc("/postpic", postpic)
	http.HandleFunc("/getpostpage", getpostpage)
	http.HandleFunc("/posts", getpost)
	http.HandleFunc("/posts/users", getallpost)
	http.HandleFunc("/getallpostpage", getallpostpage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

    handleRequests()
}