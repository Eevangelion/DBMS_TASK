import React from "react";
import { Link, useLocation } from 'react-router-dom';
import { useGetUserByNameQuery } from "../../services/User";
import "./Profile.css";

const linkStyle = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    textAlign: "center",
    width: "380px",
    height: "30px",
    borderRadius: "45px",
    backgroundColor: "#00d",
    textDecoration : "none",
    borderColor: "transparent",
    color: "white",
    fontWeight: "bold",
    fontFamily: "Arial, Helvetica, sans-serif",
}

const Profile = (username) => {
    const location = useLocation();
    const {
        data: user,
        isLoading: loadingUser,
    } = useGetUserByNameQuery(username);
    // const {
    //     data: addedToFavorite,
    //     isLoading: loadingFavoriteJokes,
    // } = useGetFavoritesByNameQuery(username);
    const loadingFrame = <div className="profile-block">Загрузка...</div>;
    const noUserFrame = <div className="profile-block">Пользователя с таким именем не существует</div>;
    if (loadingUser) {
        return loadingFrame;
    }
    if (!user) {
        return noUserFrame;
    }
    
    const reports = user.reports,
          lastUnbanDate = user.unban_date,
          role = user.role;
    return (
        <div className="profile-block">
            <strong>{username}</strong>
            <strong style={{color: "#999"}}>user/{username}</strong>
            <div className="settings-link">
                <Link to={`/settings`} style={linkStyle}>
                    <strong>Настройки</strong>
                </Link>
            </div>
            <div className="profile-info">
                Роль: {role}<br/>
                Жалобы: {reports}<br/>
                {/* Добавлено в избранное: {addedToFavorite.length} <br/> */}
                Последняя дата разблокировки: {lastUnbanDate}
            </div>
            <div className="post-joke">
                <Link to={`/create_joke`} 
                      style={linkStyle}
                      state={{ backgroundLocation: location }}>
                    <strong>Создать шутку</strong>
                </Link>
            </div>
        </div>
    );
}

export default Profile;