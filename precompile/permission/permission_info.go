package permission

type PermissionInfo struct {
	tableName   string  `json:"table_name"`
	address     string
	enableNum   string  `json:enable_num`
}

func (p *PermissionInfo) GetTableName() string {
	return p.tableName
}

func (p *PermissionInfo) SetTableName(name string) {
	p.tableName = name
}

func (p *PermissionInfo) GetAddress() string {
	return p.address
}

func (p *PermissionInfo) SetAddress(addr string) {
	p.address = addr
}

func (p *PermissionInfo) GetEnableNum() string {
	return p.enableNum
}

func (p *PermissionInfo) SetEnableNum(enable string) {
	p.enableNum = enable
}
