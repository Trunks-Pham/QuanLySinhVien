package main

import (
	"fmt"
	"strings"
)

// Định nghĩa struct SinhVien
type SinhVien struct {
	ID     string
	Name   string
	Gender string
	Age    int
}

// Khởi tạo một map để quản lý tập hợp sinh viên
var sinhVienMap = make(map[string]SinhVien)

// Hàm thêm sinh viên
func themSinhVien(sv SinhVien) {
	sinhVienMap[sv.ID] = sv
}

// Hàm tìm kiếm sinh viên theo ID
func timKiemSinhVien(id string) (SinhVien, bool) {
	sv, exists := sinhVienMap[id]
	return sv, exists
}

// Hàm sửa thông tin sinh viên
func suaSinhVien(id string, sv SinhVien) {
	if _, exists := sinhVienMap[id]; exists {
		sinhVienMap[id] = sv
	} else {
		fmt.Println("Sinh viên không tồn tại!")
	}
}

// Hàm xóa sinh viên
func xoaSinhVien(id string) {
	if _, exists := sinhVienMap[id]; exists {
		delete(sinhVienMap, id)
	} else {
		fmt.Println("Sinh viên không tồn tại!")
	}
}

// Hàm xuất danh sách sinh viên theo giới tính
func xuatSinhVienTheoGioiTinh(gender string) {
	fmt.Printf("\nDanh sách sinh viên giới tính %s:\n", gender)
	for _, sv := range sinhVienMap {
		if strings.ToLower(sv.Gender) == strings.ToLower(gender) {
			fmt.Println(sv)
		}
	}
}

func main() {
	// Thêm sinh viên
	themSinhVien(SinhVien{ID: "001", Name: "Phạm Minh Thảo", Gender: "Nam", Age: 20})
	themSinhVien(SinhVien{ID: "002", Name: "Lê Thị Một", Gender: "Nữ", Age: 20})
	themSinhVien(SinhVien{ID: "003", Name: "Nguyễn Văn Hai", Gender: "Nam", Age: 22})
	themSinhVien(SinhVien{ID: "004", Name: "Nguyễn Thi Nguyễn", Gender: "Nữ", Age: 22})
	themSinhVien(SinhVien{ID: "005", Name: "Nguyễn Thị Ngọc Ngân", Gender: "Nữ", Age: 20})

	// Tìm kiếm sinh viên
	sv, found := timKiemSinhVien("001")
	if found {
		fmt.Println("Tìm thấy sinh viên:", sv)
	} else {
		fmt.Println("Không tìm thấy sinh viên!")
	}

	// Sửa sinh viên
	suaSinhVien("002", SinhVien{ID: "002", Name: "Lê Thị Một Úp Dát", Gender: "Nữ", Age: 22})

	// Xóa sinh viên
	xoaSinhVien("003")

	// Xuất sinh viên theo giới tính
	xuatSinhVienTheoGioiTinh("Nam")
	xuatSinhVienTheoGioiTinh("Nữ")

	fmt.Println("\n")
}
