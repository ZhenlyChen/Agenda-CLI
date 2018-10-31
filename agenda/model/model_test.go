package model

import (
	"testing"
	"time"
)

func TestDataInit(t *testing.T) {
	if err := InitDB(DataFile{
		User:   	"./data/user.json",
		Meeting:  	"./data/meeting.json",
		Status:  	"./data/status.json",
	}); err != nil {
		t.Error(err)
	}
	ClearModel()
}

func TestUserAdd(t *testing.T) {
	result := User().Add(UserData{
		Name:		"test1",
		Password: 	"123",
		Email:    	"666@666.com",
		Tel:      	"13899992323",
	})
	if result == nil {
		t.Log("Add user success")
	} else {
		t.Error("Add user failed", result)
	}
}

func TestUserGetByName(t *testing.T) {
	result := User().GetByName("test1")
	if result.Name == "test1" {
		t.Log("Find user 'test1' success")
	} else {
		t.Error("Find user 'test1' failed")
	}
}

func TestUserExist(t *testing.T) {
	result := User().Exist("test1")
	if result == true {
		t.Log("'test1' exists")
	} else {
		t.Error("'test1' doesn't exist")
	}
}

func TestStatusSetUser(t *testing.T) {
	Status().SetUser("test1")
	result := Status().GetStatus()
	if result.User == "test1" {
		t.Log("Set Status 'test1' success")
	} else {
		t.Error("Set Status 'test1' failed")
	}
}

func TestRefreshTime(t *testing.T) {
	nowTime := Status().GetStatus().Expires
	time.Sleep(time.Second)
	Status().RefreshTime()
	result := Status().GetStatus()
	if result.Expires != nowTime {
		t.Log("Refresh time success")
	} else {
		t.Error("Refresh time failed", nowTime, result.Expires)
	}
}

func TestClearStatus(t *testing.T) {
	Status().ClearStatus()
	result := Status().GetStatus()
	if result.User == "" && result.Expires == 0 {
		t.Log("Clear status success")
	} else {
		t.Error("Clear status failed")
	}
}

func TestMeetingAdd(t *testing.T) {
	Status().SetUser("test1")
	User().Add(UserData{Name: "test2",  Password: "123",})
	start_time, _ := time.Parse("2006/01/02-15:04", "2019/10/01-13:33")
	end_time, _ := time.Parse("2006/01/02-15:04", "2019/10/01-15:33")
	err := Meeting().Add(MeetingData{
		Title: 			"Meeting1",
		Presenter:		 "test1",
		Participator: 	[]string{"test2"},
		Start:			start_time.Unix(),
		End:			end_time.Unix(),
	})
	if err == nil {
		t.Log("Add 'meeting1' success")
	} else {
		t.Error(err)
	}
}

func TestMeetingExist(t *testing.T) {
	result := Meeting().Exist("Meeting1")
	if result == true {
		t.Log("Test meeting exist success")
	} else {
		t.Error("Test meeting exist failed")
	}
}

func TestMeetingQuery(t *testing.T) {
	start_time, _ := time.Parse("2006/01/02-15:04", "2019/10/01-13:33")
	end_time, _ := time.Parse("2006/01/02-15:04", "2019/10/01-15:33")
	data := Meeting().Query("test1", start_time.Unix(), end_time.Unix())
	if len(data) > 0 {
		t.Log("Meeting query success: ", data[0].Title)
	} else {
		t.Error("Meeting query failed")
	}
}

func TestMeetingAddParticipator(t *testing.T) {
	User().Add(UserData{Name: "test3",  Password: "123",})
	err := Meeting().AddParticipator("Meeting1", []string{"test3"})
	if err == nil {
		t.Log("Add 'test2' to 'meeting1' success")
	} else {
		t.Error(err)
	}
}

func TestMeetingIsParticipator(t *testing.T) {
	result := Meeting().IsParticipator("Meeting1", "test2")
	if result == true {
		t.Log("Test IsParticipator success")
	} else {
		t.Error("Test IsParticipator failed")
	}
}


func TestMeetingIsPresenter(t *testing.T) {
	result := Meeting().IsPresenter("Meeting1", "test1")
	if result == true {
		t.Log("Test IsPresenter success")
	} else {
		t.Error("Test IsPresenter failed")
	}
}

func TestMeetingGetMeetingByTitle(t *testing.T) {
	data := Meeting().GetMeetingByTitle("Meeting1")
	if data.Title == "Meeting1"{
		t.Log("Get meeting by name success")
	} else {
		t.Error("Get meeting by name failed")
	}
}

func TestMeetingGetMeetingAsPresenter(t *testing.T) {
	data := Meeting().GetMeetingAsPresenter("test1")
	if len(data) > 0 {
		t.Log("Get meeting as presenter: ", data[0].Title)
	} else {
		t.Error("Get meeting as presenter failed")
	}
}

func TestMeetingGetMeetingAsParticipator(t *testing.T) {
	data := Meeting().GetMeetingAsParticipator("test2")
	if len(data) > 0 {
		t.Log("Get meeting as participator success: ", data[0].Title)
	} else {
		t.Error("Get meeting as participator failed")
	}
}

func TestMeetingGetMeetingByName(t *testing.T) {
	start_time, _ := time.Parse("2006/01/02-15:04", "2019/11/01-13:33")
	end_time, _ := time.Parse("2006/01/02-15:04", "2019/11/01-15:33")
	Meeting().Add(MeetingData{
		Title: 			"Meeting2",
		Presenter:		 "test2",
		Participator: 	[]string{"test1"},
		Start:			start_time.Unix(),
		End:			end_time.Unix(),
	})
	data := Meeting().GetMeetingByName("test2")
	if len(data) > 1 {
		t.Log("Get meeting by name success: ", data[0].Title, data[1].Title)
	} else {
		t.Error("Get meeting by name failed")
	}
}

func TestMeetingRemoveParticipator(t *testing.T) {
	err := Meeting().RemoveParticipator("Meeting1", "test2")
	data := Meeting().GetMeetingByName("test2")
	if err == nil && len(data) < 2 {
		t.Log("Remove test2 from Meeting1 success")
	} else {
		t.Error("Remove test2 from Meeting1 failed")
	}
}

func TestMeetingDelete(t *testing.T) {
	err := Meeting().Delete("Meeting1")
	data := Meeting().GetMeetingByName("test1")
	if err == nil && len(data) < 2 {
		t.Log("Delete Meeting1 success")
	} else {
		t.Error("Delete Meeting1 failed")
	}
}



