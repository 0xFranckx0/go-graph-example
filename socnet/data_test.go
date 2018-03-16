package socnet

import (
	"fmt"
	"testing"
)

func TestGetDistance(t *testing.T) {
	distance, path := GetDistance("40", "382")

	if distance == 0 {
		fmt.Println(distance, path)
	} else if distance == -1 {
		t.Error("GetDistanceFailed, Error while processing Path")
	} else {
		fmt.Println(distance, path)
	}
}

func TestDistanceNoPath(t *testing.T) {
	distance, path := GetDistance("10", "4780")
	if distance == 0 {
		fmt.Println(distance, path)
	} else if distance == -1 {
		t.Error("GetDistanceFailed, Error while processing Path")
	} else {
		fmt.Println(distance, path)
	}
}

func TestDistanceInvalidID(t *testing.T) {
	distance, path := GetDistance("34", "9999999")
	if distance == 0 {
		fmt.Println(distance, path)
	} else if distance == -1 {
		t.Error("GetDistanceFailed, Error while processing Path")
	} else {
		fmt.Println(distance, path)
	}
}

func TestGetCommonFriends(t *testing.T) {
	friends := GetCommonFriends("4017", "3980")
	if friends == nil {
		t.Error("GetCommonFriends")
	}
	fmt.Println(friends)
}

func TestNoCommonFriends(t *testing.T) {
	friends := GetCommonFriends("6764", "3980")
	if friends == nil {
		t.Error("GetCommonFriends")
	}
	fmt.Println(friends)
}

func TestCommonFriendsInvalidID(t *testing.T) {
	friends := GetCommonFriends("4017", "9999999")
	if friends == nil {
		t.Error("GetCommonFriends")
	}
	fmt.Println(friends)
}
