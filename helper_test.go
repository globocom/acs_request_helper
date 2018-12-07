package acsrequesthelper

import (
	"testing"
)

func TestBuildQueryString(t *testing.T) {
	/*
		http://localhost:8080/client/api
		?command=listUsers
		&response=json
		&apiKey=plgWJfZK4gyS3mOMTVmjUVg-X-jlWlnfaUJ9GAbBbf9EdM-kAYMmAiLqzzq1ElZLYq_u38zCm0bewzGUdP66mg
		&signature=TTpdDq%2F7j%2FJ58XCRHomKoQXEQds%3D
	*/

	params := map[string]string{
		"command":  "listUsers",
		"response": "json",
	}

	ro := BuildRequestObject("http://localhost", 8080,
		"plgWJfZK4gyS3mOMTVmjUVg-X-jlWlnfaUJ9GAbBbf9EdM-kAYMmAiLqzzq1ElZLYq_u38zCm0bewzGUdP66mg",
		"VDaACYb0LV9eNjTetIOElcVQkvJck_J_QljX_FcHRj87ZKiy0z0ty0ZsYBkoXkY9b7eq1EhwJaw7FF3akA3KBQ",
		params,
	)

	qs := ro.BuildQueryString()

	expected := "apiKey=plgWJfZK4gyS3mOMTVmjUVg-X-jlWlnfaUJ9GAbBbf9EdM-kAYMmAiLqzzq1ElZLYq_u38zCm0bewzGUdP66mg&command=listUsers&response=json&signature=TTpdDq%2F7j%2FJ58XCRHomKoQXEQds%3D"
	if qs != expected {
		t.Errorf("Expected %s, but got %s", expected, qs)
	}
}
