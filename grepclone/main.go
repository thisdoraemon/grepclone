package main

import (
	"fmt"
	"grepclone/grep-app/worker"
	"grepclone/grep-app/worklist"
	"os"
	"path/filepath"
	"sync"
)

func GetAllFiles(wl *worklist.Worklist, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("readdir error:", err)
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			nextPath := filepath.Join(path, entry.Name())
			GetAllFiles(wl, nextPath)
		} else {
			wl.Add(worklist.NewJob(filepath.Join(path, entry.Name())))
		}
	}
}

func main() {
	var workerWg sync.WaitGroup
	var mu sync.Mutex

	wl := worklist.New(100)

	var results []worker.Result

	numWorkers := 10

	workerWg.Add(1)
	// Get all files
	go func() {
		defer workerWg.Done()
		GetAllFiles(&wl, os.Args[2])
		wl.Finalize(numWorkers)
	}()

	// Find matches
	for i := 0; i < numWorkers; i++ {
		workerWg.Add(1)
		go func() {
			defer workerWg.Done()
			for {
				workEntry := wl.Next()
				if workEntry.Path != "" {
					workerResult := worker.FindInFile(workEntry.Path, os.Args[1])
					if workerResult != nil {
						mu.Lock()
						results = append(results, workerResult.Inner...)
						mu.Unlock()
					}
				} else {
					return
				}
			}
		}()
	}

	workerWg.Wait()
	for _, r := range results {
		fmt.Printf("%v[%v]:%v\n", r.Path, r.LineNum, r.Line)
	}

}
