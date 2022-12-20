package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "********"
	dbname   = "db_name"
)

func checkError(err error) { // прописываем ошибки
	if err != nil {
		panic(err)
	}
}

func connectTo() string { // connect to db

	sqlConn := fmt.Sprintf("host= %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)
	return sqlConn
}

func openDb() *sql.DB { // open db

	db, err := sql.Open("postgres", connectTo())
	checkError(err)

	return db
}

func show_cluch_dom() { // выводим запись по ключу для таблицы дом
	var etazhnost int
	var adres, otop string
	fmt.Print("Введите адрес:\n")
	fmt.Fscan(os.Stdin, &adres)
	rows, err := openDb().Query(`SELECT * from "BdHouse"."дом" where "адрес"=$1`, &adres)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&etazhnost, &adres, &otop) // сканируем (количество) записи(строки)
		checkError(err)

		fmt.Println("Этажность дома:", etazhnost, "Индивидуальное отопление: ", otop)
	}
	checkError(err)
}

func show_cluch_kvartira() { // выводим запись по ключу для таблицы квартира
	var ploshad, etaz, number int
	fmt.Print("Введите номер квартиры:\n")
	fmt.Fscan(os.Stdin, &number)
	rows, err := openDb().Query(`SELECT * from "BdHouse"."квартира" where "номер"=$1`, &number)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&ploshad, &etaz, &number) // сканируем (количество) записи(строки)
		checkError(err)

		fmt.Println("Площадь квартиры: ", ploshad, "Этаж квартиры: ", etaz)
	}
	checkError(err)
}

func show_cluch_comp() { // выводим запись по ключу для таблицы УК
	var name_cp, ur_adres, inN string
	fmt.Print("Введите ИНН компании:\n")
	fmt.Fscan(os.Stdin, &inN)
	rows, err := openDb().Query(`SELECT * from "BdHouse"."УК" where "ИНН"=$1`, &inN)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&name_cp, &ur_adres, &inN) // сканируем (количество) записи(строки)
		checkError(err)

		fmt.Println("Название компании: ", name_cp, "Юр.адрес компании: ", ur_adres)
	}
	checkError(err)
}

func show_cluch_plata() { // выводим запись по ключу для таблицы квартплата
	var oplata int
	var time string
	fmt.Print("Введите день оплаты:\n")
	fmt.Fscan(os.Stdin, &time)
	rows, err := openDb().Query(`SELECT * from "BdHouse"."квартплата" where "день_оплаты"=$1`, &time)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&time, &oplata) // сканируем (количество) записи(строки)
		checkError(err)

		fmt.Println("Сумма оплаты: ", oplata)
	}
	checkError(err)
}

func show_cluch_jilec() { // выводим запись по ключу для таблицы жилец
	var pasport int
	var fio, telefon, lgot string
	fmt.Print("Введите номер паспорта:\n")
	fmt.Fscan(os.Stdin, &pasport)
	rows, err := openDb().Query(`SELECT * from "BdHouse"."жилец" where "с-н_паспорта"=$1`, &pasport)
	checkError(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&pasport, &fio, &telefon, &lgot) // сканируем (количество) записи(строки)
		checkError(err)

		fmt.Println("ФИО жильца: ", fio, "Телефон жильца: ", telefon, "Льготы: ", lgot)
	}
	checkError(err)
}

func show_table_dom() { // показываем таблицу дом
	rows, err := openDb().Query(`SELECT "этажность", "адрес", "индивидуальное_отопление" FROM "BdHouse"."дом"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var etazhnost int
		var adres, otop string

		err = rows.Scan(&etazhnost, &adres, &otop)
		checkError(err)

		fmt.Println("Адрес дома: ", adres, "Этажность дома: ", etazhnost, " Есть ли отопление: ", otop)
	}

	checkError(err)
}

func show_table_kvartira() { // показываем таблицу квартира
	rows, err := openDb().Query(`SELECT "площадь", "этаж", "номер" FROM "BdHouse"."квартира"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var ploshad, etaz, number int

		err = rows.Scan(&ploshad, &etaz, &number)
		checkError(err)

		fmt.Println("Номер квартиры:", number, "Площадь квартиры: ", ploshad, "Этаж квартиры: ", etaz)
	}

	checkError(err)
}

func show_table_comp() { // показываем таблицу компания
	rows, err := openDb().Query(`SELECT "название_кп", "юр_адрес", "ИНН" FROM "BdHouse"."УК"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {

		var name_cp, ur_adres, inN string

		err = rows.Scan(&name_cp, &ur_adres, &inN)
		checkError(err)

		fmt.Println("ИНН компании", inN, "Название компании", name_cp, "Юр.адрес компании", ur_adres)
	}

	checkError(err)
}

func show_table_plata() { // показываем таблицу квартплата
	rows, err := openDb().Query(`SELECT "день_оплаты", "цена" FROM "BdHouse"."квартплата"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var plata int
		var time string

		err = rows.Scan(&time, &plata)
		checkError(err)

		fmt.Println("День оплаты: ", time, "Цена: ", plata)
	}

	checkError(err)
}

func show_table_jilec() { // показываем таблицу жилец
	rows, err := openDb().Query(`SELECT "с-н_паспорта", "ФИО", "телефон", "льготы" FROM "BdHouse"."жилец"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		var pasport int
		var fio, telefon, lgot string

		err = rows.Scan(&pasport, &fio, &telefon, &lgot)
		checkError(err)

		fmt.Println("номер паспорта жильца: ", pasport, "ФИО жильца: ", fio, "Номер телефона жильца:", telefon, "Льготы: ", lgot)
	}

	checkError(err)
}

func delete_key_dom() { // удаляем из таблицы дом
	var check string
	fmt.Print("какую запись удалить? введите адрес дома:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdHouse"."дом" where "адрес"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delete_key_kvartira() { // удаляем из таблицы квартира
	var check int
	fmt.Print("какую запись удалить? введите номер квартиры:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdHouse"."квартира" where "номер"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delete_key_comp() { // удаляем из таблицы компания
	var check string
	fmt.Print("какую запись удалить? введите ИНН компании:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdHouse"."УК" where "ИНН"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delete_key_plata() { // удаляем из таблицы квартплата
	var check int
	fmt.Print("какую запись удалить? День оплаты:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdHouse"."квартплата" where "день_оплаты"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func delete_key_jilec() { // удаляем из таблицы жилец
	var check int
	fmt.Print("какую запись удалить? введите номер паспорта:\n")
	fmt.Fscan(os.Stdin, &check)

	deleteS := `delete from "BdHouse"."жилец" where "с-н_паспорта"=$1`
	_, e := openDb().Exec(deleteS, &check)
	checkError(e)
}

func update_table_kvartira() { // обновляем запись в таблице квартира
	var ploshad, etaz, number int

	fmt.Print("Введите номер квартиры для обновления:\n")
	fmt.Fscan(os.Stdin, &number)

	fmt.Print("Введите новую площадь:\n")
	fmt.Fscan(os.Stdin, &ploshad)

	fmt.Print("введите новый этаж: \n")
	fmt.Fscan(os.Stdin, &etaz)

	updateStm := `update "BdHouse"."квартира" set "площадь"=$1, "этаж"=$2 where "номер"=$3`
	_, e := openDb().Exec(updateStm, &ploshad, &etaz, &number)
	checkError(e)
}

func update_table_dom() { // обновляем запись в таблице дом
	var adres, otop string
	var etazhnost int

	fmt.Print("Введите адрес для обновления:\n")
	fmt.Fscan(os.Stdin, &adres)

	fmt.Print("Введите новую этажность:\n")
	fmt.Fscan(os.Stdin, &etazhnost)

	fmt.Print("есть ли отопление? да/нет: \n")
	fmt.Fscan(os.Stdin, &otop)

	updateStm := `update "BdHouse"."дом" set "этажность"=$1, "индивидуальное_отопление"=$2 where "адрес"=$3`
	_, e := openDb().Exec(updateStm, &etazhnost, &otop, &adres)
	checkError(e)
}

func update_table_company() { // обновляем запись в таблице компания
	var nameCp, urAdres, inN string

	fmt.Print("Введите ИНН для обновления:\n")
	fmt.Fscan(os.Stdin, &inN)

	fmt.Print("Введите новое название компании:\n")
	fmt.Fscan(os.Stdin, &nameCp)

	fmt.Print("введите новыый юр-адрес: \n")
	fmt.Fscan(os.Stdin, &urAdres)

	updateStm := `update "BdHouse"."УК" set "название_кп"=$1, "юр_адрес"=$2 where "ИНН"=$3`
	_, e := openDb().Exec(updateStm, &nameCp, &urAdres, &inN)
	checkError(e)
}

func update_table_Plata() { // обновляем запись в таблице квартплата
	var date string
	var plata int

	fmt.Print("Введите день оплаты:\n")
	fmt.Fscan(os.Stdin, &date)

	fmt.Print("Введите цену:\n")
	fmt.Fscan(os.Stdin, &plata)

	updateStm := `update "BdHouse"."квартплата" set "цена"=$1 where "день_оплаты"=$2`
	_, e := openDb().Exec(updateStm, &plata, &date)
	checkError(e)
}

func update_table_Jilec() { // обновляем запись в таблице жилец
	var fio, lgot string
	var telefon, pasport int

	fmt.Print("Введите номер паспорта:\n")
	fmt.Fscan(os.Stdin, &pasport)

	fmt.Print("Введите новый номер телефона:\n")
	fmt.Fscan(os.Stdin, &telefon)

	fmt.Print("Введите новые льготы: \n")
	fmt.Fscan(os.Stdin, &lgot)

	fmt.Print("Введите новые ФИО: \n")
	fmt.Fscan(os.Stdin, &fio)

	updateStm := `update "BdHouse"."жилец" set "ФИО"=$1, "телефон"=$2, "льготы"=$3 where "с-н_паспорта"=$4`
	_, e := openDb().Exec(updateStm, &fio, &telefon, &lgot, &pasport)
	checkError(e)
}

func add_znach_to_dom() { // insert function enter value to dom table
	var adres, otop string
	var etazhnost int
	fmt.Print("Введите адрес:\n")
	fmt.Fscan(os.Stdin, &adres)

	fmt.Print("Введите этажность:\n")
	fmt.Fscan(os.Stdin, &etazhnost)

	fmt.Print("есть ли отопление? да/нет: \n")
	fmt.Fscan(os.Stdin, &otop)

	insertToDyn := `insert into "BdHouse"."дом"("этажность","адрес","индивидуальное_отопление") values($1, $2, $3)`
	_, e := openDb().Exec(insertToDyn, &etazhnost, &adres, &otop)
	checkError(e)

}

func add_znach_to_kvartira() { // insert function enter value to kvartira table
	var ploshad, etaz, nomer int
	fmt.Print("Введите площадь квартиры: \n")
	fmt.Fscan(os.Stdin, &ploshad)

	fmt.Print("Введите этаж квартиры: \n")
	fmt.Fscan(os.Stdin, &etaz)

	fmt.Print("Введите номер квартиры: \n")
	fmt.Fscan(os.Stdin, &nomer)

	insertToDyn := `insert into "BdHouse"."квартира"("площадь","этаж","номер") values($1, $2, $3)`
	_, e := openDb().Exec(insertToDyn, &ploshad, &etaz, &nomer)
	checkError(e)

}

func add_znach_to_company() { // insert function enter value to company table
	var name_cp, ur_adres, inN string
	fmt.Print("Введите название компании: \n")
	fmt.Fscan(os.Stdin, &name_cp)

	fmt.Print("Введите юр.адрес: \n")
	fmt.Fscan(os.Stdin, &ur_adres)

	fmt.Print("Введите Инн: \n")
	fmt.Fscan(os.Stdin, &inN)

	insertToDyn := `insert into "BdHouse"."УК"("название_кп","юр_адрес","ИНН") values($1, $2, $3)`
	_, e := openDb().Exec(insertToDyn, &name_cp, &ur_adres, &inN)
	checkError(e)

}

func add_znach_to_plata() { // insert function enter value to квартплата table
	var  date string
	var plata int
	fmt.Print("День оплаты: \n")
	fmt.Fscan(os.Stdin, &date)

	fmt.Print("Цена: \n")
	fmt.Fscan(os.Stdin, &plata)

	insertToDyn := `insert into "BdHouse"."квартплата"("день_оплаты","цена") values($1, $2)`
	_, e := openDb().Exec(insertToDyn, &date, &plata)
	checkError(e)

}

func add_key_to_jilec() { // insert function enter value to жилец table
	var lgot, fio string
	var number_tel, pasport int
	fmt.Print("серия номер паспорта: \n")
	fmt.Fscan(os.Stdin, &pasport)

	fmt.Print("ФИО: \n")
	fmt.Fscan(os.Stdin, &fio)

	fmt.Print("телефон: \n")
	fmt.Fscan(os.Stdin, &number_tel)

	fmt.Print("льготы: \n")
	fmt.Fscan(os.Stdin, &lgot)

	insertToDyn := `insert into "BdHouse"."жилец"("с-н_паспорта","ФИО","телефон","льготы") values($1, $2, $3, $4)`
	_, e := openDb().Exec(insertToDyn, &pasport, &fio, &number_tel, &lgot)
	checkError(e)

}

func add_switch_case() { // функция выбора таблицы для добавления записи
	var vibor string
	fmt.Print("выберите дейсвтие: add_home -  чтобы добавить данные в таблицу дом\n add_kvartira - добавить данные в таблицу квартира\n add_company добавить данные в таблицу квартира\n add_plata - добавить данные в таблицу квартплата\n add_jilec - добавить данные в таблицу жилец\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "add_home":
		add_znach_to_dom()
	case "add_kvartira":
		add_znach_to_kvartira()
	case "add_company":
		add_znach_to_company()
	case "add_plata":
		add_znach_to_plata()
	case "add_jilec":
		add_key_to_jilec()
	default:
		fmt.Println("Неверная команда...")
	}
}

func update_switch_case() { // функция выбора таблицы для обеовления
	var vibor string
	fmt.Print("выберите дейсвтие: updateHome -  чтобы обновить данные в таблице дом\n updateKvartira - обновить данные в таблице квартира\n updateCompany обновить данные в таблице компания\n updatePlata - обновить данные в таблице квартплата\n updateJilec - обновить данные в таблице жилец\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "updateHome":
		update_table_dom()
	case "updateKvartira":
		update_table_kvartira()
	case "updateCompany":
		update_table_company()
	case "updatePlata":
		update_table_Plata()
	case "updateJilec":
		update_table_Jilec()
	default:
		fmt.Println("Неверная команда...")
	}
}

func delete_switch_case() { // функция выбора таблицы для удаления записи
	var vibor string
	fmt.Print("выберите дейсвтие: delDom -  чтобы удалить данные из таблицы дом\n delKvartira - удалить данные из таблицы квартира\n delCompany удалить данные из таблицы компания\n delPlata - удалить данные из таблицы квартплата\n delJilec - удалить данные из таблицы жилец\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "delDom":
		delete_key_dom()
	case "delKvartira":
		delete_key_kvartira()
	case "delCompany":
		delete_key_comp()
	case "delPlata":
		delete_key_plata()
	case "delJilec":
		delete_key_jilec()
	default:
		fmt.Println("Неверная команда...")
	}
}

func show_switch_case() { // функция выбора таблицы для выводы данных
	var vibor string
	fmt.Print("выберите дейсвтие: showDom -  чтобы показать данные из таблицы дом\n showKvartira - показать данные из таблицы квартира\n showCompany показать данные из таблицы компания\n showPlata - показать данные из таблицы квартплата\n showJilec - показать данные из таблицы жилец\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "showDom":
		show_table_dom()
	case "showKvartira":
		show_table_kvartira()
	case "showCompany":
		show_table_comp()
	case "showPlata":
		show_table_plata()
	case "showJilec":
		show_table_jilec()
	default:
		fmt.Println("Неверная команда...")
	}
}

func show_klutch_switch_case() {
	var vibor string
	fmt.Print("выберите дейсвтие: showCluchDom -  чтобы показать данные из таблицы дом\n showCluchKvartira - показать данные из таблицы квартира\n showCluchCompany показать данные из таблицы компания\n showCluchPlata - показать данные из таблицы квартплата\n showCluchJilec - показать данные из таблицы жилец\n")
	fmt.Scanf("%s\n", &vibor)

	switch vibor {
	case "showCluchDom":
		show_cluch_dom()
	case "showCluchKvartira":
		show_cluch_kvartira()
	case "showCluchCompany":
		show_cluch_comp()
	case "showCluchPlata":
		show_cluch_plata()
	case "showCluchJilec":
		show_cluch_jilec()
	default:
		fmt.Println("Неверная команда...")
	}
}

func main() {
	var v1 string
	//close db
	defer openDb().Close()
	//check db
	err := openDb().Ping()
	checkError(err)

	fmt.Print("Что вы хотите сделать?\n Чтобы добавить значения в таблицу введите addTable\n Чтобы обновить запись введите updateTable\n Чтобы удалить запись из таблицы введите deleteFromTable\n Чтобы показать данные в таблице введите showTable\n Чтобы показать запись по ключу введите showKluch\n")
	fmt.Scanf("%s\n", &v1)

	switch v1 {
	case "addTable":
		add_switch_case()
	case "updateTable":
		update_switch_case()
	case "deleteFromTable":
		delete_switch_case()
	case "showTable":
		show_switch_case()
	case "showKluch":
		show_klutch_switch_case()
	default:
		fmt.Println("Неверная команда...")
	}
}