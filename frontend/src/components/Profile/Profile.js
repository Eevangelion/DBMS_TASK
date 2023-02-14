import React from "react";
import { Link, useLocation } from 'react-router-dom';
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
        <div className="profile-block">
            <strong>{props.username}</strong>
            <strong style={{color: "#999"}}>user/{props.username}</strong>
            <div className={"settings-link"}>
                {
                    userAccount ?
                        <Link   to={`/settings`} 
                                style={linkStyle}
                                state={{ backgroundLocation: location }}>
                            <strong>Настройки</strong>
                        </Link> : 
                        <Link 
                                to={`/settings`}
                                style={disabledLinkStyle}
                                onClick={ (event) => event.preventDefault() }
                                state={{ backgroundLocation: location }}>
                            <strong>Настройки</strong>
                        </Link>
                }
            </div>
            <div className="profile-info">
                Роль: {role}<br/>
                Жалобы: {reports}<br/>
                {/* Добавлено в избранное: {addedToFavorite.length} <br/> */}
                Последняя дата разблокировки: {(Date.now() - Date.parse(lastUnbanDate))/1000}
            </div>
            <div className="post-joke">
                {
                    userAccount ?
                        <Link to={`/create_joke`} 
                            style={linkStyle}
                            state={{ backgroundLocation: location }}>
                            <strong>Создать шутку</strong>
                        </Link> : 
                        <Link to={`/create_joke`} 
                            style={disabledLinkStyle}
                            state={{ backgroundLocation: location }}
                            onClick={ (event) => event.preventDefault() }>
                            <strong>Создать шутку</strong>
                        </Link>
                } 
            </div>
        </div>
    );
}

export default Profile;