import React from "react";
import {useGetUserByIDQuery} from "../../services/service";
import "./TopPanelButtons.css";

const TopPanelButtons = (props) => {
    const username = localStorage.getItem("userName");
    if (props.isAuth === false) {
        return (<div className="auth">
                    <div className="login">
                        <a className="login-button" href="/login">Войти</a> 
                    </div>
                    <div className="register">
                        <a className="register-button" href="/register">Зарегистрироваться</a>
                    </div>
                </div>)
    } else {
        return (<div className="profile">
                    <a className="profile-button" href={"/user/" + username}>
                        <strong style={{marginLeft: "1vw"}}>Мой профиль</strong><span style={{color: '#999', marginRight: "1vw"}}>{" (" + username + ")"}</span>
                    </a> 
                </div>);
    }
}

export default TopPanelButtons;