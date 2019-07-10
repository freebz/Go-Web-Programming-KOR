// 예제 9-22. combine 함수

func combine(r image.Rectangle, c1, c2, c3, c4 <-chan image.Image) <-chan string {
	c := make(chan string)

	go func() {
		var wg sync.WaitGroup
		img := image.NewNRGBA(r)
		copy := func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
			draw.Draw(dst, r, src, sp, draw.Src)
			wg.Done()
		}
		wg.Add(4)
		var s1, s2, s3, s4 image.Image
		var ok1, ok2, ok3, ok4 bool
		for {
			select {
			case s1, ok1 = <-c1:
				go copy(img, s1.Bounds(), s1,
					image.Point{r.Min.X, r.Min.Y})
			case s2, ok2 = <-c2:
				go copy(img, s2.Bounds(), s2,
					image.Point{r.Max.X/2, r.Min.Y})
			case s3, ok3 = <-c3:
				go copy(img, s3.Bounds(), s3,
					image.Point{r.Min.X, r.Max.Y/2})
			case s4, ok4 = <-c4:
				go copy(img, s4.Bounds(), s4,
					image.Point{r.Max.X/2, r.Max.Y/2})
			}
			if (ok1 && ok2 && ok3 && ok4) {
				break
			}
		}

		wg.Wait()
		buf2 := new(bytes.Buffer)
		jpeg.Encode(buf2, img, nil)
		c <- base64.StdEncoding.EncodeToString(buf2.Bytes())
	}()
	return c
}
