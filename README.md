# DBMS_TASK
Тестовое задание для дисциплин "Введение в системы баз данных" и "Технологии Интернет и Web-программирования"

# Описание
Урезанный аналог Reddit, направленный на публикацию шуток/анекдотов.

## Наименование
FunnyJokes

## Предметная область, данные
Шутки, анекдоты

## Данные и их ограничения
https://drawsql.app/teams/test-team-30/diagrams/jokesdiagram

## Общие ограничения целостности
- Тэги могут добавлять в базу данных/удалять из базы данных только администраторы
- Каждый пользователь может добавлять шутки только в своё избранное
- Рейтинг - неотрицательное целое число

# Пользовательские роли
- Администратор
- Пользователь
- Неавторизированный пользователь

## Для каждой роли - наименование, ответственность, количество пользователей в этой роли?
- Администратор - рассматривает жалобы пользователей посредством интерфейса на своей странице. Может добавлять новые тэги и блокировать пользователей (заблокированные пользователи не смогут публиковать свои шутки на протяжении 7 дней с начала бана)
- Пользователь - может добавлять/удалять собственные шутки на сайт, просматривать шутки других, добавлять их в избранное или отправлять на них жалобы (максимум 3 раза за день)
- Неавторизированный пользователь - не имеет доступ ни к какому контенту, пока не авторизуется/зарегистрируется (любой запрос перенаправляет на страницу с регистрацией/авторизацией)

# UI 
- Страница ленты, где будут публиковаться шутки людей, на которых пользователь подписан (будет возможность отсортировать шутки по рейтингу за 24 часа, 1 неделя, 1 месяц, всё время)
- Поиск шуток/людей (интерфейс такой же, как у ленты, меняется только содержимое)
- Личная страница, где пользователь может опубликовать свои шутки (или же страница другого пользователя. отличие лишь в том, что на чужой странице будет ограничен доступ к публикации шуток)
- Страница с настройками. Для администраторов будет добавлена отдельная ссылка на интерфейс, который позволяет просматривать жалобы на других пользователей и добавлять тэги
- Создание шутки (всплывающее окно в личной странице)
- Создание жалобы

# Технологии разработки
- [Gorilla/Mux](https://github.com/gorilla/mux) для API
- [Pgx](https://github.com/jackc/pgx) для интеграции с PostgreSQL (не ORM)
- [React](https://ru.reactjs.org/)
- [Redux](https://redux.js.org/)
## Язык программирования
- Golang (Backend)
- JavaScript (Frontend)

## СУБД
PostgreSQL

# Тестирование
Postman
