package main

//SortBooks sort books of a library
func (lib *LibraryT) SortBooks() {

}

//SortLibraries Gets de order of the libraries
func SortLibraries(libraries []LibraryT) (librariesOrder []uint32) {
	librariesOrder = make([]uint32, 0, len(libraries))
	for i, library := range libraries {
		library.SortBooks()
		librariesOrder = append(librariesOrder, uint32(i))
	}
	return
}
