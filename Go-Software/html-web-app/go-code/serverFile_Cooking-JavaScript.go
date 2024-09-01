# Cooking JavaScript  ~

package main



import (
    
    
)


// ,  ° . +
func addEventToList(ctx context.Context, client *firestore.Client, addEvent) error {
  _EventListData := addEvent

  _, err := client.Collection("EventList").Doc("EventList").Set(ctx, _EventListData)
        
  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
} // , - • ~

// ,  ° . +
func addInventoryToList(ctx context.Context, client *firestore.Client, addInventory) error {
  _InventoryListData := addInventory

  _, err := client.Collection("_InventoryList").Doc("_InventoryList").Set(ctx, _InventoryListData)
        
  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
} // , - • ~

// ,  ° . +
func addItineraryToList(ctx context.Context, client *firestore.Client, addItinerary) error {
  _ItineraryListData := addItinerary

  _, err := client.Collection("_ItineraryList").Doc("_ItineraryList").Set(ctx, _ItineraryListData)
        
  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
} // , - • ~

// ,  ° . +
func addQuestToList(ctx context.Context, client *firestore.Client, addQuestList) error {
  _QuestListData := addQuestList

  _, err := client.Collection("_QuestList").Doc("_QuestList").Set(ctx, _QuestListData)
        
  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
} // , - • ~

// ,  ° . +
func addTravelLog(ctx context.Context, client *firestore.Client, addTravelLog) error {
  _TravelLogData := addTravelLog

  _, err := client.Collection("_TravelLog").Doc("_TravelLog").Set(ctx, _TravelLogData)
        
  if err != nil {
    // Handle any errors
    log.Printf("An error has occurred: %s", err)
  }
  return err
} // , - • ~





func app_welcome_center_page() {
    
    
}




// . appHandler, ~ for Data Pages °
func appHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/app" {
    	http.NotFound(w, r)
    	return
    }
    
// ,
  pageTitle := "~ Cooking JavaScript  - // - Website App"
  pagePath := r.URL.Path
  
  pageType := ".."
  
// ,  ° . +
  pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  if pagePath == "/settings" {
      pageTitle = "Settings Page"
      pageList = pageList
  }
  
  // ,  ° . +
  if pagePath == "/inventory" {
      pageTitle = "Inventory Page"
      pageList = pageList
  }
  // ,  ° . +
  if pagePath == "/itinerary" {
      pageTitle = "Itinerary Page"
      pageList = pageList
  }
  // ,  ° . +
  if pagePath == "/quest_list" {
      pageTitle = "Quest List Page"
      pageList = pageList
  }
  // ,  ° . +
  if pagePath == "/travel_log" {
      pageTitle = "Travel Log Page"
      pageList = pageList
  }
  
  
// ,  ° . + 
  pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  appHandler




// . pageHandler, ~ for User Pages °
func pageHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +

  pageTitle := "~ Cooking JavaScript  - // - Website App"
  pagePath := r.URL.Path
  pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  // ,  ° . +
  if pagePath == "/user" {
      pageTitle = "User Page"
      pageList = pageList
  }
  
  // ,  ° . +
  if pagePath == "/user" {
      pageTitle = "User Page"
      pageList = pageList
  }
  
  if pagePath == "/account" {
      pageTitle = "Account Page"
      pageList = pageList
  }
  
  if pagePath == "/profile" {
      pageTitle = "Profile Page"
      pageList = pageList
  }
  
// ,  ° . +
  if pagePath == "/portfolio" {
      pageTitle = "Portfolio Page"
      pageList = pageList
  }
  
  if pagePath == "/resume" {
      pageTitle = "Resume Page"
      pageList = pageList
  }


// ,  ° . +
pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  pageHandler



// . indexHandler,  ~ for Public Pages °
func indexHandler(w http.ResponseWriter, r *http.Request) {
// ,  ° . +
    if r.URL.Path != "/" {
    	http.NotFound(w, r)
    	return
    }

// , ° . +
  pageTitle := "~ Cooking JavaScript  - // - Website App"
  pagePath := r.URL.Path
  pageType := ".."


// ,  ° . +
pageData := htmlPageData {
      pageTitle: pageTitle,
      pagePath: pagePath,
      
      pageList: []pageNav {
          { pageTitle: "one", pageLink: "one"},
          { pageTitle: "two", pageLink: "two"},
          { pageTitle: "three", pageLink: "three"},
      },
  	
  }  //. .  pageData
  
  
  if pagePath == "/" {
      pageTitle = "Index Page"
      pageList = pageList
  }
  
    if pagePath == "/front" {
      pageTitle = "Front Page"
      pageList = pageList
  }
  
    if pagePath == "/main" {
      pageTitle = "Main Page"
      pageList = pageList
  }
  
    if pagePath == "/home" {
      pageTitle = "Home Page"
      pageList = pageList
  }
  
    if pagePath == "/start" {
      pageTitle = "Start Page"
      pageList = pageList
  }

// ,  ° . +
  pageFilePath := template.Must(template.ParseFiles("layout_main_page.html"))
  pageFilePath.Execute(w, pageData)
  
}  //  .  indexHandler



//  .  html url routes 
//  .  as well as terminal cli logs

func main() {
// ,  ° . +
  appName := "~ Cooking JavaScript - // - Website App"

// ,  ° . +
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/front", indexHandler)
  http.HandleFunc("/main", indexHandler)
  http.HandleFunc("/home", indexHandler)
  http.HandleFunc("/start", indexHandler)
  
  // ,  ° . +
  http.HandleFunc("/user", pageHandler)
  http.HandleFunc("/account", pageHandler)
  http.HandleFunc("/profile", pageHandler)
  
  http.HandleFunc("/portfolio", pageHandler)
  http.HandleFunc("/resume", pageHandler)
  
  // ,  ° . +
  http.HandleFunc("/settings", appHandler)
  
    // ,  ° . +
  http.HandleFunc("/inventory", pageHandler)
    // ,  ° . +
  http.HandleFunc("/itinerary", pageHandler)
    // ,  ° . +
  http.HandleFunc("/quest_list", pageHandler)
    // ,  ° . +
  http.HandleFunc("/travel_log", pageHandler)
    // ,  ° . +
  http.HandleFunc("/quest_needs", pageHandler)
    // ,  ° . +
  http.HandleFunc("/quest_people", pageHandler)
  




// -- -
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
    log.Printf("Loading _webapp with default port")
  }
  
// ,  ° . +
  log.Printf("_webapp is active and Listening on port %s", port)

  log.Printf("// -- - %s", appName)
  log.Printf("_webapp now loaded and running at http://localhost:%s", port)

// -- - 
  if err := http.ListenAndServe(":"+port, nil); err != nil {
    log.Fatal("Error Starting the HTTP Server :", err)
    return
  }