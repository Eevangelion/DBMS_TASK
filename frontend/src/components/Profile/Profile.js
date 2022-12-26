import React from "react";
import { Link, useLocation } from 'react-router-dom';
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

function Profile() {
    const location = useLocation();
    const username = "Ivan";
    return (
        <div className="profile-block">
            <strong>Ivan</strong>
            <strong style={{color: "#999"}}>user/{username}</strong>
            <div className="settings">
                <Link to={`/settings`} style={linkStyle}>
                    <strong>Настройки</strong>
                </Link>
            </div>
            <div className="profile-info">
                Роль: пользователь<br/>
                Жалобы: 2 <br/>
                Добавлено в избранное: 10 <br/> 
                Последняя дата разблокировки: 17.11.2022
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