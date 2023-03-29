package app

const samplePath = "D:/Desktop/sample.pcap"

func (it *T) SwitchToDev() error {
	it.session.filename = samplePath
	return it.StartReader("")
	//r := reader.New(it.log, samplePath, it.v, cache.NewInMem)
	//_, err := r.Start("")
	//if err != nil {
	//	it.log.Error(err.Error())
	//	return err
	//}
}
