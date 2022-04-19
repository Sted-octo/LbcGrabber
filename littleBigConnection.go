package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"jaytaylor.com/html2text"
)

func littleBigConnectionResultObjectToStringMapConverter(lbcResult *LbcResponse, missionsDetails map[string]*LbcDetail) ([][]string, error) {
	var records [][]string = make([][]string, len(lbcResult.Results)+1)

	indexLine := 0
	head := []string{"\xEF\xBB\xBFcompany", "encodedId", "title", "city", "workUnit", "WorkUnitType", "averageDailyRate", "publicationDate", "missionDuration", "remoteWork", "contactName", "description", "goals", "skills", "deliverables"}
	records[indexLine] = head
	indexLine++

	for _, mission := range lbcResult.Results {
		indexColumn := 0
		records[indexLine] = make([]string, len(head))
		records[indexLine][indexColumn] = mission.Company.Name
		indexColumn++
		records[indexLine][indexColumn] = mission.RequestForProposal.EncodedID
		indexColumn++
		records[indexLine][indexColumn] = mission.RequestForProposal.Title
		indexColumn++
		records[indexLine][indexColumn] = mission.RequestForProposal.Location.City
		indexColumn++
		records[indexLine][indexColumn] = strconv.FormatInt(mission.RequestForProposal.WorkUnit.WorkUnits, 10)
		indexColumn++
		records[indexLine][indexColumn] = string(mission.RequestForProposal.WorkUnit.WorkUnitType)
		indexColumn++
		records[indexLine][indexColumn] = strconv.FormatFloat(mission.RequestForProposal.AverageDailyRate, 'f', 2, 64)
		indexColumn++
		records[indexLine][indexColumn] = mission.RequestForProposal.PublicationDate
		indexColumn++
		records[indexLine][indexColumn] = strconv.FormatInt(mission.RequestForProposal.MissionDuration, 10)
		indexColumn++
		records[indexLine][indexColumn] = mission.RequestForProposal.RemoteWork
		indexColumn++

		if missionDetail, found := missionsDetails[mission.RequestForProposal.EncodedID]; found {
			records[indexLine][indexColumn] = missionDetail.ContactName
			indexColumn++
			records[indexLine][indexColumn] = missionDetail.Description
			indexColumn++
			records[indexLine][indexColumn] = missionDetail.Goals
			indexColumn++
			records[indexLine][indexColumn] = missionDetail.Skills
			indexColumn++
			records[indexLine][indexColumn] = missionDetail.Deliverables
			indexColumn++
		} else {
			for i := 0; i < 5; i++ {
				records[indexLine][indexColumn] = ""
				indexColumn++
			}

		}

		indexLine++
	}

	return records, nil
}

func jsonToLittleBigConnectionResultObjectConverter(body []byte) (*LbcResponse, error) {
	var lbcResult LbcResponse

	err := json.Unmarshal(body, &lbcResult)

	if err != nil {
		return nil, err
	}

	return &lbcResult, nil
}

func littleBigConnectionApiListGetter(startIndex int64, resultLength int64, token string, cookie string, keywords string) ([]byte, error) {
	var requestFilter LbcFilters

	keyWordsArray := strings.Split(keywords, "|")

	requestFilter.SortedBy = "publicationDate"
	requestFilter.Keywords = make([]string, len(keyWordsArray))
	for indx, keyword := range keyWordsArray {
		requestFilter.Keywords[indx] = keyword
	}
	requestFilter.PublicationDate = "anytime"

	postBodyJson, _ := json.Marshal(requestFilter)
	postBodyArray := bytes.NewBuffer(postBodyJson)

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	urlApi := "https://www.littlebigconnection.com/api/v1.0/supplier/requestsForProposal/search?startAt=" + strconv.FormatInt(startIndex, 10) + "&maxResults=" + strconv.FormatInt(resultLength, 10)

	fmt.Println(urlApi)

	request, err := http.NewRequest("POST", urlApi, postBodyArray)

	if err != nil {
		return nil, err
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("cookie", cookie)
	request.Header.Add("csrfToken", token)

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func littlebigconnectionMissionsInformationsGetter(missionsHeader []Result, token string, cookie string) map[string]*LbcDetail {
	details := make(map[string]*LbcDetail)
	for _, mission := range missionsHeader {
		htmlMissionContent, err := littleBigConnectionMissionDetailsGetter(mission.RequestForProposal.EncodedID, token, cookie)
		if err != nil {
			continue
		}
		missionDetail := LbcDetail{}
		missionDetail.EncodedID = mission.RequestForProposal.EncodedID
		missionDetail.missionTextToDetailObjectConverter(htmlMissionContent)
		details[missionDetail.EncodedID] = &missionDetail
	}
	return details
}

func littleBigConnectionMissionDetailsGetter(misionEncodedID string, token string, cookie string) (string, error) {
	fileName := "MissionDetails/" + misionEncodedID + ".txt"

	if _, err := os.Stat(fileName); err == nil {
		fmt.Println("Reading file " + fileName)
		content, err := ioutil.ReadFile(fileName)

		if err != nil {
			return "", err
		}

		return string(content), nil
	}

	url := "https://www.littlebigconnection.com/rfp/" + misionEncodedID + "#/"

	httpClient := http.Client{
		Timeout: time.Duration(10 * time.Second),
	}

	fmt.Println(url)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	request.Header.Add("content-type", "application/json")
	request.Header.Add("cookie", cookie)
	request.Header.Add("csrfToken", token)

	response, err := httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	text, err := html2text.FromString(string(body), html2text.Options{PrettyTables: false})
	if err != nil {
		return "", err
	}

	f, err := os.Create(fileName)

	if err != nil {
		return "", err
	}

	defer f.Close()

	_, err2 := f.WriteString(text)

	if err2 != nil {
		return "", err
	}

	return text, nil
}
