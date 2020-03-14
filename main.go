package main

import (
    "log"
    "net/http"
    "fmt"
    "encoding/json"
    "strconv"
)

var (
    ARR = [6][8]string{
        {"#","#","#","#","#","#","#","#"},
        {"#",".",".",".",".",".",".","#"},
        {"#",".","#","#","#",".",".","#"},
        {"#",".",".",".","#",".","#","#"},
        {"#","X","#",".",".",".",".","#"},
        {"#","#","#","#","#","#","#","#"}}
    YP int = 4
    XP int = 1

    PRODUCT_NAME string= "T-SHIRT A"
    PRODUCT_ID int = 1
    PRODUCT_QUANTITY int = 3
)

type Response struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Data Data `json:"data"`
}


type Data struct {
    ProductId int `json:"productId"`
    ProductName string `json:"productName"`
    ProductQuantity int`json:"productQuantity"`
}


func kitaraStatus(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":
        data := Data{ProductId: PRODUCT_ID, ProductName: PRODUCT_NAME, ProductQuantity: PRODUCT_QUANTITY}
        response := &Response{Code: 200, Message:"SUCCESS", Data: data}
        e, err := json.Marshal(response)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(string(e))
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(string(e)))
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
    }
}

func getParam(r *http.Request, key string) string {
    vals, ok := r.URL.Query()[key]
    if !ok || len(vals[0]) < 1 {
        log.Println("Url Param '"+key+"' is missing")
        return ""
    }else{
        return vals[0]
    }
}

func kitaraRequest(w http.ResponseWriter, r *http.Request) {
    var productId string = getParam(r, "productId")
    var quantity string = getParam(r, "quantity")
    log.Println(productId+" "+quantity)

    id, err := strconv.Atoi(productId)
    if err != nil {log.Println(err)}
    q, err1 := strconv.Atoi(quantity)
    if err1 != nil {log.Println(err1)}
    

    var resp []byte
    if id != PRODUCT_ID {
        resp = []byte(`{"code":404,"message":"ProductId not found"}`)
    }else if q > 0 && PRODUCT_QUANTITY >= q{
        PRODUCT_QUANTITY = PRODUCT_QUANTITY - q
        resp = []byte(`{"code":200,"message":"SUCCESS"}`)
    }else {
        resp = []byte(`{"code":400,"message":"Product stock is empty"}`)
    }

    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "GET":
        w.WriteHeader(http.StatusOK)
        w.Write(resp)
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
    }
}

type ResponseMag struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Data []Key `json:"data"`
}
type Magazine struct {
    Data [][]int
}
type Key struct {
    Magazine string `json:"magazine"`
    Verified string `json:"verified"`
}

func verify(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var data Magazine
    err := decoder.Decode(&data)
    if err != nil {
        panic(err)
    }

    var ls []Key
    mag := data.Data
    for i:= 0;i<len(mag); i++{
        var v bool = true
        for j:= 0;j<len(mag[i]); j++{
            if mag[i][j] != 1{
                v = false
            }
        }
        fmt.Println(mag[i], v)
        ls = append(ls, Key{Magazine:fmt.Sprint(mag[i]), Verified:fmt.Sprint(v)})
    }

    resp := &ResponseMag{Code: 200, Message:"SUCCESS", Data:ls}

    w.Header().Set("Content-Type", "application/json")
    switch r.Method {
    case "POST":
        e, err := json.Marshal(resp)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(string(e))
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(string(e)))
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
    }
}


func findKey(){
    fmt.Printf("X in [%d,%d]\n\n",YP, XP)
    // var key bool = false
    for i:=0; i<6; i++{
        fmt.Println(ARR[i])
    }
    utara := true
    timur := true
    selatan := true
    incy := 0
    incx := 0
    incz := 0
    y := YP
    x := XP
    z := YP
    for utara{
        incy++
        y = y - 1
        // fmt.Printf("X -> utara [%d,%d] = %s\n",y, XP,ARR[y][XP])
        if(ARR[y][XP] == "#"){
            utara = false
        }else{
            timur = true
            x = XP
            incx = 0
            for timur {
                incx++
                x = x + 1
                // fmt.Printf("\t-> timur [%d,%d] = %s\n",y, x,ARR[y][x])
                if(ARR[y][x] == "#"){
                    timur = false
                }else{
                    selatan = true
                    z = y
                    incz = 0
                    for selatan {
                        incz++
                        z = z + 1
                        // fmt.Printf("\t\t-> selatan [%d,%d] = %s\n",z, x,ARR[z][x])
                        if(ARR[z][x] == "#"){
                            selatan = false
                        }else{
                            ARR[z][x] = "K"
                            fmt.Printf("%d langkah ke utara, %d langkah ke timur dan %d langkah ke selatan\n",incy, incx, incz)
                        }
                    }
                }
            }
        }
    } 

    for i:=0; i<6; i++{
        fmt.Println(ARR[i])
    }
    
}

func main() {
    findKey()

    http.HandleFunc("/kitara-store", kitaraStatus)
    http.HandleFunc("/kitara-store/request", kitaraRequest)
    http.HandleFunc("/soldier/verify", verify)


    log.Fatal(http.ListenAndServe(":8080", nil))
}