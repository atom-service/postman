package dao

// MustInitTables MustInitTables
func MustInitTables() {
	var err error
	err = InitTables()
	if err != nil {
		panic(err)
	}
}

// InitTables InitTables
func InitTables() error {
	var err error
	err = createLabelTable()
	if err != nil {
		return err
	}

	return nil
}

// MustTruncateTables MustTruncateTables
func MustTruncateTables() {
	var err error
	err = TruncateTables()
	if err != nil {
		panic(err)
	}
}

// TruncateTables TruncateTables
func TruncateTables() error {
	var err error
	err = truncateLabelTable()
	if err != nil {
		return err
	}
	err = truncateLabelMappingTable()
	if err != nil {
		return err
	}
	return nil
}
