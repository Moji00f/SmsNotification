package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Moji00f/SmsNotification/api/entities"
	"github.com/Moji00f/SmsNotification/config"
	"github.com/gin-gonic/gin"
	ptime "github.com/yaa110/go-persian-calendar"
)

type NotificationHandler struct {
}

func NewNotificationHandler() *NotificationHandler {
	return &NotificationHandler{}
}

func (h *NotificationHandler) Alert(c *gin.Context) {

	var AlertData entities.Data

	if err := c.ShouldBindJSON(&AlertData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	message := h.CreateMessage(AlertData)

	h.SendMessage(message)

}

func (h *NotificationHandler) CreateMessage(data entities.Data) string {
	st := ptime.New(data.Alerts[0].StartsAt)
	et := ptime.New(data.Alerts[0].EndsAt)

	fmt.Println(data.Alerts[0].StartsAt)
	fmt.Println(st)

	msg := &entities.Message{
		AlertName:   data.GroupLabels["alertname"],
		Severity:    data.Alerts[0].Labels["severity"],
		Instance:    data.Alerts[0].Labels["instance"],
		Status:      data.Alerts[0].Status,
		StartedAt:   st.Format("yyyy/MM/dd-hh:mm:ss"),
		EndedAt:     et.Format("yyyy/MM/dd-hh:mm:ss"),
		Description: data.Alerts[0].Annotations["description"],
	}
	prefix := "**"
	if data.Alerts[0].Status == "firing" {
		prefix = "** PTF2 PROBLEM ALERT **"
	}
	if data.Alerts[0].Status == "resolved" {
		prefix = "PTF2 RECOVERY ALERT"
		msg.StartedAt = msg.EndedAt
		msg.Severity = "oK"
	}

	sendMsg := prefix + "\n" +
		msg.AlertName + "\n" +
		"status: " + strings.ToUpper(msg.Status) + "\n" +
		"severity: " + msg.Severity + "\n" +
		"Instance: " + msg.Instance + "\n" +
		"Event time: " + msg.StartedAt + "\n" +
		msg.Description

	return sendMsg
}

func (h *NotificationHandler) SendMessage(message string) {
	config := config.GetConfig()
	var sendMsg entities.BankMessage
	for _, phoneNumber := range config.Contacts {
		sendMsg.PhoneNumber = phoneNumber
		sendMsg.Message = message

		var buffer bytes.Buffer

		err := json.NewEncoder(&buffer).Encode(sendMsg)
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &http.Client{}
		req, err := http.NewRequest("POST", config.SmsGateWay.Url, &buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		// req.Header.Add("UserName", config.SmsGateWay.UserName)
		// req.Header.Add("Password", config.SmsGateWay.Password)
		req.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(body))

	}

}
