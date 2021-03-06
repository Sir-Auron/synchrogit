package main

import (
	"fmt"
	"synchrogit/helper"
	"sync"
)

func main() {
	fmt.Println("Synchrogit is warming up. Sit down, watch the world burn")
	defer fmt.Println("Synchrogit is calling it a day, see you next time - copycat")
	synchroGitSettings, _ := helper.ParseConfig()
	var wg sync.WaitGroup
	wg.Add(len(synchroGitSettings.Syncs))
	for index, sync := range synchroGitSettings.Syncs {
		go helper.SyncRepo(index,sync, &wg)
	}
	wg.Wait()
}
