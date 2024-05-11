package repository

import (
	"fmt"
	"strings"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

func (r *ProfilePostgres) CreateProfile(userId int, profile chat.Profile) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var email string
	query := fmt.Sprintf("SELECT tl.email FROM %s tl WHERE tl.id=$1", usersTable)
	if err := r.db.Get(&email, query, userId); err != nil {
		fmt.Println("ЛСП")
		tx.Rollback()
		return 0, err
	}

	var id int
	query = fmt.Sprintf("INSERT INTO %s (name, surname, email, photo, city, birthday) values ($1, $2, $3, $4, $5, $6) RETURNING id", usersProfileTable)
	row := r.db.QueryRow(query, profile.Name, profile.Surname, email, profile.Photo, profile.City, profile.Birthday)
	if err := row.Scan(&id); err != nil {
		fmt.Println("ФИЛИПП КИРКОРОВ")
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, profile_id) VALUES ($1, $2)", usersProfileListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		fmt.Println("ТИМАТИ")
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ProfilePostgres) GetProfileId(userId int) (int, error) {
	var prof_id int

	query := fmt.Sprintf(`SELECT tl.profile_id FROM %s tl WHERE tl.user_id=$1`, usersProfileListsTable)
	err := r.db.Get(&prof_id, query, userId)

	return prof_id, err
}

func (r *ProfilePostgres) GetProfile(userId, profileId int) (chat.Profile, error) {
	var profile chat.Profile

	query := fmt.Sprintf(`SELECT tl.id, tl.name, tl.surname, tl.email, tl.photo, tl.city, tl.birthday FROM %s tl INNER JOIN %s ul on tl.id = ul.profile_id WHERE ul.user_id = $1 AND ul.profile_id = $2`,
		usersProfileTable, usersProfileListsTable)
	err := r.db.Get(&profile, query, userId, profileId)

	return profile, err
}

func (r *ProfilePostgres) EditProfile(userId, profileId int, input chat.UpdateProfile) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *input.Surname)
		argId++
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *input.Email)
		argId++
		query := fmt.Sprintf("UPDATE %s tl SET email=$1 WHERE tl.id=$2", usersTable)
		if _, err := r.db.Exec(query, input.Email, userId); err != nil {
			return err
		}
	}

	if input.Photo != nil {
		setValues = append(setValues, fmt.Sprintf("photo=$%d", argId))
		args = append(args, *input.Photo)
		argId++
	}

	if input.City != nil {
		setValues = append(setValues, fmt.Sprintf("city=$%d", argId))
		args = append(args, *input.City)
		argId++
	}

	if input.Birthday != nil {
		setValues = append(setValues, fmt.Sprintf("birthday=$%d", argId))
		args = append(args, *input.Birthday)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.profile_id AND ul.profile_id=$%d AND ul.user_id=$%d",
		usersProfileTable, setQuery, usersProfileListsTable, argId, argId+1)
	args = append(args, profileId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

// func (r *ProfilePostgres) UploadAvatar(profileId int, directory string) error {
// 	tx, err := r.db.Begin()
// 	if err != nil {
// 		return err
// 	}

// 	createUsersListQuery := fmt.Sprintf("INSERT INTO %s photo VALUES $1", usersProfileTable)
// 	_, err = tx.Exec(createUsersListQuery, directory)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	return tx.Commit()
// }

func (r *ProfilePostgres) CreateHobby(profId int, hobbies map[string][]chat.UserHobbyInput) ([]int, error) {
	var list_id []int
	var id int

	tx, err := r.db.Begin()
	if err != nil {
		return list_id, err
	}
	fmt.Println("11111111111111 ", hobbies)

	createListQuery := fmt.Sprintf("SELECT tl.id FROM %s tl WHERE tl.description=$1", userHobbyTable)
	list_hobbies := make([]string, 0, len(hobbies))
	for _, values := range hobbies {
		for i := 0; i < len(values); i++ {
			list_hobbies = append(list_hobbies, string(values[i].Description))
		}
	}
	var lengthOfHobbies = len(list_hobbies)
	for i := 0; i < lengthOfHobbies; i++ {
		if err := r.db.Get(&id, createListQuery, list_hobbies[i]); err != nil {
			tx.Rollback()
			return list_id, err
		}
		list_id = append(list_id, id)
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (prof_id, userhobby_id) VALUES ($1, $2)", usersHobbyListsTable)
	for i := 0; i < lengthOfHobbies; i++ {
		fmt.Println(profId, "::", list_id[i])
		_, err = tx.Exec(createUsersListQuery, profId, list_id[i])
		if err != nil {
			fmt.Println("ОШИБКА в ХОБИБИ 2")
			tx.Rollback()
			return list_id, err
		}
	}

	return list_id, tx.Commit()
}

func (r *ProfilePostgres) GetAllHobby(profId int) ([]chat.UserHobby, error) {
	var hobbylist []chat.UserHobby

	query := fmt.Sprintf("SELECT tl.id, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.userhobby_id WHERE ul.prof_id = $1",
		userHobbyTable, usersHobbyListsTable)
	err := r.db.Select(&hobbylist, query, profId)

	return hobbylist, err
}

func (r *ProfilePostgres) DeleteHobby(profId, hobbyId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.userhobby_id AND ul.prof_id=$1 AND ul.userhobby_id=$2",
		userHobbyTable, usersHobbyListsTable)
	_, err := r.db.Exec(query, profId, hobbyId)

	return err
}

func (r *ProfilePostgres) InitAllHobbies() error {

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	createListsHobbies := fmt.Sprintf("INSERT INTO %s (description) VALUES ($1)", userHobbyTable)

	var hobbies []chat.UserHobbyInput
	var hobby chat.UserHobbyInput
	{
		hobby.Description = "NFL"
		hobbies = append(hobbies, hobby)
		hobby.Description = "NBA"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Мировые новости"
		hobbies = append(hobbies, hobby)
		hobby.Description = "ChatGPT"
		hobbies = append(hobbies, hobby)
		hobby.Description = "One Piece"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Midjourney"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Cплетни о знаменитостях"
		hobbies = append(hobbies, hobby)

		hobby.Description = "Call of Duty"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Baldur’s Gate 3"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Minecraft"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Playstation"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Genshin Impact"
		hobbies = append(hobbies, hobby)
		hobby.Description = "GTA"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Sims"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Terraria"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Red Dead Redemption"
		hobbies = append(hobbies, hobby)

		hobby.Description = "Фотография"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Видеосъемка"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Дизайн"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Макияж"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Рукоделие"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Танцы"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Пения"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Музыка"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Ведение блога"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Рисование"
		hobbies = append(hobbies, hobby)

		hobby.Description = "Космос"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Природа"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Новости про науку"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Крутые гайды"
		hobbies = append(hobbies, hobby)
		hobby.Description = "История"
		hobbies = append(hobbies, hobby)
		hobby.Description = "География"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Философия"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Математика"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Computer Science"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Экономика"
		hobbies = append(hobbies, hobby)

		hobby.Description = "Бег"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Фитнес"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Велосипед"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Верховая езда"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Лыжи"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Йога"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Пилатес"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Сноуборд"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Ролики"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Скейтборд"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Самокат"
		hobbies = append(hobbies, hobby)

		hobby.Description = "Кинотеатр"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Концерты и шоу"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Музеи и галереи"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Театры"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Отдых на природе"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Фестивали"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Волонтество"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Тусовки и клубы"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Шопинг"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Путешествия"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Встречи с друзьями"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Искусство"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Активный отдых"
		hobbies = append(hobbies, hobby)

		hobby.Description = "Футбол"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Плавание"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Велоспорт"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Волейбол"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Баскетбол"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Хоккей"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Легкая атлетика"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Бокс"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Фигурное катание"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Гимнастика"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Лыжные гонки"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Киберспорт"
		hobbies = append(hobbies, hobby)

		hobby.Description = "Комедии"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Мультфильмы"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Исторические"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Детективы"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Приключения"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Ужасы"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Драмы"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Мелодрамы"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Триллеры"
		hobbies = append(hobbies, hobby)
		hobby.Description = "Боевики"
		hobbies = append(hobbies, hobby)
	}

	var lengthOfHobbies = len(hobbies)
	for i := 0; i < lengthOfHobbies; i++ {
		_, err = tx.Exec(createListsHobbies, hobbies[i].Description)
		fmt.Print(hobbies[i])
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
