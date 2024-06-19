func readLastTile(w io.Writer, projectID, instanceID string, tableName string) error {
	filter := bigtable.ValueFilter("data_plan_0000000000")
	rowFilter := bigtable.RowFilter(filter)
	return readWithFilter(w, projectID, instanceID, tableName, rowFilter)
}
  
