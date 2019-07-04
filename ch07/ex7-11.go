// 예제 7-11. Decoder를 사용한 JSON 파싱

jsonFile, err := os.Open("post.json")
if err != nil {
	fmt.Println("Error opening JSON file:", err)
	return
}
defer jsonFile.Close()

decoder := json.NewDecoder(jsonFile)

for {
	var post Post
	err := decoder.Decode(&post)
	if err == io.EOF {
		break
	}
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println(post)
}
