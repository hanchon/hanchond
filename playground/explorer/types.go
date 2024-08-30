package explorer

type Transaction struct {
	cosmoshash string
	ethhash    string
	typeURL    string
	sender     string
	height     int64
}

type Block struct {
	height  int64
	txcount int64
	hash    string
	txns    []Transaction
}

func NewBlock(height, txcount int64, hash string) *Block {
	return &Block{
		height:  height,
		txcount: txcount,
		hash:    hash,
		txns:    make([]Transaction, txcount),
	}
}

func (b *Block) AddTransaction(index int, cosmosHash, ethHash, typeURL, sender string) {
	b.txns[index] = Transaction{
		cosmoshash: cosmosHash,
		ethhash:    ethHash,
		typeURL:    typeURL,
		sender:     sender,
		height:     b.height,
	}
}
