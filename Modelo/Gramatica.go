package Modelo


type Gramatica struct {
    Terminales   []string `json:"Terminales"`
    NoTerminales []string    `json:"NoTerminales"`
    Expresiones  map[string][][]string `json:"Expresiones"`
}