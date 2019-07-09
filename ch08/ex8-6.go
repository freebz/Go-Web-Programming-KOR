// 예제 8-6. JSON 데이터에 대한 언마샬링

func unmarshal(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file;", err)
		return
	}
	json.Unmarshal(jsonData, &post)
	return
}
