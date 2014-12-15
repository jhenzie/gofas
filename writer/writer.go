package writer

// Writer is responsible for the Atomic, Durable, Isolated and Consisten, write of new facts to the sytem
// For a given fact store there is one writer, whilst many might see this as a weakness in the design, the ACID
// characterisitcs are worth it.  If you have a write heavy system, only benchmarking will determine if gofas
// is suitable for your application, however, the pipeline nature of the system means that much can be accomplished
// with a signle writer, and multiple readers.

type writer struct {
	transactionID uint64
	ingestQueue   chan<- interface{}
}

func (w *writer) receive(data []byte) uint64 {
	decodedData := w.decode(data)
	transactionID := w.journal(decodedData)

	return transactionID
}

func (w *writer) decode(data []byte) interface{} {
	return data
}

func (w *writer) journal(data interface{}) uint64 {
	return w.nextTransactionID()
}

func (w *writer) nextTransactionID() uint64 {
	w.transactionID += 1
	return w.transactionID
}
