package main

func main() {
	set := Constructor()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Remove(2)
	set.Contains(1)
}

type MyHashSet struct {
    data map[int]bool
}


func Constructor() MyHashSet {
    return MyHashSet {data: make(map[int]bool)}
}


func (this *MyHashSet) Add(key int)  {
    this.data[key] = true
}


func (this *MyHashSet) Remove(key int)  {
    if this.data[key] {
        delete(this.data, key)
    }
}


func (this *MyHashSet) Contains(key int) bool {
    _, ok := this.data[key]
    return ok
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
