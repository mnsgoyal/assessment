package employee

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetEmployeeDetailHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	// getting each one
	id, err := strconv.Atoi(vars["ID"])
	var respData []byte
	if err != nil {
		errorString := Error{}
		errorString.Error = "Invalid Parameter ID err:" + err.Error()
		respData, _ = json.Marshal(errorString)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(respData)
		return
	}

	if empDetail, okay := EmpDetailsDB[id]; okay {
		respData, _ = json.Marshal(empDetail)
		w.WriteHeader(http.StatusOK)

	} else {
		respData = []byte("No record Found")
		w.WriteHeader(http.StatusNotFound)
	}

	w.Write(respData)
	return
}

func AddEmployeeDetailHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var empDetail EmpDetails
	var respData []byte
	err := decoder.Decode(&empDetail)
	if err != nil {
		errorString := Error{}
		errorString.Error = "Invalid Employee Data, err:" + err.Error()
		respData, _ = json.Marshal(errorString)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(respData)
		return
	}

	if empDetail.Name == "" {
		errorString := Error{}
		errorString.Error = "Invalid Employee Name"
		respData, _ = json.Marshal(errorString)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respData)
		return
	}

	if empDetail.Position == "" {
		errorString := Error{}
		errorString.Error = "Invalid Employee Position"
		respData, _ = json.Marshal(errorString)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respData)
		return
	}

	EmpIDCounter++

	empDetail.ID = EmpIDCounter
	EmpDetailsDB[EmpIDCounter] = empDetail

	respData = []byte("Record Added successfully!")
	w.Write(respData)
	return
}

func DeleteEmployeeDetailHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	// getting each one
	id, err := strconv.Atoi(vars["ID"])
	var respData []byte
	if err != nil {
		errorString := Error{}
		errorString.Error = "Invalid Parameter ID err:" + err.Error()
		respData, _ = json.Marshal(errorString)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(respData)
		return
	}

	if _, okay := EmpDetailsDB[id]; okay {
		delete(EmpDetailsDB, id)
		respData = []byte("Record has been deleted")
		w.WriteHeader(http.StatusOK)

	} else {
		respData = []byte("No record Found")
		w.WriteHeader(http.StatusNotFound)
	}

	w.Write(respData)
	return
}

func UpdateEmployeeDetailHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var empDetail EmpDetails
	var respData []byte
	err := decoder.Decode(&empDetail)
	if err != nil {
		errorString := Error{}
		errorString.Error = "Invalid Employee Data, err:" + err.Error()
		respData, _ = json.Marshal(errorString)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(respData)
		return
	}

	if empDetail.Name == "" {
		errorString := Error{}
		errorString.Error = "Invalid Employee Name"
		respData, _ = json.Marshal(errorString)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respData)
		return
	}

	if empDetail.Position == "" {
		errorString := Error{}
		errorString.Error = "Invalid Employee Position"
		respData, _ = json.Marshal(errorString)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(respData)
		return
	}

	if _, okay := EmpDetailsDB[empDetail.ID]; !okay {
		respData = []byte("Invalid employee ID")
		w.WriteHeader(http.StatusNotFound)
		w.Write(respData)
		return
	}

	EmpDetailsDB[empDetail.ID] = empDetail

	respData = []byte("Record has been updated successfully!")
	w.WriteHeader(http.StatusOK)
	w.Write(respData)
	return
}

func ListEmployeeDetailHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var respData []byte
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		errorString := Error{}
		errorString.Error = "Invalid page number, err:" + err.Error()
		respData, _ = json.Marshal(errorString)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(respData)
		return
	}

	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		errorString := Error{}
		errorString.Error = "Invalid data limit, err:" + err.Error()
		respData, _ = json.Marshal(errorString)

		w.WriteHeader(http.StatusBadRequest)
		w.Write(respData)
		return
	}

	if len(EmpDetailsDB) > 0 {
		var empList []EmpDetails
		var objEmployeeList EmployeeList
		var objPaginationDetails PaginationDetails
		for _, data := range EmpDetailsDB {
			empList = append(empList, data)
		}

		totalPages := len(empList) / limit
		if len(empList)%limit != 0 {
			totalPages++
		}

		if totalPages < page || page == 0 {
			respData = []byte("Invalid page number")
			w.WriteHeader(http.StatusNotFound)
			w.Write(respData)
			return
		}

		if totalPages != page {
			objPaginationDetails.Next = true
		}

		if totalPages != 1 && page != 1 {
			objPaginationDetails.Previous = true
		}

		objPaginationDetails.TotalPages = totalPages
		objPaginationDetails.Limit = limit
		objPaginationDetails.Page = page
		objEmployeeList.PaginationDetails = objPaginationDetails
		start := (page - 1) * limit
		end := page * limit
		if end > len(empList) {
			end = len(empList)
		}
		objEmployeeList.EmpDetails = empList[start:end]
		respData, _ = json.Marshal(objEmployeeList)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respData)
		return
	}

	respData = []byte("No record Found")
	w.WriteHeader(http.StatusNotFound)
	w.Write(respData)
	return
}
