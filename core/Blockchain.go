/*
@Time : 2020/3/3 12:50
@Author : pasiyu
@File : Blockchain.go
*/
package core

import (
	"fmt"
	"log"
)

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	gensisBlock := GenerateGenesisBlock()
	blockchain := Blockchain{}
	blockchain.ApendBlock(&gensisBlock)
	return &blockchain
}

func (bc *Blockchain)SendData(data string)  {
	preBlock := bc.Blocks[len(bc.Blocks) - 1]
	newBlock := GenerateNewBlock(*preBlock,data)
	bc.ApendBlock(&newBlock)
}

//添加到区块链中
func (bc *Blockchain)ApendBlock(newBlock *Block)  {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks,newBlock)
		return
	}
	if isValid(*newBlock,*bc.Blocks[len(bc.Blocks) - 1]) {
		bc.Blocks = append(bc.Blocks,newBlock)
	} else {
		log.Fatal("无效块")
	}
}

func (bc *Blockchain) Print()  {
	for _,block := range bc.Blocks {
		fmt.Printf("Index: %d\n",block.Index)
		fmt.Printf("Prev.Hash: %s\n",block.PrevBlockHash)
		fmt.Printf("Curr.Hash: %s\n",block.Hash)
		fmt.Printf("Data: %s\n",block.Data)
		fmt.Printf("Timestamp: %d\n",block.Timestamp)
		fmt.Println()
	}
}

//验证区块是否正确
func isValid(newBlock Block,oldBlock Block) bool {
	if newBlock.Index - 1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}