package karate

type Chopper interface {
    Chop(needle int) (index int)
    Init(haystack []int)
}
