package service

import (
	"github.com/ZhenlyChen/Agenda-CLI/agenda/model"
	"testing"
)

func TestDataInit(t *testing.T) {
	if err := model.InitDB(model.DataFile{
		User:   	"../model/data/user.json",
		Meeting:  	"../model/data/meeting.json",
		Status:  	"../model/data/status.json",
	}); err != nil {
		t.Error(err)
	}
	model.ClearModel()
}

func TestUserRegister_normal(t *testing.T) {
	err := User().Register(model.UserData{
		Name:		"test1",
		Password: 	"321",
		Email:    	"123@163.com",
		Tel:      	"13866669999",
	})
	if err == nil {
		t.Log("Register test1 success")
	} else {
		t.Error(err)
	}
}

func TestUserRegister_IllegalName(t *testing.T) {
	err := User().Register(model.UserData{
		Name:		"!@#$%^&*",
		Password: 	"321",
	})
	if err == ErrorRegisterIllegalName {
		t.Log("Test illegal name success")
	} else {
		t.Error("Test illegal name failed", err)
	}
}

func TestUserRegister_DuplicateName(t *testing.T) {
	err := User().Register(model.UserData{
		Name:		"test1",
		Password: 	"321",
	})
	if err == ErrorRegisterDuplicateName {
		t.Log("Test duplicate name success")
	} else {
		t.Error("Test duplicate name failed", err)
	}
}

func TestUserLogin_normal(t *testing.T) {
	err := User().Login("test1", "321")
	if err == nil {
		t.Log("Login test1 success")
	} else {
		t.Error(err)
	}
}

func TestUserLogin_NullUser(t *testing.T) {
	err := User().Login("", "321")
	if err == ErrorLoginNullUser {
		t.Log("Test null user success")
	} else {
		t.Error("Test null user failed", err)
	}
}

func TestUserLogin_ErrorPassword(t *testing.T) {
	err := User().Login("test1", "666")
	if err == ErrorLoginErrorPassword {
		t.Log("Test error password success")
	} else {
		t.Error("Test error password failed", err)
	}
}

func TestStatusGetLoginUser(t *testing.T) {
	result := Status().GetLoginUser()
	if result == "test1" {
		t.Log("Get login user success")
	} else {
		t.Error("Get login user failed")
	}
}

func TestStatusClearStatus(t *testing.T) {
	Status().ClearStatus()
	result := Status().GetLoginUser()
	if result == "" {
		t.Log("Clear status success")
	} else {
		t.Error("Clear status failed")
	}
}

func TestMeetingCreate(t *testing.T) {
	//Create some users
	User().Register(model.UserData{Name: "test2", Password: "123",})
	User().Register(model.UserData{Name: "test3", Password: "123",})
	User().Register(model.UserData{Name: "test4", Password: "123",})
	User().Register(model.UserData{Name: "test5", Password: "123",})
	User().Register(model.UserData{Name: "test6", Password: "123",})
	Status().ClearStatus()
	User().Login("test1", "321")

	_, err := Meeting().Create(model.MeetingData{
		Title: "Meeting1",
		Presenter: "test1",
		Participator: []string{"test2", "test3"},
	}, "2019/10/01-13:33", "2019/10/01-15:33")
	if err == nil {
		t.Log("Create 'meeting1' success")
	} else {
		t.Error(err)
	}
}

func TestMeetingCreate_TimeOutOfRange(t *testing.T) {
	//error month
	_, err := Meeting().Create(model.MeetingData{Title: "Meeting2", Presenter: "test1", Participator: []string{"test2", "test3"},
	}, "2019/13/01-13:33", "2019/10/01-15:33")
	if err == ErrorTimeOutOfRange {
		t.Log("Test time out of range success")
	} else {
		t.Error("Test time out of range failed", err)
	}
}

func TestMeetingCreate_EndTimeEarly(t *testing.T) {
	_, err := Meeting().Create(model.MeetingData{Title: "Meeting2", Presenter: "test1", Participator: []string{"test2", "test3"},
	}, "2019/12/01-13:33", "2019/10/01-15:33")
	if err == ErrorTimeEndTimeEarly {
		t.Log("Test early end time success")
	} else {
		t.Error("Test early end time failed", err)
	}
}

func TestMeetingCreate_DuplicateTitle(t *testing.T) {
	_, err := Meeting().Create(model.MeetingData{Title: "Meeting1", Presenter: "test1", Participator: []string{"test2", "test3"},
	}, "2019/10/01-13:33", "2019/10/01-15:33")
	if err == ErrorMeetingDuplicateTitle {
		t.Log("Test duplicate title success")
	} else {
		t.Error("Test duplicate title failed", err)
	}
}

func TestMeetingCreate_UserNotExist(t *testing.T) {
	u, err := Meeting().Create(model.MeetingData{Title: "Meeting2", Presenter: "test1", Participator: []string{"test9", "test10"},
	}, "2019/10/01-13:33", "2019/10/01-15:33")
	if err == ErrorUserNotExist {
		t.Log("Test user not exist success: ", u)
	} else {
		t.Error("Test user not exist failed", err)
	}
}

func TestMeetingCreate_BothPresenterAndParticipator(t *testing.T) {
	_, err := Meeting().Create(model.MeetingData{Title: "Meeting2", Presenter: "test1", Participator: []string{"test1", "test2"},
	}, "2019/10/01-13:33", "2019/10/01-15:33")
	if err == ErrorBothPresenterAndParticipator {
		t.Log("Test both presenter and participator success")
	} else {
		t.Error("Test both presenter and participator failed", err)
	}
}

func TestMeetingCreate_Overlap(t *testing.T) {
	u, err := Meeting().Create(model.MeetingData{Title: "Meeting2", Presenter: "test1", Participator: []string{"test2", "test3"},
	}, "2019/10/01-14:33", "2019/10/01-16:33")
	if err == ErrorMeetingOverlap {
		t.Log("Test meeting overlap success: ", u)
	} else {
		t.Error("Test meeting overlap failed", err)
	}
}

func TestMeetingAddParticipator(t *testing.T) {
	err := Meeting().AddParticipator("Meeting1", []string{"test5"})
	if err == nil {
		t.Log("Add test5 to meeting1 success")
	} else {
		t.Error(err)
	}
}

func TestMeetingAddParticipator_NotExist(t *testing.T) {
	err := Meeting().AddParticipator("Meeting_err", []string{"test6"})
	if err == ErrorMeetingNotExist {
		t.Log("Test meeting not exist success")
	} else {
		t.Error("Test meeting not exist failed", err)
	}
}

func TestMeetingAddParticipator_NotPresenter(t *testing.T) {
	Status().ClearStatus()
	User().Login("test2", "123")
	err := Meeting().AddParticipator("Meeting1", []string{"test6"})
	if err == ErrorNotPresenter {
		t.Log("Test not presenter success")
	} else {
		t.Error("Test not presenter failed", err)
	}
}

func TestMeetingAddParticipator_ParticipatorExist(t *testing.T) {
	Status().ClearStatus()
	User().Login("test1", "321")
	err := Meeting().AddParticipator("Meeting1", []string{"test2"})
	if err == ErrorParticipatorExist {
		t.Log("Test participator exist success")
	} else {
		t.Error("Test participator exist failed", err)
	}
}

func TestMeetingAddParticipator_Overlap(t *testing.T) {
	Status().ClearStatus()
	User().Login("test4", "123")
	Meeting().Create(model.MeetingData{
		Title: "Meeting3",
		Presenter: "test4",
		Participator: []string{},
	}, "2019/10/01-14:33", "2019/10/01-16:33")

	Status().ClearStatus()
	User().Login("test1", "321")
	err := Meeting().AddParticipator("Meeting1", []string{"test4"})
	if err == ErrorMeetingOverlap {
		t.Log("Test participator time overlap success")
	} else {
		t.Error("Test participator time overlap failed", err)
	}
}

func TestMeetingRemoveParticipator_NotExist(t *testing.T) {
	err := Meeting().RemoveParticipator("Meeting2", []string{})
	if err == ErrorMeetingNotExist {
		t.Log("Test meeting not exist success")
	} else {
		t.Error("Test meeting not exist failed", err)
	}
}

func TestMeetingRemoveParticipator_NotPresenter(t *testing.T) {
	err := Meeting().RemoveParticipator("Meeting3", []string{"test4"})
	if err == ErrorNotPresenter {
		t.Log("Test not presenter success")
	} else {
		t.Error("Test not presenter failed", err)
	}
}

func TestMeetingRemoveParticipator_ParticipatorNotExist(t *testing.T) {
	err := Meeting().RemoveParticipator("Meeting1", []string{"test4"})
	if err == ErrorParticipatorNotExist {
		t.Log("Test not participator success")
	} else {
		t.Error("Test not participator failed", err)
	}
}

func TestMeetingRemoveParticipator(t *testing.T) {
	err := Meeting().RemoveParticipator("Meeting1", []string{"test5"})
	if err == nil {
		t.Log("Remove test5 from Meeting1 success")
	} else {
		t.Error(err)
	}
}

func TestMeetingQuery(t *testing.T) {
	data, err := Meeting().Query("test1", "2019/10/01-11:33", "2019/10/01-16:33")
	if err == nil && len(data) > 0 {
		t.Log("Query meeting success: ", data[0].Title)
	} else {
		t.Error(err)
	}
}

func TestMeetingDelete_NotExist(t *testing.T) {
	err := Meeting().Delete("Meeting_error")
	if err == ErrorMeetingNotExist {
		t.Log("Test meeting not exist success")
	} else {
		t.Error("Test meeting not exist failed", err)
	}
}

func TestMeetingDelete_NotPresenter(t *testing.T) {
	err := Meeting().Delete("Meeting3")
	if err == ErrorNotPresenter {
		t.Log("Test not presenter success")
	} else {
		t.Error("Test not presenter failed", err)
	}
}

func TestMeetingDelete(t *testing.T) {
	Status().ClearStatus()
	User().Login("test4", "123")
	err := Meeting().Delete("Meeting3")
	if err == nil {
		t.Log("Delete Meeting3 success")
	} else {
		t.Error(err)
	}
}

func TestMeetingQuit_NotExist(t *testing.T) {
	err := Meeting().Quit("Meeting_error")
	if err == ErrorMeetingNotExist {
		t.Log("Test meeting not exist success")
	} else {
		t.Error("Test meeting not exist failed", err)
	}
}

func TestMeetingQuit_ParticipatorNotExist(t *testing.T) {
	err := Meeting().Quit("Meeting1")
	if err == ErrorParticipatorNotExist {
		t.Log("Test not participator success")
	} else {
		t.Error("Test not participator failed", err)
	}
}

func TestMeetingQuit(t *testing.T) {
	Status().ClearStatus()
	User().Login("test3", "123")
	err := Meeting().Quit("Meeting1")
	if err == nil {
		t.Log("Test3 quit Meeting1 success")
	} else {
		t.Error(err)
	}
}

func TestMeetingClear(t *testing.T) {
	Status().ClearStatus()
	User().Login("test1", "321")
	err := Meeting().Clear()
	data, err2 := Meeting().Query("test1", "2019/10/01-11:33", "2019/10/01-16:33")
	if err == nil && err2 == nil && len(data) == 0 {
		t.Log("Clear success")
	} else {
		t.Error("Clear failed")
	}
}















