package main

import (
	"fmt"
	"sort"
	"sync"
)

func RunPipeline(cmds ...cmd) {
	if len(cmds) == 0{
		return
	}

	var wg sync.WaitGroup
	in := make(chan interface{})

	for i, f := range cmds{
		wg.Add(1)
		out := make(chan interface{})

		go func(f cmd, in, out chan interface{}, isLast bool){
			defer func(){
				close(out)
				wg.Done()
			}()
			f(in, out)
		}(f, in, out,  i == len(cmds) - 1)
		in = out
	}

	wg.Wait()
}

func SelectUsers(in, out chan interface{}) {
	// 	in - string
	// 	out - User
	// defer close(out)

	var mu sync.Mutex
	users := make(map[uint64]bool, 100)
	var wg sync.WaitGroup

	for email := range in{
		wg.Add(1)

		go func(email string){
			defer wg.Done()

			user := GetUser(email)
			mu.Lock()
			defer mu.Unlock()

			if users[user.ID]{
				return
			}

			users[user.ID] = true
			out <- user
		}(email.(string))
	}

	wg.Wait()
}

func SelectMessages(in, out chan interface{}) {
	// defer close(out)

	var wg sync.WaitGroup
	batch := make([]User, 0, GetMessagesMaxUsersBatch)

	processBatch := func(users []User) {
		wg.Add(1)
		usersCopy := make([]User, len(users))
		copy(usersCopy, users)
		
		go func() {
			defer wg.Done()
			msgIDs, err := GetMessages(usersCopy...)
			if err != nil {
				return
			}
			for _, id := range msgIDs {
				out <- id
			}
		}()
	}

	for user := range in {
		batch = append(batch, user.(User))
		if len(batch) >= GetMessagesMaxUsersBatch {
			processBatch(batch)
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		processBatch(batch)
	}

	wg.Wait()
}


func CheckSpam(in, out chan interface{}) {
	// in - MsgID
	// out - MsgData
	// defer close(out)

	tmp := make(chan struct{}, HasSpamMaxAsyncRequests)
	var wg sync.WaitGroup

	for msgId := range in{
		wg.Add(1)

		go func(msgId MsgID){
			defer wg.Done()

			tmp <- struct{}{}
			defer func() {<- tmp}()
			
			isSpam, err := HasSpam(msgId)

			if err != nil{
				return
			}

			out <- MsgData{msgId, isSpam}
		}(msgId.(MsgID))
	}

	wg.Wait()

}

func CombineResults(in, out chan interface{}) {
	// in - MsgData
	// out - string
	// defer close(out)

	data := make([]MsgData, 0, 100)
	for msgData := range in{
		data = append(data, msgData.(MsgData))
	}

	sort.Slice(data, func(i, j int) bool {
		if data[i].HasSpam != data[j].HasSpam{
			return data[i].HasSpam
		}
		return data[i].ID < data[j].ID
	})

	for _, msg := range data{
		out <- fmt.Sprintf("%v %d", msg.HasSpam, msg.ID)
	}
}
