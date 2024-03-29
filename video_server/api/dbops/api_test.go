package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// init(dblogin,truncate tables) -> run tests->clear data(truncate tables)
var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")

}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)

}
func testAddUser(t *testing.T) {
	err := AddUserCredential("test", "1")
	if err != nil {
		t.Errorf("error of AddUser: %v", err)
	}

}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if pwd != "1" || err != nil {
		t.Errorf("error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("test", "1")
	if err != nil {
		t.Errorf("error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("test")
	if err != nil {
		t.Errorf("error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("deleteing user test failed")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("error of RegetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)

}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"
	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("error of AddComments: %v", err)
	}

}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("error of ListComments: %v", err)
	}
	for i, ele := range res {
		fmt.Printf("comment:%d,%v\n", i, ele)
	}
}
