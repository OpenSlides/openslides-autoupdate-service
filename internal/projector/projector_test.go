package projector_test

// func TestLiveProjectorDoesNotExist(t *testing.T) {
// 	closed := make(chan struct{})
// 	defer close(closed)

// 	ds := test.NewMockDatastore(closed, nil)

// 	p := projector.New(ds, testSlides())
// 	buf := new(bytes.Buffer)

// 	if err := p.Live(context.Background(), 1, lineWriter{maxLines: 1, wr: buf}, []int{1}); err != nil {
// 		if !errors.Is(err, errWriterFull) {
// 			t.Fatalf("Live returned unexpected error: %v", err)
// 		}
// 	}

// 	expect := []byte(`{"1":null}` + "\n")
// 	if got := buf.Bytes(); !bytes.Equal(got, expect) {
// 		t.Errorf("Got `%s`, expected `%s`", got, expect)
// 	}
// }

// func TestLiveExistingProjector(t *testing.T) {
// 	closed := make(chan struct{})
// 	defer close(closed)

// 	ds := test.NewMockDatastore(closed, map[string]string{
// 		"projector/1/current_projection_ids": "[1]",
// 		"projection/1/stable":                "true",
// 		"projection/1/content_object_id":     `"test_model/1"`,
// 	})
// 	p := projector.New(ds, testSlides())
// 	buf := new(bytes.Buffer)

// 	if err := p.Live(context.Background(), 1, lineWriter{maxLines: 1, wr: buf}, []int{1}); err != nil {
// 		if !errors.Is(err, errWriterFull) {
// 			t.Fatalf("Live returned unexpected error: %v", err)
// 		}
// 	}

// 	expect := `{"1":{"1":{"data":"test_model","stable":true,"content_object_id":"test_model/1"}}}` + "\n"
// 	require.JSONEq(t, expect, string(buf.Bytes()))
// }

// func TestLiveProjectionWithType(t *testing.T) {
// 	closed := make(chan struct{})
// 	defer close(closed)

// 	ds := test.NewMockDatastore(closed, map[string]string{
// 		"projector/1/current_projection_ids": "[1]",
// 		"projection/1/content_object_id":     `"test_model/1"`,
// 		"projection/1/type":                  `"test1"`,
// 	})
// 	p := projector.New(ds, testSlides())
// 	buf := new(bytes.Buffer)

// 	if err := p.Live(context.Background(), 1, lineWriter{maxLines: 1, wr: buf}, []int{1}); err != nil {
// 		if !errors.Is(err, errWriterFull) {
// 			t.Fatalf("Live returned unexpected error: %v", err)
// 		}
// 	}

// 	expect := `{"1":{"1":{"data":"abc","type":"test1","stable":false,"content_object_id":"test_model/1"}}}` + "\n"
// 	require.JSONEq(t, expect, string(buf.Bytes()))
// }

// func TestLiveProjectionKeepOpenUntilContextCloses(t *testing.T) {
// 	t.Skip() // TODO
// 	closed := make(chan struct{})
// 	defer close(closed)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	ds := test.NewMockDatastore(closed, map[string]string{})
// 	p := projector.New(ds, testSlides())
// 	buf := new(bytes.Buffer)

// 	var err error
// 	done := make(chan struct{})
// 	go func() {
// 		defer close(done)
// 		err = p.Live(ctx, 1, buf, []int{1})
// 	}()

// 	timer := time.NewTimer(10 * time.Microsecond)
// 	defer timer.Stop()
// 	select {
// 	case <-done:
// 		t.Errorf("Live() closed to early")
// 	case <-timer.C:
// 	}

// 	// Close the connection.
// 	cancel()

// 	timer.Reset(10 * time.Microsecond)
// 	select {
// 	case <-done:
// 	case <-timer.C:
// 		t.Errorf("Live() did not close")
// 	}
// 	assert.NoErrorf(t, err, "Live returned unexpected error: %w", err)
// }

// func TestLiveProjectionKeepOpenUntilServiceCloses(t *testing.T) {
// 	t.Skip()
// 	// TODO
// 	closed := make(chan struct{})

// 	ds := test.NewMockDatastore(closed, map[string]string{})
// 	p := projector.New(ds, testSlides())
// 	buf := new(bytes.Buffer)

// 	var err error
// 	done := make(chan struct{})
// 	go func() {
// 		defer close(done)
// 		err = p.Live(context.Background(), 1, buf, []int{1})
// 	}()

// 	timer := time.NewTimer(100 * time.Microsecond)
// 	defer timer.Stop()
// 	select {
// 	case <-done:
// 		t.Errorf("Live() closed to early")
// 	case <-timer.C:
// 	}

// 	// Close the Service.
// 	close(closed)

// 	timer.Reset(100 * time.Microsecond)
// 	select {
// 	case <-done:
// 	case <-timer.C:
// 		t.Errorf("Live() did not close")
// 	}
// 	assert.NoErrorf(t, err, "Live returned unexpected error: %w", err)
// }

// // func TestLiveUpdatedData(t *testing.T) {
// // 	closed := make(chan struct{})
// // 	defer close(closed)

// // 	ds := test.NewMockDatastore(map[string]string{
// // 		"projector/1/current_projection_ids": "[1]",
// // 		"projection/1/content_object_id":     `"test_model/1"`,
// // 		"projection/1/type":                  `"test1"`,
// // 	})
// // 	p := projector.New(ds, testSlides(), closed)
// // 	buf := new(bytes.Buffer)
// // 	ch := make(chan string, 1)
// // 	w := channelWriter{ch: ch, wr: buf}
// // 	ctx, cancel := context.WithCancel(context.Background())
// // 	defer cancel()

// // 	go func() {
// // 		p.Live(ctx, 1, w, []int{1})
// // 	}()

// // 	msg := <-ch
// // 	expect := `{"1":{"1":{"data":"abc","stable":false,"content_object_id":"test_model/1","type":"test1"}}}` + "\n"
// // 	assert.JSONEqf(t, expect, msg, "First response")

// // 	ds.SendValues(map[string]string{"projection/1/type": "test_model"})
// // 	msg = <-ch
// // 	expect = `{"1":{"1":{"data":"test_model","stable":false,"content_object_id":"test_model/1","type":"test_model"}}}` + "\n"
// // 	assert.JSONEqf(t, expect, msg, "Second response")
// // }

// func testSlides() *projector.SlideStore {
// 	s := new(projector.SlideStore)
// 	s.AddFunc("test1", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
// 		return []byte(`"abc"`), nil, nil
// 	})
// 	s.AddFunc("test_model", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
// 		return []byte(`"test_model"`), nil, nil
// 	})
// 	return s
// }

// var errWriterFull = errors.New("first line full")

// // lineWriter fails after the first newline
// type lineWriter struct {
// 	maxLines int
// 	wr       io.Writer
// 	count    int
// }

// func (w lineWriter) Write(p []byte) (int, error) {
// 	if w.count >= w.maxLines {
// 		return 0, errWriterFull
// 	}

// 	idx := bytes.IndexByte(p, '\n')
// 	if idx != -1 {
// 		w.count++
// 		n, err := w.wr.Write(p[:idx+1])
// 		if err != nil {
// 			return n, err
// 		}
// 		return n, errWriterFull
// 	}

// 	return w.wr.Write(p)
// }

// type channelWriter struct {
// 	ch  chan<- string
// 	wr  io.Writer
// 	buf bytes.Buffer
// }

// func (w channelWriter) Write(p []byte) (int, error) {
// 	idx := bytes.IndexByte(p, '\n')
// 	if idx == -1 {
// 		return w.buf.Write(p)
// 	}

// 	n, err := w.buf.Write(p[:idx+1])
// 	if err != nil {
// 		return n, err
// 	}

// 	w.ch <- w.buf.String()
// 	w.buf.Reset()

// 	n2, err := w.Write(p[idx+1:])
// 	return n + n2, err
// }
