import React from "react";
import { Link, useLocation } from 'react-router-dom';
import { useSelector } from 'react-redux';
import Button from "react-bootstrap/esm/Button";
import { useGetUserByNameQuery } from "../../services/Joke";
import "./Profile.css";

const linkStyle = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    textAlign: "center",
    width: "20vw",
    height: "2vh",
    borderRadius: "45vh",
    backgroundColor: "#00d",
    textDecoration : "none",
    borderColor: "transparent",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
    fontSize: "1.4vh",
}

const disabledLinkStyle = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    textAlign: "center",
    width: "20vw",
    height: "2vh",
    borderRadius: "45vh",
    backgroundColor: "#bbb",
    textDecoration : "none",
    borderColor: "transparent",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
    fontSize: "1.4vh",
}


const Profile = (props) => {
    const userPageIsActive = useSelector(state => state.pagesReducer.userPageIsActive);
    const location = useLocation();
    const {
        data: user,
        isLoading: loadingUser,
    } = useGetUserByNameQuery(props.username);
    const loadingFrame = <div className="profile-block">Загрузка...</div>;
    const noUserFrame = <div className="profile-block">Пользователя с таким именем не существует</div>;
    if (loadingUser) {
        return loadingFrame;
    }
    if (!user) {
        return noUserFrame;
    }
    let userAccount = true;
    if (props.username !== localStorage.getItem("userName")) {
        userAccount = false;
    }
    
    const reports = user.reports,
          lastUnbanDate = user.unban_date,
          role = user.role;
    return (
        <div className="profile-block" style={userPageIsActive ? {} : {backgroundColor: "#767676", border: "0.1vh solid #555"}}>
            <strong>{props.username}</strong>
            <strong style={{color: "#999"}}>user/{props.username}</strong>
            {
                userAccount ?
                <div className={"settings-link"}>
                    <Link   to={`/settings`} 
                            style={userPageIsActive ? linkStyle : disabledLinkStyle}
                            state={{ backgroundLocation: location }}
                            onClick={(event) => {if (!userPageIsActive) event.preventDefault()}}
                    >
                        <strong>Настройки</strong>
                    </Link> 
                </div> : <></>
            }
            <div className="profile-info">
                Роль: {role === "admin" ? "Администратор" : "Пользователь"}<br/>
                Жалобы: {reports}<br/>
                {/* Добавлено в избранное: {addedToFavorite.length} <br/> */}
                Последняя дата разблокировки: {(Date.now() - Date.parse(lastUnbanDate))/1000}
            </div>
            {
                userAccount ?
                <div className="post-joke">
                    <Link to={`/create_joke`} 
                        style={userPageIsActive ? linkStyle : disabledLinkStyle}
                        state={{ backgroundLocation: location }}
                        onClick={(event) => {if (!userPageIsActive) event.preventDefault()}}
                    >
                        <strong>Создать шутку</strong>
                    </Link>
                </div> : 
                <div className="subscribe">
                    <Link   to={`/subscribe/${user.id}`}
                            style={userPageIsActive ? linkStyle : disabledLinkStyle}
                            state={{ backgroundLocation: location}}
                            onClick={(event) => {if (!userPageIsActive) event.preventDefault()}}
                    >
                        <strong>Подписаться</strong>
                    </Link> 
                </div>
            }
        </div>
    );
}

export default Profile;