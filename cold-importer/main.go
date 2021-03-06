package main

import (
	"flag"

	lib "github.com/ipfs/go-ipld-eth-import/lib"
)

/*
## EXAMPLE USAGE

make cold && ./build/bin/cold-importer \
	--geth-db-filepath /Users/hj/Documents/tmp/geth-data/geth/chaindata \
	--ipfs-repo-path ~/.ipfs \
	--block-number 0
*/

func main() {
	var (
		blockNumber  uint64
		ipfsRepoPath string
		dbFilePath   string
	)

	// Command line options
	flag.Uint64Var(&blockNumber, "block-number", 0, "Canonical number of the block state to import")
	flag.StringVar(&ipfsRepoPath, "ipfs-repo-path", "~/.ipfs", "IPFS repository path")
	flag.StringVar(&dbFilePath, "geth-db-filepath", "", "Path to the Go-Ethereum Database")
	flag.Parse()

	// IPFS
	ipfs := lib.IpfsInit(ipfsRepoPath)

	// Cold Database
	db := lib.GethDBInit(dbFilePath)
	defer db.Stop()

	// Launch State Traversal
	ts := lib.NewTrieStack(blockNumber)
	defer ts.Close()

	ts.TraverseStateTrie(db, ipfs, blockNumber)

	// Print the metrics
	printReport()
}
