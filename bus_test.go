package eventbus

import (
	"testing"
	"time"
)

type testEventA struct {
	Value int
}

type testEventB struct {
	Value int
}

func recv[E any](t *testing.T, ch <-chan E) E {
	t.Helper()

	select {
	case v := <-ch:
		return v
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for event")
		var zero E
		return zero
	}
}

func assertNoEvent[E any](t *testing.T, ch <-chan E) {
	t.Helper()

	select {
	case v := <-ch:
		t.Fatalf("received unexpected event: %#v", v)
	case <-time.After(20 * time.Millisecond):
	}
}

func TestSubscribeDeliversMatchingTypeOnly(t *testing.T) {
	b := NewBus()
	defer b.Close()

	sub := Subscribe(b, testEventA{})

	Publish(b, testEventB{Value: 2})
	assertNoEvent(t, sub.Event())

	go Publish(b, testEventA{Value: 1})

	got := recv(t, sub.Event())
	if got.Value != 1 {
		t.Fatalf("got event value %d, want 1", got.Value)
	}
}

func TestMultipleSubscribersReceiveSameEvent(t *testing.T) {
	b := NewBus()
	defer b.Close()

	sub1 := Subscribe(b, testEventA{})
	sub2 := Subscribe(b, testEventA{})

	go Publish(b, testEventA{Value: 10})

	if got := recv(t, sub1.Event()); got.Value != 10 {
		t.Fatalf("sub1 got event value %d, want 10", got.Value)
	}
	if got := recv(t, sub2.Event()); got.Value != 10 {
		t.Fatalf("sub2 got event value %d, want 10", got.Value)
	}
}

func TestSubscriberReceivesEventsInPublishOrder(t *testing.T) {
	b := NewBus()
	defer b.Close()

	sub := Subscribe(b, testEventA{})

	go func() {
		Publish(b, testEventA{Value: 1})
		Publish(b, testEventA{Value: 2})
		Publish(b, testEventA{Value: 3})
	}()

	for want := 1; want <= 3; want++ {
		if got := recv(t, sub.Event()); got.Value != want {
			t.Fatalf("got event value %d, want %d", got.Value, want)
		}
	}
}

func TestSlowSubscriberBlocksPump(t *testing.T) {
	b := NewBus()
	defer b.Close()

	sub := Subscribe(b, testEventA{})

	firstPublished := make(chan struct{})
	go func() {
		Publish(b, testEventA{Value: 1})
		close(firstPublished)
	}()
	<-firstPublished

	secondPublished := make(chan struct{})
	go func() {
		Publish(b, testEventB{Value: 2})
		close(secondPublished)
	}()

	select {
	case <-secondPublished:
		t.Fatal("second publish completed while first event was blocked in subscriber dispatch")
	case <-time.After(20 * time.Millisecond):
	}

	if got := recv(t, sub.Event()); got.Value != 1 {
		t.Fatalf("got event value %d, want 1", got.Value)
	}

	select {
	case <-secondPublished:
	case <-time.After(time.Second):
		t.Fatal("second publish did not complete after subscriber received first event")
	}
}
