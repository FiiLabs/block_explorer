package tasks

import "github.com/FiiLabs/block_explorer/monitor"

func Start(synctask SyncTask) {
	go synctask.StartCreateTask()
	go synctask.StartExecuteTask()
	go monitor.Start()
}
