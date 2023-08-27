package main

var subscribers = []func(v any){}

type  ObservableMethods interface {
    next();
    subscribe(fn func(v any)) func() bool;
}


type Observable struct  {    
    ObservableMethods
}

func (observable Observable) next(data interface{}){
    for _, member := range subscribers {
        member(data)
    }

}

func unsubscribe( index int) bool {
     newArr := make([]func(v any), len(subscribers) - 2)
     currentIndex := 0
        for i, fn := range subscribers {
            if(i != index){
               newArr[currentIndex] = fn
               currentIndex++
            }
        }
      subscribers = newArr
        return true
    }

func (observable Observable) subscribe(fn func(v interface{})) func(){    
    index :=  len(subscribers) - 1
    subscribers = append(subscribers, fn);

    return func ()  {
        unsubscribe(index)
    } 

}